// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto

package envoy_config_resource_monitor_fixed_heap_v2alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FixedHeapConfig struct {
	MaxHeapSizeBytes     uint64   `protobuf:"varint,1,opt,name=max_heap_size_bytes,json=maxHeapSizeBytes,proto3" json:"max_heap_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FixedHeapConfig) Reset()         { *m = FixedHeapConfig{} }
func (m *FixedHeapConfig) String() string { return proto.CompactTextString(m) }
func (*FixedHeapConfig) ProtoMessage()    {}
func (*FixedHeapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ee15b3c15e2df, []int{0}
}

func (m *FixedHeapConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedHeapConfig.Unmarshal(m, b)
}
func (m *FixedHeapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedHeapConfig.Marshal(b, m, deterministic)
}
func (m *FixedHeapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedHeapConfig.Merge(m, src)
}
func (m *FixedHeapConfig) XXX_Size() int {
	return xxx_messageInfo_FixedHeapConfig.Size(m)
}
func (m *FixedHeapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedHeapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FixedHeapConfig proto.InternalMessageInfo

func (m *FixedHeapConfig) GetMaxHeapSizeBytes() uint64 {
	if m != nil {
		return m.MaxHeapSizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*FixedHeapConfig)(nil), "envoy.config.resource_monitor.fixed_heap.v2alpha.FixedHeapConfig")
}

func init() {
	proto.RegisterFile("envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto", fileDescriptor_141ee15b3c15e2df)
}

var fileDescriptor_141ee15b3c15e2df = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0x8d, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xd2, 0x4f, 0xcb, 0xac, 0x48, 0x4d, 0x89,
	0xcf, 0x48, 0x4d, 0x2c, 0xd0, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x44, 0x12, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x00, 0x1b, 0xa1, 0x07, 0x31, 0x42, 0x0f, 0xdd, 0x08, 0x3d,
	0x24, 0xf5, 0x50, 0x23, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61,
	0x0c, 0x88, 0x51, 0x4a, 0x9e, 0x5c, 0xfc, 0x6e, 0x20, 0xe5, 0x1e, 0xa9, 0x89, 0x05, 0xce, 0x60,
	0xf3, 0x84, 0xcc, 0xb8, 0x84, 0x73, 0x13, 0x2b, 0xc0, 0xfa, 0xe3, 0x8b, 0x33, 0xab, 0x52, 0xe3,
	0x93, 0x2a, 0x4b, 0x52, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x9c, 0xd8, 0x7f, 0x39, 0xb1,
	0x18, 0x31, 0x29, 0x30, 0x04, 0x09, 0xe4, 0x26, 0x56, 0x80, 0x34, 0x05, 0x67, 0x56, 0xa5, 0x3a,
	0x81, 0x14, 0x38, 0x05, 0x71, 0xd9, 0x65, 0xe6, 0xeb, 0x81, 0x9d, 0x56, 0x50, 0x94, 0x5f, 0x51,
	0xa9, 0x47, 0xaa, 0x2b, 0x9d, 0xf8, 0xe0, 0x4e, 0x09, 0x00, 0x39, 0x2e, 0x80, 0x31, 0x89, 0x0d,
	0xec, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xf3, 0x0c, 0xc4, 0x35, 0x01, 0x00,
	0x00,
}
