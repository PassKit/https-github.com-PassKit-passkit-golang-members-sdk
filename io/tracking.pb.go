// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/tracking.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Standard Facebook Pixel events that we support
type FbPixelStandardEvent int32

const (
	FbPixelStandardEvent_FBP_STANDARD_EVENT_DO_NOT_USE FbPixelStandardEvent = 0
	// This is the default pixel tracking page visits.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_PAGE_VIEW FbPixelStandardEvent = 1
	// When a registration form is completed.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_COMPLETE_REGISTRATION FbPixelStandardEvent = 2
	// When a sign up is completed.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_LEAD FbPixelStandardEvent = 3
	// When a purchase is made or checkout flow is completed.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_PURCHASE FbPixelStandardEvent = 4
	// When a person books an appointment to visit one of your locations.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_SCHEDULE FbPixelStandardEvent = 5
	// When a person starts a free trial of a product or service you offer.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_START_TRIAL FbPixelStandardEvent = 6
	// When a person applies for a product, service, or program you offer.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_SUBMIT_APPLICATION FbPixelStandardEvent = 7
	// When a person applies to a start a paid subscription for a product or service you offer.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_SUBSCRIBE FbPixelStandardEvent = 8
	// A visit to a web page you care about (for example, a product page or landing page). ViewContent tells you if someone visits a web page's URL, but not what they see or do on that page.
	FbPixelStandardEvent_FBP_STANDARD_EVENT_VIEW_CONTENT FbPixelStandardEvent = 9
)

var FbPixelStandardEvent_name = map[int32]string{
	0: "FBP_STANDARD_EVENT_DO_NOT_USE",
	1: "FBP_STANDARD_EVENT_PAGE_VIEW",
	2: "FBP_STANDARD_EVENT_COMPLETE_REGISTRATION",
	3: "FBP_STANDARD_EVENT_LEAD",
	4: "FBP_STANDARD_EVENT_PURCHASE",
	5: "FBP_STANDARD_EVENT_SCHEDULE",
	6: "FBP_STANDARD_EVENT_START_TRIAL",
	7: "FBP_STANDARD_EVENT_SUBMIT_APPLICATION",
	8: "FBP_STANDARD_EVENT_SUBSCRIBE",
	9: "FBP_STANDARD_EVENT_VIEW_CONTENT",
}

var FbPixelStandardEvent_value = map[string]int32{
	"FBP_STANDARD_EVENT_DO_NOT_USE":            0,
	"FBP_STANDARD_EVENT_PAGE_VIEW":             1,
	"FBP_STANDARD_EVENT_COMPLETE_REGISTRATION": 2,
	"FBP_STANDARD_EVENT_LEAD":                  3,
	"FBP_STANDARD_EVENT_PURCHASE":              4,
	"FBP_STANDARD_EVENT_SCHEDULE":              5,
	"FBP_STANDARD_EVENT_START_TRIAL":           6,
	"FBP_STANDARD_EVENT_SUBMIT_APPLICATION":    7,
	"FBP_STANDARD_EVENT_SUBSCRIBE":             8,
	"FBP_STANDARD_EVENT_VIEW_CONTENT":          9,
}

func (x FbPixelStandardEvent) String() string {
	return proto.EnumName(FbPixelStandardEvent_name, int32(x))
}

func (FbPixelStandardEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{0}
}

type GoogleAnalyticsDataCollectionPageEvent int32

const (
	GoogleAnalyticsDataCollectionPageEvent_GA_DC_EVENT_DO_NOT_USE GoogleAnalyticsDataCollectionPageEvent = 0
	// This is the default GA even tracking page visits.
	GoogleAnalyticsDataCollectionPageEvent_GA_DC_EVENT_PAGE_VIEW GoogleAnalyticsDataCollectionPageEvent = 1
	// When a user submits the data collection form.
	GoogleAnalyticsDataCollectionPageEvent_GA_DC_EVENT_SUBMIT_FORM GoogleAnalyticsDataCollectionPageEvent = 2
)

var GoogleAnalyticsDataCollectionPageEvent_name = map[int32]string{
	0: "GA_DC_EVENT_DO_NOT_USE",
	1: "GA_DC_EVENT_PAGE_VIEW",
	2: "GA_DC_EVENT_SUBMIT_FORM",
}

var GoogleAnalyticsDataCollectionPageEvent_value = map[string]int32{
	"GA_DC_EVENT_DO_NOT_USE":  0,
	"GA_DC_EVENT_PAGE_VIEW":   1,
	"GA_DC_EVENT_SUBMIT_FORM": 2,
}

func (x GoogleAnalyticsDataCollectionPageEvent) String() string {
	return proto.EnumName(GoogleAnalyticsDataCollectionPageEvent_name, int32(x))
}

func (GoogleAnalyticsDataCollectionPageEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{1}
}

type GoogleAnalyticsPassRenderPageEvent int32

const (
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_DO_NOT_USE GoogleAnalyticsPassRenderPageEvent = 0
	// This is the default GA event tracking page visits.
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_PAGE_VIEW GoogleAnalyticsPassRenderPageEvent = 1
	// When a user clicks the Add to Apple Wallet button.
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_ADD_TO_APPLE_WALLET_CLICK GoogleAnalyticsPassRenderPageEvent = 2
	// When a user clicks the view pass in Apple Wallet button.
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_VIEW_PASS_IN_APPLE_WALLET_CLICK GoogleAnalyticsPassRenderPageEvent = 3
	// When a user clicks the Save to Google Pay button.
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_SAVE_TO_GOOGLE_PAY_CLICK GoogleAnalyticsPassRenderPageEvent = 4
	// When a user switches the language.
	GoogleAnalyticsPassRenderPageEvent_GA_PP_EVENT_CHANGE_LANGUAGE GoogleAnalyticsPassRenderPageEvent = 5
)

var GoogleAnalyticsPassRenderPageEvent_name = map[int32]string{
	0: "GA_PP_EVENT_DO_NOT_USE",
	1: "GA_PP_EVENT_PAGE_VIEW",
	2: "GA_PP_EVENT_ADD_TO_APPLE_WALLET_CLICK",
	3: "GA_PP_EVENT_VIEW_PASS_IN_APPLE_WALLET_CLICK",
	4: "GA_PP_EVENT_SAVE_TO_GOOGLE_PAY_CLICK",
	5: "GA_PP_EVENT_CHANGE_LANGUAGE",
}

var GoogleAnalyticsPassRenderPageEvent_value = map[string]int32{
	"GA_PP_EVENT_DO_NOT_USE":                      0,
	"GA_PP_EVENT_PAGE_VIEW":                       1,
	"GA_PP_EVENT_ADD_TO_APPLE_WALLET_CLICK":       2,
	"GA_PP_EVENT_VIEW_PASS_IN_APPLE_WALLET_CLICK": 3,
	"GA_PP_EVENT_SAVE_TO_GOOGLE_PAY_CLICK":        4,
	"GA_PP_EVENT_CHANGE_LANGUAGE":                 5,
}

func (x GoogleAnalyticsPassRenderPageEvent) String() string {
	return proto.EnumName(GoogleAnalyticsPassRenderPageEvent_name, int32(x))
}

func (GoogleAnalyticsPassRenderPageEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{2}
}

type FacebookPixelSettings struct {
	// The Facebook Pixel ID of the account.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	PixelId string `protobuf:"bytes,1,opt,name=pixelId,proto3" json:"pixelId,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// List of standard events for the data collection page
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	DataCollectionPageStandardEvents []*StandardEvent `protobuf:"bytes,2,rep,name=dataCollectionPageStandardEvents,proto3" json:"dataCollectionPageStandardEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// List of custom events for the data collection page
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	DataCollectionPageCustomEvents []*CustomEvent `protobuf:"bytes,3,rep,name=dataCollectionPageCustomEvents,proto3" json:"dataCollectionPageCustomEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// List of standard events for the pass page
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	PassPageStandardEvents []*StandardEvent `protobuf:"bytes,4,rep,name=passPageStandardEvents,proto3" json:"passPageStandardEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// List of custom events for the pass page
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	PassPageCustomEvents []*CustomEvent `protobuf:"bytes,5,rep,name=passPageCustomEvents,proto3" json:"passPageCustomEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FacebookPixelSettings) Reset()         { *m = FacebookPixelSettings{} }
func (m *FacebookPixelSettings) String() string { return proto.CompactTextString(m) }
func (*FacebookPixelSettings) ProtoMessage()    {}
func (*FacebookPixelSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{0}
}

func (m *FacebookPixelSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FacebookPixelSettings.Unmarshal(m, b)
}
func (m *FacebookPixelSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FacebookPixelSettings.Marshal(b, m, deterministic)
}
func (m *FacebookPixelSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FacebookPixelSettings.Merge(m, src)
}
func (m *FacebookPixelSettings) XXX_Size() int {
	return xxx_messageInfo_FacebookPixelSettings.Size(m)
}
func (m *FacebookPixelSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_FacebookPixelSettings.DiscardUnknown(m)
}

var xxx_messageInfo_FacebookPixelSettings proto.InternalMessageInfo

func (m *FacebookPixelSettings) GetPixelId() string {
	if m != nil {
		return m.PixelId
	}
	return ""
}

func (m *FacebookPixelSettings) GetDataCollectionPageStandardEvents() []*StandardEvent {
	if m != nil {
		return m.DataCollectionPageStandardEvents
	}
	return nil
}

func (m *FacebookPixelSettings) GetDataCollectionPageCustomEvents() []*CustomEvent {
	if m != nil {
		return m.DataCollectionPageCustomEvents
	}
	return nil
}

func (m *FacebookPixelSettings) GetPassPageStandardEvents() []*StandardEvent {
	if m != nil {
		return m.PassPageStandardEvents
	}
	return nil
}

func (m *FacebookPixelSettings) GetPassPageCustomEvents() []*CustomEvent {
	if m != nil {
		return m.PassPageCustomEvents
	}
	return nil
}

type StandardEvent struct {
	// Standard event
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Event FbPixelStandardEvent `protobuf:"varint,1,opt,name=event,proto3,enum=io.FbPixelStandardEvent" json:"event,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Optional event properties
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Properties           *Properties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StandardEvent) Reset()         { *m = StandardEvent{} }
func (m *StandardEvent) String() string { return proto.CompactTextString(m) }
func (*StandardEvent) ProtoMessage()    {}
func (*StandardEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{1}
}

func (m *StandardEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardEvent.Unmarshal(m, b)
}
func (m *StandardEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardEvent.Marshal(b, m, deterministic)
}
func (m *StandardEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardEvent.Merge(m, src)
}
func (m *StandardEvent) XXX_Size() int {
	return xxx_messageInfo_StandardEvent.Size(m)
}
func (m *StandardEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StandardEvent proto.InternalMessageInfo

func (m *StandardEvent) GetEvent() FbPixelStandardEvent {
	if m != nil {
		return m.Event
	}
	return FbPixelStandardEvent_FBP_STANDARD_EVENT_DO_NOT_USE
}

func (m *StandardEvent) GetProperties() *Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CustomEvent struct {
	// Custom event name
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Optional event properties
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Properties           *Properties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CustomEvent) Reset()         { *m = CustomEvent{} }
func (m *CustomEvent) String() string { return proto.CompactTextString(m) }
func (*CustomEvent) ProtoMessage()    {}
func (*CustomEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{2}
}

func (m *CustomEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomEvent.Unmarshal(m, b)
}
func (m *CustomEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomEvent.Marshal(b, m, deterministic)
}
func (m *CustomEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomEvent.Merge(m, src)
}
func (m *CustomEvent) XXX_Size() int {
	return xxx_messageInfo_CustomEvent.Size(m)
}
func (m *CustomEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CustomEvent proto.InternalMessageInfo

func (m *CustomEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *CustomEvent) GetProperties() *Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Properties struct {
	// Category of the page/product.
	ContentCategory string `protobuf:"bytes,1,opt,name=contentCategory,proto3" json:"contentCategory,omitempty"`
	// Product IDs associated with the event, such as SKUs (e.g. ['ABC123', 'XYZ789']).
	ContentIds []string `protobuf:"bytes,2,rep,name=contentIds,proto3" json:"contentIds,omitempty"`
	// Name of the page/product.
	ContentName string `protobuf:"bytes,3,opt,name=contentName,proto3" json:"contentName,omitempty"`
	// Either product or product_group based on the content_ids or contents being passed. If the IDs being passed in content_ids or contents parameter are IDs of products then the value should be product. If product group IDs are being passed, then the value should be product_group.
	ContentType string `protobuf:"bytes,4,opt,name=contentType,proto3" json:"contentType,omitempty"`
	// An array of objects that contains the quantity and the International Article Number (EAN) when applicable, or other product or content identifier(s). id and quantity are the required fields. e.g. [{'id': 'ABC123', 'quantity': 2}, {'id': 'XYZ789', 'quantity': 2}].
	Contents []*ContentObject `protobuf:"bytes,5,rep,name=Contents,proto3" json:"Contents,omitempty"`
	// The currency for the value specified.
	// @tag: json:"currency,omitempty"
	Currency string `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	// Used with InitiateCheckout event. The number of items when checkout was initiated.
	NumItems uint32 `protobuf:"varint,7,opt,name=numItems,proto3" json:"numItems,omitempty"`
	// Predicted lifetime value of a subscriber as defined by the advertiser and expressed as an exact value.
	PredictedLtv float32 `protobuf:"fixed32,8,opt,name=predictedLtv,proto3" json:"predictedLtv,omitempty"`
	// Used with the Search event. The string entered by the user for the search.
	SearchString string `protobuf:"bytes,9,opt,name=searchString,proto3" json:"searchString,omitempty"`
	// Used with the CompleteRegistration event, to show the status of the registration.
	Status bool `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	// The value of a user performing this event to the business.
	Value                float32  `protobuf:"fixed32,11,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Properties) Reset()         { *m = Properties{} }
func (m *Properties) String() string { return proto.CompactTextString(m) }
func (*Properties) ProtoMessage()    {}
func (*Properties) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{3}
}

func (m *Properties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Properties.Unmarshal(m, b)
}
func (m *Properties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Properties.Marshal(b, m, deterministic)
}
func (m *Properties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Properties.Merge(m, src)
}
func (m *Properties) XXX_Size() int {
	return xxx_messageInfo_Properties.Size(m)
}
func (m *Properties) XXX_DiscardUnknown() {
	xxx_messageInfo_Properties.DiscardUnknown(m)
}

var xxx_messageInfo_Properties proto.InternalMessageInfo

func (m *Properties) GetContentCategory() string {
	if m != nil {
		return m.ContentCategory
	}
	return ""
}

func (m *Properties) GetContentIds() []string {
	if m != nil {
		return m.ContentIds
	}
	return nil
}

func (m *Properties) GetContentName() string {
	if m != nil {
		return m.ContentName
	}
	return ""
}

func (m *Properties) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Properties) GetContents() []*ContentObject {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Properties) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Properties) GetNumItems() uint32 {
	if m != nil {
		return m.NumItems
	}
	return 0
}

func (m *Properties) GetPredictedLtv() float32 {
	if m != nil {
		return m.PredictedLtv
	}
	return 0
}

func (m *Properties) GetSearchString() string {
	if m != nil {
		return m.SearchString
	}
	return ""
}

func (m *Properties) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Properties) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ContentObject struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity             uint32   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentObject) Reset()         { *m = ContentObject{} }
func (m *ContentObject) String() string { return proto.CompactTextString(m) }
func (*ContentObject) ProtoMessage()    {}
func (*ContentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{4}
}

func (m *ContentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentObject.Unmarshal(m, b)
}
func (m *ContentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentObject.Marshal(b, m, deterministic)
}
func (m *ContentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentObject.Merge(m, src)
}
func (m *ContentObject) XXX_Size() int {
	return xxx_messageInfo_ContentObject.Size(m)
}
func (m *ContentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentObject.DiscardUnknown(m)
}

var xxx_messageInfo_ContentObject proto.InternalMessageInfo

func (m *ContentObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContentObject) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type GoogleAnalyticsSettings struct {
	// list of Google Analytics ID's to track ('UA-XXXXX-Y', 'UA-XXXXX-B', etc).
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	TrackingIds []string `protobuf:"bytes,1,rep,name=trackingIds,proto3" json:"trackingIds,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Events to send to Google on the data collection page.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	DataCollectionPageEvents []GoogleAnalyticsDataCollectionPageEvent `protobuf:"varint,2,rep,packed,name=dataCollectionPageEvents,proto3,enum=io.GoogleAnalyticsDataCollectionPageEvent" json:"dataCollectionPageEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Events to send to Google on the pass render page.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	PassRenderPageEvents []GoogleAnalyticsPassRenderPageEvent `protobuf:"varint,3,rep,packed,name=passRenderPageEvents,proto3,enum=io.GoogleAnalyticsPassRenderPageEvent" json:"passRenderPageEvents,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GoogleAnalyticsSettings) Reset()         { *m = GoogleAnalyticsSettings{} }
func (m *GoogleAnalyticsSettings) String() string { return proto.CompactTextString(m) }
func (*GoogleAnalyticsSettings) ProtoMessage()    {}
func (*GoogleAnalyticsSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{5}
}

func (m *GoogleAnalyticsSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAnalyticsSettings.Unmarshal(m, b)
}
func (m *GoogleAnalyticsSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAnalyticsSettings.Marshal(b, m, deterministic)
}
func (m *GoogleAnalyticsSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAnalyticsSettings.Merge(m, src)
}
func (m *GoogleAnalyticsSettings) XXX_Size() int {
	return xxx_messageInfo_GoogleAnalyticsSettings.Size(m)
}
func (m *GoogleAnalyticsSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAnalyticsSettings.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAnalyticsSettings proto.InternalMessageInfo

func (m *GoogleAnalyticsSettings) GetTrackingIds() []string {
	if m != nil {
		return m.TrackingIds
	}
	return nil
}

func (m *GoogleAnalyticsSettings) GetDataCollectionPageEvents() []GoogleAnalyticsDataCollectionPageEvent {
	if m != nil {
		return m.DataCollectionPageEvents
	}
	return nil
}

func (m *GoogleAnalyticsSettings) GetPassRenderPageEvents() []GoogleAnalyticsPassRenderPageEvent {
	if m != nil {
		return m.PassRenderPageEvents
	}
	return nil
}

type TrackingSettings struct {
	// Facebook Pixel settings
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	FacebookPixelSettings *FacebookPixelSettings `protobuf:"bytes,1,opt,name=facebookPixelSettings,proto3" json:"facebookPixelSettings,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Google Analytics settings
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	GoogleAnalyticsSettings *GoogleAnalyticsSettings `protobuf:"bytes,2,opt,name=googleAnalyticsSettings,proto3" json:"googleAnalyticsSettings,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *TrackingSettings) Reset()         { *m = TrackingSettings{} }
func (m *TrackingSettings) String() string { return proto.CompactTextString(m) }
func (*TrackingSettings) ProtoMessage()    {}
func (*TrackingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc91891e247a3611, []int{6}
}

func (m *TrackingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingSettings.Unmarshal(m, b)
}
func (m *TrackingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingSettings.Marshal(b, m, deterministic)
}
func (m *TrackingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingSettings.Merge(m, src)
}
func (m *TrackingSettings) XXX_Size() int {
	return xxx_messageInfo_TrackingSettings.Size(m)
}
func (m *TrackingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingSettings proto.InternalMessageInfo

func (m *TrackingSettings) GetFacebookPixelSettings() *FacebookPixelSettings {
	if m != nil {
		return m.FacebookPixelSettings
	}
	return nil
}

func (m *TrackingSettings) GetGoogleAnalyticsSettings() *GoogleAnalyticsSettings {
	if m != nil {
		return m.GoogleAnalyticsSettings
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.FbPixelStandardEvent", FbPixelStandardEvent_name, FbPixelStandardEvent_value)
	proto.RegisterEnum("io.GoogleAnalyticsDataCollectionPageEvent", GoogleAnalyticsDataCollectionPageEvent_name, GoogleAnalyticsDataCollectionPageEvent_value)
	proto.RegisterEnum("io.GoogleAnalyticsPassRenderPageEvent", GoogleAnalyticsPassRenderPageEvent_name, GoogleAnalyticsPassRenderPageEvent_value)
	proto.RegisterType((*FacebookPixelSettings)(nil), "io.FacebookPixelSettings")
	proto.RegisterType((*StandardEvent)(nil), "io.StandardEvent")
	proto.RegisterType((*CustomEvent)(nil), "io.CustomEvent")
	proto.RegisterType((*Properties)(nil), "io.Properties")
	proto.RegisterType((*ContentObject)(nil), "io.ContentObject")
	proto.RegisterType((*GoogleAnalyticsSettings)(nil), "io.GoogleAnalyticsSettings")
	proto.RegisterType((*TrackingSettings)(nil), "io.TrackingSettings")
}

func init() {
	proto.RegisterFile("io/common/tracking.proto", fileDescriptor_fc91891e247a3611)
}

var fileDescriptor_fc91891e247a3611 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x4e, 0x7f, 0x4f, 0x68, 0x6b, 0x46, 0xfd, 0xf1, 0x6e, 0xa1, 0x1b, 0xc2, 0xee, 0x2a,
	0x5b, 0x20, 0x91, 0xca, 0x25, 0x12, 0x92, 0xeb, 0x4c, 0x53, 0xab, 0xae, 0x6d, 0x8d, 0x9d, 0x56,
	0xac, 0x84, 0x2c, 0xd7, 0x9e, 0xcd, 0x9a, 0x26, 0x9e, 0x60, 0x4f, 0x2a, 0x2a, 0x71, 0xc5, 0xe3,
	0xf0, 0x08, 0xbc, 0x18, 0xe2, 0x06, 0x90, 0x1d, 0xd7, 0x38, 0xdb, 0x89, 0x8a, 0xb8, 0xcb, 0xf9,
	0xce, 0x77, 0xbe, 0xf3, 0x37, 0xc7, 0x0a, 0xa8, 0x31, 0xeb, 0x85, 0x6c, 0x32, 0x61, 0x49, 0x8f,
	0xa7, 0x41, 0x78, 0x1b, 0x27, 0xa3, 0xee, 0x34, 0x65, 0x9c, 0x21, 0x39, 0x66, 0xed, 0x3f, 0x65,
	0xd8, 0x3b, 0x0b, 0x42, 0x7a, 0xc3, 0xd8, 0xad, 0x13, 0xff, 0x4c, 0xc7, 0x2e, 0xe5, 0x3c, 0x4e,
	0x46, 0x19, 0x52, 0x61, 0x7d, 0x9a, 0x03, 0x46, 0xa4, 0x4a, 0x2d, 0xa9, 0xb3, 0x49, 0x1e, 0x4c,
	0xf4, 0x03, 0xb4, 0xa2, 0x80, 0x07, 0x3a, 0x1b, 0x8f, 0x69, 0xc8, 0x63, 0x96, 0x38, 0xc1, 0x88,
	0xba, 0x3c, 0x48, 0xa2, 0x20, 0x8d, 0xf0, 0x1d, 0x4d, 0x78, 0xa6, 0xca, 0xad, 0x46, 0xa7, 0x79,
	0xf2, 0x49, 0x37, 0x66, 0xdd, 0x05, 0x0f, 0x79, 0x32, 0x14, 0x5d, 0xc3, 0xd1, 0x63, 0x8e, 0x3e,
	0xcb, 0x38, 0x9b, 0x94, 0xe2, 0x8d, 0x42, 0x7c, 0x27, 0x17, 0xaf, 0xe1, 0xe4, 0x89, 0x30, 0x64,
	0xc0, 0xfe, 0x34, 0xc8, 0x32, 0x41, 0xb5, 0x2b, 0xcb, 0xaa, 0x5d, 0x12, 0x80, 0x74, 0xd8, 0x7d,
	0xf0, 0x2c, 0x54, 0xb6, 0x2a, 0xae, 0x4c, 0x48, 0x6e, 0x33, 0xd8, 0x5a, 0x90, 0x45, 0x5d, 0x58,
	0xa5, 0xf9, 0x8f, 0x62, 0xe0, 0xdb, 0x27, 0x6a, 0x2e, 0x73, 0x76, 0x33, 0x5f, 0xcb, 0x42, 0x59,
	0x73, 0x1a, 0xea, 0x02, 0x4c, 0x53, 0x36, 0xa5, 0x29, 0x8f, 0x69, 0x3e, 0x72, 0xa9, 0xd3, 0x3c,
	0xd9, 0xce, 0x83, 0x9c, 0x0a, 0x25, 0x35, 0x46, 0xdb, 0x85, 0x66, 0xad, 0x00, 0xb4, 0x5b, 0x4f,
	0xb7, 0xf9, 0x7f, 0x45, 0xff, 0x90, 0x01, 0xfe, 0x75, 0xa1, 0x0e, 0xec, 0x84, 0x2c, 0xe1, 0x34,
	0xe1, 0x7a, 0xc0, 0xe9, 0x88, 0xa5, 0xf7, 0xa5, 0xfc, 0x87, 0x30, 0x3a, 0x02, 0x28, 0x21, 0x23,
	0x9a, 0x3f, 0x98, 0x4d, 0x52, 0x43, 0x50, 0x0b, 0x9a, 0xa5, 0x65, 0x05, 0x13, 0xaa, 0x36, 0x0a,
	0x95, 0x3a, 0x54, 0x63, 0x78, 0xf7, 0x53, 0xaa, 0xae, 0x2c, 0x30, 0x72, 0x08, 0x7d, 0x0d, 0x1b,
	0xfa, 0xdc, 0x7c, 0xd8, 0x4d, 0xb1, 0xe4, 0x12, 0xb3, 0x6f, 0x7e, 0xa4, 0x21, 0x27, 0x15, 0x05,
	0x3d, 0x87, 0x8d, 0x70, 0x96, 0xa6, 0x34, 0x09, 0xef, 0xd5, 0xb5, 0x42, 0xad, 0xb2, 0x73, 0x5f,
	0x32, 0x9b, 0x18, 0x9c, 0x4e, 0x32, 0x75, 0xbd, 0x25, 0x75, 0xb6, 0x48, 0x65, 0xa3, 0x36, 0x7c,
	0x3c, 0x4d, 0x69, 0x14, 0x87, 0x9c, 0x46, 0x26, 0xbf, 0x53, 0x37, 0x5a, 0x52, 0x47, 0x26, 0x0b,
	0x58, 0xce, 0xc9, 0x68, 0x90, 0x86, 0xef, 0x5d, 0x9e, 0xc6, 0xc9, 0x48, 0xdd, 0x2c, 0xf4, 0x17,
	0x30, 0xb4, 0x0f, 0x6b, 0x19, 0x0f, 0xf8, 0x2c, 0x53, 0xa1, 0x25, 0x75, 0x36, 0x48, 0x69, 0xe5,
	0x9b, 0xba, 0x0b, 0xc6, 0x33, 0xaa, 0x36, 0x0b, 0xe1, 0xb9, 0xd1, 0xfe, 0x16, 0xb6, 0x16, 0x1a,
	0x41, 0xdb, 0x20, 0xc7, 0x0f, 0xd7, 0x2a, 0xc7, 0x51, 0x5e, 0xf2, 0x4f, 0xb3, 0x20, 0xe1, 0x31,
	0xbf, 0x2f, 0x16, 0xb9, 0x45, 0x2a, 0xbb, 0xfd, 0xb7, 0x04, 0x07, 0x03, 0xc6, 0x46, 0x63, 0xaa,
	0x25, 0xc1, 0xf8, 0x9e, 0xc7, 0x61, 0x56, 0x9d, 0x7e, 0x0b, 0x9a, 0x0f, 0x9f, 0x8a, 0x7c, 0x35,
	0x52, 0xb1, 0x9a, 0x3a, 0x84, 0xde, 0x81, 0xfa, 0xf8, 0xd8, 0x6a, 0xa7, 0xbf, 0x7d, 0x72, 0x9c,
	0xcf, 0xf9, 0x83, 0x04, 0x7d, 0x71, 0x08, 0x59, 0xaa, 0x85, 0xde, 0xce, 0xef, 0x8c, 0xd0, 0x24,
	0xa2, 0x69, 0x2d, 0x47, 0xa3, 0xc8, 0xf1, 0x5a, 0x90, 0xc3, 0x79, 0x4c, 0x27, 0x42, 0x8d, 0xf6,
	0xef, 0x12, 0x28, 0x5e, 0xd9, 0x53, 0xd5, 0xba, 0x0d, 0x7b, 0xef, 0x44, 0x9f, 0xc3, 0x62, 0xaa,
	0xcd, 0x93, 0x67, 0xc5, 0x49, 0x8a, 0x08, 0x44, 0x1c, 0x87, 0x86, 0x70, 0x30, 0x12, 0x8f, 0xb9,
	0xbc, 0xad, 0x43, 0x41, 0x13, 0x95, 0xe8, 0xb2, 0xd8, 0xe3, 0xbf, 0x64, 0xd8, 0x15, 0x7d, 0x1a,
	0xd0, 0xe7, 0xf0, 0xd9, 0xd9, 0xa9, 0xe3, 0xbb, 0x9e, 0x66, 0xf5, 0x35, 0xd2, 0xf7, 0xf1, 0x15,
	0xb6, 0x3c, 0xbf, 0x6f, 0xfb, 0x96, 0xed, 0xf9, 0x43, 0x17, 0x2b, 0x1f, 0xa1, 0x16, 0x7c, 0x2a,
	0xa0, 0x38, 0xda, 0x00, 0xfb, 0x57, 0x06, 0xbe, 0x56, 0x24, 0xf4, 0x15, 0x74, 0x04, 0x0c, 0xdd,
	0xbe, 0x74, 0x4c, 0xec, 0x61, 0x9f, 0xe0, 0x81, 0xe1, 0x7a, 0x44, 0xf3, 0x0c, 0xdb, 0x52, 0x64,
	0x74, 0x08, 0x07, 0x02, 0xb6, 0x89, 0xb5, 0xbe, 0xd2, 0x40, 0x2f, 0xe0, 0x50, 0x94, 0x6c, 0x48,
	0xf4, 0x73, 0xcd, 0xc5, 0xca, 0xca, 0x12, 0x82, 0xab, 0x9f, 0xe3, 0xfe, 0xd0, 0xc4, 0xca, 0x2a,
	0x6a, 0xc3, 0x91, 0x88, 0xe0, 0x69, 0xc4, 0xf3, 0x3d, 0x62, 0x68, 0xa6, 0xb2, 0x86, 0xde, 0xc0,
	0x2b, 0x11, 0x67, 0x78, 0x7a, 0x69, 0x78, 0xbe, 0xe6, 0x38, 0xa6, 0xa1, 0xcf, 0xab, 0x5d, 0x5f,
	0xd2, 0xbd, 0x3b, 0x3c, 0x75, 0x75, 0x62, 0x9c, 0x62, 0x65, 0x03, 0x7d, 0x01, 0x2f, 0x04, 0x8c,
	0x7c, 0x34, 0xbe, 0x6e, 0x5b, 0x1e, 0xb6, 0x3c, 0x65, 0xf3, 0xf8, 0x17, 0x78, 0xfd, 0xdf, 0x5e,
	0x37, 0x7a, 0x0e, 0xfb, 0x03, 0xcd, 0xef, 0xeb, 0xa2, 0x55, 0x3c, 0x83, 0xbd, 0xba, 0xaf, 0xbe,
	0x83, 0x43, 0x38, 0xa8, 0xbb, 0xca, 0x5e, 0xce, 0x6c, 0x72, 0xa9, 0xc8, 0xc7, 0xbf, 0xca, 0xd0,
	0x7e, 0xfa, 0xe1, 0x97, 0xa9, 0x1d, 0x67, 0x79, 0xea, 0xca, 0x57, 0x4f, 0xfd, 0x06, 0x5e, 0xd5,
	0x5d, 0x5a, 0xbf, 0xef, 0x7b, 0x76, 0x31, 0x46, 0xec, 0x5f, 0x6b, 0xa6, 0x89, 0x3d, 0x5f, 0x37,
	0x0d, 0xfd, 0x42, 0x91, 0x51, 0x0f, 0xbe, 0xac, 0x53, 0x8b, 0x21, 0x39, 0x9a, 0xeb, 0xfa, 0x86,
	0x25, 0x0a, 0x68, 0xa0, 0x0e, 0xbc, 0xac, 0x07, 0xb8, 0xda, 0x15, 0xce, 0xc5, 0x07, 0xb6, 0x3d,
	0x30, 0xb1, 0xef, 0x68, 0xdf, 0x97, 0xcc, 0xe2, 0x61, 0xd4, 0x99, 0xfa, 0xb9, 0x66, 0x0d, 0xb0,
	0x6f, 0x6a, 0xd6, 0x60, 0xa8, 0x0d, 0xb0, 0xb2, 0x7a, 0xfa, 0x1d, 0xec, 0xc4, 0xac, 0x9b, 0xdf,
	0xf6, 0x6d, 0xcc, 0xbb, 0xce, 0x45, 0xcc, 0xde, 0xbe, 0xcc, 0x78, 0x90, 0xbd, 0xaf, 0xb0, 0x90,
	0x4d, 0x7a, 0x31, 0xeb, 0x4d, 0x58, 0x44, 0xc7, 0xbd, 0x2c, 0xba, 0xed, 0x8d, 0x58, 0x2f, 0x66,
	0xbf, 0xc9, 0x2b, 0xce, 0x85, 0xc1, 0x6e, 0xd6, 0x8a, 0xbf, 0x41, 0xdf, 0xfc, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x01, 0xa7, 0xb4, 0x27, 0x22, 0x09, 0x00, 0x00,
}