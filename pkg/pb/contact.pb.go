// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contact.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type CheckinRequest struct {
	Sender               *Node         `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Capacity             *NodeCapacity `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Operator             *NodeOperator `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckinRequest) Reset()         { *m = CheckinRequest{} }
func (m *CheckinRequest) String() string { return proto.CompactTextString(m) }
func (*CheckinRequest) ProtoMessage()    {}
func (*CheckinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}
func (m *CheckinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckinRequest.Unmarshal(m, b)
}
func (m *CheckinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckinRequest.Marshal(b, m, deterministic)
}
func (m *CheckinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckinRequest.Merge(m, src)
}
func (m *CheckinRequest) XXX_Size() int {
	return xxx_messageInfo_CheckinRequest.Size(m)
}
func (m *CheckinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckinRequest proto.InternalMessageInfo

func (m *CheckinRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *CheckinRequest) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *CheckinRequest) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

type CheckinResponse struct {
	PingNodeSuccess      bool     `protobuf:"varint,1,opt,name=ping_node_success,json=pingNodeSuccess,proto3" json:"ping_node_success,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckinResponse) Reset()         { *m = CheckinResponse{} }
func (m *CheckinResponse) String() string { return proto.CompactTextString(m) }
func (*CheckinResponse) ProtoMessage()    {}
func (*CheckinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}
func (m *CheckinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckinResponse.Unmarshal(m, b)
}
func (m *CheckinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckinResponse.Marshal(b, m, deterministic)
}
func (m *CheckinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckinResponse.Merge(m, src)
}
func (m *CheckinResponse) XXX_Size() int {
	return xxx_messageInfo_CheckinResponse.Size(m)
}
func (m *CheckinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckinResponse proto.InternalMessageInfo

func (m *CheckinResponse) GetPingNodeSuccess() bool {
	if m != nil {
		return m.PingNodeSuccess
	}
	return false
}

func (m *CheckinResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type ContactPingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingRequest) Reset()         { *m = ContactPingRequest{} }
func (m *ContactPingRequest) String() string { return proto.CompactTextString(m) }
func (*ContactPingRequest) ProtoMessage()    {}
func (*ContactPingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}
func (m *ContactPingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingRequest.Unmarshal(m, b)
}
func (m *ContactPingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingRequest.Marshal(b, m, deterministic)
}
func (m *ContactPingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingRequest.Merge(m, src)
}
func (m *ContactPingRequest) XXX_Size() int {
	return xxx_messageInfo_ContactPingRequest.Size(m)
}
func (m *ContactPingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingRequest proto.InternalMessageInfo

type ContactPingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingResponse) Reset()         { *m = ContactPingResponse{} }
func (m *ContactPingResponse) String() string { return proto.CompactTextString(m) }
func (*ContactPingResponse) ProtoMessage()    {}
func (*ContactPingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}
func (m *ContactPingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingResponse.Unmarshal(m, b)
}
func (m *ContactPingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingResponse.Marshal(b, m, deterministic)
}
func (m *ContactPingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingResponse.Merge(m, src)
}
func (m *ContactPingResponse) XXX_Size() int {
	return xxx_messageInfo_ContactPingResponse.Size(m)
}
func (m *ContactPingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckinRequest)(nil), "contact.CheckinRequest")
	proto.RegisterType((*CheckinResponse)(nil), "contact.CheckinResponse")
	proto.RegisterType((*ContactPingRequest)(nil), "contact.ContactPingRequest")
	proto.RegisterType((*ContactPingResponse)(nil), "contact.ContactPingResponse")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xd5, 0x52, 0xb5, 0xe1, 0xa0, 0x54, 0x18, 0x10, 0x51, 0x60, 0x40, 0x61, 0x41, 0x0c,
	0x19, 0xca, 0xca, 0x44, 0x60, 0x04, 0x22, 0xb3, 0xb1, 0x44, 0x89, 0x73, 0x0a, 0x11, 0xc2, 0x67,
	0x6c, 0x77, 0xe0, 0x7b, 0xf0, 0x81, 0x51, 0x6c, 0x93, 0xaa, 0x94, 0x2d, 0x79, 0xef, 0xe7, 0x77,
	0xff, 0x60, 0x2e, 0x48, 0xda, 0x4a, 0xd8, 0x4c, 0x69, 0xb2, 0xc4, 0x66, 0xe1, 0x37, 0x81, 0x96,
	0x5a, 0xf2, 0x62, 0x02, 0x92, 0x1a, 0xf4, 0xdf, 0xe9, 0xf7, 0x08, 0x0e, 0xf2, 0x37, 0x14, 0xef,
	0x9d, 0xe4, 0xf8, 0xb9, 0x42, 0x63, 0x59, 0x0a, 0x53, 0x83, 0xb2, 0x41, 0x1d, 0x8f, 0x2e, 0x46,
	0x57, 0x7b, 0x4b, 0xc8, 0x1c, 0xff, 0x44, 0x0d, 0xf2, 0xe0, 0xb0, 0x0c, 0x22, 0x51, 0xa9, 0x4a,
	0x74, 0xf6, 0x2b, 0x1e, 0x3b, 0x8a, 0xad, 0xa9, 0x3c, 0x38, 0x7c, 0x60, 0x7a, 0x9e, 0x14, 0xea,
	0xca, 0x92, 0x8e, 0x77, 0xfe, 0xf2, 0xcf, 0xc1, 0xe1, 0x03, 0x93, 0xd6, 0xb0, 0x18, 0xba, 0x32,
	0x8a, 0xa4, 0x41, 0x76, 0x0d, 0x87, 0xaa, 0x93, 0x6d, 0xd9, 0x3f, 0x2b, 0xcd, 0x4a, 0x08, 0x34,
	0xc6, 0x75, 0x18, 0xf1, 0x45, 0x6f, 0xf4, 0x49, 0x2f, 0x5e, 0x66, 0x97, 0x30, 0x47, 0xad, 0x49,
	0x97, 0x1f, 0x68, 0x4c, 0xd5, 0xa2, 0xeb, 0x71, 0x97, 0xef, 0x3b, 0xf1, 0xd1, 0x6b, 0xe9, 0x31,
	0xb0, 0xdc, 0x6f, 0xa7, 0xe8, 0x64, 0x1b, 0xa6, 0x4f, 0x4f, 0xe0, 0x68, 0x43, 0xf5, 0xd5, 0x97,
	0x05, 0xcc, 0x82, 0xcc, 0x1e, 0x20, 0x2a, 0x42, 0x3d, 0x76, 0x96, 0xfd, 0xee, 0x7b, 0x3b, 0x2a,
	0x39, 0xff, 0xdf, 0x0c, 0x89, 0xf7, 0x30, 0x71, 0x11, 0xb7, 0x30, 0x0b, 0xa3, 0xb2, 0xd3, 0xf5,
	0x83, 0x8d, 0x93, 0x24, 0xf1, 0xb6, 0xe1, 0x53, 0xee, 0x26, 0xaf, 0x63, 0x55, 0xd7, 0x53, 0x77,
	0xcc, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x62, 0x02, 0xf6, 0xfe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactClient is the client API for Contact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactClient interface {
	PingNode(ctx context.Context, in *ContactPingRequest, opts ...grpc.CallOption) (*ContactPingResponse, error)
}

type contactClient struct {
	cc *grpc.ClientConn
}

func NewContactClient(cc *grpc.ClientConn) ContactClient {
	return &contactClient{cc}
}

func (c *contactClient) PingNode(ctx context.Context, in *ContactPingRequest, opts ...grpc.CallOption) (*ContactPingResponse, error) {
	out := new(ContactPingResponse)
	err := c.cc.Invoke(ctx, "/contact.Contact/PingNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServer is the server API for Contact service.
type ContactServer interface {
	PingNode(context.Context, *ContactPingRequest) (*ContactPingResponse, error)
}

func RegisterContactServer(s *grpc.Server, srv ContactServer) {
	s.RegisterService(&_Contact_serviceDesc, srv)
}

func _Contact_PingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).PingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.Contact/PingNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).PingNode(ctx, req.(*ContactPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contact_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contact.Contact",
	HandlerType: (*ContactServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingNode",
			Handler:    _Contact_PingNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	Checkin(ctx context.Context, in *CheckinRequest, opts ...grpc.CallOption) (*CheckinResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Checkin(ctx context.Context, in *CheckinRequest, opts ...grpc.CallOption) (*CheckinResponse, error) {
	out := new(CheckinResponse)
	err := c.cc.Invoke(ctx, "/contact.Node/Checkin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	Checkin(context.Context, *CheckinRequest) (*CheckinResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Checkin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Checkin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.Node/Checkin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Checkin(ctx, req.(*CheckinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contact.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checkin",
			Handler:    _Node_Checkin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}
