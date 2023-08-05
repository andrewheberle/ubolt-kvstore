package server

import (
	"context"
	"os"

	pb "github.com/andrewheberle/ubolt-kvstore"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"gitlab.com/andrewheberle/ubolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type KvStoreService struct {
	db *ubolt.DB
	pb.UnimplementedKeystoreServiceServer
}

func NewKvStoreService(file string) (*KvStoreService, error) {
	db, err := ubolt.Open(file)
	if err != nil {
		return nil, err
	}

	return &KvStoreService{db: db}, nil
}

func (srv *KvStoreService) Close() error {
	return srv.db.Close()
}

func (srv *KvStoreService) NewServer() (*grpc.Server, error) {
	return srv.NewTLSServer("", "")
}

func (srv *KvStoreService) NewTLSServer(cert, key string) (*grpc.Server, error) {
	var opts []grpc.ServerOption
	logger := zerolog.New(os.Stderr)

	// add TLS support
	if cert != "" && key != "" {
		cred, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			return nil, err
		}

		// set up TLS enabled server
		opts = []grpc.ServerOption{
			grpc.StreamInterceptor(
				middleware.ChainStreamServer(
					logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger)),
				),
			),
			grpc.UnaryInterceptor(
				middleware.ChainUnaryServer(
					logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger)),
				),
			),
			grpc.Creds(cred),
		}
	} else {
		// set up non-TLS enabled server
		opts = []grpc.ServerOption{
			grpc.StreamInterceptor(
				middleware.ChainStreamServer(
					logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger)),
				),
			),
			grpc.UnaryInterceptor(
				middleware.ChainUnaryServer(
					logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger)),
				),
			),
		}
	}

	s := grpc.NewServer(opts...)
	pb.RegisterKeystoreServiceServer(s, srv)

	return s, nil
}

func (srv *KvStoreService) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	data, err := srv.db.GetE([]byte(in.GetBucket()), []byte(in.GetKey()))
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{Value: data}, nil
}

func (srv *KvStoreService) Put(ctx context.Context, in *pb.PutRequest) (*pb.Empty, error) {
	err := srv.db.Put([]byte(in.GetBucket()), []byte(in.GetKey()), in.GetValue())
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *KvStoreService) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.Empty, error) {
	err := srv.db.Delete([]byte(in.GetBucket()), []byte(in.GetKey()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *KvStoreService) DeleteBucket(ctx context.Context, in *pb.BucketRequest) (*pb.Empty, error) {
	err := srv.db.DeleteBucket([]byte(in.GetBucket()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *KvStoreService) CreateBucket(ctx context.Context, in *pb.BucketRequest) (*pb.Empty, error) {
	err := srv.db.CreateBucket([]byte(in.GetBucket()))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (srv *KvStoreService) ListKeys(ctx context.Context, in *pb.BucketRequest) (*pb.ListKeysResponse, error) {
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

func (srv *KvStoreService) ListBuckets(ctx context.Context, in *pb.Empty) (*pb.ListBucketsResponse, error) {
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
