package client

import (
	"context"
	"crypto/x509"

	pb "gitlab.com/andrewheberle/ubolt-kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// KvStoreClient
type KvStoreClient struct {
	kv   pb.KeystoreServiceClient
	conn *grpc.ClientConn
}

func Connect(ctx context.Context, address, cert string, insecure bool) (*KvStoreClient, error) {
	var err error
	var creds credentials.TransportCredentials
	var opts = []grpc.DialOption{
		grpc.WithBlock(),
	}

	// connect to grpc api
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	// use provided CA cert or system certs
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

	// dial grpc server
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}

	return &KvStoreClient{pb.NewKeystoreServiceClient(conn), conn}, nil
}

func (client *KvStoreClient) Get(ctx context.Context, bucket, key string) ([]byte, error) {
	res, err := client.kv.Get(ctx, &pb.GetRequest{Bucket: bucket, Key: key}, grpc.EmptyCallOption{})
	if err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (client *KvStoreClient) Put(ctx context.Context, bucket, key string, value []byte) error {
	_, err := client.kv.Put(ctx, &pb.PutRequest{Bucket: bucket, Key: key, Value: value}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *KvStoreClient) Delete(ctx context.Context, bucket, key string) error {
	_, err := client.kv.Delete(ctx, &pb.DeleteRequest{Bucket: bucket, Key: key}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *KvStoreClient) CreateBucket(ctx context.Context, bucket string) error {
	_, err := client.kv.CreateBucket(ctx, &pb.BucketRequest{Bucket: bucket}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *KvStoreClient) DeleteBucket(ctx context.Context, bucket string) error {
	_, err := client.kv.DeleteBucket(ctx, &pb.BucketRequest{Bucket: bucket}, grpc.EmptyCallOption{})
	if err != nil {
		return err
	}

	return nil
}

func (client *KvStoreClient) Close() error {
	return client.conn.Close()
}
