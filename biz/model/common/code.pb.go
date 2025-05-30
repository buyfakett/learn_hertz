// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.29.3
// source: common/code.proto

package common

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

type Code int32

const (
	Code_Common        Code = 0
	Code_Success       Code = 200
	Code_Unauthorized  Code = 401
	Code_Err           Code = 500
	Code_DBErr         Code = 501
	Code_PasswordErr   Code = 502
	Code_AlreadyExists Code = 503
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:   "Common",
		200: "Success",
		401: "Unauthorized",
		500: "Err",
		501: "DBErr",
		502: "PasswordErr",
		503: "AlreadyExists",
	}
	Code_value = map[string]int32{
		"Common":        0,
		"Success":       200,
		"Unauthorized":  401,
		"Err":           500,
		"DBErr":         501,
		"PasswordErr":   502,
		"AlreadyExists": 503,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_common_code_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_common_code_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_common_code_proto_rawDescGZIP(), []int{0}
}

var File_common_code_proto protoreflect.FileDescriptor

var file_common_code_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x6f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0xc8, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x91, 0x03, 0x12, 0x08, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x10,
	0xf4, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x44, 0x42, 0x45, 0x72, 0x72, 0x10, 0xf5, 0x03, 0x12, 0x10,
	0x0a, 0x0b, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x10, 0xf6, 0x03,
	0x12, 0x12, 0x0a, 0x0d, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x10, 0xf7, 0x03, 0x42, 0x1d, 0x5a, 0x1b, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_code_proto_rawDescOnce sync.Once
	file_common_code_proto_rawDescData = file_common_code_proto_rawDesc
)

func file_common_code_proto_rawDescGZIP() []byte {
	file_common_code_proto_rawDescOnce.Do(func() {
		file_common_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_code_proto_rawDescData)
	})
	return file_common_code_proto_rawDescData
}

var file_common_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_code_proto_goTypes = []interface{}{
	(Code)(0), // 0: Code
}
var file_common_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_code_proto_init() }
func file_common_code_proto_init() {
	if File_common_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_code_proto_goTypes,
		DependencyIndexes: file_common_code_proto_depIdxs,
		EnumInfos:         file_common_code_proto_enumTypes,
	}.Build()
	File_common_code_proto = out.File
	file_common_code_proto_rawDesc = nil
	file_common_code_proto_goTypes = nil
	file_common_code_proto_depIdxs = nil
}
