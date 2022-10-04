//*
// Proximity
//
// Settings for lock screen notifications. Lock screen notifications are passive messages and smart phones do not vibrate when message is received.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: io/common/proximity.proto

package io

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Beacons allow you to display subtle messages to your users when they are in close proximity.
type Beacon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// System generated unique identifier for your beacon
	// @tag: validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"`
	// A valid UUID that is being broadcast from your beacon.
	// @tag: validateGeneric:"required_without=Id,uuid_rfc4122"
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty" validateGeneric:"required_without=Id,uuid_rfc4122"`
	// A friendly name used to display the beacon in the admin portal
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Major indicator.
	// @tag: validateGeneric:"required_without=Id,min=0,max=65535"
	Major uint32 `protobuf:"varint,4,opt,name=major,proto3" json:"major,omitempty" validateGeneric:"required_without=Id,min=0,max=65535"`
	// Minor indicator.
	// @tag: validateGeneric:"required_without=Id,min=0,max=65535"
	Minor uint32 `protobuf:"varint,5,opt,name=minor,proto3" json:"minor,omitempty" validateGeneric:"required_without=Id,min=0,max=65535"`
	// Message to be displayed on lock screen when user is in proximity of the beacon (iOS only).
	// @tag: validateGeneric:"required_without=Id"
	LockScreenMessage string `protobuf:"bytes,6,opt,name=lockScreenMessage,proto3" json:"lockScreenMessage,omitempty" validateGeneric:"required_without=Id"`
	// Localized lock screen message.
	LocalizedLockScreenMessage *LocalizedString `protobuf:"bytes,7,opt,name=localizedLockScreenMessage,proto3" json:"localizedLockScreenMessage,omitempty"`
	// Beacons will be added in order of their position, from lowest to highest. Position can be any value, E.g. 3 Beacons with positions 30, 10, 20 would be added 10 first, 20 second and 30 third.  If no position is provided and the number of beacons exceeds 10, there is no guarantee which beacon(s) would be excluded from the pass.
	Position uint32 `protobuf:"varint,8,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Beacon) Reset() {
	*x = Beacon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_proximity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beacon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beacon) ProtoMessage() {}

func (x *Beacon) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_proximity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beacon.ProtoReflect.Descriptor instead.
func (*Beacon) Descriptor() ([]byte, []int) {
	return file_io_common_proximity_proto_rawDescGZIP(), []int{0}
}

func (x *Beacon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Beacon) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Beacon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Beacon) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Beacon) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Beacon) GetLockScreenMessage() string {
	if x != nil {
		return x.LockScreenMessage
	}
	return ""
}

func (x *Beacon) GetLocalizedLockScreenMessage() *LocalizedString {
	if x != nil {
		return x.LocalizedLockScreenMessage
	}
	return nil
}

func (x *Beacon) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

// Locations provide a broader range using GPS to trigger a lockscreen message.
type GPSLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// System generated unique identifier for your GPS Location
	// @tag: validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"`
	// a friendly name for the location used to display the location in the admin portal
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Latitude.
	// @tag: validateGeneric:"required_without=Id,latitude"
	Lat float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty" validateGeneric:"required_without=Id,latitude"`
	// Longitude.
	// @tag: validateGeneric:"required_with=Lat,required_without=Id,longitude"
	Lon float64 `protobuf:"fixed64,4,opt,name=lon,proto3" json:"lon,omitempty" validateGeneric:"required_with=Lat,required_without=Id,longitude"`
	// Altitude in metres.
	// @tag: validateGeneric:"omitempty,max=8848"
	Alt int32 `protobuf:"varint,5,opt,name=alt,proto3" json:"alt,omitempty" validateGeneric:"omitempty,max=8848"`
	// Message to be displayed on lock screen when user is in the location (iOS only).
	// @tag: validateGeneric:"required_without=Id"
	LockScreenMessage string `protobuf:"bytes,6,opt,name=lockScreenMessage,proto3" json:"lockScreenMessage,omitempty" validateGeneric:"required_without=Id"`
	// Localized lock screen message.
	LocalizedLockScreenMessage *LocalizedString `protobuf:"bytes,7,opt,name=localizedLockScreenMessage,proto3" json:"localizedLockScreenMessage,omitempty"`
	// Locations will be added in order of their position, from lowest to highest. Position can be any value, E.g. 3 Locations with positions 30, 10, 20 would be added 10 first, 20 second and 30 third.  If no position is provided and the number of locations exceeds 10, there is no guarantee which location(s) would be excluded from the pass.
	Position uint32 `protobuf:"varint,8,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *GPSLocation) Reset() {
	*x = GPSLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_proximity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPSLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPSLocation) ProtoMessage() {}

func (x *GPSLocation) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_proximity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPSLocation.ProtoReflect.Descriptor instead.
func (*GPSLocation) Descriptor() ([]byte, []int) {
	return file_io_common_proximity_proto_rawDescGZIP(), []int{1}
}

func (x *GPSLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GPSLocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GPSLocation) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GPSLocation) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *GPSLocation) GetAlt() int32 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *GPSLocation) GetLockScreenMessage() string {
	if x != nil {
		return x.LockScreenMessage
	}
	return ""
}

func (x *GPSLocation) GetLocalizedLockScreenMessage() *LocalizedString {
	if x != nil {
		return x.LocalizedLockScreenMessage
	}
	return nil
}

func (x *GPSLocation) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

// Relevant Date in ISO8601 datetime.
type RelevantDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp. ISO8601 datetime.
	Timestamp uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RelevantDate) Reset() {
	*x = RelevantDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_proximity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelevantDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelevantDate) ProtoMessage() {}

func (x *RelevantDate) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_proximity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelevantDate.ProtoReflect.Descriptor instead.
func (*RelevantDate) Descriptor() ([]byte, []int) {
	return file_io_common_proximity_proto_rawDescGZIP(), []int{2}
}

func (x *RelevantDate) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_io_common_proximity_proto protoreflect.FileDescriptor

var file_io_common_proximity_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x6d, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69, 0x6f, 0x1a,
	0x1c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x03,
	0x0a, 0x06, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0c, 0x92, 0x41, 0x09, 0x59, 0x00, 0x00, 0x00, 0x00, 0xe0, 0xff, 0xef, 0x40, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x59, 0x00, 0x00, 0x00, 0x00, 0xe0, 0xff, 0xef,
	0x40, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x1a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0xd4, 0x01, 0x92, 0x41, 0xd0, 0x01, 0x0a, 0xcd,
	0x01, 0x2a, 0x06, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x32, 0xa7, 0x01, 0x41, 0x20, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x20,
	0x70, 0x75, 0x73, 0x68, 0x20, 0x61, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74,
	0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x6d, 0x69, 0x74, 0x79, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x20, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x20, 0x61,
	0x72, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x20, 0x6f, 0x6e, 0x6c,
	0x79, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x69,
	0x4f, 0x53, 0x2e, 0xd2, 0x01, 0x04, 0x75, 0x75, 0x69, 0x64, 0xd2, 0x01, 0x11, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe6,
	0x03, 0x0a, 0x0b, 0x47, 0x50, 0x53, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x1a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0xdd, 0x01, 0x92, 0x41, 0xd9, 0x01, 0x0a, 0xd6,
	0x01, 0x2a, 0x0c, 0x47, 0x50, 0x53, 0x20, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xa4, 0x01, 0x41, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x20, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x70, 0x61, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x70, 0x75, 0x73, 0x68, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x6e, 0x65, 0x61, 0x72, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x20, 0x46, 0x6f, 0x72, 0x20,
	0x69, 0x4f, 0x53, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x61, 0x20, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
	0x63, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x73, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x64, 0x2e, 0xd2, 0x01, 0x03, 0x6c, 0x61, 0x74, 0xd2, 0x01, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0xd2, 0x01, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65, 0x76,
	0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x47, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68,
	0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0xaa,
	0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_proximity_proto_rawDescOnce sync.Once
	file_io_common_proximity_proto_rawDescData = file_io_common_proximity_proto_rawDesc
)

func file_io_common_proximity_proto_rawDescGZIP() []byte {
	file_io_common_proximity_proto_rawDescOnce.Do(func() {
		file_io_common_proximity_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_proximity_proto_rawDescData)
	})
	return file_io_common_proximity_proto_rawDescData
}

var file_io_common_proximity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_io_common_proximity_proto_goTypes = []interface{}{
	(*Beacon)(nil),          // 0: io.Beacon
	(*GPSLocation)(nil),     // 1: io.GPSLocation
	(*RelevantDate)(nil),    // 2: io.RelevantDate
	(*LocalizedString)(nil), // 3: io.LocalizedString
}
var file_io_common_proximity_proto_depIdxs = []int32{
	3, // 0: io.Beacon.localizedLockScreenMessage:type_name -> io.LocalizedString
	3, // 1: io.GPSLocation.localizedLockScreenMessage:type_name -> io.LocalizedString
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_io_common_proximity_proto_init() }
func file_io_common_proximity_proto_init() {
	if File_io_common_proximity_proto != nil {
		return
	}
	file_io_common_localization_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_io_common_proximity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beacon); i {
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
		file_io_common_proximity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPSLocation); i {
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
		file_io_common_proximity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelevantDate); i {
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
			RawDescriptor: file_io_common_proximity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_proximity_proto_goTypes,
		DependencyIndexes: file_io_common_proximity_proto_depIdxs,
		MessageInfos:      file_io_common_proximity_proto_msgTypes,
	}.Build()
	File_io_common_proximity_proto = out.File
	file_io_common_proximity_proto_rawDesc = nil
	file_io_common_proximity_proto_goTypes = nil
	file_io_common_proximity_proto_depIdxs = nil
}
