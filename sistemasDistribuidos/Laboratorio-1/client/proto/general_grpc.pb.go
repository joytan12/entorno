// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: general.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PedidoService_GetPedidoStatus_FullMethodName = "/grupo11.PedidoService/GetPedidoStatus"
)

// PedidoServiceClient is the client API for PedidoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PedidoServiceClient interface {
	GetPedidoStatus(ctx context.Context, in *GenerarOrden, opts ...grpc.CallOption) (*RecibirNumeroSeguimiento, error)
}

type pedidoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPedidoServiceClient(cc grpc.ClientConnInterface) PedidoServiceClient {
	return &pedidoServiceClient{cc}
}

func (c *pedidoServiceClient) GetPedidoStatus(ctx context.Context, in *GenerarOrden, opts ...grpc.CallOption) (*RecibirNumeroSeguimiento, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecibirNumeroSeguimiento)
	err := c.cc.Invoke(ctx, PedidoService_GetPedidoStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PedidoServiceServer is the server API for PedidoService service.
// All implementations must embed UnimplementedPedidoServiceServer
// for forward compatibility.
type PedidoServiceServer interface {
	GetPedidoStatus(context.Context, *GenerarOrden) (*RecibirNumeroSeguimiento, error)
	mustEmbedUnimplementedPedidoServiceServer()
}

// UnimplementedPedidoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPedidoServiceServer struct{}

func (UnimplementedPedidoServiceServer) GetPedidoStatus(context.Context, *GenerarOrden) (*RecibirNumeroSeguimiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPedidoStatus not implemented")
}
func (UnimplementedPedidoServiceServer) mustEmbedUnimplementedPedidoServiceServer() {}
func (UnimplementedPedidoServiceServer) testEmbeddedByValue()                       {}

// UnsafePedidoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PedidoServiceServer will
// result in compilation errors.
type UnsafePedidoServiceServer interface {
	mustEmbedUnimplementedPedidoServiceServer()
}

func RegisterPedidoServiceServer(s grpc.ServiceRegistrar, srv PedidoServiceServer) {
	// If the following call pancis, it indicates UnimplementedPedidoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PedidoService_ServiceDesc, srv)
}

func _PedidoService_GetPedidoStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerarOrden)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PedidoServiceServer).GetPedidoStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PedidoService_GetPedidoStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PedidoServiceServer).GetPedidoStatus(ctx, req.(*GenerarOrden))
	}
	return interceptor(ctx, in, info, handler)
}

// PedidoService_ServiceDesc is the grpc.ServiceDesc for PedidoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PedidoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grupo11.PedidoService",
	HandlerType: (*PedidoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPedidoStatus",
			Handler:    _PedidoService_GetPedidoStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "general.proto",
}

const (
	CaravanaService_GetCaravanaInfo_FullMethodName = "/grupo11.CaravanaService/GetCaravanaInfo"
)

// CaravanaServiceClient is the client API for CaravanaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaravanaServiceClient interface {
	GetCaravanaInfo(ctx context.Context, in *CaravanaRequest, opts ...grpc.CallOption) (*CaravanaResponse, error)
}

type caravanaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaravanaServiceClient(cc grpc.ClientConnInterface) CaravanaServiceClient {
	return &caravanaServiceClient{cc}
}

func (c *caravanaServiceClient) GetCaravanaInfo(ctx context.Context, in *CaravanaRequest, opts ...grpc.CallOption) (*CaravanaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaravanaResponse)
	err := c.cc.Invoke(ctx, CaravanaService_GetCaravanaInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaravanaServiceServer is the server API for CaravanaService service.
// All implementations must embed UnimplementedCaravanaServiceServer
// for forward compatibility.
type CaravanaServiceServer interface {
	GetCaravanaInfo(context.Context, *CaravanaRequest) (*CaravanaResponse, error)
	mustEmbedUnimplementedCaravanaServiceServer()
}

// UnimplementedCaravanaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaravanaServiceServer struct{}

func (UnimplementedCaravanaServiceServer) GetCaravanaInfo(context.Context, *CaravanaRequest) (*CaravanaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaravanaInfo not implemented")
}
func (UnimplementedCaravanaServiceServer) mustEmbedUnimplementedCaravanaServiceServer() {}
func (UnimplementedCaravanaServiceServer) testEmbeddedByValue()                         {}

// UnsafeCaravanaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaravanaServiceServer will
// result in compilation errors.
type UnsafeCaravanaServiceServer interface {
	mustEmbedUnimplementedCaravanaServiceServer()
}

func RegisterCaravanaServiceServer(s grpc.ServiceRegistrar, srv CaravanaServiceServer) {
	// If the following call pancis, it indicates UnimplementedCaravanaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CaravanaService_ServiceDesc, srv)
}

func _CaravanaService_GetCaravanaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaravanaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaravanaServiceServer).GetCaravanaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaravanaService_GetCaravanaInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaravanaServiceServer).GetCaravanaInfo(ctx, req.(*CaravanaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaravanaService_ServiceDesc is the grpc.ServiceDesc for CaravanaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaravanaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grupo11.CaravanaService",
	HandlerType: (*CaravanaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaravanaInfo",
			Handler:    _CaravanaService_GetCaravanaInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "general.proto",
}

const (
	FinancieroService_GetTransactionStatus_FullMethodName = "/grupo11.FinancieroService/GetTransactionStatus"
)

// FinancieroServiceClient is the client API for FinancieroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinancieroServiceClient interface {
	GetTransactionStatus(ctx context.Context, in *FinancieroRequest, opts ...grpc.CallOption) (*FinancieroResponse, error)
}

type financieroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancieroServiceClient(cc grpc.ClientConnInterface) FinancieroServiceClient {
	return &financieroServiceClient{cc}
}

func (c *financieroServiceClient) GetTransactionStatus(ctx context.Context, in *FinancieroRequest, opts ...grpc.CallOption) (*FinancieroResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinancieroResponse)
	err := c.cc.Invoke(ctx, FinancieroService_GetTransactionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancieroServiceServer is the server API for FinancieroService service.
// All implementations must embed UnimplementedFinancieroServiceServer
// for forward compatibility.
type FinancieroServiceServer interface {
	GetTransactionStatus(context.Context, *FinancieroRequest) (*FinancieroResponse, error)
	mustEmbedUnimplementedFinancieroServiceServer()
}

// UnimplementedFinancieroServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFinancieroServiceServer struct{}

func (UnimplementedFinancieroServiceServer) GetTransactionStatus(context.Context, *FinancieroRequest) (*FinancieroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionStatus not implemented")
}
func (UnimplementedFinancieroServiceServer) mustEmbedUnimplementedFinancieroServiceServer() {}
func (UnimplementedFinancieroServiceServer) testEmbeddedByValue()                           {}

// UnsafeFinancieroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancieroServiceServer will
// result in compilation errors.
type UnsafeFinancieroServiceServer interface {
	mustEmbedUnimplementedFinancieroServiceServer()
}

func RegisterFinancieroServiceServer(s grpc.ServiceRegistrar, srv FinancieroServiceServer) {
	// If the following call pancis, it indicates UnimplementedFinancieroServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FinancieroService_ServiceDesc, srv)
}

func _FinancieroService_GetTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinancieroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancieroServiceServer).GetTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FinancieroService_GetTransactionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancieroServiceServer).GetTransactionStatus(ctx, req.(*FinancieroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinancieroService_ServiceDesc is the grpc.ServiceDesc for FinancieroService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinancieroService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grupo11.FinancieroService",
	HandlerType: (*FinancieroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionStatus",
			Handler:    _FinancieroService_GetTransactionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "general.proto",
}
