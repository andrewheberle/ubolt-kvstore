package client

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/andrewheberle/ubolt-kvstore/grpc"
	"gitlab.com/andrewheberle/ubolt-kvstore/server"
)

func testserver(db, addr string) error {
	srv, err := server.NewServer(db)
	if err != nil {
		return err
	}
	defer func() {
		// close db
		if err := srv.Close(); err != nil {
			panic(err)
		}

		// remove db after we are done
		if err := os.Remove(db); err != nil {
			panic(err)
		}
	}()

	server, err := grpc.NewgRPCServer(srv)
	if err != nil {
		return err
	}
	defer server.GracefulStop()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			panic(err)
		}
	}()

	if err := server.Serve(l); err != nil {
		return err
	}

	return nil
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		cert     string
		insecure bool
		wantErr  bool
	}{
		{"bad connect", "127.0.0.1:8080", "", true, true},
		{"good connect", "127.0.0.1:8081", "", true, false},
	}

	// start grpc server
	go func() {
		if err := testserver("test.db", "127.0.0.1:8081"); err != nil {
			panic(err)
		}
	}()

	for _, tt := range tests {
		_, err := Connect(tt.address, tt.cert, tt.insecure)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestGetPutDelete(t *testing.T) {
	// start grpc server
	go func() {
		if err := testserver("test.db", "127.0.0.1:8081"); err != nil {
			panic(err)
		}
	}()

	// connect
	client, err := Connect("127.0.0.1:8081", "", true)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}
	}()

	// get with empty db
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_, err := client.Get(ctx, "missing", "missing")
		assert.NotNil(t, err)
	}()

	// create bucket, put then get
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.CreateBucket(ctx, "testbucket")
		assert.Nil(t, err)
	}()
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.Put(ctx, "testbucket", "testkey", []byte("testdata"))
		assert.Nil(t, err)
	}()
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		data, err := client.Get(ctx, "testbucket", "testkey")
		assert.Nil(t, err)
		assert.Equal(t, []byte("testdata"), data)
	}()

	// test deletekey
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.Delete(ctx, "testbucket", "testkey")
		assert.Nil(t, err)
	}()
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_, err := client.Get(ctx, "testbucket", "testkey")
		assert.NotNil(t, err)
	}()
	// test deletebucket
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.DeleteBucket(ctx, "testbucket")
		assert.Nil(t, err)
	}()
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.Put(ctx, "testbucket", "testkey", []byte("testdata"))
		assert.NotNil(t, err)
	}()
}
