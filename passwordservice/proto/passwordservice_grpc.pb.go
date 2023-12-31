// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: passwordservice.proto

package passwordservice

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
	PasswordGeneratorService_Generate_FullMethodName = "/passwordservice.PasswordGeneratorService/Generate"
)

// PasswordGeneratorServiceClient is the client API for PasswordGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PasswordGeneratorServiceClient interface {
	Generate(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordResponse, error)
}

type passwordGeneratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPasswordGeneratorServiceClient(cc grpc.ClientConnInterface) PasswordGeneratorServiceClient {
	return &passwordGeneratorServiceClient{cc}
}

func (c *passwordGeneratorServiceClient) Generate(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordResponse, error) {
	out := new(PasswordResponse)
	err := c.cc.Invoke(ctx, PasswordGeneratorService_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordGeneratorServiceServer is the server API for PasswordGeneratorService service.
// All implementations must embed UnimplementedPasswordGeneratorServiceServer
// for forward compatibility
type PasswordGeneratorServiceServer interface {
	Generate(context.Context, *PasswordRequest) (*PasswordResponse, error)
	mustEmbedUnimplementedPasswordGeneratorServiceServer()
}

// UnimplementedPasswordGeneratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPasswordGeneratorServiceServer struct {
}

func (UnimplementedPasswordGeneratorServiceServer) Generate(context.Context, *PasswordRequest) (*PasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedPasswordGeneratorServiceServer) mustEmbedUnimplementedPasswordGeneratorServiceServer() {
}

// UnsafePasswordGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasswordGeneratorServiceServer will
// result in compilation errors.
type UnsafePasswordGeneratorServiceServer interface {
	mustEmbedUnimplementedPasswordGeneratorServiceServer()
}

func RegisterPasswordGeneratorServiceServer(s grpc.ServiceRegistrar, srv PasswordGeneratorServiceServer) {
	s.RegisterService(&PasswordGeneratorService_ServiceDesc, srv)
}

func _PasswordGeneratorService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordGeneratorServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PasswordGeneratorService_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordGeneratorServiceServer).Generate(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PasswordGeneratorService_ServiceDesc is the grpc.ServiceDesc for PasswordGeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PasswordGeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passwordservice.PasswordGeneratorService",
	HandlerType: (*PasswordGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _PasswordGeneratorService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passwordservice.proto",
}
