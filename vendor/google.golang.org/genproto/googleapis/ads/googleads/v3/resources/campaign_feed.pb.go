// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/campaign_feed.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A campaign feed.
type CampaignFeed struct {
	// The resource name of the campaign feed.
	// Campaign feed resource names have the form:
	//
	// `customers/{customer_id}/campaignFeeds/{campaign_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed to which the CampaignFeed belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The campaign to which the CampaignFeed belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// campaign. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v3.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the CampaignFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the campaign feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignFeed) Reset()         { *m = CampaignFeed{} }
func (m *CampaignFeed) String() string { return proto.CompactTextString(m) }
func (*CampaignFeed) ProtoMessage()    {}
func (*CampaignFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f9f25847e06da48, []int{0}
}

func (m *CampaignFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeed.Unmarshal(m, b)
}
func (m *CampaignFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeed.Marshal(b, m, deterministic)
}
func (m *CampaignFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeed.Merge(m, src)
}
func (m *CampaignFeed) XXX_Size() int {
	return xxx_messageInfo_CampaignFeed.Size(m)
}
func (m *CampaignFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeed.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeed proto.InternalMessageInfo

func (m *CampaignFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *CampaignFeed) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *CampaignFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *CampaignFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignFeed)(nil), "google.ads.googleads.v3.resources.CampaignFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/campaign_feed.proto", fileDescriptor_6f9f25847e06da48)
}

var fileDescriptor_6f9f25847e06da48 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xbb, 0xeb, 0xa2, 0x63, 0x2d, 0xdd, 0x5c, 0xc5, 0x52, 0x64, 0xab, 0x2c, 0xec, 0xd5,
	0x4c, 0xd9, 0xa8, 0x48, 0xbc, 0x31, 0x2b, 0xb6, 0x20, 0x2a, 0x4b, 0x2a, 0x8b, 0xe8, 0x6a, 0x98,
	0x26, 0xb3, 0x69, 0x68, 0xe6, 0x87, 0x4c, 0x52, 0x29, 0xd2, 0x97, 0xf1, 0x4e, 0x1f, 0xc5, 0x47,
	0xe9, 0x43, 0x88, 0x4c, 0x32, 0x99, 0xfd, 0x91, 0x6d, 0x7b, 0xf7, 0xe5, 0x9c, 0xef, 0x3b, 0x73,
	0xce, 0xf9, 0x72, 0xc0, 0xb3, 0x84, 0xf3, 0x24, 0x23, 0x08, 0xc7, 0x12, 0xd5, 0x50, 0xa1, 0x73,
	0x17, 0xe5, 0x44, 0xf2, 0x32, 0x8f, 0x88, 0x44, 0x11, 0xa6, 0x02, 0xa7, 0x09, 0x0b, 0xe7, 0x84,
	0xc4, 0x50, 0xe4, 0xbc, 0xe0, 0xf6, 0x7e, 0xcd, 0x85, 0x38, 0x96, 0xd0, 0xc8, 0xe0, 0xb9, 0x0b,
	0x8d, 0x6c, 0xf7, 0xf9, 0xa6, 0xca, 0x11, 0xa7, 0x94, 0x33, 0x44, 0x71, 0x11, 0x9d, 0xa6, 0x2c,
	0x09, 0xe7, 0x25, 0x8b, 0x8a, 0x94, 0xb3, 0xba, 0xf4, 0xee, 0xd3, 0x4d, 0x3a, 0xc2, 0x4a, 0x2a,
	0x91, 0x6a, 0x22, 0xcc, 0x52, 0x76, 0x16, 0xca, 0x02, 0x17, 0xa5, 0xbc, 0x9d, 0x4a, 0x64, 0x38,
	0x22, 0xa7, 0x3c, 0x8b, 0x49, 0x1e, 0x16, 0x17, 0x82, 0x68, 0xd5, 0xc3, 0x46, 0x25, 0x52, 0x33,
	0xb0, 0x4e, 0x3d, 0xd2, 0xa9, 0xea, 0xeb, 0xa4, 0x9c, 0xa3, 0xef, 0x39, 0x16, 0x82, 0xe4, 0xcd,
	0x83, 0x7b, 0x4b, 0x52, 0xcc, 0x18, 0x2f, 0xb0, 0x9a, 0x41, 0x67, 0x1f, 0xff, 0xea, 0x80, 0xad,
	0xd7, 0x7a, 0x6f, 0x87, 0x84, 0xc4, 0xf6, 0x13, 0xf0, 0xa0, 0x79, 0x20, 0x64, 0x98, 0x12, 0xc7,
	0xea, 0x5b, 0xc3, 0x7b, 0xc1, 0x56, 0x13, 0xfc, 0x80, 0x29, 0xb1, 0x0f, 0x40, 0x47, 0x8d, 0xe7,
	0xb4, 0xfa, 0xd6, 0xf0, 0xfe, 0x68, 0x4f, 0x6f, 0x16, 0x36, 0x2d, 0xc0, 0xe3, 0x22, 0x4f, 0x59,
	0x32, 0xc5, 0x59, 0x49, 0x82, 0x8a, 0x69, 0xbf, 0x00, 0x77, 0x1b, 0x7b, 0x9c, 0xf6, 0x2d, 0x54,
	0x86, 0x6d, 0x73, 0xd0, 0x5b, 0x5f, 0x8a, 0x74, 0x3a, 0xfd, 0xf6, 0x70, 0x7b, 0x34, 0x86, 0x9b,
	0xdc, 0xad, 0x96, 0x09, 0x27, 0x0b, 0xdd, 0xc7, 0x0b, 0x41, 0xde, 0xb0, 0x92, 0xae, 0xc7, 0x82,
	0x1d, 0xb1, 0x1a, 0x90, 0xf6, 0x57, 0xd0, 0xfb, 0xcf, 0x72, 0xe7, 0x4e, 0xd5, 0xf3, 0xc1, 0xc6,
	0x07, 0xeb, 0x7f, 0x05, 0xbe, 0xd7, 0xc2, 0x43, 0xad, 0x0b, 0x76, 0xe8, 0x5a, 0xc4, 0xfe, 0x04,
	0xba, 0xf5, 0x0f, 0xe1, 0x74, 0xfb, 0xd6, 0x70, 0x7b, 0xf4, 0xea, 0x86, 0x21, 0x94, 0x2b, 0xef,
	0x52, 0x76, 0x76, 0x5c, 0x89, 0xaa, 0x19, 0x56, 0x43, 0x81, 0xae, 0xe7, 0x7d, 0xbb, 0xf2, 0xbf,
	0x80, 0xc1, 0xa2, 0x84, 0x46, 0x22, 0x95, 0xaa, 0x3d, 0xb4, 0x62, 0xf3, 0x28, 0x2a, 0x65, 0xc1,
	0x29, 0xc9, 0x25, 0xfa, 0xd1, 0xc0, 0x4b, 0x73, 0x41, 0x8a, 0xa2, 0x12, 0xcb, 0x07, 0x75, 0x39,
	0xfe, 0x6b, 0x81, 0x41, 0xc4, 0x29, 0xbc, 0xf1, 0xa4, 0xc6, 0xbd, 0xe5, 0xb7, 0x26, 0xca, 0xdf,
	0x89, 0xf5, 0xf9, 0xad, 0xd6, 0x25, 0x3c, 0xc3, 0x2c, 0x81, 0x3c, 0x4f, 0x50, 0x42, 0x58, 0xe5,
	0x3e, 0x5a, 0xb4, 0x7a, 0xcd, 0x81, 0xbf, 0x34, 0xe8, 0x67, 0xab, 0x7d, 0xe4, 0xfb, 0xbf, 0x5b,
	0xfb, 0x47, 0x75, 0x49, 0x3f, 0x96, 0xb0, 0x86, 0x0a, 0x4d, 0x5d, 0x18, 0x34, 0xcc, 0x3f, 0x0d,
	0x67, 0xe6, 0xc7, 0x72, 0x66, 0x38, 0xb3, 0xa9, 0x3b, 0x33, 0x9c, 0xab, 0xd6, 0xa0, 0x4e, 0x78,
	0x9e, 0x1f, 0x4b, 0xcf, 0x33, 0x2c, 0xcf, 0x9b, 0xba, 0x9e, 0x67, 0x78, 0x27, 0xdd, 0xaa, 0x59,
	0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xa1, 0xf2, 0xf2, 0x8c, 0x04, 0x00, 0x00,
}
