package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server, err := NewKvStoreService("test.db")
	assert.Nil(t, err)

	grpc, err := server.NewServer()
	assert.Nil(t, err)

	grpc.GracefulStop()
}
