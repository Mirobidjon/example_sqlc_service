// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: position_service.proto

package position_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_position_service_proto protoreflect.FileDescriptor

var file_position_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xc6, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_position_service_proto_goTypes = []interface{}{
	(*CreatePositionRequest)(nil),  // 0: genproto.CreatePositionRequest
	(*PositionId)(nil),             // 1: genproto.PositionId
	(*Position)(nil),               // 2: genproto.Position
	(*GetAllPositionRequest)(nil),  // 3: genproto.GetAllPositionRequest
	(*GetAllPositionResponse)(nil), // 4: genproto.GetAllPositionResponse
	(*emptypb.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_position_service_proto_depIdxs = []int32{
	0, // 0: genproto.PositionService.Create:input_type -> genproto.CreatePositionRequest
	1, // 1: genproto.PositionService.Get:input_type -> genproto.PositionId
	2, // 2: genproto.PositionService.Update:input_type -> genproto.Position
	3, // 3: genproto.PositionService.GetAll:input_type -> genproto.GetAllPositionRequest
	1, // 4: genproto.PositionService.Delete:input_type -> genproto.PositionId
	1, // 5: genproto.PositionService.Create:output_type -> genproto.PositionId
	2, // 6: genproto.PositionService.Get:output_type -> genproto.Position
	1, // 7: genproto.PositionService.Update:output_type -> genproto.PositionId
	4, // 8: genproto.PositionService.GetAll:output_type -> genproto.GetAllPositionResponse
	5, // 9: genproto.PositionService.Delete:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_position_service_proto_init() }
func file_position_service_proto_init() {
	if File_position_service_proto != nil {
		return
	}
	file_position_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_position_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_position_service_proto_goTypes,
		DependencyIndexes: file_position_service_proto_depIdxs,
	}.Build()
	File_position_service_proto = out.File
	file_position_service_proto_rawDesc = nil
	file_position_service_proto_goTypes = nil
	file_position_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PositionServiceClient interface {
	Create(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*PositionId, error)
	Get(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*Position, error)
	Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*PositionId, error)
	GetAll(ctx context.Context, in *GetAllPositionRequest, opts ...grpc.CallOption) (*GetAllPositionResponse, error)
	Delete(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) Create(ctx context.Context, in *CreatePositionRequest, opts ...grpc.CallOption) (*PositionId, error) {
	out := new(PositionId)
	err := c.cc.Invoke(ctx, "/genproto.PositionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Get(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, "/genproto.PositionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*PositionId, error) {
	out := new(PositionId)
	err := c.cc.Invoke(ctx, "/genproto.PositionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetAll(ctx context.Context, in *GetAllPositionRequest, opts ...grpc.CallOption) (*GetAllPositionResponse, error) {
	out := new(GetAllPositionResponse)
	err := c.cc.Invoke(ctx, "/genproto.PositionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Delete(ctx context.Context, in *PositionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.PositionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
type PositionServiceServer interface {
	Create(context.Context, *CreatePositionRequest) (*PositionId, error)
	Get(context.Context, *PositionId) (*Position, error)
	Update(context.Context, *Position) (*PositionId, error)
	GetAll(context.Context, *GetAllPositionRequest) (*GetAllPositionResponse, error)
	Delete(context.Context, *PositionId) (*emptypb.Empty, error)
}

// UnimplementedPositionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (*UnimplementedPositionServiceServer) Create(context.Context, *CreatePositionRequest) (*PositionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPositionServiceServer) Get(context.Context, *PositionId) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedPositionServiceServer) Update(context.Context, *Position) (*PositionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPositionServiceServer) GetAll(context.Context, *GetAllPositionRequest) (*GetAllPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedPositionServiceServer) Delete(context.Context, *PositionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterPositionServiceServer(s *grpc.Server, srv PositionServiceServer) {
	s.RegisterService(&_PositionService_serviceDesc, srv)
}

func _PositionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PositionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Create(ctx, req.(*CreatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PositionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Get(ctx, req.(*PositionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PositionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Update(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PositionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetAll(ctx, req.(*GetAllPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PositionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Delete(ctx, req.(*PositionId))
	}
	return interceptor(ctx, in, info, handler)
}

var _PositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PositionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PositionService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PositionService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PositionService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PositionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position_service.proto",
}
