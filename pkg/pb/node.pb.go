// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_INVALID   NodeType = 0
	NodeType_SATELLITE NodeType = 1
	NodeType_STORAGE   NodeType = 2
	NodeType_UPLINK    NodeType = 3
	NodeType_BOOTSTRAP NodeType = 4
)

var NodeType_name = map[int32]string{
	0: "INVALID",
	1: "SATELLITE",
	2: "STORAGE",
	3: "UPLINK",
	4: "BOOTSTRAP",
}
var NodeType_value = map[string]int32{
	"INVALID":   0,
	"SATELLITE": 1,
	"STORAGE":   2,
	"UPLINK":    3,
	"BOOTSTRAP": 4,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{0}
}

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}
var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{1}
}

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=free_bandwidth,json=freeBandwidth,proto3" json:"free_bandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=free_disk,json=freeDisk,proto3" json:"free_disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{0}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (dst *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(dst, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// TODO move statdb.Update() stuff out of here
// Node represents a node in the overlay network
// Node is info for a updating a single storagenode, used in the Update rpc calls
type Node struct {
	Id                   NodeID            `protobuf:"bytes,1,opt,name=id,proto3,customtype=NodeID" json:"id"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=node.NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions" json:"restrictions,omitempty"`
	Reputation           *NodeStats        `protobuf:"bytes,5,opt,name=reputation" json:"reputation,omitempty"`
	Metadata             *NodeMetadata     `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	LatencyList          []int64           `protobuf:"varint,7,rep,packed,name=latency_list,json=latencyList" json:"latency_list,omitempty"`
	AuditSuccess         bool              `protobuf:"varint,8,opt,name=audit_success,json=auditSuccess,proto3" json:"audit_success,omitempty"`
	IsUp                 bool              `protobuf:"varint,9,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	UpdateLatency        bool              `protobuf:"varint,10,opt,name=update_latency,json=updateLatency,proto3" json:"update_latency,omitempty"`
	UpdateAuditSuccess   bool              `protobuf:"varint,11,opt,name=update_audit_success,json=updateAuditSuccess,proto3" json:"update_audit_success,omitempty"`
	UpdateUptime         bool              `protobuf:"varint,12,opt,name=update_uptime,json=updateUptime,proto3" json:"update_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_INVALID
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func (m *Node) GetReputation() *NodeStats {
	if m != nil {
		return m.Reputation
	}
	return nil
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetLatencyList() []int64 {
	if m != nil {
		return m.LatencyList
	}
	return nil
}

func (m *Node) GetAuditSuccess() bool {
	if m != nil {
		return m.AuditSuccess
	}
	return false
}

func (m *Node) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

func (m *Node) GetUpdateLatency() bool {
	if m != nil {
		return m.UpdateLatency
	}
	return false
}

func (m *Node) GetUpdateAuditSuccess() bool {
	if m != nil {
		return m.UpdateAuditSuccess
	}
	return false
}

func (m *Node) GetUpdateUptime() bool {
	if m != nil {
		return m.UpdateUptime
	}
	return false
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=node.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{2}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (dst *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(dst, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// NodeStats is the reputation characteristics of a node
type NodeStats struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	Latency_90           int64    `protobuf:"varint,2,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	AuditSuccessRatio    float64  `protobuf:"fixed64,3,opt,name=audit_success_ratio,json=auditSuccessRatio,proto3" json:"audit_success_ratio,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,4,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	AuditCount           int64    `protobuf:"varint,5,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	AuditSuccessCount    int64    `protobuf:"varint,6,opt,name=audit_success_count,json=auditSuccessCount,proto3" json:"audit_success_count,omitempty"`
	UptimeCount          int64    `protobuf:"varint,7,opt,name=uptime_count,json=uptimeCount,proto3" json:"uptime_count,omitempty"`
	UptimeSuccessCount   int64    `protobuf:"varint,8,opt,name=uptime_success_count,json=uptimeSuccessCount,proto3" json:"uptime_success_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{3}
}
func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (dst *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(dst, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetLatency_90() int64 {
	if m != nil {
		return m.Latency_90
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessRatio() float64 {
	if m != nil {
		return m.AuditSuccessRatio
	}
	return 0
}

func (m *NodeStats) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

func (m *NodeStats) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessCount() int64 {
	if m != nil {
		return m.AuditSuccessCount
	}
	return 0
}

func (m *NodeStats) GetUptimeCount() int64 {
	if m != nil {
		return m.UptimeCount
	}
	return 0
}

func (m *NodeStats) GetUptimeSuccessCount() int64 {
	if m != nil {
		return m.UptimeSuccessCount
	}
	return 0
}

type NodeMetadata struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_d0632cceb21d65c0, []int{4}
}
func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (dst *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(dst, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NodeMetadata) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeRestrictions)(nil), "node.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeAddress)(nil), "node.NodeAddress")
	proto.RegisterType((*NodeStats)(nil), "node.NodeStats")
	proto.RegisterType((*NodeMetadata)(nil), "node.NodeMetadata")
	proto.RegisterEnum("node.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("node.NodeTransport", NodeTransport_name, NodeTransport_value)
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_d0632cceb21d65c0) }

var fileDescriptor_node_d0632cceb21d65c0 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0x86, 0x49, 0x6c, 0x9c, 0x78, 0xec, 0xa4, 0x66, 0x40, 0xc8, 0x6a, 0xd5, 0x12, 0x82, 0xaa,
	0x46, 0x54, 0x4a, 0x29, 0x3d, 0x51, 0xf5, 0x92, 0x00, 0x42, 0x51, 0xdd, 0x10, 0x6d, 0x0c, 0x07,
	0x2e, 0x96, 0x89, 0xb7, 0x74, 0x45, 0x88, 0x2d, 0x7b, 0x2d, 0xc4, 0x1b, 0xf6, 0xd0, 0x27, 0xe8,
	0x81, 0x57, 0xe8, 0x2b, 0x54, 0xbb, 0xeb, 0x24, 0xb6, 0xaa, 0xde, 0xb2, 0xff, 0xff, 0x79, 0xc6,
	0x3b, 0xff, 0xc4, 0x00, 0x8b, 0x38, 0xa2, 0xfd, 0x24, 0x8d, 0x79, 0x8c, 0xba, 0xf8, 0xfd, 0x12,
	0xee, 0xe2, 0xbb, 0x58, 0x29, 0xdd, 0x6b, 0x70, 0xc6, 0x71, 0x44, 0x09, 0xcd, 0x78, 0xca, 0x66,
	0x9c, 0xc5, 0x8b, 0x0c, 0xdf, 0x42, 0xfb, 0x7b, 0x4a, 0x69, 0x70, 0x1b, 0x2e, 0xa2, 0x47, 0x16,
	0xf1, 0x1f, 0x6e, 0xad, 0x53, 0xeb, 0x69, 0xa4, 0x25, 0xd4, 0xe1, 0x52, 0xc4, 0x57, 0x60, 0x4a,
	0x2c, 0x62, 0xd9, 0xbd, 0x5b, 0x97, 0x44, 0x53, 0x08, 0x67, 0x2c, 0xbb, 0xef, 0xfe, 0xd1, 0x40,
	0x17, 0x85, 0xf1, 0x0d, 0xd4, 0x59, 0x24, 0x0b, 0xd8, 0xc3, 0xf6, 0xcf, 0xe7, 0xbd, 0x8d, 0xdf,
	0xcf, 0x7b, 0x86, 0x70, 0x46, 0x67, 0xa4, 0xce, 0x22, 0x7c, 0x0f, 0x8d, 0x30, 0x8a, 0x52, 0x9a,
	0x65, 0xb2, 0x86, 0x75, 0xbc, 0xd5, 0x97, 0x2f, 0x2c, 0x90, 0x81, 0x32, 0xc8, 0x92, 0xc0, 0x2e,
	0xe8, 0xfc, 0x29, 0xa1, 0xae, 0xd6, 0xa9, 0xf5, 0xda, 0xc7, 0xed, 0x35, 0xe9, 0x3f, 0x25, 0x94,
	0x48, 0x0f, 0x3f, 0x83, 0x9d, 0x96, 0x6e, 0xe3, 0xea, 0xb2, 0xea, 0xee, 0x9a, 0x2d, 0xdf, 0x95,
	0x54, 0x58, 0xfc, 0x00, 0x90, 0xd2, 0x24, 0xe7, 0xa1, 0x38, 0xba, 0x9b, 0xf2, 0xc9, 0x17, 0xeb,
	0x27, 0xa7, 0x3c, 0xe4, 0x19, 0x29, 0x21, 0xd8, 0x87, 0xe6, 0x03, 0xe5, 0x61, 0x14, 0xf2, 0xd0,
	0x35, 0x24, 0x8e, 0x6b, 0xfc, 0x5b, 0xe1, 0x90, 0x15, 0x83, 0xfb, 0x60, 0xcf, 0x43, 0x4e, 0x17,
	0xb3, 0xa7, 0x60, 0xce, 0x32, 0xee, 0x36, 0x3a, 0x5a, 0x4f, 0x23, 0x56, 0xa1, 0x79, 0x2c, 0xe3,
	0x78, 0x00, 0xad, 0x30, 0x8f, 0x18, 0x0f, 0xb2, 0x7c, 0x36, 0x13, 0x63, 0x69, 0x76, 0x6a, 0xbd,
	0x26, 0xb1, 0xa5, 0x38, 0x55, 0x1a, 0x6e, 0xc3, 0x26, 0xcb, 0x82, 0x3c, 0x71, 0x4d, 0x69, 0xea,
	0x2c, 0xbb, 0x4a, 0x44, 0x6e, 0x79, 0x12, 0x85, 0x9c, 0x06, 0x45, 0x3d, 0x17, 0xa4, 0xdb, 0x52,
	0xaa, 0xa7, 0x44, 0x3c, 0x82, 0x9d, 0x02, 0xab, 0xf6, 0xb1, 0x24, 0x8c, 0xca, 0x1b, 0x94, 0xbb,
	0x1d, 0x40, 0x51, 0x22, 0xc8, 0x13, 0xce, 0x1e, 0xa8, 0x6b, 0xab, 0x57, 0x52, 0xe2, 0x95, 0xd4,
	0xba, 0x37, 0x60, 0x95, 0x32, 0xc3, 0x8f, 0x60, 0xf2, 0x34, 0x5c, 0x64, 0x49, 0x9c, 0x72, 0x19,
	0x7f, 0xfb, 0x78, 0xbb, 0x94, 0xd7, 0xd2, 0x22, 0x6b, 0x0a, 0xdd, 0xea, 0x2a, 0x98, 0xab, 0xdc,
	0xbb, 0xbf, 0xea, 0x60, 0xae, 0x02, 0xc0, 0x77, 0xd0, 0x10, 0x85, 0x82, 0xff, 0xee, 0x95, 0x21,
	0xec, 0x51, 0x84, 0xaf, 0x01, 0x96, 0xd3, 0x3e, 0x39, 0x2a, 0x56, 0xd4, 0x2c, 0x94, 0x93, 0x23,
	0xec, 0xc3, 0x76, 0x65, 0x02, 0x41, 0x2a, 0x42, 0x95, 0xcb, 0x55, 0x23, 0x5b, 0xe5, 0x79, 0x13,
	0x61, 0x88, 0xf0, 0xd4, 0xfd, 0x0b, 0x50, 0x97, 0xa0, 0xa5, 0x34, 0x85, 0xec, 0x81, 0xa5, 0x4a,
	0xce, 0xe2, 0x7c, 0xc1, 0xe5, 0x06, 0x69, 0x04, 0xa4, 0x74, 0x2a, 0x94, 0x7f, 0x7b, 0x2a, 0xd0,
	0x90, 0x60, 0xa5, 0xa7, 0xe2, 0xd7, 0x3d, 0x15, 0xd8, 0x90, 0x60, 0xd1, 0x53, 0x21, 0x32, 0x4f,
	0x89, 0x54, 0x6b, 0x36, 0x25, 0x8a, 0xca, 0x2b, 0x17, 0xed, 0x7e, 0x01, 0xbb, 0xbc, 0x9f, 0xb8,
	0x03, 0x9b, 0xf4, 0x21, 0x64, 0x73, 0x39, 0x4e, 0x93, 0xa8, 0x03, 0xee, 0x82, 0xf1, 0x18, 0xce,
	0xe7, 0x94, 0x17, 0x69, 0x14, 0xa7, 0xc3, 0x31, 0x34, 0x97, 0x7f, 0x39, 0xb4, 0xa0, 0x31, 0x1a,
	0x5f, 0x0f, 0xbc, 0xd1, 0x99, 0xb3, 0x81, 0x2d, 0x30, 0xa7, 0x03, 0xff, 0xdc, 0xf3, 0x46, 0xfe,
	0xb9, 0x53, 0x13, 0xde, 0xd4, 0xbf, 0x24, 0x83, 0x8b, 0x73, 0xa7, 0x8e, 0x00, 0xc6, 0xd5, 0xc4,
	0x1b, 0x8d, 0xbf, 0x3a, 0x9a, 0xe0, 0x86, 0x97, 0x97, 0xfe, 0xd4, 0x27, 0x83, 0x89, 0xa3, 0x1f,
	0xee, 0x43, 0xab, 0xb2, 0x12, 0xe8, 0x80, 0xed, 0x9f, 0x4e, 0x02, 0xdf, 0x9b, 0x06, 0x17, 0x64,
	0x72, 0xea, 0x6c, 0x0c, 0xf5, 0x9b, 0x7a, 0x72, 0x7b, 0x6b, 0xc8, 0x4f, 0xd6, 0xa7, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x8a, 0x68, 0xa5, 0xd2, 0x04, 0x00, 0x00,
}
