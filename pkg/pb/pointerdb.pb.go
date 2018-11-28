// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pointerdb.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_RS RedundancyScheme_SchemeType = 0
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "RS",
}
var RedundancyScheme_SchemeType_value = map[string]int32{
	"RS": 0,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}
func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{0, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}
var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}
func (Pointer_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{3, 0}
}

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq               int32    `protobuf:"varint,2,opt,name=min_req,json=minReq,proto3" json:"min_req,omitempty"`
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	RepairThreshold      int32    `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold,proto3" json:"repair_threshold,omitempty"`
	SuccessThreshold     int32    `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold,proto3" json:"success_threshold,omitempty"`
	ErasureShareSize     int32    `protobuf:"varint,6,opt,name=erasure_share_size,json=erasureShareSize,proto3" json:"erasure_share_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedundancyScheme) Reset()         { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()    {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{0}
}
func (m *RedundancyScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedundancyScheme.Unmarshal(m, b)
}
func (m *RedundancyScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedundancyScheme.Marshal(b, m, deterministic)
}
func (dst *RedundancyScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedundancyScheme.Merge(dst, src)
}
func (m *RedundancyScheme) XXX_Size() int {
	return xxx_messageInfo_RedundancyScheme.Size(m)
}
func (m *RedundancyScheme) XXX_DiscardUnknown() {
	xxx_messageInfo_RedundancyScheme.DiscardUnknown(m)
}

var xxx_messageInfo_RedundancyScheme proto.InternalMessageInfo

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_RS
}

func (m *RedundancyScheme) GetMinReq() int32 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int32 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int32 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetErasureShareSize() int32 {
	if m != nil {
		return m.ErasureShareSize
	}
	return 0
}

type RemotePiece struct {
	PieceNum             int32    `protobuf:"varint,1,opt,name=piece_num,json=pieceNum,proto3" json:"piece_num,omitempty"`
	NodeId               NodeID   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemotePiece) Reset()         { *m = RemotePiece{} }
func (m *RemotePiece) String() string { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()    {}
func (*RemotePiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{1}
}
func (m *RemotePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemotePiece.Unmarshal(m, b)
}
func (m *RemotePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemotePiece.Marshal(b, m, deterministic)
}
func (dst *RemotePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemotePiece.Merge(dst, src)
}
func (m *RemotePiece) XXX_Size() int {
	return xxx_messageInfo_RemotePiece.Size(m)
}
func (m *RemotePiece) XXX_DiscardUnknown() {
	xxx_messageInfo_RemotePiece.DiscardUnknown(m)
}

var xxx_messageInfo_RemotePiece proto.InternalMessageInfo

func (m *RemotePiece) GetPieceNum() int32 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

type RemoteSegment struct {
	Redundancy *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy" json:"redundancy,omitempty"`
	// TODO: may want to use customtype and fixed-length byte slice
	PieceId              string         `protobuf:"bytes,2,opt,name=piece_id,json=pieceId,proto3" json:"piece_id,omitempty"`
	RemotePieces         []*RemotePiece `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces" json:"remote_pieces,omitempty"`
	MerkleRoot           []byte         `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RemoteSegment) Reset()         { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()    {}
func (*RemoteSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{2}
}
func (m *RemoteSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSegment.Unmarshal(m, b)
}
func (m *RemoteSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSegment.Marshal(b, m, deterministic)
}
func (dst *RemoteSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSegment.Merge(dst, src)
}
func (m *RemoteSegment) XXX_Size() int {
	return xxx_messageInfo_RemoteSegment.Size(m)
}
func (m *RemoteSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSegment.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSegment proto.InternalMessageInfo

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetPieceId() string {
	if m != nil {
		return m.PieceId
	}
	return ""
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type Pointer struct {
	Type          Pointer_DataType `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.Pointer_DataType" json:"type,omitempty"`
	InlineSegment []byte           `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote        *RemoteSegment   `protobuf:"bytes,4,opt,name=remote" json:"remote,omitempty"`
	// TODO: rename
	SegmentSize          int64                `protobuf:"varint,5,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	CreationDate         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	ExpirationDate       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`
	Metadata             []byte               `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pointer) Reset()         { *m = Pointer{} }
func (m *Pointer) String() string { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()    {}
func (*Pointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{3}
}
func (m *Pointer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pointer.Unmarshal(m, b)
}
func (m *Pointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pointer.Marshal(b, m, deterministic)
}
func (dst *Pointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pointer.Merge(dst, src)
}
func (m *Pointer) XXX_Size() int {
	return xxx_messageInfo_Pointer.Size(m)
}
func (m *Pointer) XXX_DiscardUnknown() {
	xxx_messageInfo_Pointer.DiscardUnknown(m)
}

var xxx_messageInfo_Pointer proto.InternalMessageInfo

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetSegmentSize() int64 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *Pointer) GetCreationDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Pointer) GetExpirationDate() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationDate
	}
	return nil
}

func (m *Pointer) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// PutRequest is a request message for the Put rpc call
type PutRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer              *Pointer `protobuf:"bytes,2,opt,name=pointer" json:"pointer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{4}
}
func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (dst *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(dst, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PutRequest) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{5}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// ListRequest is a request message for the List rpc call
type ListRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	StartAfter           string   `protobuf:"bytes,2,opt,name=start_after,json=startAfter,proto3" json:"start_after,omitempty"`
	EndBefore            string   `protobuf:"bytes,3,opt,name=end_before,json=endBefore,proto3" json:"end_before,omitempty"`
	Recursive            bool     `protobuf:"varint,4,opt,name=recursive,proto3" json:"recursive,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	MetaFlags            uint32   `protobuf:"fixed32,6,opt,name=meta_flags,json=metaFlags,proto3" json:"meta_flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{6}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListRequest) GetStartAfter() string {
	if m != nil {
		return m.StartAfter
	}
	return ""
}

func (m *ListRequest) GetEndBefore() string {
	if m != nil {
		return m.EndBefore
	}
	return ""
}

func (m *ListRequest) GetRecursive() bool {
	if m != nil {
		return m.Recursive
	}
	return false
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetMetaFlags() uint32 {
	if m != nil {
		return m.MetaFlags
	}
	return 0
}

// PutResponse is a response message for the Put rpc call
type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{7}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (dst *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(dst, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Pointer              *Pointer                  `protobuf:"bytes,1,opt,name=pointer" json:"pointer,omitempty"`
	Nodes                []*Node                   `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	Pba                  *PayerBandwidthAllocation `protobuf:"bytes,3,opt,name=pba" json:"pba,omitempty"`
	Authorization        *SignedMessage            `protobuf:"bytes,4,opt,name=authorization" json:"authorization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{8}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *GetResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *GetResponse) GetPba() *PayerBandwidthAllocation {
	if m != nil {
		return m.Pba
	}
	return nil
}

func (m *GetResponse) GetAuthorization() *SignedMessage {
	if m != nil {
		return m.Authorization
	}
	return nil
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Items                []*ListResponse_Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	More                 bool                 `protobuf:"varint,2,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{9}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*ListResponse_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListResponse) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

type ListResponse_Item struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer              *Pointer `protobuf:"bytes,2,opt,name=pointer" json:"pointer,omitempty"`
	IsPrefix             bool     `protobuf:"varint,3,opt,name=is_prefix,json=isPrefix,proto3" json:"is_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse_Item) Reset()         { *m = ListResponse_Item{} }
func (m *ListResponse_Item) String() string { return proto.CompactTextString(m) }
func (*ListResponse_Item) ProtoMessage()    {}
func (*ListResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{9, 0}
}
func (m *ListResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse_Item.Unmarshal(m, b)
}
func (m *ListResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse_Item.Marshal(b, m, deterministic)
}
func (dst *ListResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_Item.Merge(dst, src)
}
func (m *ListResponse_Item) XXX_Size() int {
	return xxx_messageInfo_ListResponse_Item.Size(m)
}
func (m *ListResponse_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse_Item.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse_Item proto.InternalMessageInfo

func (m *ListResponse_Item) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ListResponse_Item) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *ListResponse_Item) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

type DeleteRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{10}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// DeleteResponse is a response message for the Delete rpc call
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{11}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

// IterateRequest is a request message for the Iterate rpc call
type IterateRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	First                string   `protobuf:"bytes,2,opt,name=first,proto3" json:"first,omitempty"`
	Recurse              bool     `protobuf:"varint,3,opt,name=recurse,proto3" json:"recurse,omitempty"`
	Reverse              bool     `protobuf:"varint,4,opt,name=reverse,proto3" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IterateRequest) Reset()         { *m = IterateRequest{} }
func (m *IterateRequest) String() string { return proto.CompactTextString(m) }
func (*IterateRequest) ProtoMessage()    {}
func (*IterateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pointerdb_bfbd30971eb43769, []int{12}
}
func (m *IterateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IterateRequest.Unmarshal(m, b)
}
func (m *IterateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IterateRequest.Marshal(b, m, deterministic)
}
func (dst *IterateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IterateRequest.Merge(dst, src)
}
func (m *IterateRequest) XXX_Size() int {
	return xxx_messageInfo_IterateRequest.Size(m)
}
func (m *IterateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IterateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IterateRequest proto.InternalMessageInfo

func (m *IterateRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *IterateRequest) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

func (m *IterateRequest) GetRecurse() bool {
	if m != nil {
		return m.Recurse
	}
	return false
}

func (m *IterateRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func init() {
	proto.RegisterType((*RedundancyScheme)(nil), "pointerdb.RedundancyScheme")
	proto.RegisterType((*RemotePiece)(nil), "pointerdb.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "pointerdb.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "pointerdb.Pointer")
	proto.RegisterType((*PutRequest)(nil), "pointerdb.PutRequest")
	proto.RegisterType((*GetRequest)(nil), "pointerdb.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "pointerdb.ListRequest")
	proto.RegisterType((*PutResponse)(nil), "pointerdb.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "pointerdb.GetResponse")
	proto.RegisterType((*ListResponse)(nil), "pointerdb.ListResponse")
	proto.RegisterType((*ListResponse_Item)(nil), "pointerdb.ListResponse.Item")
	proto.RegisterType((*DeleteRequest)(nil), "pointerdb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pointerdb.DeleteResponse")
	proto.RegisterType((*IterateRequest)(nil), "pointerdb.IterateRequest")
	proto.RegisterEnum("pointerdb.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("pointerdb.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PointerDB service

type PointerDBClient interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type pointerDBClient struct {
	cc *grpc.ClientConn
}

func NewPointerDBClient(cc *grpc.ClientConn) PointerDBClient {
	return &pointerDBClient{cc}
}

func (c *pointerDBClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointerDBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PointerDB service

type PointerDBServer interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterPointerDBServer(s *grpc.Server, srv PointerDBServer) {
	s.RegisterService(&_PointerDB_serviceDesc, srv)
}

func _PointerDB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PointerDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pointerdb.PointerDB",
	HandlerType: (*PointerDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _PointerDB_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PointerDB_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PointerDB_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PointerDB_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pointerdb.proto",
}

func init() { proto.RegisterFile("pointerdb.proto", fileDescriptor_pointerdb_bfbd30971eb43769) }

var fileDescriptor_pointerdb_bfbd30971eb43769 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x8e, 0xff, 0xed, 0x1a, 0xdb, 0x31, 0xad, 0xb0, 0x71, 0x9c, 0xa0, 0x5d, 0x06, 0x01, 0x0b,
	0x44, 0x5e, 0x64, 0x22, 0x21, 0x11, 0x10, 0xca, 0xb2, 0x4b, 0x64, 0x29, 0x59, 0xac, 0xf6, 0x9e,
	0xb8, 0x8c, 0xda, 0x9e, 0xb2, 0xdd, 0xc2, 0x33, 0x3d, 0xdb, 0xdd, 0x13, 0xb2, 0xfb, 0x26, 0xbc,
	0x09, 0x17, 0x8e, 0x48, 0x3c, 0x03, 0x87, 0x1c, 0x78, 0x01, 0x5e, 0x80, 0x03, 0xea, 0xee, 0x19,
	0x7b, 0x9c, 0x25, 0x9b, 0x03, 0x17, 0x7b, 0xaa, 0xea, 0xab, 0xea, 0xae, 0xfa, 0xbe, 0x6a, 0xb8,
	0x9d, 0x08, 0x1e, 0x6b, 0x94, 0xe1, 0x6c, 0x98, 0x48, 0xa1, 0x05, 0x69, 0x6d, 0x1c, 0x83, 0xfd,
	0xa5, 0x10, 0xcb, 0x35, 0x1e, 0xd9, 0xc0, 0x2c, 0x5d, 0x1c, 0x69, 0x1e, 0xa1, 0xd2, 0x2c, 0x4a,
	0x1c, 0x76, 0x00, 0x4b, 0xb1, 0x14, 0xf9, 0x77, 0x2c, 0x42, 0xcc, 0xbe, 0x7b, 0x09, 0xc7, 0x39,
	0x2a, 0x2d, 0x64, 0xe6, 0xf1, 0x7f, 0x29, 0x43, 0x8f, 0x62, 0x98, 0xc6, 0x21, 0x8b, 0xe7, 0x97,
	0xd3, 0xf9, 0x0a, 0x23, 0x24, 0x5f, 0x41, 0x55, 0x5f, 0x26, 0xd8, 0x2f, 0x1d, 0x94, 0x0e, 0xbb,
	0xa3, 0x8f, 0x86, 0xdb, 0xab, 0xbc, 0x0e, 0x1d, 0xba, 0xbf, 0xf3, 0xcb, 0x04, 0xa9, 0xcd, 0x21,
	0x77, 0xa1, 0x11, 0xf1, 0x38, 0x90, 0x78, 0xd1, 0x2f, 0x1f, 0x94, 0x0e, 0x6b, 0xb4, 0x1e, 0xf1,
	0x98, 0xe2, 0x05, 0xb9, 0x03, 0x35, 0x2d, 0x34, 0x5b, 0xf7, 0x2b, 0xd6, 0xed, 0x0c, 0xf2, 0x09,
	0xf4, 0x24, 0x26, 0x8c, 0xcb, 0x40, 0xaf, 0x24, 0xaa, 0x95, 0x58, 0x87, 0xfd, 0xaa, 0x05, 0xdc,
	0x76, 0xfe, 0xf3, 0xdc, 0x4d, 0x3e, 0x83, 0x77, 0x54, 0x3a, 0x9f, 0xa3, 0x52, 0x05, 0x6c, 0xcd,
	0x62, 0x7b, 0x59, 0x60, 0x0b, 0x7e, 0x08, 0x04, 0x25, 0x53, 0xa9, 0xc4, 0x40, 0xad, 0x98, 0xf9,
	0xe5, 0x57, 0xd8, 0xaf, 0x3b, 0x74, 0x16, 0x99, 0x9a, 0xc0, 0x94, 0x5f, 0xa1, 0x7f, 0x07, 0x60,
	0xdb, 0x08, 0xa9, 0x43, 0x99, 0x4e, 0x7b, 0xb7, 0xfc, 0x29, 0x78, 0x14, 0x23, 0xa1, 0x71, 0x62,
	0xa6, 0x46, 0xee, 0x43, 0xcb, 0x8e, 0x2f, 0x88, 0xd3, 0xc8, 0x8e, 0xa6, 0x46, 0x9b, 0xd6, 0x71,
	0x96, 0x46, 0xe4, 0x63, 0x68, 0x98, 0x39, 0x07, 0x3c, 0xb4, 0x6d, 0xb7, 0x8f, 0xbb, 0x7f, 0xbc,
	0xda, 0xbf, 0xf5, 0xe7, 0xab, 0xfd, 0xfa, 0x99, 0x08, 0x71, 0x7c, 0x42, 0xeb, 0x26, 0x3c, 0x0e,
	0xfd, 0xdf, 0x4b, 0xd0, 0x71, 0x55, 0xa7, 0xb8, 0x8c, 0x30, 0xd6, 0xe4, 0x31, 0x80, 0xdc, 0x8c,
	0xd5, 0x16, 0xf6, 0x46, 0xf7, 0x6f, 0x98, 0x39, 0x2d, 0xc0, 0xc9, 0x3d, 0x70, 0x77, 0xc8, 0x0f,
	0x6e, 0xd1, 0x86, 0xb5, 0xc7, 0x21, 0x79, 0x0c, 0x1d, 0x69, 0x0f, 0x0a, 0x1c, 0xeb, 0xfd, 0xca,
	0x41, 0xe5, 0xd0, 0x1b, 0xed, 0xed, 0x94, 0xde, 0xb4, 0x47, 0xdb, 0x72, 0x6b, 0x28, 0xb2, 0x0f,
	0x5e, 0x84, 0xf2, 0xa7, 0x35, 0x06, 0x52, 0x08, 0x6d, 0x29, 0x69, 0x53, 0x70, 0x2e, 0x2a, 0x84,
	0xf6, 0xff, 0x29, 0x43, 0x63, 0xe2, 0x0a, 0x91, 0xa3, 0x1d, 0xbd, 0x14, 0xef, 0x9e, 0x21, 0x86,
	0x27, 0x4c, 0xb3, 0x82, 0x48, 0x3e, 0x84, 0x2e, 0x8f, 0xd7, 0x3c, 0xc6, 0x40, 0xb9, 0x21, 0x58,
	0x51, 0xb4, 0x69, 0xc7, 0x79, 0xf3, 0xc9, 0x7c, 0x0e, 0x75, 0x77, 0x29, 0x7b, 0xbe, 0x37, 0xea,
	0x5f, 0xbb, 0x7a, 0x86, 0xa4, 0x19, 0x8e, 0xbc, 0x0f, 0xed, 0xac, 0xa2, 0x23, 0xdc, 0xc8, 0xa3,
	0x42, 0xbd, 0xcc, 0x67, 0xb8, 0x26, 0xdf, 0x42, 0x67, 0x2e, 0x91, 0x69, 0x2e, 0xe2, 0x20, 0x64,
	0xda, 0x89, 0xc2, 0x1b, 0x0d, 0x86, 0x6e, 0xa9, 0x86, 0xf9, 0x52, 0x0d, 0xcf, 0xf3, 0xa5, 0xa2,
	0xed, 0x3c, 0xe1, 0x84, 0x69, 0x24, 0xdf, 0xc1, 0x6d, 0x7c, 0x99, 0x70, 0x59, 0x28, 0xd1, 0x78,
	0x6b, 0x89, 0xee, 0x36, 0xc5, 0x16, 0x19, 0x40, 0x33, 0x42, 0xcd, 0x42, 0xa6, 0x59, 0xbf, 0x69,
	0x7b, 0xdf, 0xd8, 0xbe, 0x0f, 0xcd, 0x7c, 0x5e, 0x04, 0xa0, 0x3e, 0x3e, 0x7b, 0x36, 0x3e, 0x3b,
	0xed, 0xdd, 0x32, 0xdf, 0xf4, 0xf4, 0xf9, 0x0f, 0xe7, 0xa7, 0xbd, 0x92, 0x7f, 0x06, 0x30, 0x49,
	0x35, 0xc5, 0x8b, 0x14, 0x95, 0x26, 0x04, 0xaa, 0x09, 0xd3, 0x2b, 0x4b, 0x40, 0x8b, 0xda, 0x6f,
	0xf2, 0x10, 0x1a, 0xd9, 0xb4, 0xac, 0x30, 0xbc, 0x11, 0xb9, 0xce, 0x0b, 0xcd, 0x21, 0xfe, 0x01,
	0xc0, 0x53, 0xbc, 0xa9, 0x9e, 0xff, 0x6b, 0x09, 0xbc, 0x67, 0x5c, 0x6d, 0x30, 0x7b, 0x50, 0x4f,
	0x24, 0x2e, 0xf8, 0xcb, 0x0c, 0x95, 0x59, 0x46, 0x39, 0x4a, 0x33, 0xa9, 0x03, 0xb6, 0xc8, 0xcf,
	0x6e, 0x51, 0xb0, 0xae, 0x27, 0xc6, 0x43, 0xde, 0x03, 0xc0, 0x38, 0x0c, 0x66, 0xb8, 0x10, 0x12,
	0x2d, 0xf1, 0x2d, 0xda, 0xc2, 0x38, 0x3c, 0xb6, 0x0e, 0xf2, 0x00, 0x5a, 0x12, 0xe7, 0xa9, 0x54,
	0xfc, 0x85, 0xe3, 0xbd, 0x49, 0xb7, 0x0e, 0xf3, 0x8a, 0xac, 0x79, 0xc4, 0x75, 0xb6, 0xf8, 0xce,
	0x30, 0x25, 0xcd, 0xf4, 0x82, 0xc5, 0x9a, 0x2d, 0x95, 0x25, 0xb4, 0x41, 0x5b, 0xc6, 0xf3, 0xbd,
	0x71, 0xf8, 0x1d, 0xf0, 0xec, 0xb0, 0x54, 0x22, 0x62, 0x85, 0xfe, 0x5f, 0x25, 0xf0, 0x6c, 0xb3,
	0xce, 0x2e, 0x4e, 0xaa, 0xf4, 0xd6, 0x49, 0x91, 0x03, 0xa8, 0x99, 0x55, 0x56, 0xfd, 0xb2, 0x5d,
	0x27, 0x18, 0xda, 0xf7, 0xd5, 0x6c, 0x39, 0x75, 0x01, 0xf2, 0x35, 0x54, 0x92, 0x19, 0xb3, 0x9d,
	0x79, 0xa3, 0x4f, 0x87, 0xdb, 0x37, 0x57, 0x8a, 0x54, 0xa3, 0x1a, 0x4e, 0xd8, 0x25, 0xca, 0x63,
	0x16, 0x87, 0x3f, 0xf3, 0x50, 0xaf, 0x9e, 0xac, 0xd7, 0x62, 0x6e, 0x85, 0x41, 0x4d, 0x1a, 0x39,
	0x85, 0x0e, 0x4b, 0xf5, 0x4a, 0x48, 0x7e, 0x65, 0xbd, 0x99, 0xf6, 0xf7, 0xaf, 0xd7, 0x99, 0xf2,
	0x65, 0x8c, 0xe1, 0x73, 0x54, 0x8a, 0x2d, 0x91, 0xee, 0x66, 0xf9, 0xbf, 0x95, 0xa0, 0xed, 0xe8,
	0xca, 0xba, 0x1c, 0x41, 0x8d, 0x6b, 0x8c, 0x54, 0xbf, 0x64, 0xef, 0xfd, 0xa0, 0xd0, 0x63, 0x11,
	0x37, 0x1c, 0x6b, 0x8c, 0xa8, 0x83, 0x1a, 0x1d, 0x44, 0x86, 0xa4, 0xb2, 0xa5, 0xc1, 0x7e, 0x0f,
	0x10, 0xaa, 0x06, 0xf2, 0xff, 0x35, 0x67, 0x1e, 0x54, 0xae, 0x82, 0x4c, 0x44, 0x15, 0x7b, 0x44,
	0x93, 0xab, 0x89, 0xb5, 0xfd, 0x0f, 0xa0, 0x73, 0x82, 0x6b, 0xd4, 0x78, 0x93, 0x26, 0x7b, 0xd0,
	0xcd, 0x41, 0x19, 0xb7, 0x12, 0xba, 0x63, 0x8d, 0x92, 0x6d, 0xf3, 0xde, 0xa4, 0xd3, 0x3b, 0x50,
	0x5b, 0x70, 0xa9, 0x74, 0xa6, 0x50, 0x67, 0x90, 0x3e, 0x34, 0x9c, 0xd8, 0x30, 0xbb, 0x51, 0x6e,
	0xba, 0xc8, 0x0b, 0x34, 0x91, 0x6a, 0x1e, 0xb1, 0xe6, 0xe8, 0xef, 0x12, 0xb4, 0xb2, 0xe6, 0x4e,
	0x8e, 0xc9, 0x23, 0xa8, 0x4c, 0x52, 0x4d, 0xde, 0x2d, 0x76, 0xbe, 0xd9, 0xd4, 0xc1, 0xde, 0xeb,
	0xee, 0x8c, 0x9d, 0x47, 0x50, 0x79, 0x8a, 0xbb, 0x59, 0xdb, 0x7d, 0xdc, 0xc9, 0x2a, 0x2a, 0xf7,
	0x4b, 0xa8, 0x1a, 0xee, 0xc8, 0xde, 0x35, 0x32, 0x5d, 0xde, 0xdd, 0x37, 0x90, 0x4c, 0xbe, 0x81,
	0xba, 0x1b, 0x1c, 0x29, 0xbe, 0xa9, 0x3b, 0x03, 0x1f, 0xdc, 0xfb, 0x8f, 0x88, 0x4b, 0x3f, 0xae,
	0xfe, 0x58, 0x4e, 0x66, 0xb3, 0xba, 0x7d, 0xe7, 0xbe, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfa,
	0x5e, 0x11, 0x71, 0xab, 0x08, 0x00, 0x00,
}
