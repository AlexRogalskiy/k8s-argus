// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/label_error.proto

package errors

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

// Enum describing possible label errors.
type LabelErrorEnum_LabelError int32

const (
	// Enum unspecified.
	LabelErrorEnum_UNSPECIFIED LabelErrorEnum_LabelError = 0
	// The received error code is not known in this version.
	LabelErrorEnum_UNKNOWN LabelErrorEnum_LabelError = 1
	// An inactive label cannot be applied.
	LabelErrorEnum_CANNOT_APPLY_INACTIVE_LABEL LabelErrorEnum_LabelError = 2
	// A label cannot be applied to a disabled ad group criterion.
	LabelErrorEnum_CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION LabelErrorEnum_LabelError = 3
	// A label cannot be applied to a negative ad group criterion.
	LabelErrorEnum_CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION LabelErrorEnum_LabelError = 4
	// Cannot apply more than 50 labels per resource.
	LabelErrorEnum_EXCEEDED_LABEL_LIMIT_PER_TYPE LabelErrorEnum_LabelError = 5
	// Labels from a manager account cannot be applied to campaign, ad group,
	// ad group ad, or ad group criterion resources.
	LabelErrorEnum_INVALID_RESOURCE_FOR_MANAGER_LABEL LabelErrorEnum_LabelError = 6
	// Label names must be unique.
	LabelErrorEnum_DUPLICATE_NAME LabelErrorEnum_LabelError = 7
	// Label names cannot be empty.
	LabelErrorEnum_INVALID_LABEL_NAME LabelErrorEnum_LabelError = 8
	// Labels cannot be applied to a draft.
	LabelErrorEnum_CANNOT_ATTACH_LABEL_TO_DRAFT LabelErrorEnum_LabelError = 9
	// Labels not from a manager account cannot be applied to the customer
	// resource.
	LabelErrorEnum_CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER LabelErrorEnum_LabelError = 10
)

var LabelErrorEnum_LabelError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "CANNOT_APPLY_INACTIVE_LABEL",
	3:  "CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION",
	4:  "CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION",
	5:  "EXCEEDED_LABEL_LIMIT_PER_TYPE",
	6:  "INVALID_RESOURCE_FOR_MANAGER_LABEL",
	7:  "DUPLICATE_NAME",
	8:  "INVALID_LABEL_NAME",
	9:  "CANNOT_ATTACH_LABEL_TO_DRAFT",
	10: "CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER",
}

var LabelErrorEnum_LabelError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_APPLY_INACTIVE_LABEL": 2,
	"CANNOT_APPLY_LABEL_TO_DISABLED_AD_GROUP_CRITERION": 3,
	"CANNOT_APPLY_LABEL_TO_NEGATIVE_AD_GROUP_CRITERION": 4,
	"EXCEEDED_LABEL_LIMIT_PER_TYPE":                     5,
	"INVALID_RESOURCE_FOR_MANAGER_LABEL":                6,
	"DUPLICATE_NAME":                                    7,
	"INVALID_LABEL_NAME":                                8,
	"CANNOT_ATTACH_LABEL_TO_DRAFT":                      9,
	"CANNOT_ATTACH_NON_MANAGER_LABEL_TO_CUSTOMER":       10,
}

func (x LabelErrorEnum_LabelError) String() string {
	return proto.EnumName(LabelErrorEnum_LabelError_name, int32(x))
}

func (LabelErrorEnum_LabelError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ab0c51c11d34350, []int{0, 0}
}

// Container for enum describing possible label errors.
type LabelErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelErrorEnum) Reset()         { *m = LabelErrorEnum{} }
func (m *LabelErrorEnum) String() string { return proto.CompactTextString(m) }
func (*LabelErrorEnum) ProtoMessage()    {}
func (*LabelErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ab0c51c11d34350, []int{0}
}

func (m *LabelErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelErrorEnum.Unmarshal(m, b)
}
func (m *LabelErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelErrorEnum.Marshal(b, m, deterministic)
}
func (m *LabelErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelErrorEnum.Merge(m, src)
}
func (m *LabelErrorEnum) XXX_Size() int {
	return xxx_messageInfo_LabelErrorEnum.Size(m)
}
func (m *LabelErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LabelErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.LabelErrorEnum_LabelError", LabelErrorEnum_LabelError_name, LabelErrorEnum_LabelError_value)
	proto.RegisterType((*LabelErrorEnum)(nil), "google.ads.googleads.v3.errors.LabelErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/label_error.proto", fileDescriptor_7ab0c51c11d34350)
}

var fileDescriptor_7ab0c51c11d34350 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x40, 0x69, 0x02, 0x2d, 0x6c, 0xa5, 0xd6, 0xda, 0x03, 0x07, 0x28, 0x05, 0x72, 0xe0, 0x82,
	0x64, 0x83, 0x2c, 0x2e, 0xe6, 0x34, 0xb1, 0x27, 0x66, 0x85, 0xb3, 0xb6, 0x36, 0x6b, 0x43, 0x51,
	0xa4, 0x95, 0x8b, 0x23, 0x2b, 0x52, 0xea, 0x8d, 0xec, 0xd0, 0x1f, 0xe0, 0x1f, 0xf8, 0x00, 0x8e,
	0x7c, 0x0a, 0x9f, 0xd2, 0xaf, 0x40, 0xf6, 0xc6, 0x29, 0x91, 0x28, 0x27, 0x8f, 0x47, 0xef, 0xcd,
	0xcc, 0x6a, 0x86, 0xbc, 0x29, 0xb5, 0x2e, 0x57, 0x0b, 0x27, 0x2f, 0x1a, 0xc7, 0x84, 0x6d, 0x74,
	0xed, 0x3a, 0x8b, 0xba, 0xd6, 0x75, 0xe3, 0xac, 0xf2, 0xcb, 0xc5, 0x4a, 0x75, 0x3f, 0xf6, 0xba,
	0xd6, 0x1b, 0x4d, 0xcf, 0x0d, 0x66, 0xe7, 0x45, 0x63, 0xef, 0x0c, 0xfb, 0xda, 0xb5, 0x8d, 0xf1,
	0xe4, 0xac, 0xaf, 0xb8, 0x5e, 0x3a, 0x79, 0x55, 0xe9, 0x4d, 0xbe, 0x59, 0xea, 0xaa, 0x31, 0xf6,
	0xe8, 0xc7, 0x90, 0x9c, 0x44, 0x6d, 0x4d, 0x6c, 0x69, 0xac, 0xbe, 0x5d, 0x8d, 0xbe, 0x0f, 0x09,
	0xb9, 0x4d, 0xd1, 0x53, 0x72, 0x9c, 0xf2, 0x59, 0x82, 0x3e, 0x9b, 0x30, 0x0c, 0xac, 0x7b, 0xf4,
	0x98, 0x1c, 0xa5, 0xfc, 0x23, 0x8f, 0x3f, 0x71, 0xeb, 0x80, 0x3e, 0x27, 0x4f, 0x7d, 0xe0, 0x3c,
	0x96, 0x0a, 0x92, 0x24, 0xba, 0x50, 0x8c, 0x83, 0x2f, 0x59, 0x86, 0x2a, 0x82, 0x31, 0x46, 0xd6,
	0x80, 0xbe, 0x23, 0x6f, 0xf7, 0x80, 0x2e, 0xaf, 0x64, 0xac, 0x02, 0x36, 0x83, 0x71, 0x84, 0x81,
	0x82, 0x40, 0x85, 0x22, 0x4e, 0x13, 0xe5, 0x0b, 0x26, 0x51, 0xb0, 0x98, 0x5b, 0xc3, 0xbb, 0x35,
	0x8e, 0x21, 0x74, 0x0d, 0xfe, 0xa1, 0xdd, 0xa7, 0x2f, 0xc9, 0x33, 0xfc, 0xec, 0x23, 0x06, 0x18,
	0x6c, 0x95, 0x88, 0x4d, 0x99, 0x54, 0x09, 0x0a, 0x25, 0x2f, 0x12, 0xb4, 0x1e, 0xd0, 0x57, 0x64,
	0xc4, 0x78, 0x06, 0x11, 0x0b, 0x94, 0xc0, 0x59, 0x9c, 0x0a, 0x1f, 0xd5, 0x24, 0x16, 0x6a, 0x0a,
	0x1c, 0x42, 0x14, 0xdb, 0xc1, 0x0f, 0x29, 0x25, 0x27, 0x41, 0x9a, 0x44, 0xcc, 0x07, 0x89, 0x8a,
	0xc3, 0x14, 0xad, 0x23, 0xfa, 0x98, 0xd0, 0xde, 0x35, 0xd5, 0xbb, 0xfc, 0x43, 0xfa, 0x82, 0x9c,
	0xf5, 0xd3, 0x4a, 0x09, 0xfe, 0x87, 0xbf, 0x5e, 0x29, 0x60, 0x22, 0xad, 0x47, 0xd4, 0x21, 0xaf,
	0xf7, 0x09, 0x1e, 0xf3, 0xfd, 0x96, 0x2d, 0xed, 0xa7, 0x33, 0x19, 0x4f, 0x51, 0x58, 0x64, 0x7c,
	0x73, 0x40, 0x46, 0x5f, 0xf5, 0x95, 0xfd, 0xff, 0xed, 0x8e, 0x4f, 0x6f, 0x37, 0x95, 0xb4, 0x0b,
	0x4d, 0x0e, 0xbe, 0x04, 0x5b, 0xa5, 0xd4, 0xab, 0xbc, 0x2a, 0x6d, 0x5d, 0x97, 0x4e, 0xb9, 0xa8,
	0xba, 0x75, 0xf7, 0x27, 0xb5, 0x5e, 0x36, 0x77, 0x5d, 0xd8, 0x7b, 0xf3, 0xf9, 0x39, 0x18, 0x86,
	0x00, 0xbf, 0x06, 0xe7, 0xa1, 0x29, 0x06, 0x45, 0x63, 0x9b, 0xb0, 0x8d, 0x32, 0xd7, 0xee, 0x5a,
	0x36, 0xbf, 0x7b, 0x60, 0x0e, 0x45, 0x33, 0xdf, 0x01, 0xf3, 0xcc, 0x9d, 0x1b, 0xe0, 0x66, 0x30,
	0x32, 0x59, 0xcf, 0x83, 0xa2, 0xf1, 0xbc, 0x1d, 0xe2, 0x79, 0x99, 0xeb, 0x79, 0x06, 0xba, 0x3c,
	0xec, 0xa6, 0x73, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x13, 0xda, 0x7d, 0x22, 0xfe, 0x02, 0x00,
	0x00,
}
