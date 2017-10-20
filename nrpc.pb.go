// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nrpc.proto

/*
Package nrpc is a generated protocol buffer package.

It is generated from these files:
	nrpc.proto

It has these top-level messages:
	Error
	RPCResponse
*/
package nrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_Type int32

const (
	Error_CLIENT Error_Type = 0
	Error_SERVER Error_Type = 1
)

var Error_Type_name = map[int32]string{
	0: "CLIENT",
	1: "SERVER",
}
var Error_Type_value = map[string]int32{
	"CLIENT": 0,
	"SERVER": 1,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}
func (Error_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Error struct {
	Type    Error_Type `protobuf:"varint,1,opt,name=type,enum=nrpc.Error_Type" json:"type,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_CLIENT
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RPCResponse struct {
	// Types that are valid to be assigned to Reply:
	//	*RPCResponse_Result
	//	*RPCResponse_Error
	Reply isRPCResponse_Reply `protobuf_oneof:"reply"`
}

func (m *RPCResponse) Reset()                    { *m = RPCResponse{} }
func (m *RPCResponse) String() string            { return proto.CompactTextString(m) }
func (*RPCResponse) ProtoMessage()               {}
func (*RPCResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isRPCResponse_Reply interface {
	isRPCResponse_Reply()
}

type RPCResponse_Result struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}
type RPCResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,oneof"`
}

func (*RPCResponse_Result) isRPCResponse_Reply() {}
func (*RPCResponse_Error) isRPCResponse_Reply()  {}

func (m *RPCResponse) GetReply() isRPCResponse_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *RPCResponse) GetResult() []byte {
	if x, ok := m.GetReply().(*RPCResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (m *RPCResponse) GetError() *Error {
	if x, ok := m.GetReply().(*RPCResponse_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RPCResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RPCResponse_OneofMarshaler, _RPCResponse_OneofUnmarshaler, _RPCResponse_OneofSizer, []interface{}{
		(*RPCResponse_Result)(nil),
		(*RPCResponse_Error)(nil),
	}
}

func _RPCResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RPCResponse)
	// reply
	switch x := m.Reply.(type) {
	case *RPCResponse_Result:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Result)
	case *RPCResponse_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RPCResponse.Reply has unexpected type %T", x)
	}
	return nil
}

func _RPCResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RPCResponse)
	switch tag {
	case 1: // reply.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Reply = &RPCResponse_Result{x}
		return true, err
	case 2: // reply.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Reply = &RPCResponse_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RPCResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RPCResponse)
	// reply
	switch x := m.Reply.(type) {
	case *RPCResponse_Result:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Result)))
		n += len(x.Result)
	case *RPCResponse_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_PackageSubject = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "nrpc.packageSubject",
	Tag:           "bytes,50000,opt,name=packageSubject",
	Filename:      "nrpc.proto",
}

var E_PackageSubjectParams = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         50001,
	Name:          "nrpc.packageSubjectParams",
	Tag:           "bytes,50001,rep,name=packageSubjectParams",
	Filename:      "nrpc.proto",
}

func init() {
	proto.RegisterType((*Error)(nil), "nrpc.Error")
	proto.RegisterType((*RPCResponse)(nil), "nrpc.RPCResponse")
	proto.RegisterEnum("nrpc.Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterExtension(E_PackageSubject)
	proto.RegisterExtension(E_PackageSubjectParams)
}

func init() { proto.RegisterFile("nrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0x4d, 0x4f, 0x3a, 0x31,
	0x10, 0xc6, 0xd9, 0x3f, 0x6f, 0x61, 0xf8, 0x87, 0x90, 0xc6, 0xc3, 0xfa, 0x12, 0xb3, 0x41, 0x0f,
	0x9c, 0xba, 0x09, 0xde, 0x38, 0x42, 0x96, 0x60, 0x62, 0x94, 0x14, 0xf4, 0xe0, 0xad, 0xbb, 0x8c,
	0x75, 0x75, 0xa1, 0x4d, 0xdb, 0x3d, 0xf0, 0x05, 0xf8, 0x7c, 0xfa, 0x8d, 0x4c, 0xbb, 0x92, 0xa8,
	0x31, 0xf1, 0x36, 0x33, 0xcf, 0xcc, 0xef, 0x99, 0x07, 0x60, 0xab, 0x55, 0x46, 0x95, 0x96, 0x56,
	0x92, 0x86, 0xab, 0x4f, 0x22, 0x21, 0xa5, 0x28, 0x30, 0xf6, 0xb3, 0xb4, 0x7c, 0x8a, 0xd7, 0x68,
	0x32, 0x9d, 0x2b, 0x2b, 0x75, 0xb5, 0x37, 0x10, 0xd0, 0x4c, 0xb4, 0x96, 0x9a, 0x5c, 0x42, 0xc3,
	0xee, 0x14, 0x86, 0x41, 0x14, 0x0c, 0x7b, 0xa3, 0x3e, 0xf5, 0x2c, 0x2f, 0xd1, 0xd5, 0x4e, 0x21,
	0xf3, 0x2a, 0x09, 0xa1, 0xbd, 0x41, 0x63, 0xb8, 0xc0, 0xf0, 0x5f, 0x14, 0x0c, 0x3b, 0xec, 0xd0,
	0x0e, 0xce, 0xa1, 0xe1, 0xf6, 0x08, 0x40, 0x6b, 0x7a, 0x73, 0x9d, 0xdc, 0xae, 0xfa, 0x35, 0x57,
	0x2f, 0x13, 0xf6, 0x90, 0xb0, 0x7e, 0x30, 0xb8, 0x87, 0x2e, 0x5b, 0x4c, 0x19, 0x1a, 0x25, 0xb7,
	0xc6, 0x81, 0x5a, 0x1a, 0x4d, 0x59, 0x58, 0x6f, 0xf8, 0x7f, 0x5e, 0x63, 0x9f, 0x3d, 0xb9, 0x80,
	0x26, 0x3a, 0x5b, 0x6f, 0xd0, 0x1d, 0x75, 0xbf, 0x7c, 0x32, 0xaf, 0xb1, 0x4a, 0x9b, 0xb4, 0xa1,
	0xa9, 0x51, 0x15, 0xbb, 0xf1, 0x0c, 0x7a, 0x8a, 0x67, 0xaf, 0x5c, 0xe0, 0xb2, 0x4c, 0x5f, 0x30,
	0xb3, 0xe4, 0x8c, 0x56, 0xa1, 0xe9, 0x21, 0x34, 0x9d, 0xe5, 0x05, 0xde, 0x29, 0x9b, 0xcb, 0xad,
	0x09, 0xdf, 0xf6, 0x75, 0xff, 0xf7, 0x8f, 0xab, 0x31, 0x83, 0xa3, 0xef, 0x93, 0x05, 0xd7, 0x7c,
	0x63, 0xfe, 0xa0, 0xbd, 0xef, 0xeb, 0x51, 0x7d, 0xd8, 0x61, 0xbf, 0xde, 0x4e, 0x4e, 0x1f, 0x8f,
	0x45, 0x6e, 0x9f, 0xcb, 0x94, 0x66, 0x72, 0x13, 0x6b, 0xae, 0xf2, 0x75, 0x21, 0xa5, 0x8a, 0x5d,
	0xa0, 0xb4, 0xe5, 0x81, 0x57, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x22, 0xea, 0xd2, 0xb5,
	0x01, 0x00, 0x00,
}
