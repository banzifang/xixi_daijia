// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/mapService/mapService.proto

package mapService

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
	MapService_GetDrivingInfo_FullMethodName = "/api.mapService.MapService/GetDrivingInfo"
)

// MapServiceClient is the client API for MapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MapServiceClient interface {
	GetDrivingInfo(ctx context.Context, in *GetDrivingInfoReq, opts ...grpc.CallOption) (*GetDrivingReply, error)
}

type mapServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMapServiceClient(cc grpc.ClientConnInterface) MapServiceClient {
	return &mapServiceClient{cc}
}

func (c *mapServiceClient) GetDrivingInfo(ctx context.Context, in *GetDrivingInfoReq, opts ...grpc.CallOption) (*GetDrivingReply, error) {
	out := new(GetDrivingReply)
	err := c.cc.Invoke(ctx, MapService_GetDrivingInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapServiceServer is the server API for MapService service.
// All implementations must embed UnimplementedMapServiceServer
// for forward compatibility
type MapServiceServer interface {
	GetDrivingInfo(context.Context, *GetDrivingInfoReq) (*GetDrivingReply, error)
	mustEmbedUnimplementedMapServiceServer()
}

// UnimplementedMapServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMapServiceServer struct {
}

func (UnimplementedMapServiceServer) GetDrivingInfo(context.Context, *GetDrivingInfoReq) (*GetDrivingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivingInfo not implemented")
}
func (UnimplementedMapServiceServer) mustEmbedUnimplementedMapServiceServer() {}

// UnsafeMapServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MapServiceServer will
// result in compilation errors.
type UnsafeMapServiceServer interface {
	mustEmbedUnimplementedMapServiceServer()
}

func RegisterMapServiceServer(s grpc.ServiceRegistrar, srv MapServiceServer) {
	s.RegisterService(&MapService_ServiceDesc, srv)
}

func _MapService_GetDrivingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDrivingInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapServiceServer).GetDrivingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MapService_GetDrivingInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapServiceServer).GetDrivingInfo(ctx, req.(*GetDrivingInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MapService_ServiceDesc is the grpc.ServiceDesc for MapService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MapService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mapService.MapService",
	HandlerType: (*MapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDrivingInfo",
			Handler:    _MapService_GetDrivingInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mapService/mapService.proto",
}