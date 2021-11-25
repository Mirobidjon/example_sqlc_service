// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: attribute.proto

package position_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attribute) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateAttributeRequest) Reset() {
	*x = CreateAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttributeRequest) ProtoMessage() {}

func (x *CreateAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttributeRequest.ProtoReflect.Descriptor instead.
func (*CreateAttributeRequest) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAttributeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAttributeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AttributeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AttributeId) Reset() {
	*x = AttributeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeId) ProtoMessage() {}

func (x *AttributeId) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeId.ProtoReflect.Descriptor instead.
func (*AttributeId) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Attribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Count      int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllAttributeResponse) Reset() {
	*x = GetAllAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeResponse) ProtoMessage() {}

func (x *GetAllAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeResponse.ProtoReflect.Descriptor instead.
func (*GetAllAttributeResponse) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllAttributeResponse) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GetAllAttributeResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit_  int32  `protobuf:"varint,1,opt,name=limit_,json=limit,proto3" json:"limit_,omitempty"`
	Offset_ int32  `protobuf:"varint,2,opt,name=offset_,json=offset,proto3" json:"offset_,omitempty"`
	Search  string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllAttributeRequest) Reset() {
	*x = GetAllAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeRequest) ProtoMessage() {}

func (x *GetAllAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeRequest.ProtoReflect.Descriptor instead.
func (*GetAllAttributeRequest) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllAttributeRequest) GetLimit_() int32 {
	if x != nil {
		return x.Limit_
	}
	return 0
}

func (x *GetAllAttributeRequest) GetOffset_() int32 {
	if x != nil {
		return x.Offset_
	}
	return 0
}

func (x *GetAllAttributeRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

var File_attribute_proto protoreflect.FileDescriptor

var file_attribute_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x09, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x40, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x64, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attribute_proto_rawDescOnce sync.Once
	file_attribute_proto_rawDescData = file_attribute_proto_rawDesc
)

func file_attribute_proto_rawDescGZIP() []byte {
	file_attribute_proto_rawDescOnce.Do(func() {
		file_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribute_proto_rawDescData)
	})
	return file_attribute_proto_rawDescData
}

var file_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_attribute_proto_goTypes = []interface{}{
	(*Attribute)(nil),               // 0: genproto.Attribute
	(*CreateAttributeRequest)(nil),  // 1: genproto.CreateAttributeRequest
	(*AttributeId)(nil),             // 2: genproto.AttributeId
	(*GetAllAttributeResponse)(nil), // 3: genproto.GetAllAttributeResponse
	(*GetAllAttributeRequest)(nil),  // 4: genproto.GetAllAttributeRequest
}
var file_attribute_proto_depIdxs = []int32{
	0, // 0: genproto.GetAllAttributeResponse.attributes:type_name -> genproto.Attribute
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_attribute_proto_init() }
func file_attribute_proto_init() {
	if File_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttributeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attribute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attribute_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_attribute_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_attribute_proto_goTypes,
		DependencyIndexes: file_attribute_proto_depIdxs,
		MessageInfos:      file_attribute_proto_msgTypes,
	}.Build()
	File_attribute_proto = out.File
	file_attribute_proto_rawDesc = nil
	file_attribute_proto_goTypes = nil
	file_attribute_proto_depIdxs = nil
}
