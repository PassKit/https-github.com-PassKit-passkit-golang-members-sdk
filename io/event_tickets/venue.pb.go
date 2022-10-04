//*
// Venue indicates where an event takes place.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: io/event_tickets/venue.proto

package event_tickets

import (
	reflect "reflect"
	sync "sync"

	io "github.com/PassKit/passkit-golang-grpc-sdk/io"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Venue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PassKit generated venue id (22 characters).
	// @tag: validateGeneric:"required_without=Uid" validateCreate:"isdefault" validateUpdate:"required_without=Uid"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"required_without=Uid" validateCreate:"isdefault" validateUpdate:"required_without=Uid"`
	// User generated venue id; unique within the user account.
	// @tag: validateGeneric:"required_without=Id" validateCreate:"omitempty" validateUpdate:"required_without=Id"
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" validateGeneric:"required_without=Id" validateCreate:"omitempty" validateUpdate:"required_without=Id"`
	// The venue name.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Localized venue name.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedName *io.LocalizedString `protobuf:"bytes,4,opt,name=localizedName,proto3" json:"localizedName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The venue address.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Localized venue address.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedAddress *io.LocalizedString `protobuf:"bytes,6,opt,name=localizedAddress,proto3" json:"localizedAddress,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Timezone applied for the event dates. e.g. America/New_York, Asia/Singapore, Europe/London.
	// @tag: validateGeneric:"omitempty,ianaTimeZone" validateCreate:"required,ianaTimeZone" validateUpdate:"required,ianaTimeZone"
	Timezone string `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty" validateGeneric:"omitempty,ianaTimeZone" validateCreate:"required,ianaTimeZone" validateUpdate:"required,ianaTimeZone"`
	// Optional GPS location details of the venue. If provided will be embedded into the ticket as the first GPS location.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	GpsCoords []*io.GPSLocation `protobuf:"bytes,8,rep,name=gpsCoords,proto3" json:"gpsCoords,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The date the venue was created. Not writable.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Created *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created,proto3" json:"created,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// The date the venue updated. Not writable.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Updated *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated,proto3" json:"updated,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
}

func (x *Venue) Reset() {
	*x = Venue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_event_tickets_venue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Venue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Venue) ProtoMessage() {}

func (x *Venue) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_venue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Venue.ProtoReflect.Descriptor instead.
func (*Venue) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_venue_proto_rawDescGZIP(), []int{0}
}

func (x *Venue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Venue) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Venue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Venue) GetLocalizedName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedName
	}
	return nil
}

func (x *Venue) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Venue) GetLocalizedAddress() *io.LocalizedString {
	if x != nil {
		return x.LocalizedAddress
	}
	return nil
}

func (x *Venue) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Venue) GetGpsCoords() []*io.GPSLocation {
	if x != nil {
		return x.GpsCoords
	}
	return nil
}

func (x *Venue) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Venue) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type VenueLimitedFieldsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PassKit generated venue id (22 characters).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User generated venue id; unique within the user account.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// The venue name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *VenueLimitedFieldsResponse) Reset() {
	*x = VenueLimitedFieldsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_event_tickets_venue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VenueLimitedFieldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VenueLimitedFieldsResponse) ProtoMessage() {}

func (x *VenueLimitedFieldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_io_event_tickets_venue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VenueLimitedFieldsResponse.ProtoReflect.Descriptor instead.
func (*VenueLimitedFieldsResponse) Descriptor() ([]byte, []int) {
	return file_io_event_tickets_venue_proto_rawDescGZIP(), []int{1}
}

func (x *VenueLimitedFieldsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VenueLimitedFieldsResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *VenueLimitedFieldsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_io_event_tickets_venue_proto protoreflect.FileDescriptor

var file_io_event_tickets_venue_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x05, 0x56, 0x65, 0x6e, 0x75,
	0x65, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x92,
	0x41, 0x02, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x67, 0x70, 0x73, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x50, 0x53, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x67, 0x70, 0x73, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x3b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41,
	0x02, 0x40, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40, 0x01,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x63, 0x92, 0x41, 0x60, 0x0a, 0x5e,
	0x2a, 0x05, 0x56, 0x65, 0x6e, 0x75, 0x65, 0x32, 0x35, 0x56, 0x65, 0x6e, 0x75, 0x65, 0x20, 0x68,
	0x6f, 0x6c, 0x64, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x20, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x20, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0xd2, 0x01,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0xd2,
	0x01, 0x0c, 0x69, 0x61, 0x6e, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0xa3,
	0x01, 0x0a, 0x1a, 0x56, 0x65, 0x6e, 0x75, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x3a, 0x4f, 0x92, 0x41, 0x4c, 0x0a, 0x4a, 0x2a, 0x11, 0x56, 0x65, 0x6e, 0x75,
	0x65, 0x20, 0x28, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x76, 0x65, 0x72, 0x29, 0x32, 0x35, 0x56,
	0x65, 0x6e, 0x75, 0x65, 0x20, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x20, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x6f, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73,
	0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x5a, 0x32, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0xaa, 0x02, 0x19, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_event_tickets_venue_proto_rawDescOnce sync.Once
	file_io_event_tickets_venue_proto_rawDescData = file_io_event_tickets_venue_proto_rawDesc
)

func file_io_event_tickets_venue_proto_rawDescGZIP() []byte {
	file_io_event_tickets_venue_proto_rawDescOnce.Do(func() {
		file_io_event_tickets_venue_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_event_tickets_venue_proto_rawDescData)
	})
	return file_io_event_tickets_venue_proto_rawDescData
}

var file_io_event_tickets_venue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_event_tickets_venue_proto_goTypes = []interface{}{
	(*Venue)(nil),                      // 0: event_tickets.Venue
	(*VenueLimitedFieldsResponse)(nil), // 1: event_tickets.VenueLimitedFieldsResponse
	(*io.LocalizedString)(nil),         // 2: io.LocalizedString
	(*io.GPSLocation)(nil),             // 3: io.GPSLocation
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_io_event_tickets_venue_proto_depIdxs = []int32{
	2, // 0: event_tickets.Venue.localizedName:type_name -> io.LocalizedString
	2, // 1: event_tickets.Venue.localizedAddress:type_name -> io.LocalizedString
	3, // 2: event_tickets.Venue.gpsCoords:type_name -> io.GPSLocation
	4, // 3: event_tickets.Venue.created:type_name -> google.protobuf.Timestamp
	4, // 4: event_tickets.Venue.updated:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_io_event_tickets_venue_proto_init() }
func file_io_event_tickets_venue_proto_init() {
	if File_io_event_tickets_venue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_io_event_tickets_venue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Venue); i {
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
		file_io_event_tickets_venue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VenueLimitedFieldsResponse); i {
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
			RawDescriptor: file_io_event_tickets_venue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_event_tickets_venue_proto_goTypes,
		DependencyIndexes: file_io_event_tickets_venue_proto_depIdxs,
		MessageInfos:      file_io_event_tickets_venue_proto_msgTypes,
	}.Build()
	File_io_event_tickets_venue_proto = out.File
	file_io_event_tickets_venue_proto_rawDesc = nil
	file_io_event_tickets_venue_proto_goTypes = nil
	file_io_event_tickets_venue_proto_depIdxs = nil
}
