// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request-service.proto

package requestservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Nodes struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4000569de199abc4, []int{0}
}

func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (m *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(m, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Asn                  int32    `protobuf:"varint,3,opt,name=asn,proto3" json:"asn,omitempty"`
	RouterIp             string   `protobuf:"bytes,4,opt,name=router_ip,json=routerIp,proto3" json:"router_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_4000569de199abc4, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetAsn() int32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *Node) GetRouterIp() string {
	if m != nil {
		return m.RouterIp
	}
	return ""
}

type NodeIds struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeIds) Reset()         { *m = NodeIds{} }
func (m *NodeIds) String() string { return proto.CompactTextString(m) }
func (*NodeIds) ProtoMessage()    {}
func (*NodeIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_4000569de199abc4, []int{2}
}

func (m *NodeIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeIds.Unmarshal(m, b)
}
func (m *NodeIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeIds.Marshal(b, m, deterministic)
}
func (m *NodeIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeIds.Merge(m, src)
}
func (m *NodeIds) XXX_Size() int {
	return xxx_messageInfo_NodeIds.Size(m)
}
func (m *NodeIds) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeIds.DiscardUnknown(m)
}

var xxx_messageInfo_NodeIds proto.InternalMessageInfo

func (m *NodeIds) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type IPv4Addresses struct {
	Ipv4Address          []string `protobuf:"bytes,1,rep,name=ipv4address,proto3" json:"ipv4address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPv4Addresses) Reset()         { *m = IPv4Addresses{} }
func (m *IPv4Addresses) String() string { return proto.CompactTextString(m) }
func (*IPv4Addresses) ProtoMessage()    {}
func (*IPv4Addresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_4000569de199abc4, []int{3}
}

func (m *IPv4Addresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPv4Addresses.Unmarshal(m, b)
}
func (m *IPv4Addresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPv4Addresses.Marshal(b, m, deterministic)
}
func (m *IPv4Addresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPv4Addresses.Merge(m, src)
}
func (m *IPv4Addresses) XXX_Size() int {
	return xxx_messageInfo_IPv4Addresses.Size(m)
}
func (m *IPv4Addresses) XXX_DiscardUnknown() {
	xxx_messageInfo_IPv4Addresses.DiscardUnknown(m)
}

var xxx_messageInfo_IPv4Addresses proto.InternalMessageInfo

func (m *IPv4Addresses) GetIpv4Address() []string {
	if m != nil {
		return m.Ipv4Address
	}
	return nil
}

type DataRate struct {
	DataRate             int64    `protobuf:"varint,1,opt,name=dataRate,proto3" json:"dataRate,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4address,proto3" json:"ipv4address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRate) Reset()         { *m = DataRate{} }
func (m *DataRate) String() string { return proto.CompactTextString(m) }
func (*DataRate) ProtoMessage()    {}
func (*DataRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4000569de199abc4, []int{4}
}

func (m *DataRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRate.Unmarshal(m, b)
}
func (m *DataRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRate.Marshal(b, m, deterministic)
}
func (m *DataRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRate.Merge(m, src)
}
func (m *DataRate) XXX_Size() int {
	return xxx_messageInfo_DataRate.Size(m)
}
func (m *DataRate) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRate.DiscardUnknown(m)
}

var xxx_messageInfo_DataRate proto.InternalMessageInfo

func (m *DataRate) GetDataRate() int64 {
	if m != nil {
		return m.DataRate
	}
	return 0
}

func (m *DataRate) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Nodes)(nil), "requestservice.Nodes")
	proto.RegisterType((*Node)(nil), "requestservice.Node")
	proto.RegisterType((*NodeIds)(nil), "requestservice.NodeIds")
	proto.RegisterType((*IPv4Addresses)(nil), "requestservice.IPv4Addresses")
	proto.RegisterType((*DataRate)(nil), "requestservice.DataRate")
}

func init() { proto.RegisterFile("request-service.proto", fileDescriptor_4000569de199abc4) }

var fileDescriptor_4000569de199abc4 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0xa9, 0x05, 0x29, 0x83, 0x1a, 0xb3, 0x91, 0xd8, 0x40, 0x4c, 0x9a, 0x9e, 0x88, 0x09,
	0xad, 0x02, 0x37, 0xe2, 0x01, 0x63, 0x82, 0x8d, 0x89, 0x31, 0x3d, 0xea, 0xc1, 0x2c, 0x74, 0xa2,
	0xab, 0xd8, 0xd6, 0xce, 0x82, 0xe1, 0x87, 0xfa, 0x7f, 0xcc, 0x6e, 0x8b, 0x81, 0x8a, 0xb7, 0x37,
	0xdf, 0xbe, 0x37, 0xdd, 0x9d, 0x0e, 0xb4, 0x32, 0xfc, 0x5c, 0x20, 0xc9, 0x1e, 0x61, 0xb6, 0x14,
	0x33, 0xf4, 0xd2, 0x2c, 0x91, 0x09, 0x3b, 0x2a, 0x70, 0x41, 0xdd, 0x01, 0xd4, 0xee, 0x93, 0x08,
	0x89, 0x9d, 0x43, 0x2d, 0x56, 0xc2, 0x36, 0x1c, 0xb3, 0xdb, 0xec, 0x9f, 0x78, 0xdb, 0x46, 0x4f,
	0xb9, 0xc2, 0xdc, 0xe2, 0x3e, 0x41, 0x55, 0x95, 0xec, 0x18, 0xcc, 0x77, 0x5c, 0xd9, 0x86, 0x63,
	0x74, 0x1b, 0xa1, 0x92, 0x8c, 0x41, 0x35, 0xe6, 0x1f, 0x68, 0xef, 0x69, 0xa4, 0xb5, 0x72, 0x71,
	0x8a, 0x6d, 0xd3, 0x31, 0xba, 0xb5, 0x50, 0x49, 0xd6, 0x81, 0x46, 0x96, 0x2c, 0x24, 0x66, 0xcf,
	0x22, 0xb5, 0xab, 0xda, 0x6a, 0xe5, 0x20, 0x48, 0xdd, 0x0e, 0xd4, 0x55, 0xf3, 0x20, 0x22, 0x95,
	0x14, 0x51, 0x7e, 0xa3, 0x46, 0xa8, 0xa4, 0x7b, 0x09, 0x87, 0xc1, 0xc3, 0x72, 0x38, 0x8e, 0xa2,
	0x0c, 0x89, 0x90, 0x98, 0x03, 0x4d, 0x91, 0x2e, 0x87, 0x3c, 0x07, 0x85, 0x75, 0x13, 0xb9, 0xb7,
	0x60, 0xdd, 0x70, 0xc9, 0x43, 0x2e, 0x91, 0xb5, 0xc1, 0x8a, 0x0a, 0xad, 0x6f, 0x6d, 0x86, 0xbf,
	0x75, 0xb9, 0x53, 0xfe, 0x82, 0x4d, 0xd4, 0xff, 0x36, 0x00, 0xc6, 0xa9, 0x98, 0x70, 0x89, 0x5f,
	0x7c, 0xc5, 0x46, 0x50, 0x9f, 0xa0, 0xd4, 0x83, 0x38, 0xdd, 0x35, 0xad, 0x20, 0xa2, 0x76, 0x6b,
	0xd7, 0x01, 0xb9, 0x15, 0x76, 0x05, 0x56, 0x11, 0xa6, 0xff, 0xd3, 0x3b, 0x7f, 0x82, 0x5b, 0xb9,
	0x30, 0xd8, 0x1d, 0x1c, 0x4c, 0x50, 0xae, 0xdf, 0x45, 0xec, 0xac, 0xec, 0xdc, 0x9a, 0x52, 0xdb,
	0x2e, 0x1f, 0xaf, 0x93, 0xaa, 0xd9, 0x35, 0x3e, 0xce, 0x5e, 0x84, 0x9c, 0xf3, 0xa9, 0x97, 0x90,
	0xf4, 0x66, 0xaf, 0xbe, 0x88, 0xc9, 0x7f, 0xe3, 0x73, 0x9e, 0x62, 0x9c, 0xf4, 0x78, 0x2a, 0xfc,
	0xd2, 0x2e, 0xf9, 0x7a, 0x97, 0xfe, 0xd0, 0x52, 0x3d, 0xda, 0xfe, 0xe6, 0x74, 0x5f, 0xa7, 0x06,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xf2, 0x13, 0x6c, 0x9a, 0x02, 0x00, 0x00,
}