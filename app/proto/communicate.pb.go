// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/communicate.proto

package proto

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

type ApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment []byte `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Service    []byte `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Ingress    []byte `protobuf:"bytes,3,opt,name=ingress,proto3" json:"ingress,omitempty"`
	Secret     string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ApplyRequest) Reset() {
	*x = ApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communicate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRequest) ProtoMessage() {}

func (x *ApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communicate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRequest.ProtoReflect.Descriptor instead.
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_communicate_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyRequest) GetDeployment() []byte {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *ApplyRequest) GetService() []byte {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ApplyRequest) GetIngress() []byte {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *ApplyRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communicate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communicate_proto_msgTypes[1]
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
	return file_proto_communicate_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_communicate_proto protoreflect.FileDescriptor

var file_proto_communicate_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0c, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x38, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0d, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x61,
	0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_communicate_proto_rawDescOnce sync.Once
	file_proto_communicate_proto_rawDescData = file_proto_communicate_proto_rawDesc
)

func file_proto_communicate_proto_rawDescGZIP() []byte {
	file_proto_communicate_proto_rawDescOnce.Do(func() {
		file_proto_communicate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_communicate_proto_rawDescData)
	})
	return file_proto_communicate_proto_rawDescData
}

var file_proto_communicate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_communicate_proto_goTypes = []interface{}{
	(*ApplyRequest)(nil),   // 0: ApplyRequest
	(*StatusResponse)(nil), // 1: StatusResponse
}
var file_proto_communicate_proto_depIdxs = []int32{
	0, // 0: Communicate.Apply:input_type -> ApplyRequest
	1, // 1: Communicate.Apply:output_type -> StatusResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_communicate_proto_init() }
func file_proto_communicate_proto_init() {
	if File_proto_communicate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_communicate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRequest); i {
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
		file_proto_communicate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_communicate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_communicate_proto_goTypes,
		DependencyIndexes: file_proto_communicate_proto_depIdxs,
		MessageInfos:      file_proto_communicate_proto_msgTypes,
	}.Build()
	File_proto_communicate_proto = out.File
	file_proto_communicate_proto_rawDesc = nil
	file_proto_communicate_proto_goTypes = nil
	file_proto_communicate_proto_depIdxs = nil
}
