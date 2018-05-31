// Code generated by protoc-gen-go. DO NOT EDIT.
// source: piece_store.proto

package piecestoreroutes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PieceStore struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Ttl                  int64    `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
	StoreOffset          int64    `protobuf:"varint,4,opt,name=storeOffset" json:"storeOffset,omitempty"`
	Content              []byte   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceStore) Reset()         { *m = PieceStore{} }
func (m *PieceStore) String() string { return proto.CompactTextString(m) }
func (*PieceStore) ProtoMessage()    {}
func (*PieceStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{0}
}
func (m *PieceStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceStore.Unmarshal(m, b)
}
func (m *PieceStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceStore.Marshal(b, m, deterministic)
}
func (dst *PieceStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceStore.Merge(dst, src)
}
func (m *PieceStore) XXX_Size() int {
	return xxx_messageInfo_PieceStore.Size(m)
}
func (m *PieceStore) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceStore.DiscardUnknown(m)
}

var xxx_messageInfo_PieceStore proto.InternalMessageInfo

func (m *PieceStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceStore) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceStore) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *PieceStore) GetStoreOffset() int64 {
	if m != nil {
		return m.StoreOffset
	}
	return 0
}

func (m *PieceStore) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceHash struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceHash) Reset()         { *m = PieceHash{} }
func (m *PieceHash) String() string { return proto.CompactTextString(m) }
func (*PieceHash) ProtoMessage()    {}
func (*PieceHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{1}
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

func (m *PieceHash) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PieceSummary struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Expiration           int64    `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceSummary) Reset()         { *m = PieceSummary{} }
func (m *PieceSummary) String() string { return proto.CompactTextString(m) }
func (*PieceSummary) ProtoMessage()    {}
func (*PieceSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{2}
}
func (m *PieceSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceSummary.Unmarshal(m, b)
}
func (m *PieceSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceSummary.Marshal(b, m, deterministic)
}
func (dst *PieceSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceSummary.Merge(dst, src)
}
func (m *PieceSummary) XXX_Size() int {
	return xxx_messageInfo_PieceSummary.Size(m)
}
func (m *PieceSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceSummary proto.InternalMessageInfo

func (m *PieceSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceSummary) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceSummary) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type PieceRetrieval struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	StoreOffset          int64    `protobuf:"varint,3,opt,name=storeOffset" json:"storeOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceRetrieval) Reset()         { *m = PieceRetrieval{} }
func (m *PieceRetrieval) String() string { return proto.CompactTextString(m) }
func (*PieceRetrieval) ProtoMessage()    {}
func (*PieceRetrieval) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{3}
}
func (m *PieceRetrieval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceRetrieval.Unmarshal(m, b)
}
func (m *PieceRetrieval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceRetrieval.Marshal(b, m, deterministic)
}
func (dst *PieceRetrieval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceRetrieval.Merge(dst, src)
}
func (m *PieceRetrieval) XXX_Size() int {
	return xxx_messageInfo_PieceRetrieval.Size(m)
}
func (m *PieceRetrieval) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceRetrieval.DiscardUnknown(m)
}

var xxx_messageInfo_PieceRetrieval proto.InternalMessageInfo

func (m *PieceRetrieval) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceRetrieval) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceRetrieval) GetStoreOffset() int64 {
	if m != nil {
		return m.StoreOffset
	}
	return 0
}

type PieceRetrievalStream struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceRetrievalStream) Reset()         { *m = PieceRetrievalStream{} }
func (m *PieceRetrievalStream) String() string { return proto.CompactTextString(m) }
func (*PieceRetrievalStream) ProtoMessage()    {}
func (*PieceRetrievalStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{4}
}
func (m *PieceRetrievalStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceRetrievalStream.Unmarshal(m, b)
}
func (m *PieceRetrievalStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceRetrievalStream.Marshal(b, m, deterministic)
}
func (dst *PieceRetrievalStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceRetrievalStream.Merge(dst, src)
}
func (m *PieceRetrievalStream) XXX_Size() int {
	return xxx_messageInfo_PieceRetrievalStream.Size(m)
}
func (m *PieceRetrievalStream) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceRetrievalStream.DiscardUnknown(m)
}

var xxx_messageInfo_PieceRetrievalStream proto.InternalMessageInfo

func (m *PieceRetrievalStream) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceRetrievalStream) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceDelete struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceDelete) Reset()         { *m = PieceDelete{} }
func (m *PieceDelete) String() string { return proto.CompactTextString(m) }
func (*PieceDelete) ProtoMessage()    {}
func (*PieceDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{5}
}
func (m *PieceDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceDelete.Unmarshal(m, b)
}
func (m *PieceDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceDelete.Marshal(b, m, deterministic)
}
func (dst *PieceDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceDelete.Merge(dst, src)
}
func (m *PieceDelete) XXX_Size() int {
	return xxx_messageInfo_PieceDelete.Size(m)
}
func (m *PieceDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceDelete.DiscardUnknown(m)
}

var xxx_messageInfo_PieceDelete proto.InternalMessageInfo

func (m *PieceDelete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PieceDeleteSummary struct {
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceDeleteSummary) Reset()         { *m = PieceDeleteSummary{} }
func (m *PieceDeleteSummary) String() string { return proto.CompactTextString(m) }
func (*PieceDeleteSummary) ProtoMessage()    {}
func (*PieceDeleteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{6}
}
func (m *PieceDeleteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceDeleteSummary.Unmarshal(m, b)
}
func (m *PieceDeleteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceDeleteSummary.Marshal(b, m, deterministic)
}
func (dst *PieceDeleteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceDeleteSummary.Merge(dst, src)
}
func (m *PieceDeleteSummary) XXX_Size() int {
	return xxx_messageInfo_PieceDeleteSummary.Size(m)
}
func (m *PieceDeleteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceDeleteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceDeleteSummary proto.InternalMessageInfo

func (m *PieceDeleteSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PieceStoreSummary struct {
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	TotalReceived        int64    `protobuf:"varint,3,opt,name=totalReceived" json:"totalReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceStoreSummary) Reset()         { *m = PieceStoreSummary{} }
func (m *PieceStoreSummary) String() string { return proto.CompactTextString(m) }
func (*PieceStoreSummary) ProtoMessage()    {}
func (*PieceStoreSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_3a09434784aedb0c, []int{7}
}
func (m *PieceStoreSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceStoreSummary.Unmarshal(m, b)
}
func (m *PieceStoreSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceStoreSummary.Marshal(b, m, deterministic)
}
func (dst *PieceStoreSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceStoreSummary.Merge(dst, src)
}
func (m *PieceStoreSummary) XXX_Size() int {
	return xxx_messageInfo_PieceStoreSummary.Size(m)
}
func (m *PieceStoreSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceStoreSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceStoreSummary proto.InternalMessageInfo

func (m *PieceStoreSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PieceStoreSummary) GetTotalReceived() int64 {
	if m != nil {
		return m.TotalReceived
	}
	return 0
}

func init() {
	proto.RegisterType((*PieceStore)(nil), "piecestoreroutes.PieceStore")
	proto.RegisterType((*PieceHash)(nil), "piecestoreroutes.PieceHash")
	proto.RegisterType((*PieceSummary)(nil), "piecestoreroutes.PieceSummary")
	proto.RegisterType((*PieceRetrieval)(nil), "piecestoreroutes.PieceRetrieval")
	proto.RegisterType((*PieceRetrievalStream)(nil), "piecestoreroutes.PieceRetrievalStream")
	proto.RegisterType((*PieceDelete)(nil), "piecestoreroutes.PieceDelete")
	proto.RegisterType((*PieceDeleteSummary)(nil), "piecestoreroutes.PieceDeleteSummary")
	proto.RegisterType((*PieceStoreSummary)(nil), "piecestoreroutes.PieceStoreSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PieceStoreRoutes service

type PieceStoreRoutesClient interface {
	Piece(ctx context.Context, in *PieceHash, opts ...grpc.CallOption) (*PieceSummary, error)
	Retrieve(ctx context.Context, in *PieceRetrieval, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error)
	Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error)
	Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error)
}

type pieceStoreRoutesClient struct {
	cc *grpc.ClientConn
}

func NewPieceStoreRoutesClient(cc *grpc.ClientConn) PieceStoreRoutesClient {
	return &pieceStoreRoutesClient{cc}
}

func (c *pieceStoreRoutesClient) Piece(ctx context.Context, in *PieceHash, opts ...grpc.CallOption) (*PieceSummary, error) {
	out := new(PieceSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Piece", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) Retrieve(ctx context.Context, in *PieceRetrieval, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[0], c.cc, "/piecestoreroutes.PieceStoreRoutes/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PieceStoreRoutes_RetrieveClient interface {
	Recv() (*PieceRetrievalStream, error)
	grpc.ClientStream
}

type pieceStoreRoutesRetrieveClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesRetrieveClient) Recv() (*PieceRetrievalStream, error) {
	m := new(PieceRetrievalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[1], c.cc, "/piecestoreroutes.PieceStoreRoutes/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesStoreClient{stream}
	return x, nil
}

type PieceStoreRoutes_StoreClient interface {
	Send(*PieceStore) error
	CloseAndRecv() (*PieceStoreSummary, error)
	grpc.ClientStream
}

type pieceStoreRoutesStoreClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesStoreClient) Send(m *PieceStore) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreClient) CloseAndRecv() (*PieceStoreSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PieceStoreSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error) {
	out := new(PieceDeleteSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PieceStoreRoutes service

type PieceStoreRoutesServer interface {
	Piece(context.Context, *PieceHash) (*PieceSummary, error)
	Retrieve(*PieceRetrieval, PieceStoreRoutes_RetrieveServer) error
	Store(PieceStoreRoutes_StoreServer) error
	Delete(context.Context, *PieceDelete) (*PieceDeleteSummary, error)
}

func RegisterPieceStoreRoutesServer(s *grpc.Server, srv PieceStoreRoutesServer) {
	s.RegisterService(&_PieceStoreRoutes_serviceDesc, srv)
}

func _PieceStoreRoutes_Piece_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Piece(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Piece",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Piece(ctx, req.(*PieceHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PieceRetrieval)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PieceStoreRoutesServer).Retrieve(m, &pieceStoreRoutesRetrieveServer{stream})
}

type PieceStoreRoutes_RetrieveServer interface {
	Send(*PieceRetrievalStream) error
	grpc.ServerStream
}

type pieceStoreRoutesRetrieveServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesRetrieveServer) Send(m *PieceRetrievalStream) error {
	return x.ServerStream.SendMsg(m)
}

func _PieceStoreRoutes_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PieceStoreRoutesServer).Store(&pieceStoreRoutesStoreServer{stream})
}

type PieceStoreRoutes_StoreServer interface {
	SendAndClose(*PieceStoreSummary) error
	Recv() (*PieceStore, error)
	grpc.ServerStream
}

type pieceStoreRoutesStoreServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesStoreServer) SendAndClose(m *PieceStoreSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreServer) Recv() (*PieceStore, error) {
	m := new(PieceStore)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PieceStoreRoutes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Delete(ctx, req.(*PieceDelete))
	}
	return interceptor(ctx, in, info, handler)
}

var _PieceStoreRoutes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "piecestoreroutes.PieceStoreRoutes",
	HandlerType: (*PieceStoreRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Piece",
			Handler:    _PieceStoreRoutes_Piece_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PieceStoreRoutes_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Retrieve",
			Handler:       _PieceStoreRoutes_Retrieve_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Store",
			Handler:       _PieceStoreRoutes_Store_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "piece_store.proto",
}

func init() { proto.RegisterFile("piece_store.proto", fileDescriptor_piece_store_3a09434784aedb0c) }

var fileDescriptor_piece_store_3a09434784aedb0c = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6f, 0xaa, 0x40,
	0x10, 0x15, 0xf0, 0xe3, 0x3a, 0x7a, 0x8d, 0x4e, 0xee, 0x03, 0xc1, 0xab, 0x21, 0x7b, 0xcd, 0x0d,
	0x4f, 0xa4, 0x69, 0xff, 0x82, 0x69, 0xfa, 0x54, 0x9b, 0x35, 0x69, 0xfa, 0xd6, 0x50, 0x1d, 0xdb,
	0x4d, 0x40, 0x0c, 0xac, 0xa6, 0xed, 0x43, 0xff, 0x65, 0xff, 0x4f, 0xc3, 0x82, 0x88, 0x1a, 0xd4,
	0xb7, 0xdd, 0x73, 0x86, 0xb3, 0x67, 0xe6, 0x0c, 0xd0, 0x5b, 0x09, 0x9a, 0xd1, 0x73, 0x2c, 0xc3,
	0x88, 0xdc, 0x55, 0x14, 0xca, 0x10, 0xbb, 0x0a, 0x52, 0x48, 0x14, 0xae, 0x25, 0xc5, 0xec, 0x0b,
	0xe0, 0x21, 0xc1, 0xa6, 0x09, 0x86, 0x1d, 0xd0, 0xc5, 0xdc, 0xd4, 0x6c, 0xcd, 0x69, 0x72, 0x5d,
	0xcc, 0x11, 0xa1, 0x1a, 0x8b, 0x4f, 0x32, 0x75, 0x5b, 0x73, 0x0c, 0xae, 0xce, 0xd8, 0x05, 0x43,
	0x4a, 0xdf, 0x34, 0x14, 0x94, 0x1c, 0xd1, 0x86, 0x96, 0x92, 0x9c, 0x2c, 0x16, 0x31, 0x49, 0xb3,
	0xaa, 0x98, 0x22, 0x84, 0x26, 0x34, 0x66, 0xe1, 0x52, 0xd2, 0x52, 0x9a, 0x35, 0x5b, 0x73, 0xda,
	0x7c, 0x7b, 0x65, 0x7d, 0x68, 0xaa, 0xf7, 0xef, 0xbc, 0xf8, 0xed, 0xf0, 0x79, 0xc6, 0xa1, 0x9d,
	0x9a, 0x5b, 0x07, 0x81, 0x17, 0x7d, 0x5c, 0x64, 0x6f, 0x08, 0x40, 0xef, 0x2b, 0x11, 0x79, 0x52,
	0x84, 0xcb, 0xcc, 0x65, 0x01, 0x61, 0x8f, 0xd0, 0x51, 0x9a, 0x9c, 0x64, 0x24, 0x68, 0xe3, 0xf9,
	0x17, 0xa9, 0x1e, 0xb4, 0x68, 0x1c, 0xb5, 0xc8, 0xc6, 0xf0, 0x67, 0x5f, 0x77, 0x2a, 0x23, 0xf2,
	0x82, 0x5c, 0x4d, 0x2b, 0xa8, 0x15, 0xc6, 0xa1, 0xef, 0x8f, 0x63, 0x00, 0x2d, 0xa5, 0x32, 0x26,
	0x9f, 0xe4, 0x51, 0x1e, 0xcc, 0x05, 0x2c, 0xd0, 0xdb, 0xb1, 0x98, 0xd0, 0x08, 0x28, 0x8e, 0xbd,
	0xd7, 0xd4, 0x73, 0x93, 0x6f, 0xaf, 0x6c, 0x0a, 0xbd, 0x5d, 0xba, 0x67, 0xcb, 0x71, 0x04, 0xbf,
	0x65, 0x28, 0x3d, 0x9f, 0xd3, 0x8c, 0xc4, 0x86, 0xe6, 0x59, 0x9f, 0xfb, 0xe0, 0xf5, 0xb7, 0x0e,
	0xdd, 0x9d, 0x2a, 0x57, 0x7b, 0x84, 0xb7, 0x50, 0x53, 0x18, 0xf6, 0xdd, 0xc3, 0x1d, 0x73, 0xf3,
	0x80, 0xad, 0x61, 0x09, 0x99, 0x59, 0x63, 0x15, 0x7c, 0x82, 0x5f, 0xd9, 0x04, 0x09, 0xed, 0x92,
	0xea, 0x7c, 0xc4, 0xd6, 0xff, 0x73, 0x15, 0x69, 0x08, 0xac, 0x72, 0xa5, 0xe1, 0x3d, 0xd4, 0xd2,
	0x25, 0xff, 0x5b, 0x66, 0x22, 0x01, 0xac, 0x7f, 0xa7, 0xd8, 0xdc, 0xa7, 0xa3, 0xe1, 0x04, 0xea,
	0x59, 0x4a, 0x83, 0x92, 0x4f, 0x52, 0xda, 0x1a, 0x9d, 0xa4, 0x73, 0xc9, 0x97, 0xba, 0xfa, 0x47,
	0x6f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x00, 0xce, 0x39, 0xb8, 0x03, 0x00, 0x00,
}
