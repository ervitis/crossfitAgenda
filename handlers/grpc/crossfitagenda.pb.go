//
//Crossfit Agenda
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 1.0.0
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: handlers/grpc/crossfitagenda.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Process Status
type Status_StatusEnum int32

const (
	Status_StatusEnum_WORKING  Status_StatusEnum = 0
	Status_StatusEnum_FINISHED Status_StatusEnum = 1
	Status_StatusEnum_FAILED   Status_StatusEnum = 2
)

// Enum value maps for Status_StatusEnum.
var (
	Status_StatusEnum_name = map[int32]string{
		0: "StatusEnum_WORKING",
		1: "StatusEnum_FINISHED",
		2: "StatusEnum_FAILED",
	}
	Status_StatusEnum_value = map[string]int32{
		"StatusEnum_WORKING":  0,
		"StatusEnum_FINISHED": 1,
		"StatusEnum_FAILED":   2,
	}
)

func (x Status_StatusEnum) Enum() *Status_StatusEnum {
	p := new(Status_StatusEnum)
	*p = x
	return p
}

func (x Status_StatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_StatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_handlers_grpc_crossfitagenda_proto_enumTypes[0].Descriptor()
}

func (Status_StatusEnum) Type() protoreflect.EnumType {
	return &file_handlers_grpc_crossfitagenda_proto_enumTypes[0]
}

func (x Status_StatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_StatusEnum.Descriptor instead.
func (Status_StatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_handlers_grpc_crossfitagenda_proto_rawDescGZIP(), []int{1, 0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  float32 `protobuf:"fixed32,355610639,opt,name=status,proto3" json:"status,omitempty"`
	Message string  `protobuf:"bytes,418054152,opt,name=message,proto3" json:"message,omitempty"`
	Date    string  `protobuf:"bytes,3076014,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_handlers_grpc_crossfitagenda_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetStatus() float32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64             `protobuf:"varint,3355,opt,name=id,proto3" json:"id,omitempty"`
	Date     int32             `protobuf:"varint,3076014,opt,name=date,proto3" json:"date,omitempty"`
	Detail   string            `protobuf:"bytes,261482417,opt,name=detail,proto3" json:"detail,omitempty"`
	Status   Status_StatusEnum `protobuf:"varint,355610639,opt,name=status,proto3,enum=crossfitagenda.Status_StatusEnum" json:"status,omitempty"`
	Complete bool              `protobuf:"varint,62574280,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_handlers_grpc_crossfitagenda_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Status) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Status) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Status) GetStatus() Status_StatusEnum {
	if x != nil {
		return x.Status
	}
	return Status_StatusEnum_WORKING
}

func (x *Status) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Status `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handlers_grpc_crossfitagenda_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_handlers_grpc_crossfitagenda_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResponse) GetData() *Status {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_handlers_grpc_crossfitagenda_proto protoreflect.FileDescriptor

var file_handlers_grpc_crossfitagenda_proto_rawDesc = []byte{
	0x0a, 0x22, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x66, 0x69, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x66, 0x69, 0x74, 0x61, 0x67,
	0x65, 0x6e, 0x64, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x8f, 0xe0, 0xc8, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x88, 0x80, 0xac, 0xc7, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0xae, 0xdf, 0xbb,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x9b, 0x1a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0xae, 0xdf, 0xbb, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0xb1, 0xcf, 0xd7, 0x7c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x8f, 0xe0, 0xc8, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x66, 0x69, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0xc8, 0x9d, 0xeb, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x6e, 0x75, 0x6d, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x3c, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x66, 0x69, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x99, 0x01, 0x0a, 0x0e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x66, 0x69, 0x74, 0x41,
	0x67, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x66,
	0x69, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handlers_grpc_crossfitagenda_proto_rawDescOnce sync.Once
	file_handlers_grpc_crossfitagenda_proto_rawDescData = file_handlers_grpc_crossfitagenda_proto_rawDesc
)

func file_handlers_grpc_crossfitagenda_proto_rawDescGZIP() []byte {
	file_handlers_grpc_crossfitagenda_proto_rawDescOnce.Do(func() {
		file_handlers_grpc_crossfitagenda_proto_rawDescData = protoimpl.X.CompressGZIP(file_handlers_grpc_crossfitagenda_proto_rawDescData)
	})
	return file_handlers_grpc_crossfitagenda_proto_rawDescData
}

var file_handlers_grpc_crossfitagenda_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_handlers_grpc_crossfitagenda_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_handlers_grpc_crossfitagenda_proto_goTypes = []interface{}{
	(Status_StatusEnum)(0), // 0: crossfitagenda.Status.StatusEnum
	(*Error)(nil),          // 1: crossfitagenda.Error
	(*Status)(nil),         // 2: crossfitagenda.Status
	(*StatusResponse)(nil), // 3: crossfitagenda.StatusResponse
	(*emptypb.Empty)(nil),  // 4: google.protobuf.Empty
}
var file_handlers_grpc_crossfitagenda_proto_depIdxs = []int32{
	0, // 0: crossfitagenda.Status.status:type_name -> crossfitagenda.Status.StatusEnum
	2, // 1: crossfitagenda.StatusResponse.data:type_name -> crossfitagenda.Status
	4, // 2: crossfitagenda.DefaultService.StartCrossfitAgenda:input_type -> google.protobuf.Empty
	4, // 3: crossfitagenda.DefaultService.Status:input_type -> google.protobuf.Empty
	4, // 4: crossfitagenda.DefaultService.StartCrossfitAgenda:output_type -> google.protobuf.Empty
	3, // 5: crossfitagenda.DefaultService.Status:output_type -> crossfitagenda.StatusResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_handlers_grpc_crossfitagenda_proto_init() }
func file_handlers_grpc_crossfitagenda_proto_init() {
	if File_handlers_grpc_crossfitagenda_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handlers_grpc_crossfitagenda_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_handlers_grpc_crossfitagenda_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_handlers_grpc_crossfitagenda_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
			RawDescriptor: file_handlers_grpc_crossfitagenda_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handlers_grpc_crossfitagenda_proto_goTypes,
		DependencyIndexes: file_handlers_grpc_crossfitagenda_proto_depIdxs,
		EnumInfos:         file_handlers_grpc_crossfitagenda_proto_enumTypes,
		MessageInfos:      file_handlers_grpc_crossfitagenda_proto_msgTypes,
	}.Build()
	File_handlers_grpc_crossfitagenda_proto = out.File
	file_handlers_grpc_crossfitagenda_proto_rawDesc = nil
	file_handlers_grpc_crossfitagenda_proto_goTypes = nil
	file_handlers_grpc_crossfitagenda_proto_depIdxs = nil
}
