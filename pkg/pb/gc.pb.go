// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gc.proto

package pb

import (
	context "context"
	fmt "fmt"
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

type BloomFilter struct {
	Seed                 uint32   `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	K                    uint32   `protobuf:"varint,2,opt,name=k,proto3" json:"k,omitempty"`
	BitsPerElement       uint32   `protobuf:"varint,3,opt,name=bits_per_element,json=bitsPerElement,proto3" json:"bits_per_element,omitempty"`
	Table                []byte   `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BloomFilter) Reset()         { *m = BloomFilter{} }
func (m *BloomFilter) String() string { return proto.CompactTextString(m) }
func (*BloomFilter) ProtoMessage()    {}
func (*BloomFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5502b0b1493f7734, []int{0}
}
func (m *BloomFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BloomFilter.Unmarshal(m, b)
}
func (m *BloomFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BloomFilter.Marshal(b, m, deterministic)
}
func (m *BloomFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BloomFilter.Merge(m, src)
}
func (m *BloomFilter) XXX_Size() int {
	return xxx_messageInfo_BloomFilter.Size(m)
}
func (m *BloomFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_BloomFilter.DiscardUnknown(m)
}

var xxx_messageInfo_BloomFilter proto.InternalMessageInfo

func (m *BloomFilter) GetSeed() uint32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *BloomFilter) GetK() uint32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *BloomFilter) GetBitsPerElement() uint32 {
	if m != nil {
		return m.BitsPerElement
	}
	return 0
}

func (m *BloomFilter) GetTable() []byte {
	if m != nil {
		return m.Table
	}
	return nil
}

type DeleteRequest struct {
	CreationDate         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Filter               *BloomFilter         `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5502b0b1493f7734, []int{1}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetCreationDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *DeleteRequest) GetFilter() *BloomFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5502b0b1493f7734, []int{2}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BloomFilter)(nil), "gc.BloomFilter")
	proto.RegisterType((*DeleteRequest)(nil), "gc.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "gc.DeleteResponse")
}

func init() { proto.RegisterFile("gc.proto", fileDescriptor_5502b0b1493f7734) }

var fileDescriptor_5502b0b1493f7734 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xbe, 0x7d, 0x2b, 0x74, 0xfd, 0x43, 0x6b, 0x31, 0x44, 0x59, 0xa8, 0xb2, 0x90,
	0xc9, 0x91, 0xc2, 0x07, 0x40, 0x2a, 0x01, 0x56, 0x14, 0x31, 0xb1, 0x44, 0x76, 0x7a, 0xb5, 0xa2,
	0x3a, 0x71, 0xb0, 0xaf, 0x03, 0xdf, 0x1e, 0xc5, 0x26, 0x52, 0xd9, 0xec, 0xdf, 0x9d, 0x9f, 0xe7,
	0x67, 0xb8, 0x51, 0x0d, 0x1f, 0xac, 0x21, 0xc3, 0x66, 0xaa, 0x49, 0xee, 0x95, 0x31, 0x4a, 0x63,
	0xee, 0x89, 0xbc, 0x9c, 0x72, 0x6a, 0x3b, 0x74, 0x24, 0xba, 0x21, 0x2c, 0xa5, 0x06, 0x96, 0x07,
	0x6d, 0x4c, 0xf7, 0xda, 0x6a, 0x42, 0xcb, 0x18, 0xcc, 0x1d, 0xe2, 0x31, 0x8e, 0xf6, 0x51, 0xb6,
	0xae, 0xfc, 0x99, 0xad, 0x20, 0x3a, 0xc7, 0x33, 0x0f, 0xa2, 0x33, 0xcb, 0x60, 0x2b, 0x5b, 0x72,
	0xf5, 0x80, 0xb6, 0x46, 0x8d, 0x1d, 0xf6, 0x14, 0xff, 0xf3, 0xc3, 0xcd, 0xc8, 0xdf, 0xd1, 0xbe,
	0x04, 0xca, 0xee, 0xe0, 0x3f, 0x09, 0xa9, 0x31, 0x9e, 0xef, 0xa3, 0x6c, 0x55, 0x85, 0x4b, 0xfa,
	0x0d, 0xeb, 0x12, 0x35, 0x12, 0x56, 0xf8, 0x75, 0x41, 0x47, 0xec, 0x09, 0xd6, 0x8d, 0x45, 0x41,
	0xad, 0xe9, 0xeb, 0xa3, 0x20, 0xf4, 0xdd, 0xcb, 0x22, 0xe1, 0x41, 0x9d, 0x4f, 0xea, 0xfc, 0x63,
	0x52, 0xaf, 0x56, 0xd3, 0x83, 0x52, 0x10, 0xb2, 0x07, 0x58, 0x9c, 0xbc, 0xbd, 0x97, 0x5c, 0x16,
	0xb7, 0x5c, 0x35, 0xfc, 0xea, 0x53, 0xd5, 0xef, 0x38, 0xdd, 0xc2, 0x66, 0xaa, 0x76, 0x83, 0xe9,
	0x1d, 0x16, 0x25, 0xec, 0xde, 0x84, 0x95, 0x42, 0xe1, 0xb3, 0xd1, 0x1a, 0x9b, 0x31, 0x93, 0xe5,
	0xb0, 0x08, 0x6b, 0x6c, 0x37, 0x26, 0xfd, 0xb1, 0x4d, 0xd8, 0x35, 0x0a, 0x29, 0x87, 0xf9, 0xe7,
	0x6c, 0x90, 0x72, 0xe1, 0x45, 0x1f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x3e, 0xbf, 0x18,
	0x81, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GarbageCollectionClient is the client API for GarbageCollection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GarbageCollectionClient interface {
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type garbageCollectionClient struct {
	cc *grpc.ClientConn
}

func NewGarbageCollectionClient(cc *grpc.ClientConn) GarbageCollectionClient {
	return &garbageCollectionClient{cc}
}

func (c *garbageCollectionClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/gc.GarbageCollection/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GarbageCollectionServer is the server API for GarbageCollection service.
type GarbageCollectionServer interface {
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterGarbageCollectionServer(s *grpc.Server, srv GarbageCollectionServer) {
	s.RegisterService(&_GarbageCollection_serviceDesc, srv)
}

func _GarbageCollection_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GarbageCollectionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gc.GarbageCollection/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GarbageCollectionServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GarbageCollection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gc.GarbageCollection",
	HandlerType: (*GarbageCollectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _GarbageCollection_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gc.proto",
}
