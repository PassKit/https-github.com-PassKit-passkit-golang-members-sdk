// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: io/flights/cabin_codes/cabin_codes.proto

package cabincodes

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

// Compartment Code descriptions represent most frequently used classes. Airlines have discretion on how they use these codes.
type CompartmentCode int32

const (
	// Compartment code not provided
	CompartmentCode_NONE CompartmentCode = 0
	// First class suite (previously supersonic)
	CompartmentCode_R CompartmentCode = 1
	// First class premium
	CompartmentCode_P CompartmentCode = 2
	// First class
	CompartmentCode_F CompartmentCode = 3
	// First class discounted
	CompartmentCode_A CompartmentCode = 4
	// Business class premium
	CompartmentCode_J CompartmentCode = 5
	// Business class
	CompartmentCode_C CompartmentCode = 6
	// Business class discounted
	CompartmentCode_D CompartmentCode = 7
	// Business class discounted
	CompartmentCode_I CompartmentCode = 8
	// Business class discounted
	CompartmentCode_Z CompartmentCode = 9
	// Premium economy
	CompartmentCode_W CompartmentCode = 10
	// Premium economy discounted
	CompartmentCode_E CompartmentCode = 11
	// Economy
	CompartmentCode_S CompartmentCode = 12
	// Economy
	CompartmentCode_Y CompartmentCode = 13
	// Economy discounted
	CompartmentCode_B CompartmentCode = 14
	// Economy discounted
	CompartmentCode_H CompartmentCode = 15
	// Economy discounted
	CompartmentCode_K CompartmentCode = 16
	// Economy discounted
	CompartmentCode_L CompartmentCode = 17
	// Economy discounted
	CompartmentCode_M CompartmentCode = 18
	// Economy discounted
	CompartmentCode_N CompartmentCode = 19
	// Economy discounted
	CompartmentCode_O CompartmentCode = 20
	// Economy discounted
	CompartmentCode_Q CompartmentCode = 21
	// Premium economy discounted / Economy discounted
	CompartmentCode_T CompartmentCode = 22
	// Economy discounted
	CompartmentCode_V CompartmentCode = 23
	// Economy discounted
	CompartmentCode_X CompartmentCode = 24
	// Conditional reservation
	CompartmentCode_G CompartmentCode = 25
	// Shuttle service
	CompartmentCode_U CompartmentCode = 26
)

// Enum value maps for CompartmentCode.
var (
	CompartmentCode_name = map[int32]string{
		0:  "NONE",
		1:  "R",
		2:  "P",
		3:  "F",
		4:  "A",
		5:  "J",
		6:  "C",
		7:  "D",
		8:  "I",
		9:  "Z",
		10: "W",
		11: "E",
		12: "S",
		13: "Y",
		14: "B",
		15: "H",
		16: "K",
		17: "L",
		18: "M",
		19: "N",
		20: "O",
		21: "Q",
		22: "T",
		23: "V",
		24: "X",
		25: "G",
		26: "U",
	}
	CompartmentCode_value = map[string]int32{
		"NONE": 0,
		"R":    1,
		"P":    2,
		"F":    3,
		"A":    4,
		"J":    5,
		"C":    6,
		"D":    7,
		"I":    8,
		"Z":    9,
		"W":    10,
		"E":    11,
		"S":    12,
		"Y":    13,
		"B":    14,
		"H":    15,
		"K":    16,
		"L":    17,
		"M":    18,
		"N":    19,
		"O":    20,
		"Q":    21,
		"T":    22,
		"V":    23,
		"X":    24,
		"G":    25,
		"U":    26,
	}
)

func (x CompartmentCode) Enum() *CompartmentCode {
	p := new(CompartmentCode)
	*p = x
	return p
}

func (x CompartmentCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompartmentCode) Descriptor() protoreflect.EnumDescriptor {
	return file_io_flights_cabin_codes_cabin_codes_proto_enumTypes[0].Descriptor()
}

func (CompartmentCode) Type() protoreflect.EnumType {
	return &file_io_flights_cabin_codes_cabin_codes_proto_enumTypes[0]
}

func (x CompartmentCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompartmentCode.Descriptor instead.
func (CompartmentCode) EnumDescriptor() ([]byte, []int) {
	return file_io_flights_cabin_codes_cabin_codes_proto_rawDescGZIP(), []int{0}
}

var File_io_flights_cabin_codes_cabin_codes_proto protoreflect.FileDescriptor

var file_io_flights_cabin_codes_cabin_codes_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x63, 0x61, 0x62,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x62, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x61, 0x62, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0xd1, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x52, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01,
	0x50, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x46, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10,
	0x04, 0x12, 0x05, 0x0a, 0x01, 0x4a, 0x10, 0x05, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x06, 0x12,
	0x05, 0x0a, 0x01, 0x44, 0x10, 0x07, 0x12, 0x05, 0x0a, 0x01, 0x49, 0x10, 0x08, 0x12, 0x05, 0x0a,
	0x01, 0x5a, 0x10, 0x09, 0x12, 0x05, 0x0a, 0x01, 0x57, 0x10, 0x0a, 0x12, 0x05, 0x0a, 0x01, 0x45,
	0x10, 0x0b, 0x12, 0x05, 0x0a, 0x01, 0x53, 0x10, 0x0c, 0x12, 0x05, 0x0a, 0x01, 0x59, 0x10, 0x0d,
	0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x0e, 0x12, 0x05, 0x0a, 0x01, 0x48, 0x10, 0x0f, 0x12, 0x05,
	0x0a, 0x01, 0x4b, 0x10, 0x10, 0x12, 0x05, 0x0a, 0x01, 0x4c, 0x10, 0x11, 0x12, 0x05, 0x0a, 0x01,
	0x4d, 0x10, 0x12, 0x12, 0x05, 0x0a, 0x01, 0x4e, 0x10, 0x13, 0x12, 0x05, 0x0a, 0x01, 0x4f, 0x10,
	0x14, 0x12, 0x05, 0x0a, 0x01, 0x51, 0x10, 0x15, 0x12, 0x05, 0x0a, 0x01, 0x54, 0x10, 0x16, 0x12,
	0x05, 0x0a, 0x01, 0x56, 0x10, 0x17, 0x12, 0x05, 0x0a, 0x01, 0x58, 0x10, 0x18, 0x12, 0x05, 0x0a,
	0x01, 0x47, 0x10, 0x19, 0x12, 0x05, 0x0a, 0x01, 0x55, 0x10, 0x1a, 0x42, 0x80, 0x01, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x62, 0x69, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x5a, 0x37, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x2f, 0x63, 0x61, 0x62, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0xaa, 0x02, 0x1f, 0x50,
	0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x43, 0x61, 0x62, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_flights_cabin_codes_cabin_codes_proto_rawDescOnce sync.Once
	file_io_flights_cabin_codes_cabin_codes_proto_rawDescData = file_io_flights_cabin_codes_cabin_codes_proto_rawDesc
)

func file_io_flights_cabin_codes_cabin_codes_proto_rawDescGZIP() []byte {
	file_io_flights_cabin_codes_cabin_codes_proto_rawDescOnce.Do(func() {
		file_io_flights_cabin_codes_cabin_codes_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_flights_cabin_codes_cabin_codes_proto_rawDescData)
	})
	return file_io_flights_cabin_codes_cabin_codes_proto_rawDescData
}

var file_io_flights_cabin_codes_cabin_codes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_flights_cabin_codes_cabin_codes_proto_goTypes = []interface{}{
	(CompartmentCode)(0), // 0: cabin_codes.CompartmentCode
}
var file_io_flights_cabin_codes_cabin_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_io_flights_cabin_codes_cabin_codes_proto_init() }
func file_io_flights_cabin_codes_cabin_codes_proto_init() {
	if File_io_flights_cabin_codes_cabin_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_flights_cabin_codes_cabin_codes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_flights_cabin_codes_cabin_codes_proto_goTypes,
		DependencyIndexes: file_io_flights_cabin_codes_cabin_codes_proto_depIdxs,
		EnumInfos:         file_io_flights_cabin_codes_cabin_codes_proto_enumTypes,
	}.Build()
	File_io_flights_cabin_codes_cabin_codes_proto = out.File
	file_io_flights_cabin_codes_cabin_codes_proto_rawDesc = nil
	file_io_flights_cabin_codes_cabin_codes_proto_goTypes = nil
	file_io_flights_cabin_codes_cabin_codes_proto_depIdxs = nil
}
