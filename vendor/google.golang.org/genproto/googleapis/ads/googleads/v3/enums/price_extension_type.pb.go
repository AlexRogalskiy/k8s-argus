// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/price_extension_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Price extension type.
type PriceExtensionTypeEnum_PriceExtensionType int32

const (
	// Not specified.
	PriceExtensionTypeEnum_UNSPECIFIED PriceExtensionTypeEnum_PriceExtensionType = 0
	// Used for return value only. Represents value unknown in this version.
	PriceExtensionTypeEnum_UNKNOWN PriceExtensionTypeEnum_PriceExtensionType = 1
	// The type for showing a list of brands.
	PriceExtensionTypeEnum_BRANDS PriceExtensionTypeEnum_PriceExtensionType = 2
	// The type for showing a list of events.
	PriceExtensionTypeEnum_EVENTS PriceExtensionTypeEnum_PriceExtensionType = 3
	// The type for showing locations relevant to your business.
	PriceExtensionTypeEnum_LOCATIONS PriceExtensionTypeEnum_PriceExtensionType = 4
	// The type for showing sub-regions or districts within a city or region.
	PriceExtensionTypeEnum_NEIGHBORHOODS PriceExtensionTypeEnum_PriceExtensionType = 5
	// The type for showing a collection of product categories.
	PriceExtensionTypeEnum_PRODUCT_CATEGORIES PriceExtensionTypeEnum_PriceExtensionType = 6
	// The type for showing a collection of related product tiers.
	PriceExtensionTypeEnum_PRODUCT_TIERS PriceExtensionTypeEnum_PriceExtensionType = 7
	// The type for showing a collection of services offered by your business.
	PriceExtensionTypeEnum_SERVICES PriceExtensionTypeEnum_PriceExtensionType = 8
	// The type for showing a collection of service categories.
	PriceExtensionTypeEnum_SERVICE_CATEGORIES PriceExtensionTypeEnum_PriceExtensionType = 9
	// The type for showing a collection of related service tiers.
	PriceExtensionTypeEnum_SERVICE_TIERS PriceExtensionTypeEnum_PriceExtensionType = 10
)

var PriceExtensionTypeEnum_PriceExtensionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BRANDS",
	3:  "EVENTS",
	4:  "LOCATIONS",
	5:  "NEIGHBORHOODS",
	6:  "PRODUCT_CATEGORIES",
	7:  "PRODUCT_TIERS",
	8:  "SERVICES",
	9:  "SERVICE_CATEGORIES",
	10: "SERVICE_TIERS",
}

var PriceExtensionTypeEnum_PriceExtensionType_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"BRANDS":             2,
	"EVENTS":             3,
	"LOCATIONS":          4,
	"NEIGHBORHOODS":      5,
	"PRODUCT_CATEGORIES": 6,
	"PRODUCT_TIERS":      7,
	"SERVICES":           8,
	"SERVICE_CATEGORIES": 9,
	"SERVICE_TIERS":      10,
}

func (x PriceExtensionTypeEnum_PriceExtensionType) String() string {
	return proto.EnumName(PriceExtensionTypeEnum_PriceExtensionType_name, int32(x))
}

func (PriceExtensionTypeEnum_PriceExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d28cbae1bf68a407, []int{0, 0}
}

// Container for enum describing types for a price extension.
type PriceExtensionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceExtensionTypeEnum) Reset()         { *m = PriceExtensionTypeEnum{} }
func (m *PriceExtensionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PriceExtensionTypeEnum) ProtoMessage()    {}
func (*PriceExtensionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28cbae1bf68a407, []int{0}
}

func (m *PriceExtensionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceExtensionTypeEnum.Unmarshal(m, b)
}
func (m *PriceExtensionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceExtensionTypeEnum.Marshal(b, m, deterministic)
}
func (m *PriceExtensionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceExtensionTypeEnum.Merge(m, src)
}
func (m *PriceExtensionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PriceExtensionTypeEnum.Size(m)
}
func (m *PriceExtensionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceExtensionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PriceExtensionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PriceExtensionTypeEnum_PriceExtensionType", PriceExtensionTypeEnum_PriceExtensionType_name, PriceExtensionTypeEnum_PriceExtensionType_value)
	proto.RegisterType((*PriceExtensionTypeEnum)(nil), "google.ads.googleads.v3.enums.PriceExtensionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/price_extension_type.proto", fileDescriptor_d28cbae1bf68a407)
}

var fileDescriptor_d28cbae1bf68a407 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6e, 0xd4, 0x30,
	0x18, 0x85, 0xc9, 0x14, 0xa6, 0xad, 0x4b, 0x85, 0xf1, 0xa2, 0x48, 0x88, 0x2e, 0xda, 0x03, 0x38,
	0x8b, 0x6c, 0x90, 0x59, 0x39, 0x89, 0x49, 0x23, 0x90, 0x1d, 0xc5, 0x99, 0x20, 0xa1, 0x48, 0xa3,
	0xd0, 0x58, 0x51, 0xa4, 0x8e, 0x1d, 0x8d, 0xd3, 0x8a, 0x5e, 0x87, 0x25, 0x47, 0x61, 0xc5, 0x39,
	0x80, 0x43, 0x20, 0x27, 0xcd, 0x08, 0x69, 0x44, 0x37, 0xd1, 0xcb, 0xff, 0xbf, 0xef, 0x29, 0x79,
	0x3f, 0x78, 0xdb, 0x1a, 0xd3, 0xde, 0x28, 0xbf, 0x6e, 0xac, 0x3f, 0x49, 0xa7, 0xee, 0x02, 0x5f,
	0xe9, 0xdb, 0x8d, 0xf5, 0xfb, 0x6d, 0x77, 0xad, 0xd6, 0xea, 0xeb, 0xa0, 0xb4, 0xed, 0x8c, 0x5e,
	0x0f, 0xf7, 0xbd, 0xc2, 0xfd, 0xd6, 0x0c, 0x06, 0x9d, 0x4f, 0x76, 0x5c, 0x37, 0x16, 0xef, 0x48,
	0x7c, 0x17, 0xe0, 0x91, 0x7c, 0xfd, 0x66, 0x0e, 0xee, 0x3b, 0xbf, 0xd6, 0xda, 0x0c, 0xf5, 0xd0,
	0x19, 0x6d, 0x27, 0xf8, 0xf2, 0xb7, 0x07, 0xce, 0x32, 0x97, 0xcd, 0xe6, 0xe8, 0xe2, 0xbe, 0x57,
	0x4c, 0xdf, 0x6e, 0x2e, 0x7f, 0x7a, 0x00, 0xed, 0xaf, 0xd0, 0x0b, 0x70, 0xb2, 0xe2, 0x32, 0x63,
	0x51, 0xfa, 0x3e, 0x65, 0x31, 0x7c, 0x82, 0x4e, 0xc0, 0xe1, 0x8a, 0x7f, 0xe0, 0xe2, 0x13, 0x87,
	0x1e, 0x02, 0x60, 0x19, 0xe6, 0x94, 0xc7, 0x12, 0x2e, 0x9c, 0x66, 0x25, 0xe3, 0x85, 0x84, 0x07,
	0xe8, 0x14, 0x1c, 0x7f, 0x14, 0x11, 0x2d, 0x52, 0xc1, 0x25, 0x7c, 0x8a, 0x5e, 0x82, 0x53, 0xce,
	0xd2, 0xe4, 0x2a, 0x14, 0xf9, 0x95, 0x10, 0xb1, 0x84, 0xcf, 0xd0, 0x19, 0x40, 0x59, 0x2e, 0xe2,
	0x55, 0x54, 0xac, 0x23, 0x5a, 0xb0, 0x44, 0xe4, 0x29, 0x93, 0x70, 0xe9, 0xac, 0xf3, 0xbc, 0x48,
	0x59, 0x2e, 0xe1, 0x21, 0x7a, 0x0e, 0x8e, 0x24, 0xcb, 0xcb, 0x34, 0x62, 0x12, 0x1e, 0x39, 0xf0,
	0xe1, 0xed, 0x5f, 0xf0, 0xd8, 0x81, 0xf3, 0x7c, 0x02, 0x41, 0xf8, 0xc7, 0x03, 0x17, 0xd7, 0x66,
	0x83, 0x1f, 0x6d, 0x2c, 0x7c, 0xb5, 0xff, 0xd7, 0x99, 0x2b, 0x2b, 0xf3, 0x3e, 0x87, 0x0f, 0x64,
	0x6b, 0x6e, 0x6a, 0xdd, 0x62, 0xb3, 0x6d, 0xfd, 0x56, 0xe9, 0xb1, 0xca, 0xf9, 0x6a, 0x7d, 0x67,
	0xff, 0x73, 0xc4, 0x77, 0xe3, 0xf3, 0xdb, 0xe2, 0x20, 0xa1, 0xf4, 0xfb, 0xe2, 0x3c, 0x99, 0xa2,
	0x68, 0x63, 0xf1, 0x24, 0x9d, 0x2a, 0x03, 0xec, 0xca, 0xb7, 0x3f, 0xe6, 0x7d, 0x45, 0x1b, 0x5b,
	0xed, 0xf6, 0x55, 0x19, 0x54, 0xe3, 0xfe, 0xd7, 0xe2, 0x62, 0x1a, 0x12, 0x42, 0x1b, 0x4b, 0xc8,
	0xce, 0x41, 0x48, 0x19, 0x10, 0x32, 0x7a, 0xbe, 0x2c, 0xc7, 0x0f, 0x0b, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x13, 0xf3, 0xb8, 0x7a, 0x5c, 0x02, 0x00, 0x00,
}
