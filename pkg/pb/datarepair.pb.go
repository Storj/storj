// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datarepair.proto

package pb

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

// QueueItem
type QueueItem struct {
	Remote               *RemoteSegment `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	PointerPath          string         `protobuf:"bytes,2,opt,name=pointerPath,proto3" json:"pointerPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueueItem) Reset()         { *m = QueueItem{} }
func (m *QueueItem) String() string { return proto.CompactTextString(m) }
func (*QueueItem) ProtoMessage()    {}
func (*QueueItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_datarepair_33a28393bb467d98, []int{0}
}
func (m *QueueItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueItem.Unmarshal(m, b)
}
func (m *QueueItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueItem.Marshal(b, m, deterministic)
}
func (dst *QueueItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueItem.Merge(dst, src)
}
func (m *QueueItem) XXX_Size() int {
	return xxx_messageInfo_QueueItem.Size(m)
}
func (m *QueueItem) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueItem.DiscardUnknown(m)
}

var xxx_messageInfo_QueueItem proto.InternalMessageInfo

func (m *QueueItem) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *QueueItem) GetPointerPath() string {
	if m != nil {
		return m.PointerPath
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueItem)(nil), "pb.QueueItem")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_datarepair_33a28393bb467d98) }

var fileDescriptor_datarepair_33a28393bb467d98 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0x4a, 0x2d, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x92, 0xe2, 0x2f, 0xc8, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x49, 0x82, 0x08, 0x2a, 0x45, 0x70,
	0x71, 0x06, 0x96, 0xa6, 0x96, 0xa6, 0x7a, 0x96, 0xa4, 0xe6, 0x0a, 0x69, 0x72, 0xb1, 0x15, 0xa5,
	0xe6, 0xe6, 0x97, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xea, 0x15, 0x24, 0xe9,
	0x05, 0x81, 0x45, 0x82, 0x53, 0xd3, 0x73, 0x53, 0xf3, 0x4a, 0x82, 0xa0, 0x0a, 0x84, 0x14, 0xb8,
	0xb8, 0xa1, 0x46, 0x05, 0x24, 0x96, 0x64, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b,
	0x25, 0xb1, 0x81, 0x2d, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x95, 0xfb, 0x73, 0x6d, 0x89,
	0x00, 0x00, 0x00,
}
