// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meta.proto

package streams

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetaStreamInfo struct {
	NumberOfSegments     int64    `protobuf:"varint,1,opt,name=NumberOfSegments" json:"NumberOfSegments,omitempty"`
	SegmengsSize         int64    `protobuf:"varint,2,opt,name=SegmengsSize" json:"SegmengsSize,omitempty"`
	LastSegmentSize      int64    `protobuf:"varint,3,opt,name=LastSegmentSize" json:"LastSegmentSize,omitempty"`
	MetaData             []byte   `protobuf:"bytes,4,opt,name=MetaData,proto3" json:"MetaData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaStreamInfo) Reset()         { *m = MetaStreamInfo{} }
func (m *MetaStreamInfo) String() string { return proto.CompactTextString(m) }
func (*MetaStreamInfo) ProtoMessage()    {}
func (*MetaStreamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_meta_4766b81157dce5f8, []int{0}
}
func (m *MetaStreamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaStreamInfo.Unmarshal(m, b)
}
func (m *MetaStreamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaStreamInfo.Marshal(b, m, deterministic)
}
func (dst *MetaStreamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaStreamInfo.Merge(dst, src)
}
func (m *MetaStreamInfo) XXX_Size() int {
	return xxx_messageInfo_MetaStreamInfo.Size(m)
}
func (m *MetaStreamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaStreamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MetaStreamInfo proto.InternalMessageInfo

func (m *MetaStreamInfo) GetNumberOfSegments() int64 {
	if m != nil {
		return m.NumberOfSegments
	}
	return 0
}

func (m *MetaStreamInfo) GetSegmengsSize() int64 {
	if m != nil {
		return m.SegmengsSize
	}
	return 0
}

func (m *MetaStreamInfo) GetLastSegmentSize() int64 {
	if m != nil {
		return m.LastSegmentSize
	}
	return 0
}

func (m *MetaStreamInfo) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func init() {
	proto.RegisterType((*MetaStreamInfo)(nil), "streams.MetaStreamInfo")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_meta_4766b81157dce5f8) }

var fileDescriptor_meta_4766b81157dce5f8 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2d, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x2d, 0x56,
	0x5a, 0xc6, 0xc8, 0xc5, 0xe7, 0x9b, 0x5a, 0x92, 0x18, 0x0c, 0xe6, 0x7b, 0xe6, 0xa5, 0xe5, 0x0b,
	0x69, 0x71, 0x09, 0xf8, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0xf9, 0xa7, 0x05, 0xa7, 0xa6, 0xe7, 0xa6,
	0xe6, 0x95, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x61, 0x88, 0x0b, 0x29, 0x71, 0xf1,
	0x40, 0xd8, 0xe9, 0xc5, 0xc1, 0x99, 0x55, 0xa9, 0x12, 0x4c, 0x60, 0x75, 0x28, 0x62, 0x42, 0x1a,
	0x5c, 0xfc, 0x3e, 0x89, 0xc5, 0x25, 0x50, 0x3d, 0x60, 0x65, 0xcc, 0x60, 0x65, 0xe8, 0xc2, 0x42,
	0x52, 0x5c, 0x1c, 0x20, 0xb7, 0xb8, 0x24, 0x96, 0x24, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x04,
	0xc1, 0xf9, 0x49, 0x6c, 0x60, 0x87, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5d, 0x14,
	0x34, 0xc6, 0x00, 0x00, 0x00,
}
