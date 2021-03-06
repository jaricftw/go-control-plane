// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rbac/v2/rbac.proto

package envoy_config_filter_http_rbac_v2

import (
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2"
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

type RBAC struct {
	Rules                *v2.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v2.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d628c6558085a7, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v2.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v2.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d628c6558085a7, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.http.rbac.v2.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.config.filter.http.rbac.v2.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rbac/v2/rbac.proto", fileDescriptor_15d628c6558085a7)
}

var fileDescriptor_15d628c6558085a7 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x70, 0x4b, 0xaa, 0x64, 0x28, 0x5e, 0x5a, 0xbc, 0x34, 0x64, 0x28, 0x85, 0xc2,
	0xa9, 0xb8, 0x5b, 0x21, 0x43, 0xdd, 0xad, 0x93, 0xd1, 0xd8, 0xa5, 0xc8, 0x96, 0xd2, 0x08, 0x44,
	0xce, 0xc8, 0x8a, 0xd2, 0xfc, 0xfb, 0x22, 0x5d, 0x3a, 0x98, 0x0c, 0x9e, 0x24, 0xa4, 0xef, 0x7d,
	0xc7, 0x3d, 0xf6, 0xac, 0xf7, 0x01, 0x4f, 0xbc, 0xc3, 0xfd, 0xd6, 0xfc, 0xf0, 0xad, 0xb1, 0x5e,
	0x3b, 0xbe, 0xf3, 0xbe, 0xe7, 0xae, 0x95, 0x1d, 0x0f, 0x55, 0x3a, 0xa1, 0x77, 0xe8, 0xb1, 0x58,
	0x25, 0x18, 0x08, 0x06, 0x82, 0x21, 0xc2, 0x90, 0xa0, 0x50, 0x95, 0x0f, 0x23, 0xdd, 0xa5, 0xa2,
	0xbc, 0x0b, 0xd2, 0x1a, 0x25, 0xbd, 0xe6, 0xff, 0x17, 0xfa, 0x58, 0x1f, 0x59, 0x2e, 0xea, 0xf7,
	0x8f, 0xe2, 0x85, 0x5d, 0xb9, 0x83, 0xd5, 0xc3, 0x7d, 0xb6, 0xca, 0x9e, 0x16, 0x55, 0x09, 0xa3,
	0x99, 0xe7, 0x39, 0x10, 0x51, 0x41, 0x60, 0xb1, 0x61, 0xcb, 0x61, 0x27, 0x15, 0x1e, 0xbf, 0x29,
	0x38, 0x9b, 0x0c, 0x2e, 0x88, 0x17, 0x11, 0x5f, 0x7f, 0xb1, 0x65, 0x7c, 0x6c, 0xb4, 0x13, 0x78,
	0xf0, 0xba, 0x78, 0x63, 0x79, 0x84, 0xcf, 0x9a, 0x47, 0x98, 0xda, 0x99, 0x94, 0x29, 0xf3, 0x99,
	0xcf, 0xb3, 0xdb, 0x99, 0x98, 0x2b, 0x33, 0xc8, 0xd6, 0x6a, 0x55, 0x6f, 0x18, 0x18, 0x24, 0x43,
	0xef, 0xf0, 0xf7, 0x34, 0x29, 0xab, 0x6f, 0x44, 0x2b, 0xbb, 0x26, 0x36, 0xd2, 0x64, 0xed, 0x75,
	0xaa, 0xe6, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x19, 0x66, 0x2f, 0x07, 0xa5, 0x01, 0x00, 0x00,
}
