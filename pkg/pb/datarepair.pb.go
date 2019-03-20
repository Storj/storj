// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datarepair.proto

package pb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// InjuredSegment is the queue item used for the data repair queue
type InjuredSegment struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LostPieces           []int32  `protobuf:"varint,2,rep,packed,name=lost_pieces,json=lostPieces" json:"lost_pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjuredSegment) Reset()         { *m = InjuredSegment{} }
func (m *InjuredSegment) String() string { return proto.CompactTextString(m) }
func (*InjuredSegment) ProtoMessage()    {}
func (*InjuredSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_datarepair_51f2d6c047f4da42, []int{0}
}
func (m *InjuredSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjuredSegment.Unmarshal(m, b)
}
func (m *InjuredSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjuredSegment.Marshal(b, m, deterministic)
}
func (dst *InjuredSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjuredSegment.Merge(dst, src)
}
func (m *InjuredSegment) XXX_Size() int {
	return xxx_messageInfo_InjuredSegment.Size(m)
}
func (m *InjuredSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_InjuredSegment.DiscardUnknown(m)
}

var xxx_messageInfo_InjuredSegment proto.InternalMessageInfo

func (m *InjuredSegment) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InjuredSegment) GetLostPieces() []int32 {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

func init() {
	proto.RegisterType((*InjuredSegment)(nil), "repair.InjuredSegment")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_datarepair_51f2d6c047f4da42) }

var fileDescriptor_datarepair_51f2d6c047f4da42 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0x4a, 0x2d, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0,
	0x94, 0x5c, 0xb9, 0xf8, 0x3c, 0xf3, 0xb2, 0x4a, 0x8b, 0x52, 0x53, 0x82, 0x53, 0xd3, 0x73, 0x53,
	0xf3, 0x4a, 0x84, 0x84, 0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x21, 0x79, 0x2e, 0xee, 0x9c, 0xfc, 0xe2, 0x92, 0xf8, 0x82, 0xcc, 0xd4, 0xe4,
	0xd4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x2e, 0x90, 0x50, 0x00, 0x58, 0xc4, 0x89,
	0x25, 0x8a, 0xa9, 0x20, 0x29, 0x89, 0x0d, 0x6c, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13,
	0xff, 0xff, 0x1e, 0x6f, 0x00, 0x00, 0x00,
}
