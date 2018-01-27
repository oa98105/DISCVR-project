// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mul.proto

/*
Package multiplier is a generated protocol buffer package.

It is generated from these files:
	mul.proto

It has these top-level messages:
	Operands
	Response
*/
package multiplier

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

type Operands struct {
	A uint64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B uint64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *Operands) Reset()                    { *m = Operands{} }
func (m *Operands) String() string            { return proto.CompactTextString(m) }
func (*Operands) ProtoMessage()               {}
func (*Operands) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Operands) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Operands) GetB() uint64 {
	if m != nil {
		return m.B
	}
	return 0
}

type Response struct {
	Result uint64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Operands)(nil), "multiplier.Operands")
	proto.RegisterType((*Response)(nil), "multiplier.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Multiplication service

type MultiplicationClient interface {
	Mul(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*Response, error)
}

type multiplicationClient struct {
	cc *grpc.ClientConn
}

func NewMultiplicationClient(cc *grpc.ClientConn) MultiplicationClient {
	return &multiplicationClient{cc}
}

func (c *multiplicationClient) Mul(ctx context.Context, in *Operands, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/multiplier.multiplication/Mul", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Multiplication service

type MultiplicationServer interface {
	Mul(context.Context, *Operands) (*Response, error)
}

func RegisterMultiplicationServer(s *grpc.Server, srv MultiplicationServer) {
	s.RegisterService(&_Multiplication_serviceDesc, srv)
}

func _Multiplication_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplicationServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiplier.multiplication/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplicationServer).Mul(ctx, req.(*Operands))
	}
	return interceptor(ctx, in, info, handler)
}

var _Multiplication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multiplier.multiplication",
	HandlerType: (*MultiplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mul",
			Handler:    _Multiplication_Mul_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mul.proto",
}

func init() { proto.RegisterFile("mul.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0xcd, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2d, 0xcd, 0x29, 0xc9, 0x2c, 0xc8, 0xc9, 0x4c,
	0x2d, 0x52, 0x52, 0xe3, 0xe2, 0xf0, 0x2f, 0x48, 0x2d, 0x4a, 0xcc, 0x4b, 0x29, 0x16, 0xe2, 0xe1,
	0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98,
	0x20, 0xbc, 0x24, 0x25, 0x25, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x62, 0x28, 0xcf, 0xc8, 0x95, 0x8b,
	0x0f, 0x66, 0x72, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x90, 0x31, 0x17, 0xb3, 0x6f, 0x69, 0x8e,
	0x90, 0x88, 0x1e, 0xc2, 0x46, 0x3d, 0x98, 0x75, 0x52, 0x28, 0xa2, 0x30, 0xc3, 0x95, 0x18, 0x92,
	0xd8, 0xc0, 0xae, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x3f, 0x49, 0xe9, 0xb2, 0x00,
	0x00, 0x00,
}
