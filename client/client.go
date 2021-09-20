package client

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	pb "gitlab.com/andrewheberle/ubolt-kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	address  string
	cert     string
	insecure bool
}

type Result interface {
	IsError() bool
	Byte() []byte
	String() string
	Status() int
	JSON() []byte
}

type MessageResult struct {
	message string
	status  int
	err     bool
}

type KeyResult struct {
	data []byte
}

type KeyListResult struct {
	data []string
}

type KeyRequest struct {
	Value []byte `json:"value"`
}

func NewClient(address string) (*Client, error) {
	return &Client{address: address, insecure: true, cert: ""}, nil
}

func NewTLSClient(address, cert string) (*Client, error) {
	return &Client{address: address, insecure: false, cert: cert}, nil
}

func NewError(msg string, status int) *MessageResult {
	return &MessageResult{msg, status, true}
}

func NewMessage(msg string) *MessageResult {
	return &MessageResult{msg, http.StatusOK, false}
}

func (r *MessageResult) IsError() bool {
	return r.err
}

func (r *MessageResult) String() string {
	return r.message
}

func (r *MessageResult) Byte() []byte {
	return []byte(r.message)
}

func (r *MessageResult) Status() int {
	if r.status == 0 {
		return http.StatusBadRequest
	}

	return r.status
}

func (r *MessageResult) JSON() (b []byte) {
	if r.err {
		b, _ = json.Marshal(struct {
			Error string `json:"error"`
		}{r.message})
	} else {
		b, _ = json.Marshal(struct {
			Message string `json:"message"`
		}{r.message})
	}
	return
}

func NewKey(data []byte) *KeyResult {
	return &KeyResult{data: data}
}

func (r *KeyResult) IsError() bool {
	return false
}

func (r *KeyResult) String() string {
	return fmt.Sprintf("%s", r.data)
}

func (r *KeyResult) Byte() []byte {
	return r.data
}

func (r *KeyResult) Status() int {
	return http.StatusOK
}

func (r *KeyResult) JSON() []byte {
	b, _ := json.Marshal(struct {
		Data []byte `json:"value"`
	}{r.data})
	return b
}

func NewKeylist(data []string) *KeyListResult {
	return &KeyListResult{data: data}
}

func (r *KeyListResult) IsError() bool {
	return false
}

func (r *KeyListResult) String() string {
	return fmt.Sprintf("%s", r.data)
}

func (r *KeyListResult) Byte() []byte {
	return []byte(r.data[0])
}

func (r *KeyListResult) Status() int {
	return http.StatusOK
}

func (r *KeyListResult) JSON() []byte {
	b, _ := json.Marshal(struct {
		Data []string `json:"value"`
	}{r.data})
	return b
}

func (c *Client) Connect() (*grpc.ClientConn, error) {
	var err error
	var creds credentials.TransportCredentials

	// connect to grpc api
	if c.insecure {
		return grpc.Dial(c.address, grpc.WithInsecure(), grpc.WithBlock())
	}

	if c.cert == "" {
		cp, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		creds = credentials.NewClientTLSFromCert(cp, "")
	} else {
		creds, err = credentials.NewClientTLSFromFile(c.cert, "")
		if err != nil {
			return nil, err
		}
	}

	return grpc.Dial(c.address, grpc.WithBlock(), grpc.WithTransportCredentials(creds))
}

func (c *Client) HandleKey(w http.ResponseWriter, r *http.Request) {
	var result Result
	var key KeyRequest

	// grab provided vars from request
	vars := mux.Vars(r)

	// connect to grpc api
	conn, err := c.Connect()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	kv := pb.NewKeystoreServiceClient(conn)

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// handle request method
	switch r.Method {
	case "GET":
		// handle key get
		r, err := kv.Get(ctx, &pb.GetRequest{Bucket: vars["bucket"], Key: vars["key"]})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewKey(r.GetValue())
		}
	case "PUT":
		// handle key update/create
		if r.Body != nil {
			defer r.Body.Close()

			// decode body as json
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&key); err == nil {
				// check if value was provided
				if key.Value != nil {
					// put into kv store
					_, err := kv.Put(ctx, &pb.PutRequest{Bucket: vars["bucket"], Key: vars["key"], Value: key.Value})
					if err != nil {
						result = NewError(err.Error(), http.StatusInternalServerError)
					} else {
						result = NewMessage("key updated or set")
					}
				} else {
					// request is invalid without "value" being provided
					result = NewError("invalid request", http.StatusBadRequest)
				}
			} else {
				// error decoding json
				result = NewError("invalid request", http.StatusBadRequest)
			}
		} else {
			// no request body
			result = NewError("missing request body", http.StatusBadRequest)
		}
	case "DELETE":
		// handle key deletion
		_, err := kv.Delete(ctx, &pb.DeleteRequest{Bucket: vars["bucket"], Key: vars["key"]})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewMessage("key deleted")
		}
	default:
		// handle unsupported request
		result = NewError("Unsupported request", http.StatusBadRequest)
	}

	// respond back to client
	w.WriteHeader(result.Status())
	w.Write(result.JSON())
}

func (c *Client) HandleBucket(w http.ResponseWriter, r *http.Request) {
	var result Result

	// grab provided vars from request
	vars := mux.Vars(r)

	// connect to grpc api
	conn, err := grpc.Dial(c.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	kv := pb.NewKeystoreServiceClient(conn)

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// handle request method
	switch r.Method {
	case "GET":
		// handle GET
		res, err := kv.ListKeys(ctx, &pb.BucketRequest{Bucket: vars["bucket"]})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewKeylist(res.GetKeys())
		}
	case "PUT":
		// handle bucket creation
		_, err := kv.CreateBucket(ctx, &pb.BucketRequest{Bucket: vars["bucket"]})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewMessage("bucket created")
		}
	case "DELETE":
		// handle bucket deletion
		_, err := kv.DeleteBucket(ctx, &pb.BucketRequest{Bucket: vars["bucket"]})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewMessage("bucket deleted")
		}
	default:
		// handle unsupported request
		result = NewError("Unsupported request", http.StatusBadRequest)
	}

	// respond back to client
	w.WriteHeader(result.Status())
	w.Write(result.JSON())
}

func (c *Client) ListBuckets(w http.ResponseWriter, r *http.Request) {
	var result Result

	// connect to grpc api
	conn, err := grpc.Dial(c.address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	kv := pb.NewKeystoreServiceClient(conn)

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if r.Method == "GET" {
		res, err := kv.ListBuckets(ctx, &pb.Empty{})
		if err != nil {
			result = NewError(err.Error(), http.StatusInternalServerError)
		} else {
			result = NewKeylist(res.GetBuckets())
		}
	} else {
		// handle unsupported request
		result = NewError("Unsupported request", http.StatusBadRequest)
	}

	// respond back to client
	w.WriteHeader(result.Status())
	w.Write(result.JSON())
}
