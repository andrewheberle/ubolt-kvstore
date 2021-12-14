package client

import (
	"context"
	"crypto/x509"
	"time"

	pb "gitlab.com/andrewheberle/ubolt-kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	address  string
	cert     string
	insecure bool
	kv       pb.KeystoreServiceClient
	conn     *grpc.ClientConn
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

func Connect(address, cert string, insecure bool) (*Client, error) {
	var err error
	var creds credentials.TransportCredentials
	var opts = []grpc.DialOption{
		grpc.WithBlock(),
	}

	// connect to grpc api
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		if cert == "" {
			cp, err := x509.SystemCertPool()
			if err != nil {
				return nil, err
			}
			creds = credentials.NewClientTLSFromCert(cp, "")
		} else {
			creds, err = credentials.NewClientTLSFromFile(cert, "")
			if err != nil {
				return nil, err
			}
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// dial grpc server
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{address: address, cert: cert, insecure: insecure, kv: pb.NewKeystoreServiceClient(conn), conn: conn}, nil
}

func (client *Client) Get(ctx context.Context, bucket, key string) ([]byte, error) {
	res, err := client.kv.Get(ctx, &pb.GetRequest{Bucket: bucket, Key: key}, grpc.EmptyCallOption{})
	if err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (client *Client) Put(ctx context.Context, bucket, key string, value []byte) error {
	_, err := client.kv.Put(ctx, &pb.PutRequest{Bucket: bucket, Key: key, Value: value}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Delete(ctx context.Context, bucket, key string) error {
	_, err := client.kv.Delete(ctx, &pb.DeleteRequest{Bucket: bucket, Key: key}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) CreateBucket(ctx context.Context, bucket string) error {
	_, err := client.kv.CreateBucket(ctx, &pb.BucketRequest{Bucket: bucket}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) DeleteBucket(ctx context.Context, bucket string) error {
	_, err := client.kv.DeleteBucket(ctx, &pb.BucketRequest{Bucket: bucket}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Close() error {
	return client.conn.Close()
}
