// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pointerdb.proto

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
	return fileDescriptor_75fef806d28fc810, []int{0, 0}
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
	return fileDescriptor_75fef806d28fc810, []int{3, 0}
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
	return fileDescriptor_75fef806d28fc810, []int{0}
}
func (m *RedundancyScheme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedundancyScheme.Unmarshal(m, b)
}
func (m *RedundancyScheme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedundancyScheme.Marshal(b, m, deterministic)
}
func (m *RedundancyScheme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedundancyScheme.Merge(m, src)
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
	PieceNum             int32      `protobuf:"varint,1,opt,name=piece_num,json=pieceNum,proto3" json:"piece_num,omitempty"`
	NodeId               NodeID     `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	Hash                 *PieceHash `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemotePiece) Reset()         { *m = RemotePiece{} }
func (m *RemotePiece) String() string { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()    {}
func (*RemotePiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{1}
}
func (m *RemotePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemotePiece.Unmarshal(m, b)
}
func (m *RemotePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemotePiece.Marshal(b, m, deterministic)
}
func (m *RemotePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemotePiece.Merge(m, src)
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

func (m *RemotePiece) GetHash() *PieceHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

type RemoteSegment struct {
	Redundancy           *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	RootPieceId          PieceID           `protobuf:"bytes,2,opt,name=root_piece_id,json=rootPieceId,proto3,customtype=PieceID" json:"root_piece_id"`
	RemotePieces         []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces,proto3" json:"remote_pieces,omitempty"`
	MerkleRoot           []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteSegment) Reset()         { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()    {}
func (*RemoteSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{2}
}
func (m *RemoteSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSegment.Unmarshal(m, b)
}
func (m *RemoteSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSegment.Marshal(b, m, deterministic)
}
func (m *RemoteSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSegment.Merge(m, src)
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
	Type                 Pointer_DataType     `protobuf:"varint,1,opt,name=type,proto3,enum=pointerdb.Pointer_DataType" json:"type,omitempty"`
	InlineSegment        []byte               `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote               *RemoteSegment       `protobuf:"bytes,4,opt,name=remote,proto3" json:"remote,omitempty"`
	SegmentSize          int64                `protobuf:"varint,5,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	CreationDate         *timestamp.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	ExpirationDate       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	Metadata             []byte               `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pointer) Reset()         { *m = Pointer{} }
func (m *Pointer) String() string { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()    {}
func (*Pointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{3}
}
func (m *Pointer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pointer.Unmarshal(m, b)
}
func (m *Pointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pointer.Marshal(b, m, deterministic)
}
func (m *Pointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pointer.Merge(m, src)
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
	return fileDescriptor_75fef806d28fc810, []int{4}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
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

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Items                []*ListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	More                 bool                 `protobuf:"varint,2,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{5}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
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
	Pointer              *Pointer `protobuf:"bytes,2,opt,name=pointer,proto3" json:"pointer,omitempty"`
	IsPrefix             bool     `protobuf:"varint,3,opt,name=is_prefix,json=isPrefix,proto3" json:"is_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse_Item) Reset()         { *m = ListResponse_Item{} }
func (m *ListResponse_Item) String() string { return proto.CompactTextString(m) }
func (*ListResponse_Item) ProtoMessage()    {}
func (*ListResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fef806d28fc810, []int{5, 0}
}
func (m *ListResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse_Item.Unmarshal(m, b)
}
func (m *ListResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse_Item.Marshal(b, m, deterministic)
}
func (m *ListResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_Item.Merge(m, src)
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

func init() {
	proto.RegisterEnum("pointerdb.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("pointerdb.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
	proto.RegisterType((*RedundancyScheme)(nil), "pointerdb.RedundancyScheme")
	proto.RegisterType((*RemotePiece)(nil), "pointerdb.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "pointerdb.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "pointerdb.Pointer")
	proto.RegisterType((*ListRequest)(nil), "pointerdb.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "pointerdb.ListResponse")
	proto.RegisterType((*ListResponse_Item)(nil), "pointerdb.ListResponse.Item")
}

func init() { proto.RegisterFile("pointerdb.proto", fileDescriptor_75fef806d28fc810) }

var fileDescriptor_75fef806d28fc810 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x8e, 0xff, 0xc6, 0x9e, 0x1a, 0x3b, 0xf1, 0xb6, 0x56, 0xd9, 0x91, 0x77, 0x51, 0xc2, 0x48,
	0x0b, 0x41, 0xac, 0x26, 0x68, 0xf6, 0x80, 0xc4, 0x1e, 0x10, 0xc1, 0x41, 0x58, 0x5a, 0x42, 0xd4,
	0xce, 0x89, 0xcb, 0xa8, 0xed, 0x29, 0xdb, 0x2d, 0x3c, 0xd3, 0x93, 0xee, 0x36, 0xda, 0xe4, 0x4d,
	0x78, 0x13, 0x2e, 0xdc, 0x79, 0x06, 0x0e, 0xcb, 0x8b, 0x70, 0x40, 0xdd, 0x3d, 0x63, 0x3b, 0x44,
	0xb0, 0x17, 0xbb, 0xeb, 0xab, 0xaf, 0xeb, 0xe7, 0xab, 0xea, 0x81, 0xa3, 0x52, 0xf0, 0x42, 0xa3,
	0xcc, 0x66, 0x71, 0x29, 0x85, 0x16, 0xc4, 0xdf, 0x02, 0xa3, 0x93, 0xa5, 0x10, 0xcb, 0x35, 0x9e,
	0x5b, 0xc7, 0x6c, 0xb3, 0x38, 0xd7, 0x3c, 0x47, 0xa5, 0x59, 0x5e, 0x3a, 0xee, 0x08, 0x96, 0x62,
	0x29, 0xea, 0x73, 0x21, 0x32, 0xac, 0xce, 0xc3, 0x92, 0xe3, 0x1c, 0x95, 0x16, 0xb2, 0x46, 0xfa,
	0x42, 0x66, 0x28, 0x95, 0xb3, 0xa2, 0x5f, 0x9b, 0x30, 0xa4, 0x98, 0x6d, 0x8a, 0x8c, 0x15, 0xf3,
	0xbb, 0xe9, 0x7c, 0x85, 0x39, 0x92, 0xaf, 0xa0, 0xad, 0xef, 0x4a, 0x0c, 0x1b, 0xa7, 0x8d, 0xb3,
	0xc3, 0xe4, 0x93, 0x78, 0x57, 0xd8, 0xbf, 0xa9, 0xb1, 0xfb, 0xbb, 0xb9, 0x2b, 0x91, 0xda, 0x3b,
	0xe4, 0x19, 0x74, 0x73, 0x5e, 0xa4, 0x12, 0x6f, 0xc3, 0xe6, 0x69, 0xe3, 0xac, 0x43, 0xbd, 0x9c,
	0x17, 0x14, 0x6f, 0xc9, 0x53, 0xe8, 0x68, 0xa1, 0xd9, 0x3a, 0x6c, 0x59, 0xd8, 0x19, 0xe4, 0x33,
	0x18, 0x4a, 0x2c, 0x19, 0x97, 0xa9, 0x5e, 0x49, 0x54, 0x2b, 0xb1, 0xce, 0xc2, 0xb6, 0x25, 0x1c,
	0x39, 0xfc, 0xa6, 0x86, 0xc9, 0xe7, 0xf0, 0x44, 0x6d, 0xe6, 0x73, 0x54, 0x6a, 0x8f, 0xdb, 0xb1,
	0xdc, 0x61, 0xe5, 0xd8, 0x91, 0x5f, 0x01, 0x41, 0xc9, 0xd4, 0x46, 0x62, 0xaa, 0x56, 0xcc, 0xfc,
	0xf2, 0x7b, 0x0c, 0x3d, 0xc7, 0xae, 0x3c, 0x53, 0xe3, 0x98, 0xf2, 0x7b, 0x8c, 0x9e, 0x02, 0xec,
	0x1a, 0x21, 0x1e, 0x34, 0xe9, 0x74, 0x78, 0x10, 0xdd, 0x43, 0x40, 0x31, 0x17, 0x1a, 0xaf, 0x8d,
	0x86, 0xe4, 0x39, 0xf8, 0x56, 0xcc, 0xb4, 0xd8, 0xe4, 0x56, 0x9a, 0x0e, 0xed, 0x59, 0xe0, 0x6a,
	0x93, 0x93, 0x4f, 0xa1, 0x6b, 0x54, 0x4f, 0x79, 0x66, 0xdb, 0xee, 0x5f, 0x1c, 0xfe, 0xf1, 0xfe,
	0xe4, 0xe0, 0xcf, 0xf7, 0x27, 0xde, 0x95, 0xc8, 0x70, 0x32, 0xa6, 0x9e, 0x71, 0x4f, 0x32, 0xf2,
	0x12, 0xda, 0x2b, 0xa6, 0x56, 0x56, 0x85, 0x20, 0x79, 0x12, 0x57, 0xd3, 0xb0, 0x29, 0xbe, 0x67,
	0x6a, 0x45, 0xad, 0x3b, 0xfa, 0xab, 0x01, 0x03, 0x97, 0x7c, 0x8a, 0xcb, 0x1c, 0x0b, 0x4d, 0xde,
	0x00, 0xc8, 0xad, 0xfa, 0x36, 0x7f, 0x90, 0x3c, 0xff, 0x9f, 0xd1, 0xd0, 0x3d, 0x3a, 0x79, 0x0d,
	0x03, 0x29, 0x84, 0x4e, 0x5d, 0x03, 0xdb, 0x22, 0x8f, 0xaa, 0x22, 0xbb, 0x36, 0xfd, 0x64, 0x4c,
	0x03, 0xc3, 0x72, 0x46, 0x46, 0xde, 0xc0, 0x40, 0xda, 0x12, 0xdc, 0x35, 0x15, 0xb6, 0x4e, 0x5b,
	0x67, 0x41, 0x72, 0xfc, 0x20, 0xe9, 0x56, 0x1f, 0xda, 0x97, 0x3b, 0x43, 0x91, 0x13, 0x08, 0x72,
	0x94, 0x3f, 0xaf, 0x31, 0x35, 0x21, 0xed, 0x4c, 0xfb, 0x14, 0x1c, 0x44, 0x85, 0xd0, 0xd1, 0xdf,
	0x4d, 0xe8, 0x5e, 0xbb, 0x40, 0xe4, 0xfc, 0xc1, 0xc2, 0xed, 0x77, 0x55, 0x31, 0xe2, 0x31, 0xd3,
	0x6c, 0x6f, 0xcb, 0x5e, 0xc2, 0x21, 0x2f, 0xd6, 0xbc, 0xc0, 0x54, 0x39, 0x79, 0xac, 0x9e, 0x7d,
	0x3a, 0x70, 0x68, 0xad, 0xd9, 0x17, 0xe0, 0xb9, 0xa2, 0x6c, 0xfe, 0x20, 0x09, 0x1f, 0x95, 0x5e,
	0x31, 0x69, 0xc5, 0x23, 0x1f, 0x43, 0xbf, 0x8a, 0xe8, 0x36, 0xc6, 0xec, 0x57, 0x8b, 0x06, 0x15,
	0x66, 0x96, 0x85, 0x7c, 0x0d, 0x83, 0xb9, 0x44, 0xa6, 0xb9, 0x28, 0xd2, 0x8c, 0x69, 0xb7, 0x55,
	0x41, 0x32, 0x8a, 0xdd, 0x1b, 0x8d, 0xeb, 0x37, 0x1a, 0xdf, 0xd4, 0x6f, 0x94, 0xf6, 0xeb, 0x0b,
	0x63, 0xa6, 0x91, 0x7c, 0x0b, 0x47, 0xf8, 0xae, 0xe4, 0x72, 0x2f, 0x44, 0xf7, 0x83, 0x21, 0x0e,
	0x77, 0x57, 0x6c, 0x90, 0x11, 0xf4, 0x72, 0xd4, 0x2c, 0x63, 0x9a, 0x85, 0x3d, 0xdb, 0xfb, 0xd6,
	0x8e, 0x22, 0xe8, 0xd5, 0x7a, 0x11, 0x00, 0x6f, 0x72, 0xf5, 0x76, 0x72, 0x75, 0x39, 0x3c, 0x30,
	0x67, 0x7a, 0xf9, 0xc3, 0x8f, 0x37, 0x97, 0xc3, 0x46, 0xf4, 0x5b, 0x03, 0x82, 0xb7, 0x5c, 0x69,
	0x8a, 0xb7, 0x1b, 0x54, 0x9a, 0x1c, 0x83, 0x57, 0x4a, 0x5c, 0xf0, 0x77, 0x76, 0x08, 0x3e, 0xad,
	0x2c, 0x33, 0x47, 0xa5, 0x99, 0xd4, 0x29, 0x5b, 0x68, 0x94, 0x76, 0x6f, 0x7c, 0x0a, 0x16, 0xfa,
	0xc6, 0x20, 0xe4, 0x23, 0x00, 0x2c, 0xb2, 0x74, 0x86, 0x0b, 0x21, 0xd1, 0x8e, 0xc1, 0xa7, 0x3e,
	0x16, 0xd9, 0x85, 0x05, 0xc8, 0x0b, 0xf0, 0x25, 0xce, 0x37, 0x52, 0xf1, 0x5f, 0xdc, 0x14, 0x7a,
	0x74, 0x07, 0x98, 0x8f, 0xc2, 0x9a, 0xe7, 0x5c, 0x57, 0xef, 0xd8, 0x19, 0x26, 0xa4, 0xe9, 0x25,
	0x5d, 0xac, 0xd9, 0x52, 0x59, 0x79, 0xbb, 0xd4, 0x37, 0xc8, 0x77, 0x06, 0x88, 0x7e, 0x6f, 0x40,
	0xdf, 0x95, 0xae, 0x4a, 0x51, 0x28, 0x24, 0x09, 0x74, 0xb8, 0xc6, 0x5c, 0x85, 0x0d, 0xbb, 0xa0,
	0x2f, 0xf6, 0xa6, 0xbc, 0xcf, 0x8b, 0x27, 0x1a, 0x73, 0xea, 0xa8, 0x84, 0x40, 0x3b, 0x37, 0x05,
	0x37, 0x6d, 0x49, 0xf6, 0x3c, 0x42, 0x68, 0x1b, 0x8a, 0xf1, 0x95, 0x4c, 0xaf, 0x2a, 0x25, 0xec,
	0x99, 0xbc, 0x82, 0x6e, 0x15, 0xd5, 0x5e, 0x09, 0x12, 0xf2, 0x78, 0x4b, 0x69, 0x4d, 0x31, 0xdf,
	0x0a, 0xae, 0xd2, 0x4a, 0xd0, 0x96, 0x4d, 0xd1, 0xe3, 0xea, 0xda, 0xda, 0xc9, 0x18, 0xfc, 0xea,
	0xc2, 0xf8, 0x82, 0x7c, 0x09, 0x6d, 0x53, 0x23, 0x39, 0x7e, 0x54, 0xb4, 0x9d, 0xcb, 0xe8, 0xd9,
	0x7f, 0x34, 0x73, 0xd1, 0xfe, 0xa9, 0x59, 0xce, 0x66, 0x9e, 0x5d, 0x95, 0xd7, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0xea, 0xe6, 0x76, 0x3d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PointerDBClient is the client API for PointerDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PointerDBClient interface {
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type pointerDBClient struct {
	cc *grpc.ClientConn
}

func NewPointerDBClient(cc *grpc.ClientConn) PointerDBClient {
	return &pointerDBClient{cc}
}

func (c *pointerDBClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pointerdb.PointerDB/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointerDBServer is the server API for PointerDB service.
type PointerDBServer interface {
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterPointerDBServer(s *grpc.Server, srv PointerDBServer) {
	s.RegisterService(&_PointerDB_serviceDesc, srv)
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

var _PointerDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pointerdb.PointerDB",
	HandlerType: (*PointerDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PointerDB_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pointerdb.proto",
}
