// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/bidding_source.proto

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

// Indicates where a bid or target is defined. For example, an ad group
// criterion may define a cpc bid directly, or it can inherit its cpc bid from
// the ad group.
type BiddingSourceEnum_BiddingSource int32

const (
	// Not specified.
	BiddingSourceEnum_UNSPECIFIED BiddingSourceEnum_BiddingSource = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingSourceEnum_UNKNOWN BiddingSourceEnum_BiddingSource = 1
	// Effective bid or target is inherited from campaign bidding strategy.
	BiddingSourceEnum_CAMPAIGN_BIDDING_STRATEGY BiddingSourceEnum_BiddingSource = 5
	// The bid or target is defined on the ad group.
	BiddingSourceEnum_AD_GROUP BiddingSourceEnum_BiddingSource = 6
	// The bid or target is defined on the ad group criterion.
	BiddingSourceEnum_AD_GROUP_CRITERION BiddingSourceEnum_BiddingSource = 7
)

var BiddingSourceEnum_BiddingSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "CAMPAIGN_BIDDING_STRATEGY",
	6: "AD_GROUP",
	7: "AD_GROUP_CRITERION",
}

var BiddingSourceEnum_BiddingSource_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"CAMPAIGN_BIDDING_STRATEGY": 5,
	"AD_GROUP":                  6,
	"AD_GROUP_CRITERION":        7,
}

func (x BiddingSourceEnum_BiddingSource) String() string {
	return proto.EnumName(BiddingSourceEnum_BiddingSource_name, int32(x))
}

func (BiddingSourceEnum_BiddingSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6307df4d287cd0d, []int{0, 0}
}

// Container for enum describing possible bidding sources.
type BiddingSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingSourceEnum) Reset()         { *m = BiddingSourceEnum{} }
func (m *BiddingSourceEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingSourceEnum) ProtoMessage()    {}
func (*BiddingSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6307df4d287cd0d, []int{0}
}

func (m *BiddingSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingSourceEnum.Unmarshal(m, b)
}
func (m *BiddingSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingSourceEnum.Marshal(b, m, deterministic)
}
func (m *BiddingSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingSourceEnum.Merge(m, src)
}
func (m *BiddingSourceEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingSourceEnum.Size(m)
}
func (m *BiddingSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.BiddingSourceEnum_BiddingSource", BiddingSourceEnum_BiddingSource_name, BiddingSourceEnum_BiddingSource_value)
	proto.RegisterType((*BiddingSourceEnum)(nil), "google.ads.googleads.v3.enums.BiddingSourceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/bidding_source.proto", fileDescriptor_e6307df4d287cd0d)
}

var fileDescriptor_e6307df4d287cd0d = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0xb5, 0x13, 0x37, 0xc9, 0x14, 0x6b, 0x2e, 0x04, 0xc5, 0x5d, 0x6c, 0x0f, 0x90, 0x82, 0xbd,
	0x8b, 0x57, 0xe9, 0x5a, 0x4b, 0x11, 0xb3, 0xd2, 0xfd, 0x88, 0x52, 0x28, 0xdd, 0x52, 0x42, 0x61,
	0x4b, 0x46, 0xb3, 0xed, 0x15, 0x7c, 0x0f, 0x2f, 0x7d, 0x14, 0x5f, 0x44, 0xf0, 0x29, 0xa4, 0x89,
	0x1d, 0xec, 0x42, 0x6f, 0xc2, 0xc9, 0xf7, 0x9d, 0x73, 0xf8, 0xce, 0x01, 0x77, 0x5c, 0x4a, 0xbe,
	0x2c, 0x9c, 0x9c, 0x29, 0xc7, 0xc0, 0x1a, 0xed, 0x5c, 0xa7, 0x10, 0xdb, 0x95, 0x72, 0xe6, 0x25,
	0x63, 0xa5, 0xe0, 0x99, 0x92, 0xdb, 0x6a, 0x51, 0xa0, 0x75, 0x25, 0x37, 0x12, 0xf6, 0x0c, 0x11,
	0xe5, 0x4c, 0xa1, 0xbd, 0x06, 0xed, 0x5c, 0xa4, 0x35, 0x37, 0xb7, 0x8d, 0xe5, 0xba, 0x74, 0x72,
	0x21, 0xe4, 0x26, 0xdf, 0x94, 0x52, 0x28, 0x23, 0x1e, 0xbc, 0x59, 0xe0, 0xd2, 0x33, 0xae, 0x63,
	0x6d, 0x1a, 0x88, 0xed, 0x6a, 0x50, 0x81, 0xf3, 0x83, 0x21, 0xbc, 0x00, 0xdd, 0x29, 0x1d, 0xc7,
	0xc1, 0x30, 0x7a, 0x88, 0x02, 0xdf, 0x3e, 0x82, 0x5d, 0xd0, 0x99, 0xd2, 0x47, 0x3a, 0x7a, 0xa6,
	0xb6, 0x05, 0x7b, 0xe0, 0x7a, 0x48, 0x9e, 0x62, 0x12, 0x85, 0x34, 0xf3, 0x22, 0xdf, 0x8f, 0x68,
	0x98, 0x8d, 0x27, 0x09, 0x99, 0x04, 0xe1, 0x8b, 0x7d, 0x02, 0xcf, 0xc0, 0x29, 0xf1, 0xb3, 0x30,
	0x19, 0x4d, 0x63, 0xbb, 0x0d, 0xaf, 0x00, 0x6c, 0x7e, 0xd9, 0x30, 0x89, 0x26, 0x41, 0x12, 0x8d,
	0xa8, 0xdd, 0xf1, 0xbe, 0x2c, 0xd0, 0x5f, 0xc8, 0x15, 0xfa, 0x37, 0x8d, 0x07, 0x0f, 0xee, 0x8a,
	0xeb, 0x0c, 0xb1, 0xf5, 0xea, 0xfd, 0x8a, 0xb8, 0x5c, 0xe6, 0x82, 0x23, 0x59, 0x71, 0x87, 0x17,
	0x42, 0x27, 0x6c, 0x6a, 0x5c, 0x97, 0xea, 0x8f, 0x56, 0xef, 0xf5, 0xfb, 0xde, 0x3a, 0x0e, 0x09,
	0xf9, 0x68, 0xf5, 0x42, 0x63, 0x45, 0x98, 0x42, 0x06, 0xd6, 0x68, 0xe6, 0xa2, 0xba, 0x18, 0xf5,
	0xd9, 0xec, 0x53, 0xc2, 0x54, 0xba, 0xdf, 0xa7, 0x33, 0x37, 0xd5, 0xfb, 0xef, 0x56, 0xdf, 0x0c,
	0x31, 0x26, 0x4c, 0x61, 0xbc, 0x67, 0x60, 0x3c, 0x73, 0x31, 0xd6, 0x9c, 0x79, 0x5b, 0x1f, 0xe6,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x97, 0x6b, 0xfe, 0x2a, 0xed, 0x01, 0x00, 0x00,
}
