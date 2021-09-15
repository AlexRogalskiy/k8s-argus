// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/shared_set_type.proto

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

// Enum listing the possible shared set types.
type SharedSetTypeEnum_SharedSetType int32

const (
	// Not specified.
	SharedSetTypeEnum_UNSPECIFIED SharedSetTypeEnum_SharedSetType = 0
	// Used for return value only. Represents value unknown in this version.
	SharedSetTypeEnum_UNKNOWN SharedSetTypeEnum_SharedSetType = 1
	// A set of keywords that can be excluded from targeting.
	SharedSetTypeEnum_NEGATIVE_KEYWORDS SharedSetTypeEnum_SharedSetType = 2
	// A set of placements that can be excluded from targeting.
	SharedSetTypeEnum_NEGATIVE_PLACEMENTS SharedSetTypeEnum_SharedSetType = 3
)

var SharedSetTypeEnum_SharedSetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NEGATIVE_KEYWORDS",
	3: "NEGATIVE_PLACEMENTS",
}

var SharedSetTypeEnum_SharedSetType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"NEGATIVE_KEYWORDS":   2,
	"NEGATIVE_PLACEMENTS": 3,
}

func (x SharedSetTypeEnum_SharedSetType) String() string {
	return proto.EnumName(SharedSetTypeEnum_SharedSetType_name, int32(x))
}

func (SharedSetTypeEnum_SharedSetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f04b2780a66f344, []int{0, 0}
}

// Container for enum describing types of shared sets.
type SharedSetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetTypeEnum) Reset()         { *m = SharedSetTypeEnum{} }
func (m *SharedSetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetTypeEnum) ProtoMessage()    {}
func (*SharedSetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f04b2780a66f344, []int{0}
}

func (m *SharedSetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetTypeEnum.Unmarshal(m, b)
}
func (m *SharedSetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetTypeEnum.Marshal(b, m, deterministic)
}
func (m *SharedSetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetTypeEnum.Merge(m, src)
}
func (m *SharedSetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetTypeEnum.Size(m)
}
func (m *SharedSetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SharedSetTypeEnum_SharedSetType", SharedSetTypeEnum_SharedSetType_name, SharedSetTypeEnum_SharedSetType_value)
	proto.RegisterType((*SharedSetTypeEnum)(nil), "google.ads.googleads.v3.enums.SharedSetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/shared_set_type.proto", fileDescriptor_3f04b2780a66f344)
}

var fileDescriptor_3f04b2780a66f344 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x1d, 0x28, 0x64, 0x88, 0x5d, 0x45, 0x04, 0x71, 0x87, 0xed, 0x01, 0xd2, 0x43, 0x6e,
	0xf1, 0x94, 0x6d, 0x71, 0x8c, 0x69, 0x57, 0xec, 0xd6, 0xa1, 0x54, 0x46, 0xb5, 0x21, 0x0e, 0xb6,
	0xa4, 0x34, 0xd9, 0x60, 0xaf, 0xe3, 0xd1, 0x47, 0xf1, 0x45, 0x04, 0x9f, 0x42, 0x9a, 0xd8, 0xc2,
	0x0e, 0x7a, 0x09, 0x1f, 0xbf, 0xef, 0x0f, 0x5f, 0x3e, 0x80, 0xb8, 0x94, 0x7c, 0xcd, 0xfc, 0x34,
	0x53, 0xbe, 0x85, 0x25, 0xda, 0x21, 0x9f, 0x89, 0xed, 0x46, 0xf9, 0xea, 0x2d, 0x2d, 0x58, 0xb6,
	0x54, 0x4c, 0x2f, 0xf5, 0x3e, 0x67, 0x30, 0x2f, 0xa4, 0x96, 0x5e, 0xc7, 0x2a, 0x61, 0x9a, 0x29,
	0x58, 0x9b, 0xe0, 0x0e, 0x41, 0x63, 0xba, 0xba, 0xae, 0x32, 0xf3, 0x95, 0x9f, 0x0a, 0x21, 0x75,
	0xaa, 0x57, 0x52, 0x28, 0x6b, 0xee, 0x15, 0xa0, 0x1d, 0x99, 0xd4, 0x88, 0xe9, 0xd9, 0x3e, 0x67,
	0x54, 0x6c, 0x37, 0xbd, 0x67, 0x70, 0x7a, 0x70, 0xf4, 0xce, 0x40, 0x6b, 0x1e, 0x44, 0x21, 0x1d,
	0x8c, 0x6f, 0xc7, 0x74, 0xe8, 0x1e, 0x79, 0x2d, 0x70, 0x32, 0x0f, 0x26, 0xc1, 0x74, 0x11, 0xb8,
	0x0d, 0xef, 0x02, 0xb4, 0x03, 0x3a, 0x22, 0xb3, 0x71, 0x4c, 0x97, 0x13, 0xfa, 0xb8, 0x98, 0x3e,
	0x0c, 0x23, 0xd7, 0xf1, 0x2e, 0xc1, 0x79, 0x7d, 0x0e, 0xef, 0xc8, 0x80, 0xde, 0xd3, 0x60, 0x16,
	0xb9, 0xcd, 0xfe, 0x57, 0x03, 0x74, 0x5f, 0xe5, 0x06, 0xfe, 0xdb, 0xbb, 0xef, 0x1d, 0x54, 0x08,
	0xcb, 0xb6, 0x61, 0xe3, 0xa9, 0xff, 0x6b, 0xe2, 0x72, 0x9d, 0x0a, 0x0e, 0x65, 0xc1, 0x7d, 0xce,
	0x84, 0xf9, 0x4b, 0xb5, 0x58, 0xbe, 0x52, 0x7f, 0x0c, 0x78, 0x63, 0xde, 0x77, 0xa7, 0x39, 0x22,
	0xe4, 0xc3, 0xe9, 0x8c, 0x6c, 0x14, 0xc9, 0x14, 0xb4, 0xb0, 0x44, 0x31, 0x82, 0xe5, 0x06, 0xea,
	0xb3, 0xe2, 0x13, 0x92, 0xa9, 0xa4, 0xe6, 0x93, 0x18, 0x25, 0x86, 0xff, 0x76, 0xba, 0xf6, 0x88,
	0x31, 0xc9, 0x14, 0xc6, 0xb5, 0x02, 0xe3, 0x18, 0x61, 0x6c, 0x34, 0x2f, 0xc7, 0xa6, 0x18, 0xfa,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x69, 0xb8, 0x8e, 0xf5, 0xd8, 0x01, 0x00, 0x00,
}
