// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: group.proto

package schedule_service

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
	GroupService_Create_FullMethodName  = "/schedule_service.GroupService/Create"
	GroupService_GetByID_FullMethodName = "/schedule_service.GroupService/GetByID"
	GroupService_GetList_FullMethodName = "/schedule_service.GroupService/GetList"
	GroupService_Update_FullMethodName  = "/schedule_service.GroupService/Update"
	GroupService_Delete_FullMethodName  = "/schedule_service.GroupService/Delete"
)

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	Create(ctx context.Context, in *CreateGroup, opts ...grpc.CallOption) (*GroupPrimaryKey, error)
	GetByID(ctx context.Context, in *GroupPrimaryKey, opts ...grpc.CallOption) (*Group, error)
	GetList(ctx context.Context, in *GetListGroupRequest, opts ...grpc.CallOption) (*GetListGroupResponse, error)
	Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*GRMessage, error)
	Delete(ctx context.Context, in *GroupPrimaryKey, opts ...grpc.CallOption) (*GRMessage, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) Create(ctx context.Context, in *CreateGroup, opts ...grpc.CallOption) (*GroupPrimaryKey, error) {
	out := new(GroupPrimaryKey)
	err := c.cc.Invoke(ctx, GroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetByID(ctx context.Context, in *GroupPrimaryKey, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, GroupService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetList(ctx context.Context, in *GetListGroupRequest, opts ...grpc.CallOption) (*GetListGroupResponse, error) {
	out := new(GetListGroupResponse)
	err := c.cc.Invoke(ctx, GroupService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*GRMessage, error) {
	out := new(GRMessage)
	err := c.cc.Invoke(ctx, GroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) Delete(ctx context.Context, in *GroupPrimaryKey, opts ...grpc.CallOption) (*GRMessage, error) {
	out := new(GRMessage)
	err := c.cc.Invoke(ctx, GroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	Create(context.Context, *CreateGroup) (*GroupPrimaryKey, error)
	GetByID(context.Context, *GroupPrimaryKey) (*Group, error)
	GetList(context.Context, *GetListGroupRequest) (*GetListGroupResponse, error)
	Update(context.Context, *UpdateGroupRequest) (*GRMessage, error)
	Delete(context.Context, *GroupPrimaryKey) (*GRMessage, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) Create(context.Context, *CreateGroup) (*GroupPrimaryKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupServiceServer) GetByID(context.Context, *GroupPrimaryKey) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedGroupServiceServer) GetList(context.Context, *GetListGroupRequest) (*GetListGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedGroupServiceServer) Update(context.Context, *UpdateGroupRequest) (*GRMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupServiceServer) Delete(context.Context, *GroupPrimaryKey) (*GRMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Create(ctx, req.(*CreateGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetByID(ctx, req.(*GroupPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetList(ctx, req.(*GetListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Update(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).Delete(ctx, req.(*GroupPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_service.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GroupService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _GroupService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _GroupService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
