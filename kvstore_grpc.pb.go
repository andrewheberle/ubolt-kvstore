// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kvstore

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KeystoreServiceClient is the client API for KeystoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeystoreServiceClient interface {
	ListKeys(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*ListKeysResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteBucket(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateBucket(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*Empty, error)
	ListBuckets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListBucketsResponse, error)
}

type keystoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeystoreServiceClient(cc grpc.ClientConnInterface) KeystoreServiceClient {
	return &keystoreServiceClient{cc}
}

func (c *keystoreServiceClient) ListKeys(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*ListKeysResponse, error) {
	out := new(ListKeysResponse)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/ListKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) DeleteBucket(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) CreateBucket(ctx context.Context, in *BucketRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keystoreServiceClient) ListBuckets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, "/kvstore.KeystoreService/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeystoreServiceServer is the server API for KeystoreService service.
// All implementations must embed UnimplementedKeystoreServiceServer
// for forward compatibility
type KeystoreServiceServer interface {
	ListKeys(context.Context, *BucketRequest) (*ListKeysResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*Empty, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
	DeleteBucket(context.Context, *BucketRequest) (*Empty, error)
	CreateBucket(context.Context, *BucketRequest) (*Empty, error)
	ListBuckets(context.Context, *Empty) (*ListBucketsResponse, error)
	mustEmbedUnimplementedKeystoreServiceServer()
}

// UnimplementedKeystoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeystoreServiceServer struct {
}

func (UnimplementedKeystoreServiceServer) ListKeys(context.Context, *BucketRequest) (*ListKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeys not implemented")
}
func (UnimplementedKeystoreServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeystoreServiceServer) Put(context.Context, *PutRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKeystoreServiceServer) Delete(context.Context, *DeleteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKeystoreServiceServer) DeleteBucket(context.Context, *BucketRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedKeystoreServiceServer) CreateBucket(context.Context, *BucketRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedKeystoreServiceServer) ListBuckets(context.Context, *Empty) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedKeystoreServiceServer) mustEmbedUnimplementedKeystoreServiceServer() {}

// UnsafeKeystoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeystoreServiceServer will
// result in compilation errors.
type UnsafeKeystoreServiceServer interface {
	mustEmbedUnimplementedKeystoreServiceServer()
}

func RegisterKeystoreServiceServer(s grpc.ServiceRegistrar, srv KeystoreServiceServer) {
	s.RegisterService(&KeystoreService_ServiceDesc, srv)
}

func _KeystoreService_ListKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).ListKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/ListKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).ListKeys(ctx, req.(*BucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).DeleteBucket(ctx, req.(*BucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).CreateBucket(ctx, req.(*BucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeystoreService_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServiceServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvstore.KeystoreService/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServiceServer).ListBuckets(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// KeystoreService_ServiceDesc is the grpc.ServiceDesc for KeystoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeystoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvstore.KeystoreService",
	HandlerType: (*KeystoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKeys",
			Handler:    _KeystoreService_ListKeys_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KeystoreService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KeystoreService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KeystoreService_Delete_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _KeystoreService_DeleteBucket_Handler,
		},
		{
			MethodName: "CreateBucket",
			Handler:    _KeystoreService_CreateBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _KeystoreService_ListBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvstore.proto",
}
