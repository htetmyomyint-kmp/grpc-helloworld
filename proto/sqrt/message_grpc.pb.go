// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sqrt

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

// SqrtClient is the client API for Sqrt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqrtClient interface {
	GetSquareRoot(ctx context.Context, in *MRequest, opts ...grpc.CallOption) (*MResponse, error)
}

type sqrtClient struct {
	cc grpc.ClientConnInterface
}

func NewSqrtClient(cc grpc.ClientConnInterface) SqrtClient {
	return &sqrtClient{cc}
}

func (c *sqrtClient) GetSquareRoot(ctx context.Context, in *MRequest, opts ...grpc.CallOption) (*MResponse, error) {
	out := new(MResponse)
	err := c.cc.Invoke(ctx, "/Sqrt/GetSquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqrtServer is the server API for Sqrt service.
// All implementations should embed UnimplementedSqrtServer
// for forward compatibility
type SqrtServer interface {
	GetSquareRoot(context.Context, *MRequest) (*MResponse, error)
}

// UnimplementedSqrtServer should be embedded to have forward compatible implementations.
type UnimplementedSqrtServer struct {
}

func (UnimplementedSqrtServer) GetSquareRoot(context.Context, *MRequest) (*MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSquareRoot not implemented")
}

// UnsafeSqrtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqrtServer will
// result in compilation errors.
type UnsafeSqrtServer interface {
	mustEmbedUnimplementedSqrtServer()
}

func RegisterSqrtServer(s grpc.ServiceRegistrar, srv SqrtServer) {
	s.RegisterService(&Sqrt_ServiceDesc, srv)
}

func _Sqrt_GetSquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqrtServer).GetSquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Sqrt/GetSquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqrtServer).GetSquareRoot(ctx, req.(*MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sqrt_ServiceDesc is the grpc.ServiceDesc for Sqrt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sqrt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Sqrt",
	HandlerType: (*SqrtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSquareRoot",
			Handler:    _Sqrt_GetSquareRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}