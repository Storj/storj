// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// PieceAction is an enumeration of all possible executed actions on storage node
type PieceAction int32

const (
	PieceAction_INVALID    PieceAction = 0
	PieceAction_PUT        PieceAction = 1
	PieceAction_GET        PieceAction = 2
	PieceAction_GET_AUDIT  PieceAction = 3
	PieceAction_GET_REPAIR PieceAction = 4
	PieceAction_PUT_REPAIR PieceAction = 5
	PieceAction_DELETE     PieceAction = 6
)

var PieceAction_name = map[int32]string{
	0: "INVALID",
	1: "PUT",
	2: "GET",
	3: "GET_AUDIT",
	4: "GET_REPAIR",
	5: "PUT_REPAIR",
	6: "DELETE",
}

var PieceAction_value = map[string]int32{
	"INVALID":    0,
	"PUT":        1,
	"GET":        2,
	"GET_AUDIT":  3,
	"GET_REPAIR": 4,
	"PUT_REPAIR": 5,
	"DELETE":     6,
}

func (x PieceAction) String() string {
	return proto.EnumName(PieceAction_name, int32(x))
}

func (PieceAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{0}
}

type SettlementResponse_Status int32

const (
	SettlementResponse_INVALID  SettlementResponse_Status = 0
	SettlementResponse_OK       SettlementResponse_Status = 1
	SettlementResponse_REJECTED SettlementResponse_Status = 2
)

var SettlementResponse_Status_name = map[int32]string{
	0: "INVALID",
	1: "OK",
	2: "REJECTED",
}

var SettlementResponse_Status_value = map[string]int32{
	"INVALID":  0,
	"OK":       1,
	"REJECTED": 2,
}

func (x SettlementResponse_Status) String() string {
	return proto.EnumName(SettlementResponse_Status_name, int32(x))
}

func (SettlementResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{4, 0}
}

// OrderLimit2 is provided by satellite to execute specific action on storage node within some limits
type OrderLimit2 struct {
	// unique serial to avoid replay attacks
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// satellite who issued this order limit allowing orderer to do the specified action
	SatelliteId NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// uplink who requested or whom behalf the order limit to do an action
	UplinkId NodeID `protobuf:"bytes,3,opt,name=uplink_id,json=uplinkId,proto3,customtype=NodeID" json:"uplink_id"`
	// storage node who can reclaim the order limit specified by serial
	StorageNodeId NodeID `protobuf:"bytes,4,opt,name=storage_node_id,json=storageNodeId,proto3,customtype=NodeID" json:"storage_node_id"`
	// piece which is allowed to be touched
	PieceId PieceID `protobuf:"bytes,5,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// limit in bytes how much can be changed
	Limit                int64                `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Action               PieceAction          `protobuf:"varint,7,opt,name=action,proto3,enum=orders.PieceAction" json:"action,omitempty"`
	PieceExpiration      *timestamp.Timestamp `protobuf:"bytes,8,opt,name=piece_expiration,json=pieceExpiration,proto3" json:"piece_expiration,omitempty"`
	OrderExpiration      *timestamp.Timestamp `protobuf:"bytes,9,opt,name=order_expiration,json=orderExpiration,proto3" json:"order_expiration,omitempty"`
	SatelliteSignature   []byte               `protobuf:"bytes,10,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderLimit2) Reset()         { *m = OrderLimit2{} }
func (m *OrderLimit2) String() string { return proto.CompactTextString(m) }
func (*OrderLimit2) ProtoMessage()    {}
func (*OrderLimit2) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{0}
}
func (m *OrderLimit2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimit2.Unmarshal(m, b)
}
func (m *OrderLimit2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimit2.Marshal(b, m, deterministic)
}
func (m *OrderLimit2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimit2.Merge(m, src)
}
func (m *OrderLimit2) XXX_Size() int {
	return xxx_messageInfo_OrderLimit2.Size(m)
}
func (m *OrderLimit2) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimit2.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimit2 proto.InternalMessageInfo

func (m *OrderLimit2) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderLimit2) GetAction() PieceAction {
	if m != nil {
		return m.Action
	}
	return PieceAction_INVALID
}

func (m *OrderLimit2) GetPieceExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.PieceExpiration
	}
	return nil
}

func (m *OrderLimit2) GetOrderExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.OrderExpiration
	}
	return nil
}

func (m *OrderLimit2) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

// Order2 is a one step of fullfilling Amount number of bytes from an OrderLimit2 with SerialNumber
type Order2 struct {
	// serial of the order limit that was signed
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// amount to be signed for
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// signature
	UplinkSignature      []byte   `protobuf:"bytes,3,opt,name=uplink_signature,json=uplinkSignature,proto3" json:"uplink_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order2) Reset()         { *m = Order2{} }
func (m *Order2) String() string { return proto.CompactTextString(m) }
func (*Order2) ProtoMessage()    {}
func (*Order2) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{1}
}
func (m *Order2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order2.Unmarshal(m, b)
}
func (m *Order2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order2.Marshal(b, m, deterministic)
}
func (m *Order2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order2.Merge(m, src)
}
func (m *Order2) XXX_Size() int {
	return xxx_messageInfo_Order2.Size(m)
}
func (m *Order2) XXX_DiscardUnknown() {
	xxx_messageInfo_Order2.DiscardUnknown(m)
}

var xxx_messageInfo_Order2 proto.InternalMessageInfo

func (m *Order2) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order2) GetUplinkSignature() []byte {
	if m != nil {
		return m.UplinkSignature
	}
	return nil
}

type PieceHash struct {
	// piece id
	PieceId PieceID `protobuf:"bytes,1,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// hash of the piece that was/is uploaded
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// signature either satellite or storage node
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceHash) Reset()         { *m = PieceHash{} }
func (m *PieceHash) String() string { return proto.CompactTextString(m) }
func (*PieceHash) ProtoMessage()    {}
func (*PieceHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{2}
}
func (m *PieceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceHash.Unmarshal(m, b)
}
func (m *PieceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceHash.Marshal(b, m, deterministic)
}
func (m *PieceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceHash.Merge(m, src)
}
func (m *PieceHash) XXX_Size() int {
	return xxx_messageInfo_PieceHash.Size(m)
}
func (m *PieceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceHash.DiscardUnknown(m)
}

var xxx_messageInfo_PieceHash proto.InternalMessageInfo

func (m *PieceHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PieceHash) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SettlementRequest struct {
	Limit                *OrderLimit2 `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Order                *Order2      `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SettlementRequest) Reset()         { *m = SettlementRequest{} }
func (m *SettlementRequest) String() string { return proto.CompactTextString(m) }
func (*SettlementRequest) ProtoMessage()    {}
func (*SettlementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{3}
}
func (m *SettlementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlementRequest.Unmarshal(m, b)
}
func (m *SettlementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlementRequest.Marshal(b, m, deterministic)
}
func (m *SettlementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementRequest.Merge(m, src)
}
func (m *SettlementRequest) XXX_Size() int {
	return xxx_messageInfo_SettlementRequest.Size(m)
}
func (m *SettlementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementRequest proto.InternalMessageInfo

func (m *SettlementRequest) GetLimit() *OrderLimit2 {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *SettlementRequest) GetOrder() *Order2 {
	if m != nil {
		return m.Order
	}
	return nil
}

type SettlementResponse struct {
	SerialNumber         SerialNumber              `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	Status               SettlementResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=orders.SettlementResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SettlementResponse) Reset()         { *m = SettlementResponse{} }
func (m *SettlementResponse) String() string { return proto.CompactTextString(m) }
func (*SettlementResponse) ProtoMessage()    {}
func (*SettlementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{4}
}
func (m *SettlementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlementResponse.Unmarshal(m, b)
}
func (m *SettlementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlementResponse.Marshal(b, m, deterministic)
}
func (m *SettlementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementResponse.Merge(m, src)
}
func (m *SettlementResponse) XXX_Size() int {
	return xxx_messageInfo_SettlementResponse.Size(m)
}
func (m *SettlementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementResponse proto.InternalMessageInfo

func (m *SettlementResponse) GetStatus() SettlementResponse_Status {
	if m != nil {
		return m.Status
	}
	return SettlementResponse_INVALID
}

func init() {
	proto.RegisterEnum("orders.PieceAction", PieceAction_name, PieceAction_value)
	proto.RegisterEnum("orders.SettlementResponse_Status", SettlementResponse_Status_name, SettlementResponse_Status_value)
	proto.RegisterType((*OrderLimit2)(nil), "orders.OrderLimit2")
	proto.RegisterType((*Order2)(nil), "orders.Order2")
	proto.RegisterType((*PieceHash)(nil), "orders.PieceHash")
	proto.RegisterType((*SettlementRequest)(nil), "orders.SettlementRequest")
	proto.RegisterType((*SettlementResponse)(nil), "orders.SettlementResponse")
}

func init() { proto.RegisterFile("orders.proto", fileDescriptor_e0f5d4cf0fc9e41b) }

var fileDescriptor_e0f5d4cf0fc9e41b = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xee, 0xe6, 0xc7, 0x49, 0x26, 0x69, 0x62, 0xb6, 0x15, 0x0a, 0x11, 0x52, 0x43, 0xc4, 0x21,
	0xb4, 0x52, 0x0a, 0x46, 0x42, 0xea, 0x31, 0x25, 0x56, 0x31, 0x54, 0x6d, 0xb4, 0x71, 0x39, 0x70,
	0x89, 0x9c, 0x7a, 0x71, 0x2d, 0x1c, 0xaf, 0xf1, 0xae, 0x25, 0x9e, 0x80, 0x23, 0xaf, 0xc3, 0x2b,
	0xf0, 0x0c, 0x1c, 0xfa, 0x2c, 0xc8, 0x63, 0x27, 0x71, 0x21, 0xa8, 0x87, 0xde, 0xfc, 0xcd, 0x7c,
	0xdf, 0x7e, 0xbb, 0xf3, 0x8d, 0xa1, 0x25, 0x62, 0x97, 0xc7, 0x72, 0x14, 0xc5, 0x42, 0x09, 0xaa,
	0x65, 0xa8, 0x07, 0x9e, 0xf0, 0x44, 0x56, 0xeb, 0x1d, 0x78, 0x42, 0x78, 0x01, 0x3f, 0x46, 0xb4,
	0x48, 0x3e, 0x1f, 0x2b, 0x7f, 0xc9, 0xa5, 0x72, 0x96, 0x51, 0x46, 0x18, 0xfc, 0xa8, 0x40, 0xf3,
	0x32, 0xd5, 0x9d, 0xfb, 0x4b, 0x5f, 0x19, 0xf4, 0x04, 0x76, 0x25, 0x8f, 0x7d, 0x27, 0x98, 0x87,
	0xc9, 0x72, 0xc1, 0xe3, 0x2e, 0xe9, 0x93, 0x61, 0xeb, 0x74, 0xff, 0xd7, 0xed, 0xc1, 0xce, 0xef,
	0xdb, 0x83, 0xd6, 0x0c, 0x9b, 0x17, 0xd8, 0x63, 0x2d, 0x59, 0x40, 0xf4, 0x15, 0xb4, 0xa4, 0xa3,
	0x78, 0x10, 0xf8, 0x8a, 0xcf, 0x7d, 0xb7, 0x5b, 0x42, 0x65, 0x3b, 0x57, 0x6a, 0x17, 0xc2, 0xe5,
	0xd6, 0x84, 0x35, 0xd7, 0x1c, 0xcb, 0xa5, 0x47, 0xd0, 0x48, 0xa2, 0xc0, 0x0f, 0xbf, 0xa4, 0xfc,
	0xf2, 0x56, 0x7e, 0x3d, 0x23, 0x58, 0x2e, 0x7d, 0x03, 0x1d, 0xa9, 0x44, 0xec, 0x78, 0x7c, 0x1e,
	0x0a, 0x17, 0x2d, 0x2a, 0x5b, 0x25, 0xbb, 0x39, 0x0d, 0xa1, 0x4b, 0x0f, 0xa1, 0x1e, 0xf9, 0xfc,
	0x1a, 0x05, 0x55, 0x14, 0x74, 0x72, 0x41, 0x6d, 0x9a, 0xd6, 0xad, 0x09, 0xab, 0x21, 0xc1, 0x72,
	0xe9, 0x3e, 0x54, 0x83, 0x74, 0x10, 0x5d, 0xad, 0x4f, 0x86, 0x65, 0x96, 0x01, 0x7a, 0x04, 0x9a,
	0x73, 0xad, 0x7c, 0x11, 0x76, 0x6b, 0x7d, 0x32, 0x6c, 0x1b, 0x7b, 0xa3, 0x7c, 0xf0, 0xa8, 0x1f,
	0x63, 0x8b, 0xe5, 0x14, 0x6a, 0x82, 0x9e, 0xd9, 0xf1, 0x6f, 0x91, 0x1f, 0x3b, 0x28, 0xab, 0xf7,
	0xc9, 0xb0, 0x69, 0xf4, 0x46, 0x59, 0x1a, 0xa3, 0x55, 0x1a, 0x23, 0x7b, 0x95, 0x06, 0xeb, 0xa0,
	0xc6, 0x5c, 0x4b, 0xd2, 0x63, 0xd0, 0xa4, 0x78, 0x4c, 0xe3, 0xfe, 0x63, 0x50, 0x53, 0x38, 0xe6,
	0x18, 0xf6, 0x36, 0xa1, 0x48, 0xdf, 0x0b, 0x1d, 0x95, 0xc4, 0xbc, 0x0b, 0xe9, 0x1c, 0x18, 0x5d,
	0xb7, 0x66, 0xab, 0xce, 0xe0, 0x3b, 0x01, 0x0d, 0x17, 0xe2, 0x41, 0xbb, 0xf0, 0x18, 0x34, 0x67,
	0x29, 0x92, 0x50, 0xe1, 0x16, 0x94, 0x59, 0x8e, 0xe8, 0x0b, 0xd0, 0xf3, 0xc0, 0x37, 0x77, 0xc1,
	0xdc, 0x59, 0x27, 0xab, 0x6f, 0x2e, 0xe2, 0x43, 0x03, 0xc7, 0xfb, 0xce, 0x91, 0x37, 0x77, 0x32,
	0x24, 0xf7, 0x64, 0x48, 0xa1, 0x72, 0xe3, 0xc8, 0x9b, 0x6c, 0xff, 0x18, 0x7e, 0xd3, 0xa7, 0xd0,
	0xf8, 0xdb, 0x70, 0x53, 0x18, 0xb8, 0xf0, 0x68, 0xc6, 0x95, 0x0a, 0xf8, 0x92, 0x87, 0x8a, 0xf1,
	0xaf, 0x09, 0x97, 0xe9, 0x55, 0xf3, 0x55, 0x20, 0x38, 0xf5, 0x75, 0xe6, 0x85, 0xbf, 0x65, 0xb5,
	0x1f, 0xcf, 0xa1, 0x8a, 0x4d, 0xb4, 0x6c, 0x1a, 0xed, 0x3b, 0x54, 0x83, 0x65, 0xcd, 0xc1, 0x4f,
	0x02, 0xb4, 0x68, 0x23, 0x23, 0x11, 0x4a, 0xfe, 0x90, 0x29, 0x9f, 0x80, 0x26, 0x95, 0xa3, 0x12,
	0x89, 0xc6, 0x6d, 0xe3, 0xd9, 0xca, 0xf8, 0x5f, 0x9b, 0xd1, 0x0c, 0x89, 0x2c, 0x17, 0x0c, 0x8e,
	0x40, 0xcb, 0x2a, 0xb4, 0x09, 0x35, 0xeb, 0xe2, 0xe3, 0xf8, 0xdc, 0x9a, 0xe8, 0x3b, 0x54, 0x83,
	0xd2, 0xe5, 0x07, 0x9d, 0xd0, 0x16, 0xd4, 0x99, 0xf9, 0xde, 0x7c, 0x6b, 0x9b, 0x13, 0xbd, 0x74,
	0xe8, 0x41, 0xb3, 0xb0, 0xe9, 0x77, 0x15, 0x35, 0x28, 0x4f, 0xaf, 0x6c, 0x9d, 0xa4, 0x1f, 0x67,
	0xa6, 0xad, 0x97, 0xe8, 0x2e, 0x34, 0xce, 0x4c, 0x7b, 0x3e, 0xbe, 0x9a, 0x58, 0xb6, 0x5e, 0xa6,
	0x6d, 0x80, 0x14, 0x32, 0x73, 0x3a, 0xb6, 0x98, 0x5e, 0x49, 0xf1, 0xf4, 0x6a, 0x8d, 0xab, 0x14,
	0x40, 0x9b, 0x98, 0xe7, 0xa6, 0x6d, 0xea, 0x9a, 0x31, 0xcb, 0x77, 0x4f, 0x52, 0x0b, 0x60, 0xf3,
	0x08, 0xfa, 0x64, 0xdb, 0xc3, 0x30, 0xa6, 0x5e, 0xef, 0xff, 0x6f, 0x1e, 0xec, 0x0c, 0xc9, 0x4b,
	0x72, 0x5a, 0xf9, 0x54, 0x8a, 0x16, 0x0b, 0x0d, 0xff, 0x96, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd4, 0xc6, 0x2b, 0xc2, 0x34, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrdersClient interface {
	Settlement(ctx context.Context, opts ...grpc.CallOption) (Orders_SettlementClient, error)
}

type ordersClient struct {
	cc *grpc.ClientConn
}

func NewOrdersClient(cc *grpc.ClientConn) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) Settlement(ctx context.Context, opts ...grpc.CallOption) (Orders_SettlementClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Orders_serviceDesc.Streams[0], "/orders.Orders/Settlement", opts...)
	if err != nil {
		return nil, err
	}
	x := &ordersSettlementClient{stream}
	return x, nil
}

type Orders_SettlementClient interface {
	Send(*SettlementRequest) error
	Recv() (*SettlementResponse, error)
	grpc.ClientStream
}

type ordersSettlementClient struct {
	grpc.ClientStream
}

func (x *ordersSettlementClient) Send(m *SettlementRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ordersSettlementClient) Recv() (*SettlementResponse, error) {
	m := new(SettlementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrdersServer is the server API for Orders service.
type OrdersServer interface {
	Settlement(Orders_SettlementServer) error
}

func RegisterOrdersServer(s *grpc.Server, srv OrdersServer) {
	s.RegisterService(&_Orders_serviceDesc, srv)
}

func _Orders_Settlement_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrdersServer).Settlement(&ordersSettlementServer{stream})
}

type Orders_SettlementServer interface {
	Send(*SettlementResponse) error
	Recv() (*SettlementRequest, error)
	grpc.ServerStream
}

type ordersSettlementServer struct {
	grpc.ServerStream
}

func (x *ordersSettlementServer) Send(m *SettlementResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ordersSettlementServer) Recv() (*SettlementRequest, error) {
	m := new(SettlementRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Orders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Settlement",
			Handler:       _Orders_Settlement_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orders.proto",
}
