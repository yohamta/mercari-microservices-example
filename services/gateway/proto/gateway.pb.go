// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.2
// source: services/gateway/proto/gateway.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_gateway_proto_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_gateway_proto_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_services_gateway_proto_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *SigninRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_gateway_proto_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_gateway_proto_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_services_gateway_proto_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *SigninResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_services_gateway_proto_gateway_proto protoreflect.FileDescriptor

var file_services_gateway_proto_gateway_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30,
	0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x33, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc1, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xae, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x12, 0x44, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31,
	0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68,
	0x6f, 0x75, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x32, 0x30,
	0x32, 0x31, 0x2d, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x2d, 0x68, 0x6f, 0x75, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_gateway_proto_gateway_proto_rawDescOnce sync.Once
	file_services_gateway_proto_gateway_proto_rawDescData = file_services_gateway_proto_gateway_proto_rawDesc
)

func file_services_gateway_proto_gateway_proto_rawDescGZIP() []byte {
	file_services_gateway_proto_gateway_proto_rawDescOnce.Do(func() {
		file_services_gateway_proto_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_gateway_proto_gateway_proto_rawDescData)
	})
	return file_services_gateway_proto_gateway_proto_rawDescData
}

var file_services_gateway_proto_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_gateway_proto_gateway_proto_goTypes = []interface{}{
	(*SigninRequest)(nil),  // 0: mercari.go_conference_2021_spring_office_hour.gateway.SigninRequest
	(*SigninResponse)(nil), // 1: mercari.go_conference_2021_spring_office_hour.gateway.SigninResponse
}
var file_services_gateway_proto_gateway_proto_depIdxs = []int32{
	0, // 0: mercari.go_conference_2021_spring_office_hour.gateway.GatewayService.Signin:input_type -> mercari.go_conference_2021_spring_office_hour.gateway.SigninRequest
	1, // 1: mercari.go_conference_2021_spring_office_hour.gateway.GatewayService.Signin:output_type -> mercari.go_conference_2021_spring_office_hour.gateway.SigninResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_gateway_proto_gateway_proto_init() }
func file_services_gateway_proto_gateway_proto_init() {
	if File_services_gateway_proto_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_gateway_proto_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninRequest); i {
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
		file_services_gateway_proto_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninResponse); i {
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
			RawDescriptor: file_services_gateway_proto_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_gateway_proto_gateway_proto_goTypes,
		DependencyIndexes: file_services_gateway_proto_gateway_proto_depIdxs,
		MessageInfos:      file_services_gateway_proto_gateway_proto_msgTypes,
	}.Build()
	File_services_gateway_proto_gateway_proto = out.File
	file_services_gateway_proto_gateway_proto_rawDesc = nil
	file_services_gateway_proto_gateway_proto_goTypes = nil
	file_services_gateway_proto_gateway_proto_depIdxs = nil
}
