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
	SettlementResponse_ACCEPTED SettlementResponse_Status = 1
	SettlementResponse_REJECTED SettlementResponse_Status = 2
)

var SettlementResponse_Status_name = map[int32]string{
	0: "INVALID",
	1: "ACCEPTED",
	2: "REJECTED",
}

var SettlementResponse_Status_value = map[string]int32{
	"INVALID":  0,
	"ACCEPTED": 1,
	"REJECTED": 2,
}

func (x SettlementResponse_Status) String() string {
	return proto.EnumName(SettlementResponse_Status_name, int32(x))
}

func (SettlementResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{4, 0}
}

// OrderLimit is provided by satellite to execute specific action on storage node within some limits
type OrderLimit struct {
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
	Limit              int64                `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Action             PieceAction          `protobuf:"varint,7,opt,name=action,proto3,enum=orders.PieceAction" json:"action,omitempty"`
	PieceExpiration    *timestamp.Timestamp `protobuf:"bytes,8,opt,name=piece_expiration,json=pieceExpiration,proto3" json:"piece_expiration,omitempty"`
	OrderExpiration    *timestamp.Timestamp `protobuf:"bytes,9,opt,name=order_expiration,json=orderExpiration,proto3" json:"order_expiration,omitempty"`
	SatelliteSignature []byte               `protobuf:"bytes,10,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	// satellites aren't necessarily discoverable in kademlia. this allows
	// a storage node to find a satellite and handshake with it to get its key.
	SatelliteAddress     *NodeAddress `protobuf:"bytes,11,opt,name=satellite_address,json=satelliteAddress,proto3" json:"satellite_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrderLimit) Reset()         { *m = OrderLimit{} }
func (m *OrderLimit) String() string { return proto.CompactTextString(m) }
func (*OrderLimit) ProtoMessage()    {}
func (*OrderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{0}
}
func (m *OrderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimit.Unmarshal(m, b)
}
func (m *OrderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimit.Marshal(b, m, deterministic)
}
func (m *OrderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimit.Merge(m, src)
}
func (m *OrderLimit) XXX_Size() int {
	return xxx_messageInfo_OrderLimit.Size(m)
}
func (m *OrderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimit proto.InternalMessageInfo

func (m *OrderLimit) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderLimit) GetAction() PieceAction {
	if m != nil {
		return m.Action
	}
	return PieceAction_INVALID
}

func (m *OrderLimit) GetPieceExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.PieceExpiration
	}
	return nil
}

func (m *OrderLimit) GetOrderExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.OrderExpiration
	}
	return nil
}

func (m *OrderLimit) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

func (m *OrderLimit) GetSatelliteAddress() *NodeAddress {
	if m != nil {
		return m.SatelliteAddress
	}
	return nil
}

// Order is a one step of fullfilling Amount number of bytes from an OrderLimit with SerialNumber
type Order struct {
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

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f5d4cf0fc9e41b, []int{1}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetUplinkSignature() []byte {
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
	Limit                *OrderLimit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Order                *Order      `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
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

func (m *SettlementRequest) GetLimit() *OrderLimit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *SettlementRequest) GetOrder() *Order {
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
	proto.RegisterType((*OrderLimit)(nil), "orders.OrderLimit")
	proto.RegisterType((*Order)(nil), "orders.Order")
	proto.RegisterType((*PieceHash)(nil), "orders.PieceHash")
	proto.RegisterType((*SettlementRequest)(nil), "orders.SettlementRequest")
	proto.RegisterType((*SettlementResponse)(nil), "orders.SettlementResponse")
}

func init() { proto.RegisterFile("orders.proto", fileDescriptor_e0f5d4cf0fc9e41b) }

var fileDescriptor_e0f5d4cf0fc9e41b = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xf3, 0xe3, 0x24, 0x93, 0x3f, 0x77, 0x5b, 0xa1, 0x10, 0x21, 0x35, 0x98, 0x4b, 0x68,
	0x25, 0x97, 0x06, 0x09, 0xa9, 0x17, 0xa4, 0xb4, 0xb1, 0x8a, 0x51, 0x55, 0xa2, 0x8d, 0xcb, 0x81,
	0x4b, 0xe4, 0xd4, 0x8b, 0x6b, 0xe1, 0xd8, 0xc6, 0xbb, 0x96, 0x78, 0x01, 0x1e, 0x8b, 0x3b, 0x17,
	0x5e, 0x80, 0x43, 0x9f, 0x05, 0xed, 0xd8, 0x4e, 0x52, 0x08, 0xea, 0xa1, 0x37, 0x7f, 0x33, 0xdf,
	0x37, 0xe3, 0x99, 0xf9, 0x16, 0x5a, 0x51, 0xe2, 0xb2, 0x84, 0x1b, 0x71, 0x12, 0x89, 0x88, 0xa8,
	0x19, 0xea, 0x83, 0x17, 0x79, 0x51, 0x16, 0xeb, 0x1f, 0x78, 0x51, 0xe4, 0x05, 0xec, 0x18, 0xd1,
	0x22, 0xfd, 0x7c, 0x2c, 0xfc, 0x25, 0xe3, 0xc2, 0x59, 0xc6, 0x39, 0x01, 0xc2, 0xc8, 0x65, 0xd9,
	0xb7, 0xfe, 0xab, 0x02, 0xf0, 0x41, 0xd6, 0xb8, 0xf4, 0x97, 0xbe, 0x20, 0xa7, 0xd0, 0xe6, 0x2c,
	0xf1, 0x9d, 0x60, 0x1e, 0xa6, 0xcb, 0x05, 0x4b, 0x7a, 0xca, 0x40, 0x19, 0xb6, 0xce, 0xf6, 0x7f,
	0xde, 0x1d, 0xec, 0xfc, 0xbe, 0x3b, 0x68, 0xcd, 0x30, 0x79, 0x85, 0x39, 0xda, 0xe2, 0x1b, 0x88,
	0x9c, 0x40, 0x8b, 0x3b, 0x82, 0x05, 0x81, 0x2f, 0xd8, 0xdc, 0x77, 0x7b, 0x25, 0x54, 0x76, 0x72,
	0xa5, 0x7a, 0x15, 0xb9, 0xcc, 0x9a, 0xd0, 0xe6, 0x8a, 0x63, 0xb9, 0xe4, 0x08, 0x1a, 0x69, 0x1c,
	0xf8, 0xe1, 0x17, 0xc9, 0x2f, 0x6f, 0xe5, 0xd7, 0x33, 0x82, 0xe5, 0x92, 0x37, 0xd0, 0xe5, 0x22,
	0x4a, 0x1c, 0x8f, 0xcd, 0xe5, 0xff, 0x4b, 0x49, 0x65, 0xab, 0xa4, 0x9d, 0xd3, 0x10, 0xba, 0xe4,
	0x10, 0xea, 0xb1, 0xcf, 0x6e, 0x50, 0x50, 0x45, 0x41, 0x37, 0x17, 0xd4, 0xa6, 0x32, 0x6e, 0x4d,
	0x68, 0x0d, 0x09, 0x96, 0x4b, 0xf6, 0xa1, 0x1a, 0xc8, 0x3d, 0xf4, 0xd4, 0x81, 0x32, 0x2c, 0xd3,
	0x0c, 0x90, 0x23, 0x50, 0x9d, 0x1b, 0xe1, 0x47, 0x61, 0xaf, 0x36, 0x50, 0x86, 0x9d, 0xd1, 0x9e,
	0x91, 0xdf, 0x00, 0xf5, 0x63, 0x4c, 0xd1, 0x9c, 0x42, 0x4c, 0xd0, 0xb2, 0x76, 0xec, 0x5b, 0xec,
	0x27, 0x0e, 0xca, 0xea, 0x03, 0x65, 0xd8, 0x1c, 0xf5, 0x8d, 0xec, 0x30, 0x46, 0x71, 0x18, 0xc3,
	0x2e, 0x0e, 0x43, 0xbb, 0xa8, 0x31, 0x57, 0x12, 0x59, 0x06, 0x9b, 0x6c, 0x96, 0x69, 0x3c, 0x5c,
	0x06, 0x35, 0x1b, 0x65, 0x8e, 0x61, 0x6f, 0x7d, 0x14, 0xee, 0x7b, 0xa1, 0x23, 0xd2, 0x84, 0xf5,
	0x40, 0xee, 0x81, 0x92, 0x55, 0x6a, 0x56, 0x64, 0xc8, 0x5b, 0xd8, 0x5d, 0x0b, 0x1c, 0xd7, 0x4d,
	0x18, 0xe7, 0xbd, 0x26, 0x36, 0xde, 0x35, 0xd0, 0x37, 0x72, 0xad, 0xe3, 0x2c, 0x41, 0xb5, 0x15,
	0x37, 0x8f, 0xe8, 0xdf, 0x15, 0xa8, 0xa2, 0x9f, 0x1e, 0x63, 0xa5, 0x27, 0xa0, 0x3a, 0xcb, 0x28,
	0x0d, 0x05, 0x9a, 0xa8, 0x4c, 0x73, 0x44, 0x5e, 0x82, 0x96, 0xfb, 0x65, 0x3d, 0x0a, 0xda, 0x86,
	0x76, 0xb3, 0xf8, 0x6a, 0x0e, 0xdd, 0x87, 0x06, 0x5e, 0xe7, 0x9d, 0xc3, 0x6f, 0xef, 0x59, 0x40,
	0x79, 0xc0, 0x02, 0x04, 0x2a, 0xb7, 0x0e, 0xbf, 0xcd, 0xec, 0x4b, 0xf1, 0x9b, 0x3c, 0x83, 0xc6,
	0xdf, 0x0d, 0xd7, 0x01, 0x7d, 0x01, 0xbb, 0x33, 0x26, 0x44, 0xc0, 0x96, 0x2c, 0x14, 0x94, 0x7d,
	0x4d, 0x19, 0x17, 0x64, 0x58, 0x38, 0x49, 0xc1, 0xdd, 0x91, 0xc2, 0x32, 0xeb, 0xb7, 0x56, 0xb8,
	0xeb, 0x05, 0x54, 0x31, 0x87, 0x1d, 0x9b, 0xa3, 0xf6, 0x3d, 0x26, 0xcd, 0x72, 0xfa, 0x0f, 0x05,
	0xc8, 0x66, 0x13, 0x1e, 0x47, 0x21, 0x67, 0x8f, 0xd9, 0xf1, 0x29, 0xa8, 0x5c, 0x38, 0x22, 0xe5,
	0xd8, 0xb7, 0x33, 0x7a, 0x5e, 0xf4, 0xfd, 0xb7, 0x8d, 0x31, 0x43, 0x22, 0xcd, 0x05, 0xfa, 0x09,
	0xa8, 0x59, 0x84, 0x34, 0xa1, 0x66, 0x5d, 0x7d, 0x1c, 0x5f, 0x5a, 0x13, 0x6d, 0x87, 0xb4, 0xa0,
	0x3e, 0x3e, 0x3f, 0x37, 0xa7, 0xb6, 0x39, 0xd1, 0x14, 0x89, 0xa8, 0xf9, 0xde, 0x3c, 0x97, 0xa8,
	0x74, 0xe8, 0x41, 0x73, 0xe3, 0xb1, 0xdc, 0xd7, 0xd5, 0xa0, 0x3c, 0xbd, 0xb6, 0x35, 0x45, 0x7e,
	0x5c, 0x98, 0xb6, 0x56, 0x22, 0x6d, 0x68, 0x5c, 0x98, 0xf6, 0x7c, 0x7c, 0x3d, 0xb1, 0x6c, 0xad,
	0x4c, 0x3a, 0x00, 0x12, 0x52, 0x73, 0x3a, 0xb6, 0xa8, 0x56, 0x91, 0x78, 0x7a, 0xbd, 0xc2, 0x55,
	0x02, 0xa0, 0x4e, 0xcc, 0x4b, 0xd3, 0x36, 0x35, 0x75, 0x34, 0x03, 0x15, 0x17, 0xc7, 0x89, 0x05,
	0xb0, 0x1e, 0x85, 0x3c, 0xdd, 0x36, 0x1e, 0x9e, 0xaa, 0xdf, 0xff, 0xff, 0xe4, 0xfa, 0xce, 0x50,
	0x79, 0xa5, 0x9c, 0x55, 0x3e, 0x95, 0xe2, 0xc5, 0x42, 0xc5, 0x07, 0xf7, 0xfa, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0x80, 0xc7, 0x0e, 0x82, 0x05, 0x00, 0x00,
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
