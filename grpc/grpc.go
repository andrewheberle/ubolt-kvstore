package grpc

import (
	"os"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	pb "gitlab.com/andrewheberle/ubolt-kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewgRPCServer(addr string, srv pb.KeystoreServiceServer) (*grpc.Server, error) {
	return NewTLSgRPCServer(addr, "", "", srv)
}

func NewTLSgRPCServer(addr, cert, key string, srv pb.KeystoreServiceServer) (s *grpc.Server, err error) {
	logger := zerolog.New(os.Stderr)

	// add TLS support
	if cert != "" && key != "" {
		cred, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			return nil, err
		}

		// set up TLS enabled server
		s = grpc.NewServer(grpc.StreamInterceptor(
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
		)
	} else {
		// set up non-TLS enabled server
		// set up TLS enabled server
		s = grpc.NewServer(grpc.StreamInterceptor(
			middleware.ChainStreamServer(
				logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger)),
			),
		),
			grpc.UnaryInterceptor(
				middleware.ChainUnaryServer(
					logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger)),
				),
			),
		)
	}

	pb.RegisterKeystoreServiceServer(s, srv)

	return s, nil
}
