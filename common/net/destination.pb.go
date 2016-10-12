// Code generated by protoc-gen-go.
// source: v2ray.com/core/common/net/destination.proto
// DO NOT EDIT!

package net

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Endpoint struct {
	Network Network     `protobuf:"varint,1,opt,name=network,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	Address *IPOrDomain `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port    uint32      `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Endpoint) GetAddress() *IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "v2ray.core.common.net.Endpoint")
}

func init() { proto.RegisterFile("v2ray.com/core/common/net/destination.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x46, 0x89, 0x27, 0x9e, 0x44, 0x14, 0x09, 0x08, 0xab, 0x85, 0xac, 0x36, 0x2e, 0x08, 0x09,
	0xc4, 0x46, 0xb0, 0x3b, 0xce, 0xc2, 0x46, 0x8f, 0x2d, 0xed, 0x62, 0x32, 0x45, 0x90, 0xcc, 0x2c,
	0xb3, 0x83, 0xe2, 0x0f, 0xf1, 0xff, 0x8a, 0x97, 0x5b, 0x6c, 0xee, 0xec, 0x52, 0xbc, 0xf7, 0xc8,
	0x37, 0xfa, 0xf6, 0xc3, 0x73, 0xf8, 0xb2, 0x91, 0x8a, 0x8b, 0xc4, 0xe0, 0x22, 0x95, 0x42, 0xe8,
	0x10, 0xc4, 0x25, 0x18, 0x25, 0x63, 0x90, 0x4c, 0x68, 0x07, 0x26, 0x21, 0x73, 0x36, 0xc1, 0x0c,
	0xb6, 0x82, 0x16, 0x41, 0x2e, 0x6e, 0x76, 0x37, 0x10, 0xe4, 0x93, 0xf8, 0xbd, 0xfa, 0xff, 0x81,
	0x21, 0x25, 0x86, 0x71, 0xac, 0xe0, 0xf5, 0xb7, 0xd2, 0x87, 0x8f, 0x98, 0x06, 0xca, 0x28, 0xe6,
	0x5e, 0xcf, 0x37, 0x99, 0x46, 0xb5, 0xaa, 0x3b, 0xf1, 0x97, 0x76, 0xeb, 0x3f, 0xec, 0x73, 0xa5,
	0xfa, 0x09, 0x37, 0x0f, 0x7a, 0xbe, 0xe9, 0x36, 0x7b, 0xad, 0xea, 0x8e, 0xfc, 0xd5, 0x0e, 0xf3,
	0x69, 0xf5, 0xc2, 0x4b, 0x2a, 0x21, 0x63, 0x3f, 0x19, 0xc6, 0xe8, 0xfd, 0x81, 0x58, 0x9a, 0x59,
	0xab, 0xba, 0xe3, 0x7e, 0xfd, 0x5e, 0x78, 0x7d, 0x1e, 0xa9, 0x6c, 0x8f, 0x2c, 0x4e, 0x97, 0x7f,
	0x07, 0x5b, 0xfd, 0xce, 0x78, 0x9d, 0x21, 0xc8, 0xdb, 0xc1, 0x7a, 0xd2, 0xdd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0xbf, 0x9f, 0x8e, 0x6a, 0x01, 0x00, 0x00,
}
