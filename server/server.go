package server

import (
	"context"

	"gitlab.com/andrewheberle/ubolt"
	pb "gitlab.com/andrewheberle/ubolt-kvstore"
)

type Server struct {
	db *ubolt.DB
	pb.UnimplementedKeystoreServiceServer
}

func NewServer(file string) (*Server, error) {
	db, err := ubolt.Open(file)
	if err != nil {
		return nil, err
	}

	return &Server{db: db}, nil
}

func (srv *Server) Close() error {
	return srv.db.Close()
}

func (srv *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	data, err := srv.db.GetE([]byte(in.GetBucket()), []byte(in.GetKey()))
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{Value: data}, nil
}

func (srv *Server) Put(ctx context.Context, in *pb.PutRequest) (*pb.Empty, error) {
	err := srv.db.Put([]byte(in.GetBucket()), []byte(in.GetKey()), in.GetValue())
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.Empty, error) {
	err := srv.db.Delete([]byte(in.GetBucket()), []byte(in.GetKey()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *Server) DeleteBucket(ctx context.Context, in *pb.BucketRequest) (*pb.Empty, error) {
	err := srv.db.DeleteBucket([]byte(in.GetBucket()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *Server) CreateBucket(ctx context.Context, in *pb.BucketRequest) (*pb.Empty, error) {
	err := srv.db.CreateBucket([]byte(in.GetBucket()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *Server) ListKeys(ctx context.Context, in *pb.BucketRequest) (*pb.ListKeysResponse, error) {
	var keys []string

	// grab keys
	keyList, err := srv.db.GetKeysE([]byte(in.GetBucket()))
	if err != nil {
		return nil, err
	}

	// convert to []string
	for _, k := range keyList {
		keys = append(keys, string(k))
	}

	return &pb.ListKeysResponse{Keys: keys}, nil
}

func (srv *Server) ListBuckets(ctx context.Context, in *pb.Empty) (*pb.ListBucketsResponse, error) {
	var buckets []string

	// grab buckets
	bucketList, err := srv.db.GetBucketsE()
	if err != nil {
		return nil, err
	}

	// convert to []string
	for _, b := range bucketList {
		buckets = append(buckets, string(b))
	}

	return &pb.ListBucketsResponse{Buckets: buckets}, nil
}
