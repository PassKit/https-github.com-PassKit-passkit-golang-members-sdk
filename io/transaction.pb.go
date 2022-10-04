//*
// Transaction
//
// Transaction renders transaction history on individual pass.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: io/common/transaction.proto

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

// Some point based loyalty programs will pass on transaction information (can be used for segmenting, additional reporting, and/or showing transactions on back of the pass).
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference ID is the ID used in the system where the transaction is coming from. Needs to be unique within the program.
	ReferenceId string `protobuf:"bytes,1,opt,name=referenceId,proto3" json:"referenceId,omitempty"`
	// The total amount of all order items. Based on POS setting, the totalPrice can already include the tax amount
	TotalPrice float32 `protobuf:"fixed32,2,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	// List of order items in the  transaction
	OrderItems []*OrderItem `protobuf:"bytes,3,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
	// The total discount amount
	Discount float32 `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount,omitempty"`
	// List of discount items
	DiscountItems []*DiscountItem `protobuf:"bytes,5,rep,name=discountItems,proto3" json:"discountItems,omitempty"`
	// The service charge amount
	ServiceCharge float32 `protobuf:"fixed32,6,opt,name=serviceCharge,proto3" json:"serviceCharge,omitempty"`
	// The tax amount.
	TotalTax float32 `protobuf:"fixed32,7,opt,name=totalTax,proto3" json:"totalTax,omitempty"`
	// The final price of the transaction: (total price + service charge + (totalTax: if POS settings set to tax exclusive)) - discount
	FinalPrice float32 `protobuf:"fixed32,8,opt,name=finalPrice,proto3" json:"finalPrice,omitempty"`
	// Rounding difference
	RoundingDifference float32 `protobuf:"fixed32,9,opt,name=roundingDifference,proto3" json:"roundingDifference,omitempty"`
	// Indicates if this transaction is a refund
	IsRefunded bool `protobuf:"varint,10,opt,name=isRefunded,proto3" json:"isRefunded,omitempty"`
	// The transaction timestamp
	Timestamp *Date `protobuf:"bytes,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The transaction currency
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// GPS details of check in
	Location *GPSLocation `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	// 17 is reserved for transaction source (online, or restaurant code)
	TransactionSource string `protobuf:"bytes,14,opt,name=transactionSource,proto3" json:"transactionSource,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_io_common_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Transaction) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Transaction) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *Transaction) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Transaction) GetDiscountItems() []*DiscountItem {
	if x != nil {
		return x.DiscountItems
	}
	return nil
}

func (x *Transaction) GetServiceCharge() float32 {
	if x != nil {
		return x.ServiceCharge
	}
	return 0
}

func (x *Transaction) GetTotalTax() float32 {
	if x != nil {
		return x.TotalTax
	}
	return 0
}

func (x *Transaction) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

func (x *Transaction) GetRoundingDifference() float32 {
	if x != nil {
		return x.RoundingDifference
	}
	return 0
}

func (x *Transaction) GetIsRefunded() bool {
	if x != nil {
		return x.IsRefunded
	}
	return false
}

func (x *Transaction) GetTimestamp() *Date {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetLocation() *GPSLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Transaction) GetTransactionSource() string {
	if x != nil {
		return x.TransactionSource
	}
	return ""
}

type DiscountItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Discount code
	DiscountCode string `protobuf:"bytes,1,opt,name=discountCode,proto3" json:"discountCode,omitempty"`
	// Specific voucher code
	VoucherCode string `protobuf:"bytes,2,opt,name=voucherCode,proto3" json:"voucherCode,omitempty"`
	// The discount amount
	Amount float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// The discount item name
	ItemName string `protobuf:"bytes,5,opt,name=itemName,proto3" json:"itemName,omitempty"`
}

func (x *DiscountItem) Reset() {
	*x = DiscountItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountItem) ProtoMessage() {}

func (x *DiscountItem) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountItem.ProtoReflect.Descriptor instead.
func (*DiscountItem) Descriptor() ([]byte, []int) {
	return file_io_common_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountItem) GetDiscountCode() string {
	if x != nil {
		return x.DiscountCode
	}
	return ""
}

func (x *DiscountItem) GetVoucherCode() string {
	if x != nil {
		return x.VoucherCode
	}
	return ""
}

func (x *DiscountItem) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DiscountItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The item price
	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Tax on the item
	Tax float32 `protobuf:"fixed32,2,opt,name=tax,proto3" json:"tax,omitempty"`
	// The item name
	ItemName string `protobuf:"bytes,3,opt,name=itemName,proto3" json:"itemName,omitempty"`
	// The quantity of items ordered
	Quantity int32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The SKU as used in the system generating the transaction
	Sku string `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_common_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_io_common_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_io_common_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderItem) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *OrderItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

var File_io_common_transaction_proto protoreflect.FileDescriptor

var file_io_common_transaction_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x69,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x69, 0x6d, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x05, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x54, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x54, 0x61, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x12, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6f, 0x2e, 0x47,
	0x50, 0x53, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x3a, 0xa3, 0x01, 0x92, 0x41, 0x9f, 0x01, 0x0a, 0x9c, 0x01, 0x2a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x77, 0x61, 0x6e, 0x74,
	0x20, 0x74, 0x6f, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x20, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x6e, 0x20,
	0x62, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x20, 0x63, 0x61, 0x72, 0x64, 0x2e, 0xd2, 0x01, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0xd2, 0x01, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0xd2, 0x01, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x7f, 0x0a, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x42, 0x47, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x24,
	0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x6f, 0x2f, 0x69, 0x6f, 0xaa, 0x02, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_common_transaction_proto_rawDescOnce sync.Once
	file_io_common_transaction_proto_rawDescData = file_io_common_transaction_proto_rawDesc
)

func file_io_common_transaction_proto_rawDescGZIP() []byte {
	file_io_common_transaction_proto_rawDescOnce.Do(func() {
		file_io_common_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_common_transaction_proto_rawDescData)
	})
	return file_io_common_transaction_proto_rawDescData
}

var file_io_common_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_io_common_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),  // 0: io.Transaction
	(*DiscountItem)(nil), // 1: io.DiscountItem
	(*OrderItem)(nil),    // 2: io.OrderItem
	(*Date)(nil),         // 3: io.Date
	(*GPSLocation)(nil),  // 4: io.GPSLocation
}
var file_io_common_transaction_proto_depIdxs = []int32{
	2, // 0: io.Transaction.orderItems:type_name -> io.OrderItem
	1, // 1: io.Transaction.discountItems:type_name -> io.DiscountItem
	3, // 2: io.Transaction.timestamp:type_name -> io.Date
	4, // 3: io.Transaction.location:type_name -> io.GPSLocation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_io_common_transaction_proto_init() }
func file_io_common_transaction_proto_init() {
	if File_io_common_transaction_proto != nil {
		return
	}
	file_io_common_common_objects_proto_init()
	file_io_common_proximity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_io_common_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_io_common_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountItem); i {
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
		file_io_common_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
			RawDescriptor: file_io_common_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_common_transaction_proto_goTypes,
		DependencyIndexes: file_io_common_transaction_proto_depIdxs,
		MessageInfos:      file_io_common_transaction_proto_msgTypes,
	}.Build()
	File_io_common_transaction_proto = out.File
	file_io_common_transaction_proto_rawDesc = nil
	file_io_common_transaction_proto_goTypes = nil
	file_io_common_transaction_proto_depIdxs = nil
}
