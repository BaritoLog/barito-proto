// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: producer/producer.proto

package producer

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

const (
	Producer_Produce_FullMethodName      = "/producer.Producer/Produce"
	Producer_ProduceBatch_FullMethodName = "/producer.Producer/ProduceBatch"
)

// ProducerClient is the client API for Producer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProducerClient interface {
	Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResult, error)
	ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResult, error)
}

type producerClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerClient(cc grpc.ClientConnInterface) ProducerClient {
	return &producerClient{cc}
}

func (c *producerClient) Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResult, error) {
	out := new(ProduceResult)
	err := c.cc.Invoke(ctx, Producer_Produce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerClient) ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResult, error) {
	out := new(ProduceResult)
	err := c.cc.Invoke(ctx, Producer_ProduceBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServer is the server API for Producer service.
// All implementations must embed UnimplementedProducerServer
// for forward compatibility
type ProducerServer interface {
	Produce(context.Context, *Timber) (*ProduceResult, error)
	ProduceBatch(context.Context, *TimberCollection) (*ProduceResult, error)
	mustEmbedUnimplementedProducerServer()
}

// UnimplementedProducerServer must be embedded to have forward compatible implementations.
type UnimplementedProducerServer struct {
}

func (UnimplementedProducerServer) Produce(context.Context, *Timber) (*ProduceResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedProducerServer) ProduceBatch(context.Context, *TimberCollection) (*ProduceResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProduceBatch not implemented")
}
func (UnimplementedProducerServer) mustEmbedUnimplementedProducerServer() {}

// UnsafeProducerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProducerServer will
// result in compilation errors.
type UnsafeProducerServer interface {
	mustEmbedUnimplementedProducerServer()
}

func RegisterProducerServer(s grpc.ServiceRegistrar, srv ProducerServer) {
	s.RegisterService(&Producer_ServiceDesc, srv)
}

func _Producer_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Producer_Produce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServer).Produce(ctx, req.(*Timber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Producer_ProduceBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimberCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServer).ProduceBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Producer_ProduceBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServer).ProduceBatch(ctx, req.(*TimberCollection))
	}
	return interceptor(ctx, in, info, handler)
}

// Producer_ServiceDesc is the grpc.ServiceDesc for Producer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Producer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "producer.Producer",
	HandlerType: (*ProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Producer_Produce_Handler,
		},
		{
			MethodName: "ProduceBatch",
			Handler:    _Producer_ProduceBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "producer/producer.proto",
}
