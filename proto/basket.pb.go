// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: basket.proto

package protopb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NewOrderByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *NewOrderByUser) Reset() {
	*x = NewOrderByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrderByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrderByUser) ProtoMessage() {}

func (x *NewOrderByUser) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrderByUser.ProtoReflect.Descriptor instead.
func (*NewOrderByUser) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{0}
}

func (x *NewOrderByUser) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type NewOrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *NewOrderCreated) Reset() {
	*x = NewOrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrderCreated) ProtoMessage() {}

func (x *NewOrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrderCreated.ProtoReflect.Descriptor instead.
func (*NewOrderCreated) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{1}
}

func (x *NewOrderCreated) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type OperationConfirmed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *OperationConfirmed) Reset() {
	*x = OperationConfirmed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationConfirmed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationConfirmed) ProtoMessage() {}

func (x *OperationConfirmed) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationConfirmed.ProtoReflect.Descriptor instead.
func (*OperationConfirmed) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{2}
}

func (x *OperationConfirmed) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int64  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemsToOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64   `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Items   []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AddItemsToOrder) Reset() {
	*x = AddItemsToOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemsToOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemsToOrder) ProtoMessage() {}

func (x *AddItemsToOrder) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemsToOrder.ProtoReflect.Descriptor instead.
func (*AddItemsToOrder) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{4}
}

func (x *AddItemsToOrder) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *AddItemsToOrder) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_basket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_basket_proto_rawDescGZIP(), []int{5}
}

func (x *Amount) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_basket_proto protoreflect.FileDescriptor

var file_basket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x2b, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x24, 0x0a,
	0x12, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x22, 0x40, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x4f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc5, 0x02, 0x0a, 0x06, 0x42, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4e,
	0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f,
	0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12,
	0x42, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x54, 0x6f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4e,
	0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40,
	0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x72, 0x76, 0x69, 0x74, 0x69, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basket_proto_rawDescOnce sync.Once
	file_basket_proto_rawDescData = file_basket_proto_rawDesc
)

func file_basket_proto_rawDescGZIP() []byte {
	file_basket_proto_rawDescOnce.Do(func() {
		file_basket_proto_rawDescData = protoimpl.X.CompressGZIP(file_basket_proto_rawDescData)
	})
	return file_basket_proto_rawDescData
}

var file_basket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_basket_proto_goTypes = []interface{}{
	(*NewOrderByUser)(nil),     // 0: basket.NewOrderByUser
	(*NewOrderCreated)(nil),    // 1: basket.NewOrderCreated
	(*OperationConfirmed)(nil), // 2: basket.OperationConfirmed
	(*Item)(nil),               // 3: basket.Item
	(*AddItemsToOrder)(nil),    // 4: basket.AddItemsToOrder
	(*Amount)(nil),             // 5: basket.Amount
}
var file_basket_proto_depIdxs = []int32{
	3, // 0: basket.AddItemsToOrder.items:type_name -> basket.Item
	0, // 1: basket.Basket.Create:input_type -> basket.NewOrderByUser
	1, // 2: basket.Basket.Checkout:input_type -> basket.NewOrderCreated
	4, // 3: basket.Basket.AddToBasket:input_type -> basket.AddItemsToOrder
	1, // 4: basket.Basket.GetTotalAmount:input_type -> basket.NewOrderCreated
	1, // 5: basket.Basket.RemoveAll:input_type -> basket.NewOrderCreated
	1, // 6: basket.Basket.Create:output_type -> basket.NewOrderCreated
	2, // 7: basket.Basket.Checkout:output_type -> basket.OperationConfirmed
	2, // 8: basket.Basket.AddToBasket:output_type -> basket.OperationConfirmed
	5, // 9: basket.Basket.GetTotalAmount:output_type -> basket.Amount
	2, // 10: basket.Basket.RemoveAll:output_type -> basket.OperationConfirmed
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basket_proto_init() }
func file_basket_proto_init() {
	if File_basket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrderByUser); i {
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
		file_basket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrderCreated); i {
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
		file_basket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationConfirmed); i {
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
		file_basket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_basket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemsToOrder); i {
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
		file_basket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
			RawDescriptor: file_basket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basket_proto_goTypes,
		DependencyIndexes: file_basket_proto_depIdxs,
		MessageInfos:      file_basket_proto_msgTypes,
	}.Build()
	File_basket_proto = out.File
	file_basket_proto_rawDesc = nil
	file_basket_proto_goTypes = nil
	file_basket_proto_depIdxs = nil
}
