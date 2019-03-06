// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metainfo.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Action int32

const (
	Action_INVALID    Action = 0
	Action_PUT        Action = 1
	Action_GET        Action = 2
	Action_GET_AUDIT  Action = 3
	Action_GET_REPAIR Action = 4
	Action_PUT_REPAIR Action = 5
	Action_DELETE     Action = 6
)

var Action_name = map[int32]string{
	0: "INVALID",
	1: "PUT",
	2: "GET",
	3: "GET_AUDIT",
	4: "GET_REPAIR",
	5: "PUT_REPAIR",
	6: "DELETE",
}
var Action_value = map[string]int32{
	"INVALID":    0,
	"PUT":        1,
	"GET":        2,
	"GET_AUDIT":  3,
	"GET_REPAIR": 4,
	"PUT_REPAIR": 5,
	"DELETE":     6,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{0}
}

type OrderLimit struct {
	// unique serial to avoid replay attacks
	SerialNumber []byte `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// satellite who issued this order limit allowing orderer to do the specified action
	SatelliteId *NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id,omitempty"`
	// uplink who requested or whom behalf the order limit to do an action
	UplinkId *NodeID `protobuf:"bytes,3,opt,name=uplink_id,json=uplinkId,proto3,customtype=NodeID" json:"uplink_id,omitempty"`
	// storage node who can reclaim the order limit specified by serial
	StorageNodeId *NodeID `protobuf:"bytes,4,opt,name=storage_node_id,json=storageNodeId,proto3,customtype=NodeID" json:"storage_node_id,omitempty"`
	// piece which is allowed to be touched
	PieceId *PieceID `protobuf:"bytes,5,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id,omitempty"`
	// limit in bytes how much can be changed
	Limit                int64                `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Action               Action               `protobuf:"varint,7,opt,name=action,proto3,enum=metainfo.Action" json:"action,omitempty"`
	PieceExpiration      *timestamp.Timestamp `protobuf:"bytes,8,opt,name=piece_expiration,json=pieceExpiration,proto3" json:"piece_expiration,omitempty"`
	OrderExpiration      *timestamp.Timestamp `protobuf:"bytes,9,opt,name=order_expiration,json=orderExpiration,proto3" json:"order_expiration,omitempty"`
	SatelliteSignature   []byte               `protobuf:"bytes,10,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderLimit) Reset()         { *m = OrderLimit{} }
func (m *OrderLimit) String() string { return proto.CompactTextString(m) }
func (*OrderLimit) ProtoMessage()    {}
func (*OrderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{0}
}
func (m *OrderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimit.Unmarshal(m, b)
}
func (m *OrderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimit.Marshal(b, m, deterministic)
}
func (dst *OrderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimit.Merge(dst, src)
}
func (m *OrderLimit) XXX_Size() int {
	return xxx_messageInfo_OrderLimit.Size(m)
}
func (m *OrderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimit proto.InternalMessageInfo

func (m *OrderLimit) GetSerialNumber() []byte {
	if m != nil {
		return m.SerialNumber
	}
	return nil
}

func (m *OrderLimit) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderLimit) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INVALID
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

type Order struct {
	// serial of the order limit that was signed
	SerialNumber []byte `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
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
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{1}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetSerialNumber() []byte {
	if m != nil {
		return m.SerialNumber
	}
	return nil
}

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
	PieceId []byte `protobuf:"bytes,1,opt,name=piece_id,json=pieceId,proto3" json:"piece_id,omitempty"`
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
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{2}
}
func (m *PieceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceHash.Unmarshal(m, b)
}
func (m *PieceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceHash.Marshal(b, m, deterministic)
}
func (dst *PieceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceHash.Merge(dst, src)
}
func (m *PieceHash) XXX_Size() int {
	return xxx_messageInfo_PieceHash.Size(m)
}
func (m *PieceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceHash.DiscardUnknown(m)
}

var xxx_messageInfo_PieceHash proto.InternalMessageInfo

func (m *PieceHash) GetPieceId() []byte {
	if m != nil {
		return m.PieceId
	}
	return nil
}

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

type SegmentWriteRequest struct {
	Bucket               []byte   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Segment              int64    `protobuf:"varint,2,opt,name=segment,proto3" json:"segment,omitempty"`
	EncryptedPath        []byte   `protobuf:"bytes,3,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentWriteRequest) Reset()         { *m = SegmentWriteRequest{} }
func (m *SegmentWriteRequest) String() string { return proto.CompactTextString(m) }
func (*SegmentWriteRequest) ProtoMessage()    {}
func (*SegmentWriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{3}
}
func (m *SegmentWriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentWriteRequest.Unmarshal(m, b)
}
func (m *SegmentWriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentWriteRequest.Marshal(b, m, deterministic)
}
func (dst *SegmentWriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentWriteRequest.Merge(dst, src)
}
func (m *SegmentWriteRequest) XXX_Size() int {
	return xxx_messageInfo_SegmentWriteRequest.Size(m)
}
func (m *SegmentWriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentWriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentWriteRequest proto.InternalMessageInfo

func (m *SegmentWriteRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *SegmentWriteRequest) GetSegment() int64 {
	if m != nil {
		return m.Segment
	}
	return 0
}

func (m *SegmentWriteRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

type SegmentCommitRequest struct {
	Bucket                 []byte       `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Segment                int64        `protobuf:"varint,2,opt,name=segment,proto3" json:"segment,omitempty"`
	EncryptedPath          []byte       `protobuf:"bytes,3,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	StorageNodePiecehashes []*PieceHash `protobuf:"bytes,4,rep,name=storage_node_piecehashes,json=storageNodePiecehashes,proto3" json:"storage_node_piecehashes,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}     `json:"-"`
	XXX_unrecognized       []byte       `json:"-"`
	XXX_sizecache          int32        `json:"-"`
}

func (m *SegmentCommitRequest) Reset()         { *m = SegmentCommitRequest{} }
func (m *SegmentCommitRequest) String() string { return proto.CompactTextString(m) }
func (*SegmentCommitRequest) ProtoMessage()    {}
func (*SegmentCommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{4}
}
func (m *SegmentCommitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentCommitRequest.Unmarshal(m, b)
}
func (m *SegmentCommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentCommitRequest.Marshal(b, m, deterministic)
}
func (dst *SegmentCommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentCommitRequest.Merge(dst, src)
}
func (m *SegmentCommitRequest) XXX_Size() int {
	return xxx_messageInfo_SegmentCommitRequest.Size(m)
}
func (m *SegmentCommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentCommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentCommitRequest proto.InternalMessageInfo

func (m *SegmentCommitRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *SegmentCommitRequest) GetSegment() int64 {
	if m != nil {
		return m.Segment
	}
	return 0
}

func (m *SegmentCommitRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

func (m *SegmentCommitRequest) GetStorageNodePiecehashes() []*PieceHash {
	if m != nil {
		return m.StorageNodePiecehashes
	}
	return nil
}

type SegmentCommitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentCommitResponse) Reset()         { *m = SegmentCommitResponse{} }
func (m *SegmentCommitResponse) String() string { return proto.CompactTextString(m) }
func (*SegmentCommitResponse) ProtoMessage()    {}
func (*SegmentCommitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{5}
}
func (m *SegmentCommitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentCommitResponse.Unmarshal(m, b)
}
func (m *SegmentCommitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentCommitResponse.Marshal(b, m, deterministic)
}
func (dst *SegmentCommitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentCommitResponse.Merge(dst, src)
}
func (m *SegmentCommitResponse) XXX_Size() int {
	return xxx_messageInfo_SegmentCommitResponse.Size(m)
}
func (m *SegmentCommitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentCommitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentCommitResponse proto.InternalMessageInfo

type SegementDownloadRequest struct {
	Bucket               []byte   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	EncryptedPath        []byte   `protobuf:"bytes,2,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegementDownloadRequest) Reset()         { *m = SegementDownloadRequest{} }
func (m *SegementDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*SegementDownloadRequest) ProtoMessage()    {}
func (*SegementDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{6}
}
func (m *SegementDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegementDownloadRequest.Unmarshal(m, b)
}
func (m *SegementDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegementDownloadRequest.Marshal(b, m, deterministic)
}
func (dst *SegementDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegementDownloadRequest.Merge(dst, src)
}
func (m *SegementDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_SegementDownloadRequest.Size(m)
}
func (m *SegementDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SegementDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SegementDownloadRequest proto.InternalMessageInfo

func (m *SegementDownloadRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *SegementDownloadRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

type SegmentDeleteRequest struct {
	Bucket               []byte   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Segment              int64    `protobuf:"varint,2,opt,name=segment,proto3" json:"segment,omitempty"`
	EncryptedPath        []byte   `protobuf:"bytes,3,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentDeleteRequest) Reset()         { *m = SegmentDeleteRequest{} }
func (m *SegmentDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*SegmentDeleteRequest) ProtoMessage()    {}
func (*SegmentDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{7}
}
func (m *SegmentDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentDeleteRequest.Unmarshal(m, b)
}
func (m *SegmentDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *SegmentDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentDeleteRequest.Merge(dst, src)
}
func (m *SegmentDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_SegmentDeleteRequest.Size(m)
}
func (m *SegmentDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentDeleteRequest proto.InternalMessageInfo

func (m *SegmentDeleteRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *SegmentDeleteRequest) GetSegment() int64 {
	if m != nil {
		return m.Segment
	}
	return 0
}

func (m *SegmentDeleteRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

type OrderLimitResponse struct {
	Limits               []*OrderLimit `protobuf:"bytes,1,rep,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OrderLimitResponse) Reset()         { *m = OrderLimitResponse{} }
func (m *OrderLimitResponse) String() string { return proto.CompactTextString(m) }
func (*OrderLimitResponse) ProtoMessage()    {}
func (*OrderLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_metainfo_3ccf470fc1d7e21f, []int{8}
}
func (m *OrderLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimitResponse.Unmarshal(m, b)
}
func (m *OrderLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimitResponse.Marshal(b, m, deterministic)
}
func (dst *OrderLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimitResponse.Merge(dst, src)
}
func (m *OrderLimitResponse) XXX_Size() int {
	return xxx_messageInfo_OrderLimitResponse.Size(m)
}
func (m *OrderLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimitResponse proto.InternalMessageInfo

func (m *OrderLimitResponse) GetLimits() []*OrderLimit {
	if m != nil {
		return m.Limits
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderLimit)(nil), "metainfo.OrderLimit")
	proto.RegisterType((*Order)(nil), "metainfo.Order")
	proto.RegisterType((*PieceHash)(nil), "metainfo.PieceHash")
	proto.RegisterType((*SegmentWriteRequest)(nil), "metainfo.SegmentWriteRequest")
	proto.RegisterType((*SegmentCommitRequest)(nil), "metainfo.SegmentCommitRequest")
	proto.RegisterType((*SegmentCommitResponse)(nil), "metainfo.SegmentCommitResponse")
	proto.RegisterType((*SegementDownloadRequest)(nil), "metainfo.SegementDownloadRequest")
	proto.RegisterType((*SegmentDeleteRequest)(nil), "metainfo.SegmentDeleteRequest")
	proto.RegisterType((*OrderLimitResponse)(nil), "metainfo.OrderLimitResponse")
	proto.RegisterEnum("metainfo.Action", Action_name, Action_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PieceStoreRoutesClient is the client API for PieceStoreRoutes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PieceStoreRoutesClient interface {
	CreateSegment(ctx context.Context, in *SegmentWriteRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error)
	CommitSegment(ctx context.Context, in *SegmentCommitRequest, opts ...grpc.CallOption) (*SegmentCommitResponse, error)
	DownloadSegment(ctx context.Context, in *SegementDownloadRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error)
	DeleteSegment(ctx context.Context, in *SegmentDeleteRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error)
}

type pieceStoreRoutesClient struct {
	cc *grpc.ClientConn
}

func NewPieceStoreRoutesClient(cc *grpc.ClientConn) PieceStoreRoutesClient {
	return &pieceStoreRoutesClient{cc}
}

func (c *pieceStoreRoutesClient) CreateSegment(ctx context.Context, in *SegmentWriteRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error) {
	out := new(OrderLimitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.PieceStoreRoutes/CreateSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) CommitSegment(ctx context.Context, in *SegmentCommitRequest, opts ...grpc.CallOption) (*SegmentCommitResponse, error) {
	out := new(SegmentCommitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.PieceStoreRoutes/CommitSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) DownloadSegment(ctx context.Context, in *SegementDownloadRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error) {
	out := new(OrderLimitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.PieceStoreRoutes/DownloadSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) DeleteSegment(ctx context.Context, in *SegmentDeleteRequest, opts ...grpc.CallOption) (*OrderLimitResponse, error) {
	out := new(OrderLimitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.PieceStoreRoutes/DeleteSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PieceStoreRoutesServer is the server API for PieceStoreRoutes service.
type PieceStoreRoutesServer interface {
	CreateSegment(context.Context, *SegmentWriteRequest) (*OrderLimitResponse, error)
	CommitSegment(context.Context, *SegmentCommitRequest) (*SegmentCommitResponse, error)
	DownloadSegment(context.Context, *SegementDownloadRequest) (*OrderLimitResponse, error)
	DeleteSegment(context.Context, *SegmentDeleteRequest) (*OrderLimitResponse, error)
}

func RegisterPieceStoreRoutesServer(s *grpc.Server, srv PieceStoreRoutesServer) {
	s.RegisterService(&_PieceStoreRoutes_serviceDesc, srv)
}

func _PieceStoreRoutes_CreateSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).CreateSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.PieceStoreRoutes/CreateSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).CreateSegment(ctx, req.(*SegmentWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_CommitSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).CommitSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.PieceStoreRoutes/CommitSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).CommitSegment(ctx, req.(*SegmentCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_DownloadSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegementDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).DownloadSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.PieceStoreRoutes/DownloadSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).DownloadSegment(ctx, req.(*SegementDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_DeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).DeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.PieceStoreRoutes/DeleteSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).DeleteSegment(ctx, req.(*SegmentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PieceStoreRoutes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metainfo.PieceStoreRoutes",
	HandlerType: (*PieceStoreRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSegment",
			Handler:    _PieceStoreRoutes_CreateSegment_Handler,
		},
		{
			MethodName: "CommitSegment",
			Handler:    _PieceStoreRoutes_CommitSegment_Handler,
		},
		{
			MethodName: "DownloadSegment",
			Handler:    _PieceStoreRoutes_DownloadSegment_Handler,
		},
		{
			MethodName: "DeleteSegment",
			Handler:    _PieceStoreRoutes_DeleteSegment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metainfo.proto",
}

func init() { proto.RegisterFile("metainfo.proto", fileDescriptor_metainfo_3ccf470fc1d7e21f) }

var fileDescriptor_metainfo_3ccf470fc1d7e21f = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0x8e, 0xf9, 0x31, 0x30, 0x04, 0xb0, 0x36, 0x39, 0x89, 0x0f, 0xca, 0x39, 0x70, 0x38, 0x6a,
	0x4b, 0xab, 0x96, 0x48, 0xf4, 0x09, 0x48, 0xb0, 0x52, 0x4b, 0x29, 0x41, 0x86, 0xb4, 0x55, 0x6f,
	0xd0, 0x82, 0x37, 0xc6, 0x8a, 0xed, 0x75, 0xed, 0xb5, 0xda, 0xbe, 0x5c, 0x1f, 0x24, 0x52, 0x73,
	0xd5, 0x07, 0xa9, 0xbc, 0xfe, 0x83, 0x40, 0x44, 0x6f, 0xda, 0x3b, 0xcf, 0xcc, 0xb7, 0xdf, 0xec,
	0x7c, 0xdf, 0xac, 0xa1, 0x6e, 0x13, 0x86, 0x4d, 0xe7, 0x86, 0xf6, 0x5c, 0x8f, 0x32, 0x8a, 0xca,
	0x49, 0xdc, 0x04, 0x83, 0x1a, 0x71, 0xb6, 0xd9, 0x32, 0x28, 0x35, 0x2c, 0x72, 0xca, 0xa3, 0x79,
	0x70, 0x73, 0xca, 0x4c, 0x9b, 0xf8, 0x0c, 0xdb, 0x6e, 0x04, 0xe8, 0xfc, 0xc8, 0x03, 0x5c, 0x79,
	0x3a, 0xf1, 0x2e, 0x4d, 0xdb, 0x64, 0xe8, 0x7f, 0xa8, 0xf9, 0xc4, 0x33, 0xb1, 0x35, 0x73, 0x02,
	0x7b, 0x4e, 0x3c, 0x59, 0x68, 0x0b, 0xdd, 0x7d, 0x6d, 0x3f, 0x4a, 0x8e, 0x78, 0x0e, 0xbd, 0x82,
	0x7d, 0x1f, 0x33, 0x62, 0x59, 0x26, 0x23, 0x33, 0x53, 0x97, 0x73, 0x21, 0xe6, 0x0c, 0xee, 0xee,
	0x5b, 0xe2, 0x88, 0xea, 0x44, 0x1d, 0x6a, 0xd5, 0xb4, 0xae, 0xea, 0xe8, 0x19, 0x54, 0x02, 0xd7,
	0x32, 0x9d, 0xdb, 0x10, 0x9b, 0xdf, 0xc0, 0x96, 0xa3, 0xa2, 0xaa, 0xa3, 0x3e, 0x34, 0x7c, 0x46,
	0x3d, 0x6c, 0x90, 0x99, 0x43, 0x75, 0x4e, 0x5d, 0xd8, 0x80, 0xd7, 0x62, 0x08, 0x0f, 0x75, 0xf4,
	0x14, 0xca, 0xae, 0x49, 0x16, 0x1c, 0x5c, 0xe4, 0xe0, 0xea, 0xdd, 0x7d, 0xab, 0x34, 0x0e, 0x73,
	0xea, 0x50, 0x2b, 0xf1, 0xa2, 0xaa, 0xa3, 0x43, 0x28, 0x5a, 0xe1, 0x84, 0xb2, 0xd8, 0x16, 0xba,
	0x79, 0x2d, 0x0a, 0x50, 0x17, 0x44, 0xbc, 0x60, 0x26, 0x75, 0xe4, 0x52, 0x5b, 0xe8, 0xd6, 0xfb,
	0x52, 0x2f, 0x55, 0x75, 0xc0, 0xf3, 0x5a, 0x5c, 0x47, 0x0a, 0x48, 0x51, 0x1f, 0xf2, 0xc5, 0x35,
	0x3d, 0xcc, 0xcf, 0x94, 0xdb, 0x42, 0xb7, 0xda, 0x6f, 0xf6, 0x22, 0x8d, 0x7b, 0x89, 0xc6, 0xbd,
	0x69, 0xa2, 0xb1, 0xd6, 0xe0, 0x67, 0x94, 0xf4, 0x48, 0x48, 0x43, 0x43, 0xb5, 0x57, 0x69, 0x2a,
	0xbb, 0x69, 0xf8, 0x99, 0x15, 0x9a, 0x53, 0x38, 0xc8, 0x1c, 0xf0, 0x4d, 0xc3, 0xc1, 0x2c, 0xf0,
	0x88, 0x0c, 0xdc, 0x2c, 0x94, 0x96, 0x26, 0x49, 0xa5, 0x43, 0xa1, 0xc8, 0x5d, 0xfe, 0x35, 0x83,
	0x8f, 0x40, 0xc4, 0x36, 0x0d, 0x1c, 0xc6, 0xad, 0xcd, 0x6b, 0x71, 0x84, 0x9e, 0x83, 0x14, 0x3b,
	0x99, 0xf5, 0xe4, 0x86, 0x6a, 0x8d, 0x28, 0x9f, 0x35, 0xfc, 0x00, 0x15, 0xee, 0xc1, 0x1b, 0xec,
	0x2f, 0xd1, 0xdf, 0x2b, 0x26, 0x45, 0xfd, 0x52, 0x5f, 0x10, 0x14, 0x96, 0xd8, 0x5f, 0x46, 0x3b,
	0xa4, 0xf1, 0x6f, 0x74, 0x02, 0x95, 0x87, 0xfc, 0x59, 0xa2, 0xe3, 0xc0, 0xc1, 0x84, 0x18, 0x36,
	0x71, 0xd8, 0x7b, 0xcf, 0x64, 0x44, 0x23, 0x9f, 0x02, 0xe2, 0xb3, 0xf0, 0xce, 0xf3, 0x60, 0x71,
	0x4b, 0x58, 0xdc, 0x21, 0x8e, 0x90, 0x0c, 0x25, 0x3f, 0x82, 0xc7, 0xc3, 0x24, 0x21, 0x7a, 0x02,
	0x75, 0xe2, 0x2c, 0xbc, 0xaf, 0x2e, 0x23, 0xfa, 0xcc, 0xc5, 0x6c, 0x19, 0xf7, 0xaa, 0xa5, 0xd9,
	0x31, 0x66, 0xcb, 0xce, 0x37, 0x01, 0x0e, 0xe3, 0x86, 0xe7, 0xd4, 0xb6, 0x4d, 0xf6, 0xbb, 0x3b,
	0xa2, 0xb7, 0x20, 0xaf, 0xbd, 0x03, 0xae, 0x55, 0xa8, 0x0c, 0xf1, 0xe5, 0x42, 0x3b, 0xdf, 0xad,
	0xf6, 0x0f, 0xb2, 0x3d, 0x4d, 0x55, 0xd6, 0x8e, 0x56, 0x5e, 0xc6, 0x38, 0x3b, 0xd2, 0x39, 0x86,
	0xbf, 0x1e, 0xdc, 0xdf, 0x77, 0xa9, 0xe3, 0x87, 0x1e, 0x1d, 0x4f, 0x88, 0x41, 0xc2, 0xca, 0x90,
	0x7e, 0x76, 0x2c, 0x8a, 0xf5, 0x5d, 0xb3, 0x6d, 0x4e, 0x90, 0xdb, 0xa6, 0x19, 0x4d, 0x25, 0x1b,
	0x12, 0x8b, 0xfc, 0x01, 0x93, 0xce, 0x00, 0x65, 0x7f, 0xb1, 0x64, 0x40, 0xf4, 0x12, 0x44, 0xfe,
	0xce, 0x7d, 0x59, 0xe0, 0xb2, 0x1d, 0x66, 0xb2, 0xad, 0xa0, 0x63, 0xcc, 0x8b, 0x39, 0x88, 0xd1,
	0xa3, 0x47, 0x55, 0x28, 0xa9, 0xa3, 0x77, 0x83, 0x4b, 0x75, 0x28, 0xed, 0xa1, 0x12, 0xe4, 0xc7,
	0xd7, 0x53, 0x49, 0x08, 0x3f, 0x2e, 0x94, 0xa9, 0x94, 0x43, 0x35, 0xa8, 0x5c, 0x28, 0xd3, 0xd9,
	0xe0, 0x7a, 0xa8, 0x4e, 0xa5, 0x3c, 0xaa, 0x03, 0x84, 0xa1, 0xa6, 0x8c, 0x07, 0xaa, 0x26, 0x15,
	0xc2, 0x78, 0x7c, 0x9d, 0xc6, 0x45, 0x04, 0x20, 0x0e, 0x95, 0x4b, 0x65, 0xaa, 0x48, 0x62, 0xff,
	0x7b, 0x0e, 0x24, 0xee, 0xcd, 0x84, 0x51, 0x8f, 0x68, 0x34, 0x60, 0xc4, 0x47, 0x23, 0xa8, 0x9d,
	0x7b, 0x04, 0x33, 0x12, 0x6b, 0x86, 0xfe, 0xc9, 0xee, 0xb9, 0x65, 0xd5, 0x9b, 0x27, 0x5b, 0xc7,
	0x48, 0x5c, 0xdd, 0x43, 0x1a, 0xd4, 0x22, 0xa7, 0x13, 0xbe, 0x7f, 0x37, 0xf8, 0xd6, 0x36, 0xb9,
	0xd9, 0x7a, 0xb4, 0x9e, 0x72, 0x4e, 0xa1, 0x91, 0xec, 0x48, 0xc2, 0xfa, 0xdf, 0xda, 0xa9, 0x6d,
	0x6b, 0xb4, 0xf3, 0xa6, 0x57, 0x50, 0x8b, 0x16, 0xe4, 0xf1, 0x9b, 0xae, 0x2d, 0xd0, 0x2e, 0xc2,
	0xb3, 0xc2, 0xc7, 0x9c, 0x3b, 0x9f, 0x8b, 0xfc, 0x1f, 0xfa, 0xfa, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xda, 0xad, 0x6b, 0x52, 0x24, 0x07, 0x00, 0x00,
}
