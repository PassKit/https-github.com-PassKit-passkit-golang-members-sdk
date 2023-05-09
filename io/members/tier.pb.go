//*
// Tiers
//
// Tier allows to override certain details of the pass (colours, logo's, dynamic back fields, and labels)
//
// todo: we might need something around tier points & auto messaging to indicate they need to take action soon if they want to remain in the tier? Or auto message to tell people they are only xx points away from reaching the next tier?

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.4
// source: io/member/tier.proto

package members

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

// The Tier object; will override certain details of the pass (colours, logo's, dynamic back fields, and labels).
type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tier identifier (could just be: blue, gold, etc); needs to be lower case. Tier ID needs to be uique within the program.
	// @tag: validateGeneric:"required" validateCreate:"-" validateUpdate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"required" validateCreate:"-" validateUpdate:"required"`
	// Index of the tier; can be used for managing downgrades / upgrade messaging; needs to be unique within the program.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	TierIndex uint32 `protobuf:"varint,2,opt,name=tierIndex,proto3" json:"tierIndex,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Name of tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Localized name of tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedName *io.LocalizedString `protobuf:"bytes,4,opt,name=localizedName,proto3" json:"localizedName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Name of the secondary reward tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	SecondaryTierName string `protobuf:"bytes,5,opt,name=secondaryTierName,proto3" json:"secondaryTierName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Localized name of the secondary reward tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedSecondaryTierName *io.LocalizedString `protobuf:"bytes,6,opt,name=localizedSecondaryTierName,proto3" json:"localizedSecondaryTierName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The program id that the tier belongs to.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	ProgramId string `protobuf:"bytes,7,opt,name=programId,proto3" json:"programId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Identifier of pass template which identifies design and data elements for this tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	PassTemplateId string `protobuf:"bytes,8,opt,name=passTemplateId,proto3" json:"passTemplateId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Notification to be shown when someone upgrades tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	TierUpgradeMessage string `protobuf:"bytes,9,opt,name=tierUpgradeMessage,proto3" json:"tierUpgradeMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Localized notification to be shown when someone upgrades tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedTierUpgradeMessage *io.LocalizedString `protobuf:"bytes,10,opt,name=localizedTierUpgradeMessage,proto3" json:"localizedTierUpgradeMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Notification to be shown when someone downgrades tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	TierDowngradeMessage string `protobuf:"bytes,11,opt,name=tierDowngradeMessage,proto3" json:"tierDowngradeMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Localized notification to be shown when someone upgrades tier.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedTierDowngradeMessage *io.LocalizedString `protobuf:"bytes,12,opt,name=localizedTierDowngradeMessage,proto3" json:"localizedTierDowngradeMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The date the tier was created.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Created *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created,proto3" json:"created,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// The date the tier was updated.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Updated *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated,proto3" json:"updated,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// Set true to accept a negative point balance. Default is false.
	PointsOverdrawn bool `protobuf:"varint,15,opt,name=pointsOverdrawn,proto3" json:"pointsOverdrawn,omitempty"`
	// Set true to accept a negative secondary point balance. Default is false.
	SecondaryPointsOverdrawn bool `protobuf:"varint,16,opt,name=secondaryPointsOverdrawn,proto3" json:"secondaryPointsOverdrawn,omitempty"`
	// Expiry date setting. You can set expiry logic here. Default no expiry logic.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	ExpirySettings *io.ExpirySettings `protobuf:"bytes,17,opt,name=expirySettings,proto3" json:"expirySettings,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Timezone string in IANA timezone format. If not provided defaults to Etc/UTC.
	// @tag: validateGeneric:"omitempty,ianaTimeZone" validateCreate:"required,ianaTimeZone" validateUpdate:"omitempty,ianaTimeZone"
	Timezone string `protobuf:"bytes,18,opt,name=timezone,proto3" json:"timezone,omitempty" validateGeneric:"omitempty,ianaTimeZone" validateCreate:"required,ianaTimeZone" validateUpdate:"omitempty,ianaTimeZone"`
	// Indicates if customers can enrol into this tier via a public web form (only for public programs).
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	AllowTierEnrolment *io.PkBool `protobuf:"bytes,19,opt,name=allowTierEnrolment,proto3" json:"allowTierEnrolment,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Tier shortcode is used in public enrolment URLs, that enrol members into the tier if the program is set to public and allowTierEnrolment = true for this tier. System generated.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	ShortCode string `protobuf:"bytes,20,opt,name=shortCode,proto3" json:"shortCode,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_member_tier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_io_member_tier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_io_member_tier_proto_rawDescGZIP(), []int{0}
}

func (x *Tier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tier) GetTierIndex() uint32 {
	if x != nil {
		return x.TierIndex
	}
	return 0
}

func (x *Tier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tier) GetLocalizedName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedName
	}
	return nil
}

func (x *Tier) GetSecondaryTierName() string {
	if x != nil {
		return x.SecondaryTierName
	}
	return ""
}

func (x *Tier) GetLocalizedSecondaryTierName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedSecondaryTierName
	}
	return nil
}

func (x *Tier) GetProgramId() string {
	if x != nil {
		return x.ProgramId
	}
	return ""
}

func (x *Tier) GetPassTemplateId() string {
	if x != nil {
		return x.PassTemplateId
	}
	return ""
}

func (x *Tier) GetTierUpgradeMessage() string {
	if x != nil {
		return x.TierUpgradeMessage
	}
	return ""
}

func (x *Tier) GetLocalizedTierUpgradeMessage() *io.LocalizedString {
	if x != nil {
		return x.LocalizedTierUpgradeMessage
	}
	return nil
}

func (x *Tier) GetTierDowngradeMessage() string {
	if x != nil {
		return x.TierDowngradeMessage
	}
	return ""
}

func (x *Tier) GetLocalizedTierDowngradeMessage() *io.LocalizedString {
	if x != nil {
		return x.LocalizedTierDowngradeMessage
	}
	return nil
}

func (x *Tier) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Tier) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *Tier) GetPointsOverdrawn() bool {
	if x != nil {
		return x.PointsOverdrawn
	}
	return false
}

func (x *Tier) GetSecondaryPointsOverdrawn() bool {
	if x != nil {
		return x.SecondaryPointsOverdrawn
	}
	return false
}

func (x *Tier) GetExpirySettings() *io.ExpirySettings {
	if x != nil {
		return x.ExpirySettings
	}
	return nil
}

func (x *Tier) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Tier) GetAllowTierEnrolment() *io.PkBool {
	if x != nil {
		return x.AllowTierEnrolment
	}
	return nil
}

func (x *Tier) GetShortCode() string {
	if x != nil {
		return x.ShortCode
	}
	return ""
}

// Used to request member tier record by tier id.
type TierRequestInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Program ID.
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	ProgramId string `protobuf:"bytes,1,opt,name=programId,proto3" json:"programId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// Tier ID.
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	TierId string `protobuf:"bytes,2,opt,name=tierId,proto3" json:"tierId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
}

func (x *TierRequestInput) Reset() {
	*x = TierRequestInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_member_tier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TierRequestInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TierRequestInput) ProtoMessage() {}

func (x *TierRequestInput) ProtoReflect() protoreflect.Message {
	mi := &file_io_member_tier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TierRequestInput.ProtoReflect.Descriptor instead.
func (*TierRequestInput) Descriptor() ([]byte, []int) {
	return file_io_member_tier_proto_rawDescGZIP(), []int{1}
}

func (x *TierRequestInput) GetProgramId() string {
	if x != nil {
		return x.ProgramId
	}
	return ""
}

func (x *TierRequestInput) GetTierId() string {
	if x != nil {
		return x.TierId
	}
	return ""
}

var File_io_member_tier_proto protoreflect.FileDescriptor

var file_io_member_tier_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x08, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x54, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x54, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x54, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x1a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x54, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x69, 0x65, 0x72, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x74, 0x69, 0x65, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x55, 0x0a, 0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x54, 0x69, 0x65, 0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x69, 0x65, 0x72, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14,
	0x74, 0x69, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x69, 0x65, 0x72,
	0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x59, 0x0a, 0x1d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x69, 0x65,
	0x72, 0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x1d, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x69, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x72, 0x61, 0x77,
	0x6e, 0x12, 0x3a, 0x0a, 0x18, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x12, 0x3a, 0x0a,
	0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6f, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69,
	0x65, 0x72, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6b, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x12, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x3a,
	0x84, 0x01, 0x92, 0x41, 0x80, 0x01, 0x0a, 0x7e, 0x2a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x32, 0x41,
	0x54, 0x69, 0x65, 0x72, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x20, 0x63, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x2e, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x09, 0x74, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0xd2, 0x01, 0x09, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0xd2, 0x01,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x54, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x5f, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5a, 0x2c, 0x73, 0x74,
	0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f,
	0x69, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0xaa, 0x02, 0x14, 0x50, 0x61, 0x73,
	0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_member_tier_proto_rawDescOnce sync.Once
	file_io_member_tier_proto_rawDescData = file_io_member_tier_proto_rawDesc
)

func file_io_member_tier_proto_rawDescGZIP() []byte {
	file_io_member_tier_proto_rawDescOnce.Do(func() {
		file_io_member_tier_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_member_tier_proto_rawDescData)
	})
	return file_io_member_tier_proto_rawDescData
}

var file_io_member_tier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_member_tier_proto_goTypes = []interface{}{
	(*Tier)(nil),                  // 0: members.Tier
	(*TierRequestInput)(nil),      // 1: members.TierRequestInput
	(*io.LocalizedString)(nil),    // 2: io.LocalizedString
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*io.ExpirySettings)(nil),     // 4: io.ExpirySettings
	(*io.PkBool)(nil),             // 5: io.PkBool
}
var file_io_member_tier_proto_depIdxs = []int32{
	2, // 0: members.Tier.localizedName:type_name -> io.LocalizedString
	2, // 1: members.Tier.localizedSecondaryTierName:type_name -> io.LocalizedString
	2, // 2: members.Tier.localizedTierUpgradeMessage:type_name -> io.LocalizedString
	2, // 3: members.Tier.localizedTierDowngradeMessage:type_name -> io.LocalizedString
	3, // 4: members.Tier.created:type_name -> google.protobuf.Timestamp
	3, // 5: members.Tier.updated:type_name -> google.protobuf.Timestamp
	4, // 6: members.Tier.expirySettings:type_name -> io.ExpirySettings
	5, // 7: members.Tier.allowTierEnrolment:type_name -> io.PkBool
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_io_member_tier_proto_init() }
func file_io_member_tier_proto_init() {
	if File_io_member_tier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_io_member_tier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
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
		file_io_member_tier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TierRequestInput); i {
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
			RawDescriptor: file_io_member_tier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_member_tier_proto_goTypes,
		DependencyIndexes: file_io_member_tier_proto_depIdxs,
		MessageInfos:      file_io_member_tier_proto_msgTypes,
	}.Build()
	File_io_member_tier_proto = out.File
	file_io_member_tier_proto_rawDesc = nil
	file_io_member_tier_proto_goTypes = nil
	file_io_member_tier_proto_depIdxs = nil
}
