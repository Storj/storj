// Code generated by protoc-gen-go. DO NOT EDIT.
// source: overlay.proto

package overlay

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

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

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP",
}
var NodeTransport_value = map[string]int32{
	"TCP": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{0}
}

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{1}
}

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}
var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}
func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{11, 0}
}

type Restriction_Operand int32

const (
	Restriction_freeBandwidth Restriction_Operand = 0
	Restriction_freeDisk      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "freeBandwidth",
	1: "freeDisk",
}
var Restriction_Operand_value = map[string]int32{
	"freeBandwidth": 0,
	"freeDisk":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}
func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{11, 1}
}

// LookupRequest is is request message for the lookup rpc call
type LookupRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (dst *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(dst, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// LookupResponse is is response message for the lookup rpc call
type LookupResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{1}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (dst *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(dst, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// FindStorageNodesResponse is is response message for the FindStorageNodes rpc call
type FindStorageNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStorageNodesResponse) Reset()         { *m = FindStorageNodesResponse{} }
func (m *FindStorageNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesResponse) ProtoMessage()    {}
func (*FindStorageNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{2}
}
func (m *FindStorageNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesResponse.Unmarshal(m, b)
}
func (m *FindStorageNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesResponse.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesResponse.Merge(dst, src)
}
func (m *FindStorageNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesResponse.Size(m)
}
func (m *FindStorageNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesResponse proto.InternalMessageInfo

func (m *FindStorageNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// FindStorageNodesRequest is is request message for the FindStorageNodes rpc call
type FindStorageNodesRequest struct {
	ObjectSize           int64              `protobuf:"varint,1,opt,name=objectSize,proto3" json:"objectSize,omitempty"`
	ContractLength       *duration.Duration `protobuf:"bytes,2,opt,name=contractLength,proto3" json:"contractLength,omitempty"`
	Opts                 *OverlayOptions    `protobuf:"bytes,3,opt,name=opts,proto3" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindStorageNodesRequest) Reset()         { *m = FindStorageNodesRequest{} }
func (m *FindStorageNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesRequest) ProtoMessage()    {}
func (*FindStorageNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{3}
}
func (m *FindStorageNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesRequest.Unmarshal(m, b)
}
func (m *FindStorageNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesRequest.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesRequest.Merge(dst, src)
}
func (m *FindStorageNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesRequest.Size(m)
}
func (m *FindStorageNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesRequest proto.InternalMessageInfo

func (m *FindStorageNodesRequest) GetObjectSize() int64 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *FindStorageNodesRequest) GetContractLength() *duration.Duration {
	if m != nil {
		return m.ContractLength
	}
	return nil
}

func (m *FindStorageNodesRequest) GetOpts() *OverlayOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=overlay.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{4}
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
	return NodeTransport_TCP
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// OverlayOptions is a set of criteria that a node must meet to be considered for a storage opportunity
type OverlayOptions struct {
	MaxLatency           *duration.Duration `protobuf:"bytes,1,opt,name=maxLatency,proto3" json:"maxLatency,omitempty"`
	MinReputation        *NodeRep           `protobuf:"bytes,2,opt,name=minReputation,proto3" json:"minReputation,omitempty"`
	MinSpeedKbps         int64              `protobuf:"varint,3,opt,name=minSpeedKbps,proto3" json:"minSpeedKbps,omitempty"`
	Amount               int64              `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Restrictions         *NodeRestrictions  `protobuf:"bytes,5,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverlayOptions) Reset()         { *m = OverlayOptions{} }
func (m *OverlayOptions) String() string { return proto.CompactTextString(m) }
func (*OverlayOptions) ProtoMessage()    {}
func (*OverlayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{5}
}
func (m *OverlayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverlayOptions.Unmarshal(m, b)
}
func (m *OverlayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverlayOptions.Marshal(b, m, deterministic)
}
func (dst *OverlayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverlayOptions.Merge(dst, src)
}
func (m *OverlayOptions) XXX_Size() int {
	return xxx_messageInfo_OverlayOptions.Size(m)
}
func (m *OverlayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_OverlayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_OverlayOptions proto.InternalMessageInfo

func (m *OverlayOptions) GetMaxLatency() *duration.Duration {
	if m != nil {
		return m.MaxLatency
	}
	return nil
}

func (m *OverlayOptions) GetMinReputation() *NodeRep {
	if m != nil {
		return m.MinReputation
	}
	return nil
}

func (m *OverlayOptions) GetMinSpeedKbps() int64 {
	if m != nil {
		return m.MinSpeedKbps
	}
	return 0
}

func (m *OverlayOptions) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OverlayOptions) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

// NodeRep is the reputation characteristics of a node
type NodeRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRep) Reset()         { *m = NodeRep{} }
func (m *NodeRep) String() string { return proto.CompactTextString(m) }
func (*NodeRep) ProtoMessage()    {}
func (*NodeRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{6}
}
func (m *NodeRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRep.Unmarshal(m, b)
}
func (m *NodeRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRep.Marshal(b, m, deterministic)
}
func (dst *NodeRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRep.Merge(dst, src)
}
func (m *NodeRep) XXX_Size() int {
	return xxx_messageInfo_NodeRep.Size(m)
}
func (m *NodeRep) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRep.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRep proto.InternalMessageInfo

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=freeBandwidth,proto3" json:"freeBandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=freeDisk,proto3" json:"freeDisk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{7}
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

// Node represents a node in the overlay network
type Node struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=overlay.NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions,proto3" json:"restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{8}
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

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

type QueryRequest struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Target               *Node    `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Pingback             bool     `protobuf:"varint,4,opt,name=pingback,proto3" json:"pingback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{9}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetPingback() bool {
	if m != nil {
		return m.Pingback
	}
	return false
}

type QueryResponse struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Response             []*Node  `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{10}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (dst *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(dst, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryResponse) GetResponse() []*Node {
	if m != nil {
		return m.Response
	}
	return nil
}

type Restriction struct {
	Operator             Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand              Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value                int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Restriction) Reset()         { *m = Restriction{} }
func (m *Restriction) String() string { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()    {}
func (*Restriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_4eb454476e2140c1, []int{11}
}
func (m *Restriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restriction.Unmarshal(m, b)
}
func (m *Restriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restriction.Marshal(b, m, deterministic)
}
func (dst *Restriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restriction.Merge(dst, src)
}
func (m *Restriction) XXX_Size() int {
	return xxx_messageInfo_Restriction.Size(m)
}
func (m *Restriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Restriction.DiscardUnknown(m)
}

var xxx_messageInfo_Restriction proto.InternalMessageInfo

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_freeBandwidth
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "overlay.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "overlay.LookupResponse")
	proto.RegisterType((*FindStorageNodesResponse)(nil), "overlay.FindStorageNodesResponse")
	proto.RegisterType((*FindStorageNodesRequest)(nil), "overlay.FindStorageNodesRequest")
	proto.RegisterType((*NodeAddress)(nil), "overlay.NodeAddress")
	proto.RegisterType((*OverlayOptions)(nil), "overlay.OverlayOptions")
	proto.RegisterType((*NodeRep)(nil), "overlay.NodeRep")
	proto.RegisterType((*NodeRestrictions)(nil), "overlay.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "overlay.Node")
	proto.RegisterType((*QueryRequest)(nil), "overlay.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "overlay.QueryResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
	proto.RegisterEnum("overlay.NodeTransport", NodeTransport_name, NodeTransport_value)
	proto.RegisterEnum("overlay.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OverlayClient is the client API for Overlay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OverlayClient interface {
	// Lookup finds a nodes address from the network
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error)
}

type overlayClient struct {
	cc *grpc.ClientConn
}

func NewOverlayClient(cc *grpc.ClientConn) OverlayClient {
	return &overlayClient{cc}
}

func (c *overlayClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error) {
	out := new(FindStorageNodesResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/FindStorageNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverlayServer is the server API for Overlay service.
type OverlayServer interface {
	// Lookup finds a nodes address from the network
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(context.Context, *FindStorageNodesRequest) (*FindStorageNodesResponse, error)
}

func RegisterOverlayServer(s *grpc.Server, srv OverlayServer) {
	s.RegisterService(&_Overlay_serviceDesc, srv)
}

func _Overlay_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_FindStorageNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).FindStorageNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/FindStorageNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).FindStorageNodes(ctx, req.(*FindStorageNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Overlay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Overlay",
	HandlerType: (*OverlayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Overlay_Lookup_Handler,
		},
		{
			MethodName: "FindStorageNodes",
			Handler:    _Overlay_FindStorageNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

// NodesClient is the client API for Nodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodesClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesServer is the server API for Nodes service.
type NodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Nodes_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_overlay_4eb454476e2140c1) }

var fileDescriptor_overlay_4eb454476e2140c1 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5f, 0x6f, 0xe2, 0x46,
	0x10, 0xc7, 0xfc, 0x33, 0x0c, 0x60, 0x39, 0xab, 0x34, 0x71, 0x51, 0x5b, 0x25, 0x6e, 0xa3, 0xb6,
	0xa9, 0xc4, 0x03, 0x89, 0x22, 0xe5, 0xa1, 0x8a, 0x68, 0xa1, 0x51, 0x54, 0x1a, 0x9a, 0xc5, 0x52,
	0x9f, 0xfa, 0x60, 0xf0, 0x86, 0xf8, 0x02, 0xbb, 0xbe, 0xf5, 0x3a, 0x77, 0xdc, 0x87, 0xb8, 0x6f,
	0x70, 0x1f, 0xe0, 0xa4, 0xfb, 0x80, 0xf7, 0x74, 0x3a, 0x79, 0xbd, 0x36, 0x18, 0x92, 0xe8, 0xee,
	0xc9, 0xcc, 0xcc, 0x6f, 0x66, 0x7f, 0xfb, 0x9b, 0xd9, 0x01, 0x5a, 0xec, 0x81, 0xf0, 0xb9, 0xbb,
	0xec, 0x04, 0x9c, 0x09, 0x86, 0x74, 0x65, 0xb6, 0x0d, 0x2f, 0xe2, 0xae, 0xf0, 0x19, 0x4d, 0x02,
	0xf6, 0xcf, 0xd0, 0x1a, 0x32, 0x76, 0x1f, 0x05, 0x98, 0xbc, 0x8c, 0x48, 0x28, 0xd0, 0x1e, 0x54,
	0x29, 0xf3, 0xc8, 0x55, 0xdf, 0xd2, 0x0e, 0xb4, 0x5f, 0xea, 0x58, 0x59, 0xf6, 0x09, 0x18, 0x29,
	0x30, 0x0c, 0x18, 0x0d, 0x09, 0x3a, 0x84, 0x72, 0x1c, 0x93, 0xb8, 0x46, 0xb7, 0xd5, 0x49, 0x4f,
	0xbc, 0x66, 0x1e, 0xc1, 0x32, 0x64, 0x5f, 0x80, 0xf5, 0x97, 0x4f, 0xbd, 0xb1, 0x60, 0xdc, 0x9d,
	0x91, 0x38, 0x10, 0x66, 0xe9, 0x3f, 0x42, 0x25, 0xc6, 0x84, 0x96, 0x76, 0x50, 0xda, 0xce, 0x4f,
	0x62, 0xf6, 0x7b, 0x0d, 0xf6, 0xb7, 0x2b, 0x24, 0x4c, 0x7f, 0x00, 0x60, 0x93, 0x17, 0x64, 0x2a,
	0xc6, 0xfe, 0x9b, 0x84, 0x45, 0x09, 0xaf, 0x79, 0x50, 0x0f, 0x8c, 0x29, 0xa3, 0x82, 0xbb, 0x53,
	0x31, 0x24, 0x74, 0x26, 0xee, 0xac, 0xa2, 0x64, 0xfa, 0x6d, 0x67, 0xc6, 0xd8, 0x6c, 0x4e, 0x12,
	0x05, 0x26, 0xd1, 0x6d, 0xa7, 0xaf, 0x34, 0xc1, 0x1b, 0x09, 0xe8, 0x37, 0x28, 0xb3, 0x40, 0x84,
	0x56, 0x49, 0x26, 0xee, 0x67, 0x14, 0x47, 0xc9, 0x77, 0x14, 0xc4, 0x59, 0x21, 0x96, 0x20, 0xfb,
	0x7f, 0x68, 0xc4, 0xfc, 0x7a, 0x9e, 0xc7, 0x49, 0x18, 0xa2, 0x53, 0xa8, 0x0b, 0xee, 0xd2, 0x30,
	0x60, 0x5c, 0x48, 0x76, 0x46, 0x77, 0x2f, 0x77, 0x47, 0x27, 0x8d, 0xe2, 0x15, 0x10, 0x59, 0xa0,
	0xbb, 0x49, 0x01, 0xc9, 0xb6, 0x8e, 0x53, 0xd3, 0xfe, 0xa4, 0x81, 0x91, 0x3f, 0x17, 0x9d, 0x03,
	0x2c, 0xdc, 0xd7, 0x43, 0x57, 0x10, 0x3a, 0x5d, 0xaa, 0x3e, 0x3c, 0x73, 0xbb, 0x35, 0x30, 0x3a,
	0x83, 0xd6, 0xc2, 0xa7, 0x98, 0x04, 0x91, 0x90, 0x41, 0xa5, 0x8d, 0x99, 0xef, 0x02, 0x09, 0x70,
	0x1e, 0x86, 0x6c, 0x68, 0x2e, 0x7c, 0x3a, 0x0e, 0x08, 0xf1, 0xfe, 0x9e, 0x04, 0x89, 0x32, 0x25,
	0x9c, 0xf3, 0xc5, 0x23, 0xe4, 0x2e, 0x58, 0x44, 0x85, 0x55, 0x96, 0x51, 0x65, 0xa1, 0xdf, 0xa1,
	0xc9, 0x49, 0x28, 0xb8, 0x3f, 0x95, 0xf4, 0xad, 0x8a, 0x22, 0x9c, 0x3f, 0x72, 0x05, 0xc0, 0x39,
	0xb8, 0x5d, 0x07, 0x5d, 0x91, 0xb2, 0x1d, 0x30, 0x37, 0xc1, 0xe8, 0x27, 0x68, 0xdd, 0x72, 0x42,
	0xfe, 0x70, 0xa9, 0xf7, 0xca, 0xf7, 0xc4, 0x9d, 0x9a, 0x88, 0xbc, 0x13, 0xb5, 0xa1, 0x16, 0x3b,
	0xfa, 0x7e, 0x78, 0x2f, 0xaf, 0x5c, 0xc2, 0x99, 0x6d, 0x7f, 0xd0, 0xa0, 0x1c, 0x97, 0x45, 0x06,
	0x14, 0x7d, 0x4f, 0xcd, 0x7f, 0xd1, 0xf7, 0x50, 0x27, 0xdf, 0x94, 0x46, 0x77, 0x37, 0xc7, 0x59,
	0x75, 0x3c, 0x6b, 0x15, 0x3a, 0x82, 0xb2, 0x58, 0x06, 0x44, 0x8a, 0x63, 0x74, 0x77, 0xf2, 0x5d,
	0x5f, 0x06, 0x04, 0xcb, 0xf0, 0x96, 0x1e, 0xe5, 0xaf, 0xd3, 0xe3, 0xad, 0x06, 0xcd, 0x9b, 0x88,
	0xf0, 0x65, 0xfa, 0x20, 0x8e, 0xa0, 0x1a, 0x12, 0xea, 0x11, 0xfe, 0xf8, 0x93, 0x54, 0xc1, 0x18,
	0x26, 0x5c, 0x3e, 0x23, 0x42, 0x5d, 0x66, 0x13, 0x96, 0x04, 0xd1, 0x2e, 0x54, 0xe6, 0xfe, 0xc2,
	0x17, 0xaa, 0xc5, 0x89, 0x11, 0xeb, 0x17, 0xf8, 0x74, 0x36, 0x71, 0xa7, 0xf7, 0x92, 0x6f, 0x0d,
	0x67, 0xb6, 0xed, 0x42, 0x4b, 0xf1, 0x51, 0x4f, 0xfc, 0x0b, 0x09, 0xfd, 0x0a, 0x35, 0xae, 0x52,
	0xac, 0xe2, 0x63, 0xcb, 0x20, 0x0b, 0xdb, 0x1f, 0x35, 0x68, 0xac, 0x49, 0x82, 0xce, 0xa1, 0xc6,
	0x02, 0xc2, 0x5d, 0xc1, 0xb8, 0x7a, 0x63, 0xdf, 0x67, 0xa9, 0x6b, 0xb8, 0xce, 0x48, 0x81, 0x70,
	0x06, 0x47, 0x67, 0xa0, 0xcb, 0xdf, 0xd4, 0x93, 0x3a, 0x18, 0xdd, 0xef, 0x9e, 0xce, 0xa4, 0x1e,
	0x4e, 0xc1, 0xb1, 0x2e, 0x0f, 0xee, 0x3c, 0x22, 0xa9, 0x2e, 0xd2, 0xb0, 0x4f, 0xa1, 0x96, 0x9e,
	0x81, 0xaa, 0x50, 0x1c, 0x3a, 0x66, 0x21, 0xfe, 0x0e, 0x6e, 0x4c, 0x2d, 0xfe, 0x5e, 0x3a, 0x66,
	0x11, 0xe9, 0x50, 0x1a, 0x3a, 0x03, 0xb3, 0x14, 0xff, 0xb8, 0x74, 0x06, 0x66, 0xd9, 0x3e, 0x06,
	0x5d, 0xd5, 0x47, 0x3b, 0x1b, 0xe3, 0x6b, 0x16, 0x50, 0x73, 0x35, 0xab, 0xa6, 0x76, 0x6c, 0x41,
	0x2b, 0xb7, 0x35, 0xe2, 0x2a, 0xce, 0x9f, 0xff, 0x9a, 0x85, 0x63, 0x1b, 0x6a, 0xe9, 0x64, 0xa1,
	0x3a, 0x54, 0x7a, 0xfd, 0x7f, 0xae, 0xae, 0xcd, 0x02, 0x6a, 0x80, 0x3e, 0x76, 0x46, 0xb8, 0x77,
	0x39, 0x30, 0xb5, 0xee, 0x3b, 0x0d, 0x74, 0xb5, 0x3d, 0xd0, 0x39, 0x54, 0x93, 0x55, 0x8e, 0x56,
	0x0b, 0x29, 0xf7, 0x27, 0xd0, 0xde, 0xdf, 0xf2, 0xab, 0x8e, 0xfe, 0x07, 0xe6, 0xe6, 0x3a, 0x46,
	0x07, 0x19, 0xf8, 0x89, 0x4d, 0xdd, 0x3e, 0x7c, 0x06, 0x91, 0x14, 0xee, 0x5e, 0x40, 0x25, 0xa9,
	0x76, 0x06, 0x15, 0x39, 0x44, 0xe8, 0x9b, 0x2c, 0x69, 0x7d, 0xc8, 0xdb, 0x7b, 0x9b, 0xee, 0xa4,
	0xc0, 0xa4, 0x2a, 0xf7, 0xdd, 0xc9, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x2d, 0x6f, 0xa8,
	0xf9, 0x06, 0x00, 0x00,
}
