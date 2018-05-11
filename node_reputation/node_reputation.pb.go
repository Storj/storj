// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_reputation.proto

package nodereputation

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

type UpdateReply_ReplyType int32

const (
	UpdateReply_UPDATE_SUCCESS UpdateReply_ReplyType = 0
	UpdateReply_UPDATE_FAILED  UpdateReply_ReplyType = 1
)

var UpdateReply_ReplyType_name = map[int32]string{
	0: "UPDATE_SUCCESS",
	1: "UPDATE_FAILED",
}
var UpdateReply_ReplyType_value = map[string]int32{
	"UPDATE_SUCCESS": 0,
	"UPDATE_FAILED":  1,
}

func (x UpdateReply_ReplyType) String() string {
	return proto.EnumName(UpdateReply_ReplyType_name, int32(x))
}
func (UpdateReply_ReplyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_reputation_27bbcf2417f4c800, []int{1, 0}
}

type NodeReputationRecord struct {
	Source               string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Uptime               int64    `protobuf:"varint,4,opt,name=uptime" json:"uptime,omitempty"`
	AuditSuccess         int64    `protobuf:"varint,5,opt,name=auditSuccess" json:"auditSuccess,omitempty"`
	AuditFail            int64    `protobuf:"varint,6,opt,name=auditFail" json:"auditFail,omitempty"`
	Latency              int64    `protobuf:"varint,7,opt,name=latency" json:"latency,omitempty"`
	AmountOfDataStored   int64    `protobuf:"varint,8,opt,name=amountOfDataStored" json:"amountOfDataStored,omitempty"`
	FalseClaims          int64    `protobuf:"varint,9,opt,name=falseClaims" json:"falseClaims,omitempty"`
	ShardsModified       int64    `protobuf:"varint,10,opt,name=shardsModified" json:"shardsModified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeReputationRecord) Reset()         { *m = NodeReputationRecord{} }
func (m *NodeReputationRecord) String() string { return proto.CompactTextString(m) }
func (*NodeReputationRecord) ProtoMessage()    {}
func (*NodeReputationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_reputation_27bbcf2417f4c800, []int{0}
}
func (m *NodeReputationRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeReputationRecord.Unmarshal(m, b)
}
func (m *NodeReputationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeReputationRecord.Marshal(b, m, deterministic)
}
func (dst *NodeReputationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReputationRecord.Merge(dst, src)
}
func (m *NodeReputationRecord) XXX_Size() int {
	return xxx_messageInfo_NodeReputationRecord.Size(m)
}
func (m *NodeReputationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReputationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReputationRecord proto.InternalMessageInfo

func (m *NodeReputationRecord) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *NodeReputationRecord) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeReputationRecord) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *NodeReputationRecord) GetUptime() int64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *NodeReputationRecord) GetAuditSuccess() int64 {
	if m != nil {
		return m.AuditSuccess
	}
	return 0
}

func (m *NodeReputationRecord) GetAuditFail() int64 {
	if m != nil {
		return m.AuditFail
	}
	return 0
}

func (m *NodeReputationRecord) GetLatency() int64 {
	if m != nil {
		return m.Latency
	}
	return 0
}

func (m *NodeReputationRecord) GetAmountOfDataStored() int64 {
	if m != nil {
		return m.AmountOfDataStored
	}
	return 0
}

func (m *NodeReputationRecord) GetFalseClaims() int64 {
	if m != nil {
		return m.FalseClaims
	}
	return 0
}

func (m *NodeReputationRecord) GetShardsModified() int64 {
	if m != nil {
		return m.ShardsModified
	}
	return 0
}

type UpdateReply struct {
	BridgeName           string                `protobuf:"bytes,1,opt,name=bridgeName" json:"bridgeName,omitempty"`
	NodeName             string                `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
	Source               string                `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Status               UpdateReply_ReplyType `protobuf:"varint,4,opt,name=status,enum=nodereputation.UpdateReply_ReplyType" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateReply) Reset()         { *m = UpdateReply{} }
func (m *UpdateReply) String() string { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()    {}
func (*UpdateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_reputation_27bbcf2417f4c800, []int{1}
}
func (m *UpdateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReply.Unmarshal(m, b)
}
func (m *UpdateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReply.Marshal(b, m, deterministic)
}
func (dst *UpdateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReply.Merge(dst, src)
}
func (m *UpdateReply) XXX_Size() int {
	return xxx_messageInfo_UpdateReply.Size(m)
}
func (m *UpdateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReply proto.InternalMessageInfo

func (m *UpdateReply) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *UpdateReply) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *UpdateReply) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *UpdateReply) GetStatus() UpdateReply_ReplyType {
	if m != nil {
		return m.Status
	}
	return UpdateReply_UPDATE_SUCCESS
}

type NodeQuery struct {
	Source               string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeQuery) Reset()         { *m = NodeQuery{} }
func (m *NodeQuery) String() string { return proto.CompactTextString(m) }
func (*NodeQuery) ProtoMessage()    {}
func (*NodeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_reputation_27bbcf2417f4c800, []int{2}
}
func (m *NodeQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeQuery.Unmarshal(m, b)
}
func (m *NodeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeQuery.Marshal(b, m, deterministic)
}
func (dst *NodeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeQuery.Merge(dst, src)
}
func (m *NodeQuery) XXX_Size() int {
	return xxx_messageInfo_NodeQuery.Size(m)
}
func (m *NodeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_NodeQuery proto.InternalMessageInfo

func (m *NodeQuery) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *NodeQuery) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type NodeUpdate struct {
	Source               string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
	ColumnName           string   `protobuf:"bytes,3,opt,name=columnName" json:"columnName,omitempty"`
	ColumnValue          string   `protobuf:"bytes,4,opt,name=columnValue" json:"columnValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUpdate) Reset()         { *m = NodeUpdate{} }
func (m *NodeUpdate) String() string { return proto.CompactTextString(m) }
func (*NodeUpdate) ProtoMessage()    {}
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_reputation_27bbcf2417f4c800, []int{3}
}
func (m *NodeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUpdate.Unmarshal(m, b)
}
func (m *NodeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUpdate.Marshal(b, m, deterministic)
}
func (dst *NodeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUpdate.Merge(dst, src)
}
func (m *NodeUpdate) XXX_Size() int {
	return xxx_messageInfo_NodeUpdate.Size(m)
}
func (m *NodeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUpdate proto.InternalMessageInfo

func (m *NodeUpdate) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *NodeUpdate) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeUpdate) GetColumnName() string {
	if m != nil {
		return m.ColumnName
	}
	return ""
}

func (m *NodeUpdate) GetColumnValue() string {
	if m != nil {
		return m.ColumnValue
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeReputationRecord)(nil), "nodereputation.NodeReputationRecord")
	proto.RegisterType((*UpdateReply)(nil), "nodereputation.UpdateReply")
	proto.RegisterType((*NodeQuery)(nil), "nodereputation.NodeQuery")
	proto.RegisterType((*NodeUpdate)(nil), "nodereputation.NodeUpdate")
	proto.RegisterEnum("nodereputation.UpdateReply_ReplyType", UpdateReply_ReplyType_name, UpdateReply_ReplyType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeReputation service

type NodeReputationClient interface {
	UpdateReputation(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*UpdateReply, error)
	QueryAggregatedNodeInfo(ctx context.Context, in *NodeQuery, opts ...grpc.CallOption) (*NodeReputationRecord, error)
}

type nodeReputationClient struct {
	cc *grpc.ClientConn
}

func NewNodeReputationClient(cc *grpc.ClientConn) NodeReputationClient {
	return &nodeReputationClient{cc}
}

func (c *nodeReputationClient) UpdateReputation(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := grpc.Invoke(ctx, "/nodereputation.NodeReputation/UpdateReputation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeReputationClient) QueryAggregatedNodeInfo(ctx context.Context, in *NodeQuery, opts ...grpc.CallOption) (*NodeReputationRecord, error) {
	out := new(NodeReputationRecord)
	err := grpc.Invoke(ctx, "/nodereputation.NodeReputation/QueryAggregatedNodeInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeReputation service

type NodeReputationServer interface {
	UpdateReputation(context.Context, *NodeUpdate) (*UpdateReply, error)
	QueryAggregatedNodeInfo(context.Context, *NodeQuery) (*NodeReputationRecord, error)
}

func RegisterNodeReputationServer(s *grpc.Server, srv NodeReputationServer) {
	s.RegisterService(&_NodeReputation_serviceDesc, srv)
}

func _NodeReputation_UpdateReputation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeReputationServer).UpdateReputation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodereputation.NodeReputation/UpdateReputation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeReputationServer).UpdateReputation(ctx, req.(*NodeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeReputation_QueryAggregatedNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeReputationServer).QueryAggregatedNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodereputation.NodeReputation/QueryAggregatedNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeReputationServer).QueryAggregatedNodeInfo(ctx, req.(*NodeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeReputation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodereputation.NodeReputation",
	HandlerType: (*NodeReputationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateReputation",
			Handler:    _NodeReputation_UpdateReputation_Handler,
		},
		{
			MethodName: "QueryAggregatedNodeInfo",
			Handler:    _NodeReputation_QueryAggregatedNodeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_reputation.proto",
}

func init() {
	proto.RegisterFile("node_reputation.proto", fileDescriptor_node_reputation_27bbcf2417f4c800)
}

var fileDescriptor_node_reputation_27bbcf2417f4c800 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xad, 0xb8, 0x75, 0xa2, 0x49, 0x2b, 0xdc, 0xa1, 0x7f, 0xb6, 0x6e, 0x09, 0x46, 0xb4,
	0x25, 0x27, 0x1d, 0xdc, 0x73, 0x29, 0xc6, 0x76, 0x20, 0xd0, 0xa4, 0xad, 0x14, 0xf7, 0x54, 0x08,
	0x1b, 0xed, 0xda, 0x15, 0x48, 0x5a, 0xb1, 0x7f, 0x0e, 0xbe, 0xf6, 0xbd, 0xfa, 0x16, 0xa5, 0xcf,
	0x53, 0x76, 0xe5, 0x58, 0xb2, 0x31, 0x39, 0xf8, 0x22, 0xf8, 0x7e, 0xdf, 0x68, 0x18, 0xbe, 0x99,
	0x85, 0x17, 0xa5, 0x60, 0xfc, 0x56, 0xf2, 0xca, 0x68, 0xaa, 0x33, 0x51, 0x46, 0x95, 0x14, 0x5a,
	0x60, 0x60, 0x71, 0x43, 0xc3, 0x7f, 0x47, 0xf0, 0xfc, 0x5a, 0x30, 0x1e, 0x6f, 0x50, 0xcc, 0x53,
	0x21, 0x19, 0xbe, 0x84, 0x9e, 0x12, 0x46, 0xa6, 0x9c, 0x78, 0x43, 0xef, 0xdc, 0x8f, 0xd7, 0x0a,
	0x07, 0x70, 0x62, 0x5b, 0x5c, 0xd3, 0x82, 0x93, 0x23, 0xe7, 0x6c, 0x34, 0xbe, 0x05, 0x5f, 0x67,
	0x05, 0x57, 0x9a, 0x16, 0x15, 0xe9, 0x3a, 0xb3, 0x01, 0xb6, 0xa3, 0xa9, 0xac, 0x24, 0x8f, 0x86,
	0xde, 0x79, 0x37, 0x5e, 0x2b, 0x0c, 0xe1, 0x09, 0x35, 0x2c, 0xd3, 0x89, 0x49, 0x53, 0xae, 0x14,
	0x79, 0xec, 0xdc, 0x2d, 0x66, 0x3b, 0x3b, 0x7d, 0x41, 0xb3, 0x9c, 0xf4, 0x5c, 0x41, 0x03, 0x90,
	0xc0, 0x71, 0x4e, 0x35, 0x2f, 0xd3, 0x15, 0x39, 0x76, 0xde, 0xbd, 0xc4, 0x08, 0x90, 0x16, 0xc2,
	0x94, 0xfa, 0xeb, 0x62, 0x4a, 0x35, 0x4d, 0xb4, 0x90, 0x9c, 0x91, 0x13, 0x57, 0xb4, 0xc7, 0xc1,
	0x21, 0x9c, 0x2e, 0x68, 0xae, 0xf8, 0x24, 0xa7, 0x59, 0xa1, 0x88, 0xef, 0x0a, 0xdb, 0x08, 0x3f,
	0x40, 0xa0, 0x7e, 0x51, 0xc9, 0xd4, 0x95, 0x60, 0xd9, 0x22, 0xe3, 0x8c, 0x80, 0x2b, 0xda, 0xa1,
	0xe1, 0x5f, 0x0f, 0x4e, 0xe7, 0x15, 0xa3, 0xda, 0x46, 0x9b, 0xaf, 0xf0, 0x0c, 0xe0, 0x4e, 0x66,
	0x6c, 0x59, 0x27, 0x57, 0x67, 0xda, 0x22, 0x0f, 0xe6, 0xda, 0xec, 0xa2, 0xbb, 0xb5, 0x8b, 0x4f,
	0xd0, 0x53, 0x9a, 0x6a, 0xa3, 0x5c, 0xa2, 0xc1, 0xe8, 0x7d, 0xb4, 0xbd, 0xdd, 0xa8, 0x35, 0x40,
	0xe4, 0xbe, 0x37, 0xab, 0x8a, 0xc7, 0xeb, 0x9f, 0xc2, 0x11, 0xf8, 0x1b, 0x88, 0x08, 0xc1, 0xfc,
	0xdb, 0x74, 0x7c, 0x33, 0xbb, 0x4d, 0xe6, 0x93, 0xc9, 0x2c, 0x49, 0xfa, 0x1d, 0x7c, 0x06, 0x4f,
	0xd7, 0xec, 0x62, 0x7c, 0xf9, 0x65, 0x36, 0xed, 0x7b, 0xe1, 0x67, 0xf0, 0xed, 0xb9, 0x7c, 0x37,
	0x5c, 0xae, 0x0e, 0xb9, 0x91, 0xf0, 0xb7, 0x07, 0x60, 0x3b, 0xd4, 0xa3, 0x1d, 0x74, 0x66, 0x67,
	0x00, 0xa9, 0xc8, 0x4d, 0x51, 0x3a, 0xb7, 0x8e, 0xa4, 0x45, 0xec, 0x12, 0x6b, 0xf5, 0x83, 0xe6,
	0xa6, 0xbe, 0x36, 0x3f, 0x6e, 0xa3, 0xd1, 0x1f, 0x0f, 0x82, 0xed, 0xab, 0xc7, 0x2b, 0xe8, 0x6f,
	0xd2, 0xba, 0x67, 0x83, 0xdd, 0x3c, 0x9b, 0xc1, 0x07, 0x6f, 0x1e, 0xc8, 0x3a, 0xec, 0xe0, 0x4f,
	0x78, 0xe5, 0x32, 0x1a, 0x2f, 0x97, 0x92, 0x2f, 0xa9, 0xe6, 0xcc, 0xfe, 0x7b, 0x59, 0x2e, 0x04,
	0xbe, 0xde, 0xd7, 0xd5, 0x15, 0x0f, 0xde, 0xed, 0xb3, 0x76, 0x9f, 0x66, 0xd8, 0xb9, 0xeb, 0xb9,
	0xc7, 0xfc, 0xf1, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x6a, 0x72, 0xe4, 0xe5, 0x03, 0x00,
	0x00,
}
