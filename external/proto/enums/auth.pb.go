// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/enums/auth.proto

package enums

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

type DeleteType int32

const (
	DeleteType_Soft DeleteType = 0
	DeleteType_Hard DeleteType = 1
)

// Enum value maps for DeleteType.
var (
	DeleteType_name = map[int32]string{
		0: "Soft",
		1: "Hard",
	}
	DeleteType_value = map[string]int32{
		"Soft": 0,
		"Hard": 1,
	}
)

func (x DeleteType) Enum() *DeleteType {
	p := new(DeleteType)
	*p = x
	return p
}

func (x DeleteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_auth_proto_enumTypes[0].Descriptor()
}

func (DeleteType) Type() protoreflect.EnumType {
	return &file_proto_enums_auth_proto_enumTypes[0]
}

func (x DeleteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteType.Descriptor instead.
func (DeleteType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_auth_proto_rawDescGZIP(), []int{0}
}

var File_proto_enums_auth_proto protoreflect.FileDescriptor

var file_proto_enums_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a,
	0x20, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x6f, 0x66, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x61, 0x72, 0x64, 0x10,
	0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x79, 0x73, 0x6e, 0x69, 0x78, 0x2f, 0x61, 0x69, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_enums_auth_proto_rawDescOnce sync.Once
	file_proto_enums_auth_proto_rawDescData = file_proto_enums_auth_proto_rawDesc
)

func file_proto_enums_auth_proto_rawDescGZIP() []byte {
	file_proto_enums_auth_proto_rawDescOnce.Do(func() {
		file_proto_enums_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enums_auth_proto_rawDescData)
	})
	return file_proto_enums_auth_proto_rawDescData
}

var file_proto_enums_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_enums_auth_proto_goTypes = []interface{}{
	(DeleteType)(0), // 0: enums.DeleteType
}
var file_proto_enums_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_enums_auth_proto_init() }
func file_proto_enums_auth_proto_init() {
	if File_proto_enums_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_enums_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enums_auth_proto_goTypes,
		DependencyIndexes: file_proto_enums_auth_proto_depIdxs,
		EnumInfos:         file_proto_enums_auth_proto_enumTypes,
	}.Build()
	File_proto_enums_auth_proto = out.File
	file_proto_enums_auth_proto_rawDesc = nil
	file_proto_enums_auth_proto_goTypes = nil
	file_proto_enums_auth_proto_depIdxs = nil
}
