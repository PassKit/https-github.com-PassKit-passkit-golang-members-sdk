//*
// Expiry
//
// You can set expiry logic.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: io/common/expiry.proto

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

// Options to manage the expiry date of the digital card.
type ExpiryType int32

const (
	// Please do not use this enum. This enum do not have any effect on expiry logic.
	ExpiryType_EXPIRE_NONE ExpiryType = 0
	// Expiry date is set with year, month and date. The expiry date will be based on your timezone. The digital card will expire at 23:59:59:59.99999 of set date in fixed timezone. The expiry date is the same for all cards.
	ExpiryType_EXPIRE_ON_FIXED_DATE ExpiryType = 1
	// The digital card expires after the number of days after the digital card issuing.
	ExpiryType_EXPIRE_AFTER_X_DAYS ExpiryType = 2
	// If you want to change expiry date for each digital card, you can use this expiry type. You can set expiry date and time in fixed timezone.
	ExpiryType_EXPIRE_ON_VARIABLE_DATE_TIME ExpiryType = 3
	// The digital card will set as NULL and the pass will not expire..
	ExpiryType_EXPIRE_SET_TO_NULL ExpiryType = 4
)

// Enum value maps for ExpiryType.
var (
	ExpiryType_name = map[int32]string{
		0: "EXPIRE_NONE",
		1: "EXPIRE_ON_FIXED_DATE",
		2: "EXPIRE_AFTER_X_DAYS",
		3: "EXPIRE_ON_VARIABLE_DATE_TIME",
		4: "EXPIRE_SET_TO_NULL",
	}
	ExpiryType_value = map[string]int32{
		"EXPIRE_NONE":                  0,
		"EXPIRE_ON_FIXED_DATE":         1,
		"EXPIRE_AFTER_X_DAYS":          2,
		"EXPIRE_ON_VARIABLE_DATE_TIME": 3,
		"EXPIRE_SET_TO_NULL":           4,
	}
)

func (x ExpiryType) Enum() *ExpiryType {
	p := new(ExpiryType)
	*p = x
	return p
}

func (x ExpiryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpiryType) Descriptor() protoreflect.EnumDescriptor {
	return file_io_common_expiry_proto_enumTypes[0].Descriptor()
}

func (ExpiryType) Type() protoreflect.EnumType {
	return &file_io_common_expiry_proto_enumTypes[0]
}

func (x ExpiryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpiryType.Descriptor instead.
func (ExpiryType) EnumDescriptor() ([]byte, []int) {
	return file_io_common_expiry_proto_rawDescGZIP(), []int{0}
}

// The digital card will be expired on the expiry date. The barcode will not be rendered on digital card and the card itself will not be updated after it has been expired.
type ExpirySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpiryType ExpiryType `protobuf:"varint,1,opt,name=expiryType,proto3,enum=io.ExpiryType" json:"expiryType,omitempty"`
	// Types that are assignable to ExpiryOneof:
	//
	//	*ExpirySettings_FixedExpiryDate
	//	*ExpirySettings_ExpireAfterXDays
	ExpiryOneof isExpirySettings_ExpiryOneof `protobuf_oneof:"expiryOneof"`
}

func (x *ExpirySettings) Reset() {
	*x = ExpirySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_expiry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpirySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpirySettings) ProtoMessage() {}

func (x *ExpirySettings) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_expiry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpirySettings.ProtoReflect.Descriptor instead.
func (*ExpirySettings) Descriptor() ([]byte, []int) {
	return file_io_common_expiry_proto_rawDescGZIP(), []int{0}
}

func (x *ExpirySettings) GetExpiryType() ExpiryType {
	if x != nil {
		return x.ExpiryType
	}
	return ExpiryType_EXPIRE_NONE
}

func (m *ExpirySettings) GetExpiryOneof() isExpirySettings_ExpiryOneof {
	if m != nil {
		return m.ExpiryOneof
	}
	return nil
}

func (x *ExpirySettings) GetFixedExpiryDate() *Date {
	if x, ok := x.GetExpiryOneof().(*ExpirySettings_FixedExpiryDate); ok {
		return x.FixedExpiryDate
	}
	return nil
}

func (x *ExpirySettings) GetExpireAfterXDays() uint32 {
	if x, ok := x.GetExpiryOneof().(*ExpirySettings_ExpireAfterXDays); ok {
		return x.ExpireAfterXDays
	}
	return 0
}

type isExpirySettings_ExpiryOneof interface {
	isExpirySettings_ExpiryOneof()
}

type ExpirySettings_FixedExpiryDate struct {
	// The expiry date for digital membership card.
	// Please set expiryType as EXPIRE_ON_FIXED_DATE.
	FixedExpiryDate *Date `protobuf:"bytes,2,opt,name=fixedExpiryDate,proto3,oneof"`
}

type ExpirySettings_ExpireAfterXDays struct {
	// Number of days the digital membership card is expired after the enrolment date.
	// Please set expiryType as EXPIRE_AFTER_X_DAYS.
	ExpireAfterXDays uint32 `protobuf:"varint,3,opt,name=expireAfterXDays,proto3,oneof"`
}

func (*ExpirySettings_FixedExpiryDate) isExpirySettings_ExpiryOneof() {}

func (*ExpirySettings_ExpireAfterXDays) isExpirySettings_ExpiryOneof() {}

var File_io_common_expiry_proto protoreflect.FileDescriptor

var file_io_common_expiry_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x1a, 0x1e, 0x69, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a,
	0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x34, 0x0a, 0x0f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x58, 0x44, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x58, 0x44,
	0x61, 0x79, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x4f, 0x6e, 0x65,
	0x6f, 0x66, 0x2a, 0x8a, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x5f, 0x4f, 0x4e, 0x5f,
	0x46, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x44,
	0x41, 0x59, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x5f,
	0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x04, 0x42,
	0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b,
	0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_expiry_proto_rawDescOnce sync.Once
	file_io_common_expiry_proto_rawDescData = file_io_common_expiry_proto_rawDesc
)

func file_io_common_expiry_proto_rawDescGZIP() []byte {
	file_io_common_expiry_proto_rawDescOnce.Do(func() {
		file_io_common_expiry_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_expiry_proto_rawDescData)
	})
	return file_io_common_expiry_proto_rawDescData
}

var file_io_common_expiry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_io_common_expiry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_io_common_expiry_proto_goTypes = []interface{}{
	(ExpiryType)(0),        // 0: io.ExpiryType
	(*ExpirySettings)(nil), // 1: io.ExpirySettings
	(*Date)(nil),           // 2: io.Date
}
var file_io_common_expiry_proto_depIdxs = []int32{
	0, // 0: io.ExpirySettings.expiryType:type_name -> io.ExpiryType
	2, // 1: io.ExpirySettings.fixedExpiryDate:type_name -> io.Date
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_io_common_expiry_proto_init() }
func file_io_common_expiry_proto_init() {
	if File_io_common_expiry_proto != nil {
		return
	}
	file_io_common_common_objects_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_io_common_expiry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpirySettings); i {
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
	file_io_common_expiry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExpirySettings_FixedExpiryDate)(nil),
		(*ExpirySettings_ExpireAfterXDays)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_common_expiry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_expiry_proto_goTypes,
		DependencyIndexes: file_io_common_expiry_proto_depIdxs,
		EnumInfos:         file_io_common_expiry_proto_enumTypes,
		MessageInfos:      file_io_common_expiry_proto_msgTypes,
	}.Build()
	File_io_common_expiry_proto = out.File
	file_io_common_expiry_proto_rawDesc = nil
	file_io_common_expiry_proto_goTypes = nil
	file_io_common_expiry_proto_depIdxs = nil
}
