// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: io/member/event.proto

package members

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

// Details on the event. These info will be stored in Member Event.
type EventDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address the event took place.
	// @tag: validateCreate:"omitempty,max=255"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" validateCreate:"omitempty,max=255"`
	// Latitude the event took place.
	// @tag: validateCreate:"omitempty"
	Lat float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty" validateCreate:"omitempty"`
	// Longitude the event took place.
	// @tag: validateCreate:"omitempty"
	Lon float64 `protobuf:"fixed64,3,opt,name=lon,proto3" json:"lon,omitempty" validateCreate:"omitempty"`
	// Altitude the event took place (in metres).
	// @tag: validateCreate:"omitempty"
	Alt int32 `protobuf:"varint,4,opt,name=alt,proto3" json:"alt,omitempty" validateCreate:"omitempty"`
	// External unique ID of the event.
	// @tag: validateCreate:"omitempty,max=255"
	ExternalEventId string `protobuf:"bytes,5,opt,name=externalEventId,proto3" json:"externalEventId,omitempty" validateCreate:"omitempty,max=255"`
	// External device ID of the device that was used to capture the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalDeviceId string `protobuf:"bytes,6,opt,name=externalDeviceId,proto3" json:"externalDeviceId,omitempty" validateCreate:"omitempty,max=255"`
	// External service ID of the service that was used for capturing the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalServiceId string `protobuf:"bytes,7,opt,name=externalServiceId,proto3" json:"externalServiceId,omitempty" validateCreate:"omitempty,max=255"`
	// Any meta data (for example gathered on scanning) that is relevant to the event (# of points earner, bill spent, device meta-data, etc).
	// @tag: validateCreate:"omitempty"
	MetaData map[string]string `protobuf:"bytes,8,rep,name=metaData,proto3" json:"metaData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" validateCreate:"omitempty"`
	// Any relevant notes for the event.
	// @tag: validateCreate:"omitempty"
	Notes string `protobuf:"bytes,9,opt,name=notes,proto3" json:"notes,omitempty" validateCreate:"omitempty"`
	// External user ID of the logged in user that captured the event (for example when using an external scanning app).
	// @tag: validateCreate:"omitempty,max=255"
	ExternalUserId string `protobuf:"bytes,10,opt,name=externalUserId,proto3" json:"externalUserId,omitempty" validateCreate:"omitempty,max=255"`
}

func (x *EventDetails) Reset() {
	*x = EventDetails{}
	mi := &file_io_member_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDetails) ProtoMessage() {}

func (x *EventDetails) ProtoReflect() protoreflect.Message {
	mi := &file_io_member_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDetails.ProtoReflect.Descriptor instead.
func (*EventDetails) Descriptor() ([]byte, []int) {
	return file_io_member_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventDetails) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EventDetails) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *EventDetails) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *EventDetails) GetAlt() int32 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *EventDetails) GetExternalEventId() string {
	if x != nil {
		return x.ExternalEventId
	}
	return ""
}

func (x *EventDetails) GetExternalDeviceId() string {
	if x != nil {
		return x.ExternalDeviceId
	}
	return ""
}

func (x *EventDetails) GetExternalServiceId() string {
	if x != nil {
		return x.ExternalServiceId
	}
	return ""
}

func (x *EventDetails) GetMetaData() map[string]string {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *EventDetails) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *EventDetails) GetExternalUserId() string {
	if x != nil {
		return x.ExternalUserId
	}
	return ""
}

var File_io_member_event_proto protoreflect.FileDescriptor

var file_io_member_event_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x22, 0x9e, 0x03, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x5f, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5a, 0x2c, 0x73,
	0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f,
	0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0xaa, 0x02, 0x14, 0x50, 0x61,
	0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_member_event_proto_rawDescOnce sync.Once
	file_io_member_event_proto_rawDescData = file_io_member_event_proto_rawDesc
)

func file_io_member_event_proto_rawDescGZIP() []byte {
	file_io_member_event_proto_rawDescOnce.Do(func() {
		file_io_member_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_member_event_proto_rawDescData)
	})
	return file_io_member_event_proto_rawDescData
}

var file_io_member_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_member_event_proto_goTypes = []any{
	(*EventDetails)(nil), // 0: members.EventDetails
	nil,                  // 1: members.EventDetails.MetaDataEntry
}
var file_io_member_event_proto_depIdxs = []int32{
	1, // 0: members.EventDetails.metaData:type_name -> members.EventDetails.MetaDataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_io_member_event_proto_init() }
func file_io_member_event_proto_init() {
	if File_io_member_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_io_member_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_member_event_proto_goTypes,
		DependencyIndexes: file_io_member_event_proto_depIdxs,
		MessageInfos:      file_io_member_event_proto_msgTypes,
	}.Build()
	File_io_member_event_proto = out.File
	file_io_member_event_proto_rawDesc = nil
	file_io_member_event_proto_goTypes = nil
	file_io_member_event_proto_depIdxs = nil
}
