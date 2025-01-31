//*
// Billing
//
// Keep track of your account quota and consumption.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: io/common/billing.proto

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

type Quota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details map[uint32]*QuotaDetails `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // key is one of pk.MeteredEventType
}

func (x *Quota) Reset() {
	*x = Quota{}
	mi := &file_io_common_billing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_billing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_io_common_billing_proto_rawDescGZIP(), []int{0}
}

func (x *Quota) GetDetails() map[uint32]*QuotaDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

type QuotaDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quota  int32  `protobuf:"varint,1,opt,name=quota,proto3" json:"quota,omitempty"`
	Status uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"` // contains bitmask of pk.Status
}

func (x *QuotaDetails) Reset() {
	*x = QuotaDetails{}
	mi := &file_io_common_billing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuotaDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaDetails) ProtoMessage() {}

func (x *QuotaDetails) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_billing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaDetails.ProtoReflect.Descriptor instead.
func (*QuotaDetails) Descriptor() ([]byte, []int) {
	return file_io_common_billing_proto_rawDescGZIP(), []int{1}
}

func (x *QuotaDetails) GetQuota() int32 {
	if x != nil {
		return x.Quota
	}
	return 0
}

func (x *QuotaDetails) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_io_common_billing_proto protoreflect.FileDescriptor

var file_io_common_billing_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x22, 0x87, 0x01,
	0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x2e, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x4c, 0x0a, 0x0c, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6f, 0x2e,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3c, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa,
	0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_billing_proto_rawDescOnce sync.Once
	file_io_common_billing_proto_rawDescData = file_io_common_billing_proto_rawDesc
)

func file_io_common_billing_proto_rawDescGZIP() []byte {
	file_io_common_billing_proto_rawDescOnce.Do(func() {
		file_io_common_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_billing_proto_rawDescData)
	})
	return file_io_common_billing_proto_rawDescData
}

var file_io_common_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_io_common_billing_proto_goTypes = []any{
	(*Quota)(nil),        // 0: io.Quota
	(*QuotaDetails)(nil), // 1: io.QuotaDetails
	nil,                  // 2: io.Quota.DetailsEntry
}
var file_io_common_billing_proto_depIdxs = []int32{
	2, // 0: io.Quota.details:type_name -> io.Quota.DetailsEntry
	1, // 1: io.Quota.DetailsEntry.value:type_name -> io.QuotaDetails
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_io_common_billing_proto_init() }
func file_io_common_billing_proto_init() {
	if File_io_common_billing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_common_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_billing_proto_goTypes,
		DependencyIndexes: file_io_common_billing_proto_depIdxs,
		MessageInfos:      file_io_common_billing_proto_msgTypes,
	}.Build()
	File_io_common_billing_proto = out.File
	file_io_common_billing_proto_rawDesc = nil
	file_io_common_billing_proto_goTypes = nil
	file_io_common_billing_proto_depIdxs = nil
}
