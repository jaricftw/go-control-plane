// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v3alpha/ext_authz.proto

package envoy_config_filter_http_ext_authz_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
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

type ExtAuthz struct {
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services                  isExtAuthz_Services `protobuf_oneof:"services"`
	FailureModeAllow          bool                `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	UseAlpha                  bool                `protobuf:"varint,4,opt,name=use_alpha,json=useAlpha,proto3" json:"use_alpha,omitempty"` // Deprecated: Do not use.
	WithRequestBody           *BufferSettings     `protobuf:"bytes,5,opt,name=with_request_body,json=withRequestBody,proto3" json:"with_request_body,omitempty"`
	ClearRouteCache           bool                `protobuf:"varint,6,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	StatusOnError             *_type.HttpStatus   `protobuf:"bytes,7,opt,name=status_on_error,json=statusOnError,proto3" json:"status_on_error,omitempty"`
	MetadataContextNamespaces []string            `protobuf:"bytes,8,rep,name=metadata_context_namespaces,json=metadataContextNamespaces,proto3" json:"metadata_context_namespaces,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}            `json:"-"`
	XXX_unrecognized          []byte              `json:"-"`
	XXX_sizecache             int32               `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// Deprecated: Do not use.
func (m *ExtAuthz) GetUseAlpha() bool {
	if m != nil {
		return m.UseAlpha
	}
	return false
}

func (m *ExtAuthz) GetWithRequestBody() *BufferSettings {
	if m != nil {
		return m.WithRequestBody
	}
	return nil
}

func (m *ExtAuthz) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *ExtAuthz) GetStatusOnError() *_type.HttpStatus {
	if m != nil {
		return m.StatusOnError
	}
	return nil
}

func (m *ExtAuthz) GetMetadataContextNamespaces() []string {
	if m != nil {
		return m.MetadataContextNamespaces
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

type BufferSettings struct {
	MaxRequestBytes      uint32   `protobuf:"varint,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	AllowPartialMessage  bool     `protobuf:"varint,2,opt,name=allow_partial_message,json=allowPartialMessage,proto3" json:"allow_partial_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{1}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

func (m *BufferSettings) GetAllowPartialMessage() bool {
	if m != nil {
		return m.AllowPartialMessage
	}
	return false
}

type HttpService struct {
	ServerUri             *core.HttpUri          `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	PathPrefix            string                 `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	AuthorizationRequest  *AuthorizationRequest  `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{2}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *core.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	AllowedHeaders       *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	HeadersToAdd         []*core.HeaderValue        `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{3}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*core.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	AllowedUpstreamHeaders *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	AllowedClientHeaders   *matcher.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                   `json:"-"`
	XXX_unrecognized       []byte                     `json:"-"`
	XXX_sizecache          int32                      `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{4}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{5}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

type CheckSettings struct {
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73c121e155221d0, []int{6}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v3alpha.ExtAuthz")
	proto.RegisterType((*BufferSettings)(nil), "envoy.config.filter.http.ext_authz.v3alpha.BufferSettings")
	proto.RegisterType((*HttpService)(nil), "envoy.config.filter.http.ext_authz.v3alpha.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.config.filter.http.ext_authz.v3alpha.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.config.filter.http.ext_authz.v3alpha.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.config.filter.http.ext_authz.v3alpha.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.config.filter.http.ext_authz.v3alpha.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.config.filter.http.ext_authz.v3alpha.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v3alpha/ext_authz.proto", fileDescriptor_b73c121e155221d0)
}

var fileDescriptor_b73c121e155221d0 = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0xb4, 0x75, 0xa7, 0x3f, 0x49, 0x87, 0xb6, 0x98, 0x82, 0xd4, 0x12, 0x54, 0xa9,
	0x54, 0xc8, 0x91, 0xb6, 0x17, 0x2c, 0xbd, 0xa8, 0x88, 0x4b, 0x45, 0x29, 0x74, 0xa9, 0x5c, 0xca,
	0x0d, 0x48, 0xa3, 0xa9, 0x7d, 0x12, 0x0f, 0xeb, 0x78, 0xcc, 0xcc, 0x38, 0x9b, 0xec, 0x0d, 0xf7,
	0xbc, 0x05, 0x8f, 0xc0, 0x1d, 0x2f, 0x02, 0xaf, 0x02, 0xda, 0x2b, 0x34, 0x9e, 0x71, 0x9a, 0xae,
	0x52, 0x44, 0x76, 0xef, 0xe2, 0x73, 0xbe, 0xf3, 0x9d, 0xef, 0xfc, 0xcc, 0x09, 0x3a, 0x81, 0x6c,
	0xc8, 0xc7, 0x9d, 0x88, 0x67, 0x3d, 0xd6, 0xef, 0xf4, 0x58, 0xaa, 0x40, 0x74, 0x12, 0xa5, 0xf2,
	0x0e, 0x8c, 0x14, 0xa1, 0x85, 0x4a, 0x5e, 0x76, 0x86, 0xc7, 0x34, 0xcd, 0x13, 0x7a, 0x6f, 0xf1,
	0x73, 0xc1, 0x15, 0xc7, 0x47, 0x65, 0xac, 0x6f, 0x62, 0x7d, 0x13, 0xeb, 0xeb, 0x58, 0xff, 0x1e,
	0x69, 0x63, 0x77, 0x3f, 0x34, 0x79, 0x68, 0xce, 0x26, 0x74, 0x11, 0x17, 0xd0, 0xb9, 0xa3, 0x12,
	0x0c, 0xdd, 0xee, 0xc7, 0x8f, 0x40, 0xfa, 0x22, 0x8f, 0x88, 0x04, 0x31, 0x64, 0x51, 0x05, 0x3d,
	0x78, 0x04, 0xaa, 0x93, 0x93, 0x42, 0x30, 0x0b, 0xfb, 0xc0, 0xc0, 0xd4, 0x38, 0xb7, 0x2e, 0xa9,
	0xa8, 0x2a, 0xa4, 0xf5, 0xee, 0x4d, 0x79, 0x07, 0x54, 0x45, 0x09, 0x88, 0x8e, 0x54, 0x82, 0x65,
	0x7d, 0x0b, 0x78, 0x77, 0x48, 0x53, 0x16, 0x53, 0x05, 0x9d, 0xea, 0x87, 0x71, 0xb4, 0x7f, 0x6b,
	0x20, 0xf7, 0x7c, 0xa4, 0xba, 0xba, 0x42, 0x7c, 0x81, 0xd6, 0xa6, 0x15, 0x7a, 0xce, 0xbe, 0x73,
	0xb8, 0xfa, 0xe4, 0x23, 0xdf, 0x34, 0x87, 0xe6, 0xac, 0xea, 0x81, 0xaf, 0x25, 0xfa, 0x5f, 0x8a,
	0x3c, 0xba, 0x31, 0xd0, 0x8b, 0x85, 0x70, 0xb5, 0x7f, 0xff, 0x89, 0x7f, 0x44, 0x6b, 0x46, 0xa5,
	0x65, 0xaa, 0x97, 0x4c, 0x9f, 0xfa, 0xff, 0xbf, 0xcd, 0xfe, 0x85, 0x52, 0xf9, 0x14, 0x7b, 0x72,
	0xff, 0x89, 0x3f, 0x41, 0xb8, 0x47, 0x59, 0x5a, 0x08, 0x20, 0x03, 0x1e, 0x03, 0xa1, 0x69, 0xca,
	0x5f, 0x78, 0xb5, 0x7d, 0xe7, 0xd0, 0x0d, 0x5b, 0xd6, 0x73, 0xc5, 0x63, 0xe8, 0x6a, 0x3b, 0xde,
	0x43, 0x2b, 0x85, 0xd4, 0xa0, 0x3c, 0xa1, 0x5e, 0x43, 0x83, 0x82, 0x9a, 0xe7, 0x84, 0x6e, 0x21,
	0xa1, 0xab, 0x6d, 0xb8, 0x87, 0x36, 0x5f, 0x30, 0x95, 0x10, 0x01, 0x3f, 0x17, 0x20, 0x15, 0xb9,
	0xe3, 0xf1, 0xd8, 0x5b, 0x2c, 0x15, 0x9f, 0xcc, 0xa3, 0x38, 0x28, 0x7a, 0x3d, 0x10, 0x37, 0xa0,
	0x14, 0xcb, 0xfa, 0x32, 0x6c, 0x6a, 0xd2, 0xd0, 0x70, 0x06, 0x3c, 0x1e, 0xe3, 0x23, 0xb4, 0x19,
	0xa5, 0x40, 0x05, 0x11, 0xbc, 0x50, 0x40, 0x22, 0x1a, 0x25, 0xe0, 0x2d, 0x95, 0xaa, 0x9b, 0xa5,
	0x23, 0xd4, 0xf6, 0x33, 0x6d, 0xc6, 0xa7, 0xa8, 0x69, 0x26, 0x4c, 0x78, 0x46, 0x40, 0x08, 0x2e,
	0xbc, 0xe5, 0x52, 0xd1, 0x8e, 0x55, 0xa4, 0x67, 0x6d, 0x7a, 0x54, 0xc2, 0xc2, 0x75, 0x03, 0xff,
	0x36, 0x3b, 0xd7, 0x60, 0x7c, 0x8a, 0xde, 0x1f, 0x80, 0xa2, 0x31, 0x55, 0x94, 0x44, 0x3c, 0x53,
	0x5a, 0x6d, 0x46, 0x07, 0x20, 0x73, 0x1a, 0x81, 0xf4, 0xdc, 0xfd, 0xfa, 0xe1, 0x4a, 0xf8, 0x5e,
	0x05, 0x39, 0x33, 0x88, 0x67, 0x13, 0x40, 0x80, 0x90, 0x6b, 0x67, 0x27, 0xdb, 0x63, 0xb4, 0xf1,
	0xb0, 0x34, 0x7c, 0x8c, 0x36, 0x07, 0x74, 0x74, 0xdf, 0xb0, 0xb1, 0x02, 0x59, 0x6e, 0xcb, 0x7a,
	0xb0, 0xfc, 0x2a, 0x68, 0x1c, 0xd5, 0xf6, 0x17, 0xc2, 0xe6, 0x80, 0x8e, 0xaa, 0xea, 0xb5, 0x1f,
	0x3f, 0x41, 0xdb, 0xe5, 0xa0, 0x48, 0x4e, 0x85, 0x62, 0x34, 0x25, 0x03, 0x90, 0x92, 0xf6, 0xc1,
	0x0e, 0xee, 0x9d, 0xd2, 0x79, 0x6d, 0x7c, 0x57, 0xc6, 0xd5, 0xfe, 0xbb, 0x86, 0x56, 0xa7, 0x16,
	0x01, 0x9f, 0x22, 0xa4, 0x65, 0x81, 0xd0, 0x4f, 0xc3, 0xee, 0xe7, 0xde, 0x63, 0xfb, 0xa9, 0x03,
	0x6f, 0x05, 0x0b, 0x57, 0x4c, 0xc8, 0xad, 0x60, 0x78, 0x0f, 0xad, 0xe6, 0x54, 0x25, 0x24, 0x17,
	0xd0, 0x63, 0xa3, 0x32, 0xf3, 0x4a, 0x88, 0xb4, 0xe9, 0xba, 0xb4, 0xe0, 0x02, 0x6d, 0xeb, 0xa1,
	0x72, 0xc1, 0x5e, 0x52, 0xc5, 0x78, 0x56, 0xd5, 0x68, 0xbb, 0xff, 0xf9, 0x3c, 0xfb, 0xd0, 0x9d,
	0x26, 0xb2, 0xad, 0x08, 0xb7, 0xe8, 0x0c, 0x2b, 0x1e, 0xa1, 0x9d, 0xd7, 0xd3, 0xca, 0x9c, 0x67,
	0x12, 0x3c, 0xb7, 0xcc, 0xdb, 0x7d, 0x8b, 0xbc, 0x86, 0x28, 0xdc, 0xa6, 0xb3, 0xcc, 0x97, 0x0d,
	0xb7, 0xde, 0x6a, 0x5c, 0x36, 0xdc, 0x46, 0x6b, 0xf1, 0xb2, 0xe1, 0x2e, 0xb6, 0x96, 0x2e, 0x1b,
	0xee, 0x52, 0x6b, 0xb9, 0xfd, 0xbb, 0x83, 0xb6, 0x66, 0x15, 0x80, 0x9f, 0xa1, 0x66, 0x39, 0x29,
	0x88, 0x49, 0x02, 0x34, 0x06, 0x21, 0xed, 0x1c, 0x0e, 0xa6, 0x37, 0xd3, 0x5e, 0x21, 0xff, 0x1b,
	0x26, 0xd5, 0x4d, 0x79, 0x89, 0xae, 0x8c, 0x25, 0xdc, 0xb0, 0xd1, 0x17, 0x26, 0x18, 0x7f, 0x85,
	0x36, 0x2c, 0x0f, 0x51, 0x9c, 0xd0, 0x38, 0xf6, 0x6a, 0xfb, 0xf5, 0xff, 0x3a, 0x3b, 0x26, 0xf0,
	0x7b, 0x9a, 0x16, 0x10, 0xae, 0xd9, 0xd0, 0xef, 0x78, 0x37, 0x8e, 0xdb, 0x7f, 0x3a, 0x68, 0x7b,
	0x66, 0xf1, 0x98, 0x20, 0xaf, 0x12, 0x5d, 0xe4, 0x52, 0x09, 0xa0, 0x83, 0x37, 0x53, 0xbf, 0x63,
	0x69, 0x6e, 0x2d, 0x4b, 0x55, 0xc5, 0x0f, 0xa8, 0xf2, 0x90, 0x28, 0x65, 0x90, 0xa9, 0x09, 0x7d,
	0x6d, 0x1e, 0xfa, 0x2d, 0x4b, 0x72, 0x56, 0x72, 0x58, 0xf2, 0xf6, 0x1f, 0x0e, 0x6a, 0x55, 0x47,
	0xfa, 0x1a, 0xcc, 0x99, 0xc0, 0x07, 0xc8, 0x8d, 0x99, 0xa4, 0x77, 0x29, 0xc4, 0x65, 0x09, 0x6e,
	0xf9, 0xf4, 0x7e, 0xaa, 0xb9, 0xce, 0xc5, 0x42, 0x38, 0x71, 0xe1, 0x14, 0x6d, 0x44, 0x09, 0x44,
	0xcf, 0x89, 0xb4, 0x8f, 0xd7, 0x0a, 0xfa, 0x6c, 0x9e, 0x8d, 0x3a, 0xd3, 0x0c, 0xd5, 0xeb, 0x0f,
	0xdc, 0x57, 0xc1, 0xe2, 0xaf, 0x4e, 0xad, 0xa5, 0x13, 0xad, 0x47, 0x0f, 0x5c, 0x4d, 0xe4, 0xf2,
	0x21, 0x08, 0xc1, 0x62, 0xc0, 0xf5, 0x7f, 0x02, 0xa7, 0xfd, 0x97, 0x83, 0xd6, 0x1f, 0x44, 0xe3,
	0x5f, 0x10, 0xae, 0x0e, 0x12, 0x8c, 0x14, 0x64, 0x92, 0xf1, 0x4c, 0x0f, 0x41, 0xcf, 0xfc, 0xfa,
	0x8d, 0x45, 0xf9, 0xf6, 0x84, 0x9d, 0x4f, 0x28, 0xcf, 0x33, 0x25, 0xc6, 0xe1, 0x66, 0xf4, 0xba,
	0x7d, 0xf7, 0x0b, 0xb4, 0x33, 0x1b, 0x8c, 0x5b, 0xa8, 0xfe, 0x1c, 0xc6, 0x65, 0x37, 0x57, 0x42,
	0xfd, 0x13, 0x6f, 0xa1, 0xc5, 0xa1, 0x5e, 0x34, 0x7b, 0x29, 0xcc, 0xc7, 0x49, 0xed, 0xa9, 0x13,
	0x7c, 0x8d, 0x9e, 0x32, 0x6e, 0xe4, 0xe6, 0x82, 0x8f, 0xc6, 0x73, 0x28, 0x0f, 0xd6, 0x27, 0xc3,
	0xd4, 0xff, 0xc1, 0xd7, 0xce, 0xdd, 0x52, 0xf9, 0x67, 0x7c, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd7, 0xf8, 0x78, 0x1d, 0xc3, 0x08, 0x00, 0x00,
}
