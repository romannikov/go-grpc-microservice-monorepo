// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: example_service.proto

package example

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExampleRequest) Reset() {
	*x = ExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleRequest) ProtoMessage() {}

func (x *ExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleRequest.ProtoReflect.Descriptor instead.
func (*ExampleRequest) Descriptor() ([]byte, []int) {
	return file_example_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ExampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExampleResponse) Reset() {
	*x = ExampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleResponse) ProtoMessage() {}

func (x *ExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleResponse.ProtoReflect.Descriptor instead.
func (*ExampleResponse) Descriptor() ([]byte, []int) {
	return file_example_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExampleResponse) GetResult() uint64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_example_service_proto protoreflect.FileDescriptor

var file_example_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x22, 0x20, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x44, 0x0a,
	0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x45, 0x78, 0x61, 0x6d,
	0x12, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x6f, 0x6d, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x6f, 0x76, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_service_proto_rawDescOnce sync.Once
	file_example_service_proto_rawDescData = file_example_service_proto_rawDesc
)

func file_example_service_proto_rawDescGZIP() []byte {
	file_example_service_proto_rawDescOnce.Do(func() {
		file_example_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_service_proto_rawDescData)
	})
	return file_example_service_proto_rawDescData
}

var file_example_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_service_proto_goTypes = []interface{}{
	(*ExampleRequest)(nil),  // 0: example.ExampleRequest
	(*ExampleResponse)(nil), // 1: example.ExampleResponse
}
var file_example_service_proto_depIdxs = []int32{
	0, // 0: example.Example.Exam:input_type -> example.ExampleRequest
	1, // 1: example.Example.Exam:output_type -> example.ExampleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_service_proto_init() }
func file_example_service_proto_init() {
	if File_example_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleRequest); i {
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
		file_example_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleResponse); i {
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
			RawDescriptor: file_example_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_service_proto_goTypes,
		DependencyIndexes: file_example_service_proto_depIdxs,
		MessageInfos:      file_example_service_proto_msgTypes,
	}.Build()
	File_example_service_proto = out.File
	file_example_service_proto_rawDesc = nil
	file_example_service_proto_goTypes = nil
	file_example_service_proto_depIdxs = nil
}