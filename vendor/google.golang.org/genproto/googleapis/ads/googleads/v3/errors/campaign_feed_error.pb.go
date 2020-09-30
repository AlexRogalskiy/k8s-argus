// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/campaign_feed_error.proto

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

// Enum describing possible campaign feed errors.
type CampaignFeedErrorEnum_CampaignFeedError int32

const (
	// Enum unspecified.
	CampaignFeedErrorEnum_UNSPECIFIED CampaignFeedErrorEnum_CampaignFeedError = 0
	// The received error code is not known in this version.
	CampaignFeedErrorEnum_UNKNOWN CampaignFeedErrorEnum_CampaignFeedError = 1
	// An active feed already exists for this campaign and placeholder type.
	CampaignFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 2
	// The specified feed is removed.
	CampaignFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CampaignFeedErrorEnum_CampaignFeedError = 4
	// The CampaignFeed already exists. UPDATE should be used to modify the
	// existing CampaignFeed.
	CampaignFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 5
	// Cannot update removed campaign feed.
	CampaignFeedErrorEnum_CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED CampaignFeedErrorEnum_CampaignFeedError = 6
	// Invalid placeholder type.
	CampaignFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 7
	// Feed mapping for this placeholder type does not exist.
	CampaignFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CampaignFeedErrorEnum_CampaignFeedError = 8
	// Location CampaignFeeds cannot be created unless there is a location
	// CustomerFeed for the specified feed.
	CampaignFeedErrorEnum_NO_EXISTING_LOCATION_CUSTOMER_FEED CampaignFeedErrorEnum_CampaignFeedError = 9
)

var CampaignFeedErrorEnum_CampaignFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	4: "CANNOT_CREATE_FOR_REMOVED_FEED",
	5: "CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED",
	6: "CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED",
	7: "INVALID_PLACEHOLDER_TYPE",
	8: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	9: "NO_EXISTING_LOCATION_CUSTOMER_FEED",
}

var CampaignFeedErrorEnum_CampaignFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":     2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":               4,
	"CANNOT_CREATE_ALREADY_EXISTING_CAMPAIGN_FEED": 5,
	"CANNOT_MODIFY_REMOVED_CAMPAIGN_FEED":          6,
	"INVALID_PLACEHOLDER_TYPE":                     7,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":     8,
	"NO_EXISTING_LOCATION_CUSTOMER_FEED":           9,
}

func (x CampaignFeedErrorEnum_CampaignFeedError) String() string {
	return proto.EnumName(CampaignFeedErrorEnum_CampaignFeedError_name, int32(x))
}

func (CampaignFeedErrorEnum_CampaignFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4307471cace193a, []int{0, 0}
}

// Container for enum describing possible campaign feed errors.
type CampaignFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignFeedErrorEnum) Reset()         { *m = CampaignFeedErrorEnum{} }
func (m *CampaignFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignFeedErrorEnum) ProtoMessage()    {}
func (*CampaignFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4307471cace193a, []int{0}
}

func (m *CampaignFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeedErrorEnum.Unmarshal(m, b)
}
func (m *CampaignFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeedErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeedErrorEnum.Merge(m, src)
}
func (m *CampaignFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignFeedErrorEnum.Size(m)
}
func (m *CampaignFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.CampaignFeedErrorEnum_CampaignFeedError", CampaignFeedErrorEnum_CampaignFeedError_name, CampaignFeedErrorEnum_CampaignFeedError_value)
	proto.RegisterType((*CampaignFeedErrorEnum)(nil), "google.ads.googleads.v3.errors.CampaignFeedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/campaign_feed_error.proto", fileDescriptor_a4307471cace193a)
}

var fileDescriptor_a4307471cace193a = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0x80, 0x19, 0xf0, 0x2c, 0x28, 0x96, 0x40, 0x08, 0x8d, 0xba, 0x08, 0x12, 0xb0,
	0x18, 0x39, 0x48, 0xd9, 0xa0, 0xb0, 0xf2, 0x24, 0x4e, 0xb1, 0x48, 0xec, 0x28, 0x49, 0x03, 0x45,
	0x95, 0xac, 0x30, 0x09, 0x51, 0xa5, 0x69, 0x5c, 0xc5, 0xa5, 0x07, 0x62, 0xc9, 0x51, 0x58, 0x72,
	0x0c, 0xc4, 0x86, 0x1b, 0x20, 0xc7, 0x6d, 0x50, 0x55, 0x98, 0x55, 0x7e, 0x39, 0xdf, 0xff, 0xbf,
	0x67, 0xbf, 0x07, 0x5e, 0x37, 0x52, 0x36, 0xd7, 0xb5, 0x53, 0x56, 0xca, 0x31, 0x52, 0xab, 0xad,
	0xeb, 0xd4, 0x5d, 0x27, 0x3b, 0xe5, 0x5c, 0x95, 0xab, 0x75, 0xb9, 0x6c, 0x5a, 0xf1, 0xb9, 0xae,
	0x2b, 0xd1, 0x1f, 0xa2, 0x75, 0x27, 0x37, 0x12, 0x4e, 0x0c, 0x8e, 0xca, 0x4a, 0xa1, 0xc1, 0x89,
	0xb6, 0x2e, 0x32, 0xce, 0xa7, 0xe7, 0xfb, 0xe4, 0xf5, 0xd2, 0x29, 0xdb, 0x56, 0x6e, 0xca, 0xcd,
	0x52, 0xb6, 0xca, 0xb8, 0xed, 0x5f, 0x16, 0x78, 0xe4, 0xef, 0xb2, 0xc3, 0xba, 0xae, 0x88, 0x36,
	0x91, 0xf6, 0xcb, 0xca, 0xfe, 0x61, 0x81, 0x87, 0x47, 0x7f, 0xe0, 0x03, 0x70, 0x36, 0x63, 0x59,
	0x42, 0x7c, 0x1a, 0x52, 0x12, 0x8c, 0x6f, 0xc1, 0x33, 0x70, 0x3a, 0x63, 0xef, 0x18, 0x7f, 0xcf,
	0xc6, 0x23, 0x78, 0x01, 0x5e, 0x86, 0x84, 0x04, 0x02, 0x47, 0x29, 0xc1, 0xc1, 0x5c, 0x90, 0x0f,
	0x34, 0xcb, 0x33, 0x11, 0xf2, 0x54, 0x24, 0x11, 0xf6, 0xc9, 0x5b, 0x1e, 0x05, 0x24, 0x15, 0xf9,
	0x3c, 0x21, 0x63, 0x0b, 0xda, 0x60, 0xe2, 0x63, 0xc6, 0x78, 0x2e, 0xfc, 0x94, 0xe0, 0x9c, 0xf4,
	0x5c, 0x4a, 0x62, 0x5e, 0x90, 0x40, 0xe8, 0x9c, 0xf1, 0x1d, 0xf8, 0x0a, 0x5c, 0x1c, 0x32, 0x07,
	0xd1, 0x94, 0x4d, 0x85, 0x8f, 0xe3, 0x04, 0xd3, 0x29, 0x33, 0x8e, 0xbb, 0xf0, 0x05, 0x78, 0xb6,
	0x73, 0xc4, 0x3c, 0xa0, 0xe1, 0x7c, 0x48, 0x3c, 0x04, 0x4f, 0xe0, 0x39, 0x78, 0x42, 0x59, 0x81,
	0x23, 0x1a, 0x1c, 0x37, 0x77, 0xaa, 0xaf, 0x12, 0xd3, 0x2c, 0xd3, 0x15, 0x34, 0x1f, 0xe3, 0x24,
	0xe9, 0xf5, 0xbf, 0xae, 0x72, 0x0f, 0x3e, 0x07, 0x36, 0xe3, 0x7f, 0x7b, 0x8a, 0xb8, 0x8f, 0x73,
	0xca, 0x99, 0xf0, 0x67, 0x59, 0xce, 0x63, 0x92, 0x9a, 0x9a, 0xf7, 0x2f, 0x7f, 0x8f, 0x80, 0x7d,
	0x25, 0x57, 0xe8, 0xe6, 0x99, 0x5d, 0x3e, 0x3e, 0x7a, 0xf8, 0x44, 0x4f, 0x2b, 0x19, 0x7d, 0x0c,
	0x76, 0xce, 0x46, 0x5e, 0x97, 0x6d, 0x83, 0x64, 0xd7, 0x38, 0x4d, 0xdd, 0xf6, 0xb3, 0xdc, 0xef,
	0xcd, 0x7a, 0xa9, 0xfe, 0xb7, 0x46, 0x6f, 0xcc, 0xe7, 0xab, 0x75, 0x7b, 0x8a, 0xf1, 0x37, 0x6b,
	0x32, 0x35, 0x61, 0xb8, 0x52, 0xc8, 0x48, 0xad, 0x0a, 0x17, 0xf5, 0x25, 0xd5, 0xf7, 0x3d, 0xb0,
	0xc0, 0x95, 0x5a, 0x0c, 0xc0, 0xa2, 0x70, 0x17, 0x06, 0xf8, 0x69, 0xd9, 0xe6, 0xd4, 0xf3, 0x70,
	0xa5, 0x3c, 0x6f, 0x40, 0x3c, 0xaf, 0x70, 0x3d, 0xcf, 0x40, 0x9f, 0x4e, 0xfa, 0xee, 0xdc, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x14, 0x3d, 0x9c, 0xe3, 0x02, 0x00, 0x00,
}
