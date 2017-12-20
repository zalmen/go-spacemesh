// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeinfo.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NodeInfo struct {
	NodeId     []byte `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	TcpAddress string `protobuf:"bytes,2,opt,name=tcpAddress" json:"tcpAddress,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NodeInfo) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeInfo) GetTcpAddress() string {
	if m != nil {
		return m.TcpAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "pb.NodeInfo")
}

func init() { proto.RegisterFile("nodeinfo.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcb, 0x4f, 0x49,
	0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72,
	0xe2, 0xe2, 0xf0, 0xcb, 0x4f, 0x49, 0xf5, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe3, 0x62, 0x03, 0xa9,
	0xf0, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf2, 0x84, 0xe4, 0xb8, 0xb8, 0x4a,
	0x92, 0x0b, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0x90, 0x44, 0x9c, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x86, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x15, 0x66, 0x53, 0x6e, 0x66, 0x00, 0x00, 0x00,
}