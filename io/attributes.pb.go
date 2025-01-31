//*
// Attributes
//
// Attributes used for performance indicator.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: io/common/attributes.proto

package io

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceAttributes int32

const (
	DeviceAttributes_NoAttributes   DeviceAttributes = 0
	DeviceAttributes_Ios            DeviceAttributes = 1
	DeviceAttributes_Android        DeviceAttributes = 2
	DeviceAttributes_SupportWallet  DeviceAttributes = 4
	DeviceAttributes_WalletScanner  DeviceAttributes = 8
	DeviceAttributes_WalletDaemon   DeviceAttributes = 16
	DeviceAttributes_WalletPasses   DeviceAttributes = 32
	DeviceAttributes_Windows        DeviceAttributes = 64
	DeviceAttributes_OSX            DeviceAttributes = 128
	DeviceAttributes_Linux          DeviceAttributes = 256
	DeviceAttributes_Mobile         DeviceAttributes = 512
	DeviceAttributes_Desktop        DeviceAttributes = 1024
	DeviceAttributes_Tablet         DeviceAttributes = 2048
	DeviceAttributes_UnsupportedIos DeviceAttributes = 4096
)

// Enum value maps for DeviceAttributes.
var (
	DeviceAttributes_name = map[int32]string{
		0:    "NoAttributes",
		1:    "Ios",
		2:    "Android",
		4:    "SupportWallet",
		8:    "WalletScanner",
		16:   "WalletDaemon",
		32:   "WalletPasses",
		64:   "Windows",
		128:  "OSX",
		256:  "Linux",
		512:  "Mobile",
		1024: "Desktop",
		2048: "Tablet",
		4096: "UnsupportedIos",
	}
	DeviceAttributes_value = map[string]int32{
		"NoAttributes":   0,
		"Ios":            1,
		"Android":        2,
		"SupportWallet":  4,
		"WalletScanner":  8,
		"WalletDaemon":   16,
		"WalletPasses":   32,
		"Windows":        64,
		"OSX":            128,
		"Linux":          256,
		"Mobile":         512,
		"Desktop":        1024,
		"Tablet":         2048,
		"UnsupportedIos": 4096,
	}
)

func (x DeviceAttributes) Enum() *DeviceAttributes {
	p := new(DeviceAttributes)
	*p = x
	return p
}

func (x DeviceAttributes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceAttributes) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_attributes_proto_enumTypes[0].Descriptor()
}

func (DeviceAttributes) Type() protoreflect.EnumType {
	return &file_io_common_attributes_proto_enumTypes[0]
}

func (x DeviceAttributes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceAttributes.Descriptor instead.
func (DeviceAttributes) EnumDescriptor() ([]byte, []int) {
	return file_io_common_attributes_proto_rawDescGZIP(), []int{0}
}

type Channel int32

const (
	Channel_UnknownChannel Channel = 0
	Channel_Web            Channel = 1
	Channel_API            Channel = 2
	Channel_App            Channel = 3
)

// Enum value maps for Channel.
var (
	Channel_name = map[int32]string{
		0: "UnknownChannel",
		1: "Web",
		2: "API",
		3: "App",
	}
	Channel_value = map[string]int32{
		"UnknownChannel": 0,
		"Web":            1,
		"API":            2,
		"App":            3,
	}
)

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_attributes_proto_enumTypes[1].Descriptor()
}

func (Channel) Type() protoreflect.EnumType {
	return &file_io_common_attributes_proto_enumTypes[1]
}

func (x Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel.Descriptor instead.
func (Channel) EnumDescriptor() ([]byte, []int) {
	return file_io_common_attributes_proto_rawDescGZIP(), []int{1}
}

type Authentication int32

const (
	Authentication_Unauthenticated Authentication = 0
	Authentication_WebToken        Authentication = 1
	Authentication_BearerToken     Authentication = 2
	Authentication_Certificate     Authentication = 3
)

// Enum value maps for Authentication.
var (
	Authentication_name = map[int32]string{
		0: "Unauthenticated",
		1: "WebToken",
		2: "BearerToken",
		3: "Certificate",
	}
	Authentication_value = map[string]int32{
		"Unauthenticated": 0,
		"WebToken":        1,
		"BearerToken":     2,
		"Certificate":     3,
	}
)

func (x Authentication) Enum() *Authentication {
	p := new(Authentication)
	*p = x
	return p
}

func (x Authentication) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Authentication) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_attributes_proto_enumTypes[2].Descriptor()
}

func (Authentication) Type() protoreflect.EnumType {
	return &file_io_common_attributes_proto_enumTypes[2]
}

func (x Authentication) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Authentication.Descriptor instead.
func (Authentication) EnumDescriptor() ([]byte, []int) {
	return file_io_common_attributes_proto_rawDescGZIP(), []int{2}
}

type UserType int32

const (
	UserType_UnknownUserType UserType = 0
	UserType_User            UserType = 1
	UserType_TeamMember      UserType = 2
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "UnknownUserType",
		1: "User",
		2: "TeamMember",
	}
	UserType_value = map[string]int32{
		"UnknownUserType": 0,
		"User":            1,
		"TeamMember":      2,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_attributes_proto_enumTypes[3].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_io_common_attributes_proto_enumTypes[3]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_io_common_attributes_proto_rawDescGZIP(), []int{3}
}

var File_io_common_attributes_proto protoreflect.FileDescriptor

var file_io_common_attributes_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f,
	0x2a, 0xe4, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x6f, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x6f, 0x73, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x10, 0x20, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x73, 0x10, 0x40, 0x12, 0x08, 0x0a, 0x03, 0x4f, 0x53, 0x58, 0x10, 0x80, 0x01, 0x12, 0x0a,
	0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x75, 0x78, 0x10, 0x80, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x10, 0x80, 0x04, 0x12, 0x0c, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x6b, 0x74,
	0x6f, 0x70, 0x10, 0x80, 0x08, 0x12, 0x0b, 0x0a, 0x06, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x10,
	0x80, 0x10, 0x12, 0x13, 0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x49, 0x6f, 0x73, 0x10, 0x80, 0x20, 0x2a, 0x38, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x65, 0x62, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x10,
	0x03, 0x2a, 0x55, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x10, 0x03, 0x2a, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x10, 0x02, 0x42, 0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70,
	0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c,
	0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_attributes_proto_rawDescOnce sync.Once
	file_io_common_attributes_proto_rawDescData = file_io_common_attributes_proto_rawDesc
)

func file_io_common_attributes_proto_rawDescGZIP() []byte {
	file_io_common_attributes_proto_rawDescOnce.Do(func() {
		file_io_common_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_attributes_proto_rawDescData)
	})
	return file_io_common_attributes_proto_rawDescData
}

var file_io_common_attributes_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_io_common_attributes_proto_goTypes = []any{
	(DeviceAttributes)(0), // 0: io.DeviceAttributes
	(Channel)(0),          // 1: io.Channel
	(Authentication)(0),   // 2: io.Authentication
	(UserType)(0),         // 3: io.UserType
}
var file_io_common_attributes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_common_attributes_proto_init() }
func file_io_common_attributes_proto_init() {
	if File_io_common_attributes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_common_attributes_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_attributes_proto_goTypes,
		DependencyIndexes: file_io_common_attributes_proto_depIdxs,
		EnumInfos:         file_io_common_attributes_proto_enumTypes,
	}.Build()
	File_io_common_attributes_proto = out.File
	file_io_common_attributes_proto_rawDesc = nil
	file_io_common_attributes_proto_goTypes = nil
	file_io_common_attributes_proto_depIdxs = nil
}
