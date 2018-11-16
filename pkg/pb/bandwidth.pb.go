// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bandwidth.proto

package pb

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

type AgreementsSummary struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgreementsSummary) Reset()         { *m = AgreementsSummary{} }
func (m *AgreementsSummary) String() string { return proto.CompactTextString(m) }
func (*AgreementsSummary) ProtoMessage()    {}
func (*AgreementsSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_bandwidth_c1d6f056fe8fd53b, []int{0}
}
func (m *AgreementsSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgreementsSummary.Unmarshal(m, b)
}
func (m *AgreementsSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgreementsSummary.Marshal(b, m, deterministic)
}
func (dst *AgreementsSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgreementsSummary.Merge(dst, src)
}
func (m *AgreementsSummary) XXX_Size() int {
	return xxx_messageInfo_AgreementsSummary.Size(m)
}
func (m *AgreementsSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_AgreementsSummary.DiscardUnknown(m)
}

var xxx_messageInfo_AgreementsSummary proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AgreementsSummary)(nil), "bandwidth.AgreementsSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BandwidthClient is the client API for Bandwidth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BandwidthClient interface {
	BandwidthAgreements(ctx context.Context, opts ...grpc.CallOption) (Bandwidth_BandwidthAgreementsClient, error)
}

type bandwidthClient struct {
	cc *grpc.ClientConn
}

func NewBandwidthClient(cc *grpc.ClientConn) BandwidthClient {
	return &bandwidthClient{cc}
}

func (c *bandwidthClient) BandwidthAgreements(ctx context.Context, opts ...grpc.CallOption) (Bandwidth_BandwidthAgreementsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Bandwidth_serviceDesc.Streams[0], "/bandwidth.Bandwidth/BandwidthAgreements", opts...)
	if err != nil {
		return nil, err
	}
	x := &bandwidthBandwidthAgreementsClient{stream}
	return x, nil
}

type Bandwidth_BandwidthAgreementsClient interface {
	Send(*RenterBandwidthAllocation) error
	CloseAndRecv() (*AgreementsSummary, error)
	grpc.ClientStream
}

type bandwidthBandwidthAgreementsClient struct {
	grpc.ClientStream
}

func (x *bandwidthBandwidthAgreementsClient) Send(m *RenterBandwidthAllocation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bandwidthBandwidthAgreementsClient) CloseAndRecv() (*AgreementsSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AgreementsSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BandwidthServer is the server API for Bandwidth service.
type BandwidthServer interface {
	BandwidthAgreements(Bandwidth_BandwidthAgreementsServer) error
}

func RegisterBandwidthServer(s *grpc.Server, srv BandwidthServer) {
	s.RegisterService(&_Bandwidth_serviceDesc, srv)
}

func _Bandwidth_BandwidthAgreements_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BandwidthServer).BandwidthAgreements(&bandwidthBandwidthAgreementsServer{stream})
}

type Bandwidth_BandwidthAgreementsServer interface {
	SendAndClose(*AgreementsSummary) error
	Recv() (*RenterBandwidthAllocation, error)
	grpc.ServerStream
}

type bandwidthBandwidthAgreementsServer struct {
	grpc.ServerStream
}

func (x *bandwidthBandwidthAgreementsServer) SendAndClose(m *AgreementsSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bandwidthBandwidthAgreementsServer) Recv() (*RenterBandwidthAllocation, error) {
	m := new(RenterBandwidthAllocation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Bandwidth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bandwidth.Bandwidth",
	HandlerType: (*BandwidthServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BandwidthAgreements",
			Handler:       _Bandwidth_BandwidthAgreements_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "bandwidth.proto",
}

func init() { proto.RegisterFile("bandwidth.proto", fileDescriptor_bandwidth_c1d6f056fe8fd53b) }

var fileDescriptor_bandwidth_c1d6f056fe8fd53b = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4a, 0xcc, 0x4b,
	0x29, 0xcf, 0x4c, 0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0x09, 0x14, 0x64, 0xa6, 0x26, 0xa7, 0x16, 0x97, 0xe4, 0x17, 0xa5, 0x42, 0x24, 0x95, 0x84, 0xb9,
	0x04, 0x1d, 0xd3, 0x8b, 0x52, 0x53, 0x73, 0x53, 0xf3, 0x4a, 0x8a, 0x83, 0x4b, 0x73, 0x73, 0x13,
	0x8b, 0x2a, 0x8d, 0x0a, 0xb9, 0x38, 0x9d, 0x60, 0x7a, 0x84, 0x52, 0xb8, 0x84, 0xe1, 0x1c, 0x84,
	0x52, 0x21, 0x6d, 0x3d, 0x84, 0x59, 0x45, 0xf9, 0xa5, 0x25, 0xa9, 0xc5, 0x7a, 0x41, 0xa9, 0x79,
	0x25, 0xa9, 0x45, 0x08, 0xc5, 0x39, 0x39, 0xf9, 0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x52, 0x32,
	0x7a, 0x08, 0x47, 0x61, 0x58, 0xa7, 0xc4, 0xa0, 0xc1, 0xe8, 0xc4, 0x12, 0xc5, 0x54, 0x90, 0x94,
	0xc4, 0x06, 0x76, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x47, 0x9f, 0xcd, 0x7b, 0xc4, 0x00,
	0x00, 0x00,
}
