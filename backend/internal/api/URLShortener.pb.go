//
//Create will take in a long url and generate for it short code
//
//Get will take in a short url and return the original url

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: api/URLShortener.proto

package api

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

type OriginalURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OriginalURL) Reset() {
	*x = OriginalURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_URLShortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OriginalURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginalURL) ProtoMessage() {}

func (x *OriginalURL) ProtoReflect() protoreflect.Message {
	mi := &file_api_URLShortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalURL.ProtoReflect.Descriptor instead.
func (*OriginalURL) Descriptor() ([]byte, []int) {
	return file_api_URLShortener_proto_rawDescGZIP(), []int{0}
}

func (x *OriginalURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURLCode string `protobuf:"bytes,1,opt,name=ShortURLCode,proto3" json:"ShortURLCode,omitempty"`
}

func (x *ShortCode) Reset() {
	*x = ShortCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_URLShortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortCode) ProtoMessage() {}

func (x *ShortCode) ProtoReflect() protoreflect.Message {
	mi := &file_api_URLShortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortCode.ProtoReflect.Descriptor instead.
func (*ShortCode) Descriptor() ([]byte, []int) {
	return file_api_URLShortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortCode) GetShortURLCode() string {
	if x != nil {
		return x.ShortURLCode
	}
	return ""
}

var File_api_URLShortener_proto protoreflect.FileDescriptor

var file_api_URLShortener_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x1f, 0x0a,
	0x0b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2f,
	0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0x63, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x1a, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x52, 0x4c, 0x42, 0x1a, 0x5a, 0x18, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_URLShortener_proto_rawDescOnce sync.Once
	file_api_URLShortener_proto_rawDescData = file_api_URLShortener_proto_rawDesc
)

func file_api_URLShortener_proto_rawDescGZIP() []byte {
	file_api_URLShortener_proto_rawDescOnce.Do(func() {
		file_api_URLShortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_URLShortener_proto_rawDescData)
	})
	return file_api_URLShortener_proto_rawDescData
}

var file_api_URLShortener_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_URLShortener_proto_goTypes = []interface{}{
	(*OriginalURL)(nil), // 0: api.OriginalURL
	(*ShortCode)(nil),   // 1: api.ShortCode
}
var file_api_URLShortener_proto_depIdxs = []int32{
	0, // 0: api.URLShortener.Create:input_type -> api.OriginalURL
	1, // 1: api.URLShortener.Get:input_type -> api.ShortCode
	1, // 2: api.URLShortener.Create:output_type -> api.ShortCode
	0, // 3: api.URLShortener.Get:output_type -> api.OriginalURL
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_URLShortener_proto_init() }
func file_api_URLShortener_proto_init() {
	if File_api_URLShortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_URLShortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OriginalURL); i {
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
		file_api_URLShortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortCode); i {
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
			RawDescriptor: file_api_URLShortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_URLShortener_proto_goTypes,
		DependencyIndexes: file_api_URLShortener_proto_depIdxs,
		MessageInfos:      file_api_URLShortener_proto_msgTypes,
	}.Build()
	File_api_URLShortener_proto = out.File
	file_api_URLShortener_proto_rawDesc = nil
	file_api_URLShortener_proto_goTypes = nil
	file_api_URLShortener_proto_depIdxs = nil
}
