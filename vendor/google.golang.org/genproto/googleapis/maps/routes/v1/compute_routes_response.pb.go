// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/compute_routes_response.proto

package routes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Reasons for using fallback response.
type FallbackReason int32

const (
	// No fallback reason specified.
	FallbackReason_FALLBACK_REASON_UNSPECIFIED FallbackReason = 0
	// A server error happened while calculating routes with your preferred
	// routing mode, but we were able to return a result calculated by an
	// alternative mode.
	FallbackReason_SERVER_ERROR FallbackReason = 1
	// We were not able to finish the calculation with your preferred routing mode
	// on time, but we were able to return a result calculated by an alternative
	// mode.
	FallbackReason_LATENCY_EXCEEDED FallbackReason = 2
)

var FallbackReason_name = map[int32]string{
	0: "FALLBACK_REASON_UNSPECIFIED",
	1: "SERVER_ERROR",
	2: "LATENCY_EXCEEDED",
}

var FallbackReason_value = map[string]int32{
	"FALLBACK_REASON_UNSPECIFIED": 0,
	"SERVER_ERROR":                1,
	"LATENCY_EXCEEDED":            2,
}

func (x FallbackReason) String() string {
	return proto.EnumName(FallbackReason_name, int32(x))
}

func (FallbackReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1609c8dc8046a099, []int{0}
}

// Actual routing mode used for returned fallback response.
type FallbackRoutingMode int32

const (
	// Not used.
	FallbackRoutingMode_FALLBACK_ROUTING_MODE_UNSPECIFIED FallbackRoutingMode = 0
	// Indicates the "TRAFFIC_UNAWARE" routing mode was used to compute the
	// response.
	FallbackRoutingMode_FALLBACK_TRAFFIC_UNAWARE FallbackRoutingMode = 1
	// Indicates the "TRAFFIC_AWARE" routing mode was used to compute the
	// response.
	FallbackRoutingMode_FALLBACK_TRAFFIC_AWARE FallbackRoutingMode = 2
)

var FallbackRoutingMode_name = map[int32]string{
	0: "FALLBACK_ROUTING_MODE_UNSPECIFIED",
	1: "FALLBACK_TRAFFIC_UNAWARE",
	2: "FALLBACK_TRAFFIC_AWARE",
}

var FallbackRoutingMode_value = map[string]int32{
	"FALLBACK_ROUTING_MODE_UNSPECIFIED": 0,
	"FALLBACK_TRAFFIC_UNAWARE":          1,
	"FALLBACK_TRAFFIC_AWARE":            2,
}

func (x FallbackRoutingMode) String() string {
	return proto.EnumName(FallbackRoutingMode_name, int32(x))
}

func (FallbackRoutingMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1609c8dc8046a099, []int{1}
}

// ComputeRoutes the response message.
type ComputeRoutesResponse struct {
	// Contains an array of computed routes (up to three) when you specify
	// compute_alternatives_routes, and contains just one route when you don't.
	// When this array contains multiple entries, the first one is the most
	// recommended route. If the array is empty, then it means no route could be
	// found.
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// In some cases when the server is not able to compute the route results with
	// all of the input preferences, it may fallback to using a different way of
	// computation. When fallback mode is used, this field contains detailed info
	// about the fallback response. Otherwise this field is unset.
	FallbackInfo         *FallbackInfo `protobuf:"bytes,2,opt,name=fallback_info,json=fallbackInfo,proto3" json:"fallback_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ComputeRoutesResponse) Reset()         { *m = ComputeRoutesResponse{} }
func (m *ComputeRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeRoutesResponse) ProtoMessage()    {}
func (*ComputeRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1609c8dc8046a099, []int{0}
}

func (m *ComputeRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeRoutesResponse.Unmarshal(m, b)
}
func (m *ComputeRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeRoutesResponse.Marshal(b, m, deterministic)
}
func (m *ComputeRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeRoutesResponse.Merge(m, src)
}
func (m *ComputeRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeRoutesResponse.Size(m)
}
func (m *ComputeRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeRoutesResponse proto.InternalMessageInfo

func (m *ComputeRoutesResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *ComputeRoutesResponse) GetFallbackInfo() *FallbackInfo {
	if m != nil {
		return m.FallbackInfo
	}
	return nil
}

// Information related to how and why a fallback result was used. If this field
// is set, then it means the server used a different routing mode from your
// preferred mode as fallback.
type FallbackInfo struct {
	// Routing mode used for the response. If fallback was triggered, the mode
	// may be different from routing preference set in the original client
	// request.
	RoutingMode FallbackRoutingMode `protobuf:"varint,1,opt,name=routing_mode,json=routingMode,proto3,enum=google.maps.routes.v1.FallbackRoutingMode" json:"routing_mode,omitempty"`
	// The reason why fallback response was used instead of the original response.
	// This field is only populated when the fallback mode is triggered and the
	// fallback response is returned.
	Reason               FallbackReason `protobuf:"varint,2,opt,name=reason,proto3,enum=google.maps.routes.v1.FallbackReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FallbackInfo) Reset()         { *m = FallbackInfo{} }
func (m *FallbackInfo) String() string { return proto.CompactTextString(m) }
func (*FallbackInfo) ProtoMessage()    {}
func (*FallbackInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1609c8dc8046a099, []int{1}
}

func (m *FallbackInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FallbackInfo.Unmarshal(m, b)
}
func (m *FallbackInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FallbackInfo.Marshal(b, m, deterministic)
}
func (m *FallbackInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FallbackInfo.Merge(m, src)
}
func (m *FallbackInfo) XXX_Size() int {
	return xxx_messageInfo_FallbackInfo.Size(m)
}
func (m *FallbackInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FallbackInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FallbackInfo proto.InternalMessageInfo

func (m *FallbackInfo) GetRoutingMode() FallbackRoutingMode {
	if m != nil {
		return m.RoutingMode
	}
	return FallbackRoutingMode_FALLBACK_ROUTING_MODE_UNSPECIFIED
}

func (m *FallbackInfo) GetReason() FallbackReason {
	if m != nil {
		return m.Reason
	}
	return FallbackReason_FALLBACK_REASON_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.maps.routes.v1.FallbackReason", FallbackReason_name, FallbackReason_value)
	proto.RegisterEnum("google.maps.routes.v1.FallbackRoutingMode", FallbackRoutingMode_name, FallbackRoutingMode_value)
	proto.RegisterType((*ComputeRoutesResponse)(nil), "google.maps.routes.v1.ComputeRoutesResponse")
	proto.RegisterType((*FallbackInfo)(nil), "google.maps.routes.v1.FallbackInfo")
}

func init() {
	proto.RegisterFile("google/maps/routes/v1/compute_routes_response.proto", fileDescriptor_1609c8dc8046a099)
}

var fileDescriptor_1609c8dc8046a099 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0x40, 0x3d, 0xb8, 0xa5, 0x8a, 0x0c, 0x45, 0xa5, 0x4c, 0xa2, 0x1b, 0x9a, 0x54,
	0xf5, 0xe0, 0xa8, 0x1d, 0xb7, 0x89, 0x43, 0x9a, 0x3a, 0x25, 0xa2, 0x49, 0xab, 0xb7, 0x7f, 0x60,
	0xa8, 0x92, 0x95, 0x75, 0x6e, 0x54, 0xd1, 0xc4, 0x51, 0xd2, 0xf5, 0xab, 0x70, 0xe1, 0xc4, 0x91,
	0x13, 0x9f, 0x83, 0x4f, 0xc4, 0x11, 0x25, 0x0e, 0x6c, 0x8c, 0xa0, 0xdd, 0xec, 0xf7, 0x79, 0x7e,
	0xcf, 0x63, 0x59, 0x2f, 0x3e, 0x0b, 0xa4, 0x0c, 0x76, 0xc2, 0x08, 0xfd, 0x38, 0x35, 0x12, 0x79,
	0xbd, 0x17, 0xa9, 0x71, 0xe8, 0x19, 0x6b, 0x19, 0xc6, 0xd7, 0x7b, 0xc1, 0xd5, 0x84, 0x27, 0x22,
	0x8d, 0x65, 0x94, 0x0a, 0x1a, 0x27, 0x72, 0x2f, 0x49, 0x43, 0x41, 0x34, 0x83, 0xa8, 0xb2, 0xd0,
	0x43, 0xaf, 0x75, 0x5c, 0x9e, 0x95, 0x9f, 0x14, 0x79, 0xf2, 0x19, 0xe1, 0x86, 0xa5, 0xb2, 0x21,
	0x37, 0x40, 0x91, 0x4c, 0x5e, 0xe3, 0x8a, 0x42, 0x9a, 0xa8, 0xfd, 0xb0, 0x53, 0xed, 0x1f, 0xd1,
	0xd2, 0x12, 0x9a, 0x63, 0x50, 0x78, 0xc9, 0x5b, 0xfc, 0x78, 0xe3, 0xef, 0x76, 0x97, 0xfe, 0xfa,
	0x13, 0xdf, 0x46, 0x1b, 0xd9, 0xd4, 0xda, 0xa8, 0x53, 0xed, 0xbf, 0xfa, 0x0f, 0x6c, 0x17, 0x5e,
	0x27, 0xda, 0x48, 0xa8, 0x6d, 0x6e, 0xdd, 0x4e, 0xbe, 0x20, 0x5c, 0xbb, 0x2d, 0x13, 0x17, 0xd7,
	0x32, 0x70, 0x1b, 0x05, 0x3c, 0x94, 0x57, 0xa2, 0x89, 0xda, 0xa8, 0x53, 0xef, 0x77, 0xef, 0x49,
	0x06, 0x85, 0xb8, 0xf2, 0x4a, 0x40, 0x35, 0xb9, 0xb9, 0x90, 0x37, 0xb8, 0x92, 0x08, 0x3f, 0x95,
	0x51, 0xfe, 0xc4, 0x7a, 0xff, 0xf4, 0xbe, 0xa0, 0xdc, 0x0c, 0x05, 0xd4, 0xbd, 0xc0, 0xf5, 0xbf,
	0x15, 0xf2, 0x12, 0xbf, 0xb0, 0xcd, 0xf1, 0x78, 0x60, 0x5a, 0xef, 0x38, 0x30, 0x73, 0x36, 0xf1,
	0xf8, 0xc2, 0x9b, 0x4d, 0x99, 0xe5, 0xd8, 0x0e, 0x1b, 0xea, 0x0f, 0x88, 0x8e, 0x6b, 0x33, 0x06,
	0x4b, 0x06, 0x9c, 0x01, 0x4c, 0x40, 0x47, 0xe4, 0x29, 0xd6, 0xc7, 0xe6, 0x9c, 0x79, 0xd6, 0x05,
	0x67, 0x1f, 0x2c, 0xc6, 0x86, 0x6c, 0xa8, 0x6b, 0xdd, 0x03, 0x7e, 0x52, 0xf2, 0x7a, 0x72, 0x8a,
	0x8f, 0x6f, 0xf2, 0x27, 0x8b, 0xb9, 0xe3, 0x8d, 0xb8, 0x3b, 0x19, 0xb2, 0x3b, 0x2d, 0x47, 0xb8,
	0xf9, 0xc7, 0x36, 0x07, 0xd3, 0xb6, 0x1d, 0x8b, 0x2f, 0x3c, 0xf3, 0xbd, 0x09, 0x4c, 0x47, 0xa4,
	0x85, 0x9f, 0xfd, 0xa3, 0x2a, 0x4d, 0x1b, 0x7c, 0x47, 0xf8, 0xf9, 0x5a, 0x86, 0xe5, 0xff, 0x30,
	0x68, 0x95, 0xae, 0xc9, 0x34, 0xdb, 0xa2, 0x29, 0xfa, 0x78, 0x5e, 0x40, 0x81, 0xdc, 0xf9, 0x51,
	0x40, 0x65, 0x12, 0x18, 0x81, 0x88, 0xf2, 0x1d, 0x33, 0x94, 0xe4, 0xc7, 0xdb, 0xf4, 0xce, 0x26,
	0x9e, 0xab, 0xd3, 0x4f, 0x84, 0xbe, 0x6a, 0x8f, 0x46, 0x2e, 0xcc, 0xbe, 0x69, 0x8d, 0x91, 0xca,
	0x71, 0xb3, 0x72, 0x55, 0x45, 0x97, 0xbd, 0x1f, 0xbf, 0xe7, 0xab, 0x6c, 0xbe, 0x52, 0xf3, 0xd5,
	0xb2, 0x77, 0x59, 0xc9, 0x1b, 0xce, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xab, 0xf0, 0x50,
	0x36, 0x03, 0x00, 0x00,
}
