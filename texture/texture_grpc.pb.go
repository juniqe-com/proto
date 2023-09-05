// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: texture/texture.proto

package texture

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
	TextureService_GetTexture_FullMethodName = "/texture.TextureService/GetTexture"
)

// TextureServiceClient is the client API for TextureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextureServiceClient interface {
	GetTexture(ctx context.Context, in *TextureRequest, opts ...grpc.CallOption) (*TextureResponse, error)
}

type textureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextureServiceClient(cc grpc.ClientConnInterface) TextureServiceClient {
	return &textureServiceClient{cc}
}

func (c *textureServiceClient) GetTexture(ctx context.Context, in *TextureRequest, opts ...grpc.CallOption) (*TextureResponse, error) {
	out := new(TextureResponse)
	err := c.cc.Invoke(ctx, TextureService_GetTexture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextureServiceServer is the server API for TextureService service.
// All implementations must embed UnimplementedTextureServiceServer
// for forward compatibility
type TextureServiceServer interface {
	GetTexture(context.Context, *TextureRequest) (*TextureResponse, error)
	mustEmbedUnimplementedTextureServiceServer()
}

// UnimplementedTextureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextureServiceServer struct {
}

func (UnimplementedTextureServiceServer) GetTexture(context.Context, *TextureRequest) (*TextureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTexture not implemented")
}
func (UnimplementedTextureServiceServer) mustEmbedUnimplementedTextureServiceServer() {}

// UnsafeTextureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextureServiceServer will
// result in compilation errors.
type UnsafeTextureServiceServer interface {
	mustEmbedUnimplementedTextureServiceServer()
}

func RegisterTextureServiceServer(s grpc.ServiceRegistrar, srv TextureServiceServer) {
	s.RegisterService(&TextureService_ServiceDesc, srv)
}

func _TextureService_GetTexture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextureServiceServer).GetTexture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextureService_GetTexture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextureServiceServer).GetTexture(ctx, req.(*TextureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextureService_ServiceDesc is the grpc.ServiceDesc for TextureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "texture.TextureService",
	HandlerType: (*TextureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTexture",
			Handler:    _TextureService_GetTexture_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "texture/texture.proto",
}
