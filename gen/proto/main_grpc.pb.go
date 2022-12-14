// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: main.proto

package proto

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

// NutServiceClient is the client API for NutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NutServiceClient interface {
	Nudge(ctx context.Context, in *TaskOption, opts ...grpc.CallOption) (*DoneReply, error)
}

type nutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNutServiceClient(cc grpc.ClientConnInterface) NutServiceClient {
	return &nutServiceClient{cc}
}

func (c *nutServiceClient) Nudge(ctx context.Context, in *TaskOption, opts ...grpc.CallOption) (*DoneReply, error) {
	out := new(DoneReply)
	err := c.cc.Invoke(ctx, "/main.NutService/Nudge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NutServiceServer is the server API for NutService service.
// All implementations should embed UnimplementedNutServiceServer
// for forward compatibility
type NutServiceServer interface {
	Nudge(context.Context, *TaskOption) (*DoneReply, error)
}

// UnimplementedNutServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNutServiceServer struct {
}

func (UnimplementedNutServiceServer) Nudge(context.Context, *TaskOption) (*DoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nudge not implemented")
}

// UnsafeNutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NutServiceServer will
// result in compilation errors.
type UnsafeNutServiceServer interface {
	mustEmbedUnimplementedNutServiceServer()
}

func RegisterNutServiceServer(s grpc.ServiceRegistrar, srv NutServiceServer) {
	s.RegisterService(&NutService_ServiceDesc, srv)
}

func _NutService_Nudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NutServiceServer).Nudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.NutService/Nudge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NutServiceServer).Nudge(ctx, req.(*TaskOption))
	}
	return interceptor(ctx, in, info, handler)
}

// NutService_ServiceDesc is the grpc.ServiceDesc for NutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.NutService",
	HandlerType: (*NutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nudge",
			Handler:    _NutService_Nudge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
