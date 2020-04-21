// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/detail_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [DetailPlacementViewService.GetDetailPlacementView][google.ads.googleads.v3.services.DetailPlacementViewService.GetDetailPlacementView].
type GetDetailPlacementViewRequest struct {
	// Required. The resource name of the Detail Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailPlacementViewRequest) Reset()         { *m = GetDetailPlacementViewRequest{} }
func (m *GetDetailPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailPlacementViewRequest) ProtoMessage()    {}
func (*GetDetailPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0191e49c6a7b6d2e, []int{0}
}

func (m *GetDetailPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetDetailPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDetailPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailPlacementViewRequest.Merge(m, src)
}
func (m *GetDetailPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDetailPlacementViewRequest.Size(m)
}
func (m *GetDetailPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailPlacementViewRequest proto.InternalMessageInfo

func (m *GetDetailPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDetailPlacementViewRequest)(nil), "google.ads.googleads.v3.services.GetDetailPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/detail_placement_view_service.proto", fileDescriptor_0191e49c6a7b6d2e)
}

var fileDescriptor_0191e49c6a7b6d2e = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xd0, 0x23, 0x2a, 0xb7, 0x1c, 0x57, 0x1c, 0x57, 0xcc,
	0x80, 0x81, 0x43, 0x46, 0x0e, 0x99, 0xe5, 0x60, 0xb5, 0x59, 0x96, 0x15, 0x52, 0x48, 0x20, 0xcc,
	0x26, 0x9f, 0x71, 0x20, 0x99, 0x89, 0x99, 0xd9, 0x6c, 0x21, 0x36, 0x16, 0xfe, 0x01, 0x6b, 0x1b,
	0x4b, 0x7f, 0xca, 0xb6, 0x76, 0x56, 0x0a, 0x56, 0xfe, 0x0a, 0xc9, 0x4e, 0x26, 0xbb, 0x0b, 0x1b,
	0xb7, 0x7b, 0xe4, 0xbd, 0xef, 0xbd, 0x6f, 0xbe, 0x17, 0xef, 0x36, 0x97, 0x32, 0x2f, 0x00, 0xb3,
	0x4c, 0x61, 0x03, 0x5b, 0xd4, 0x84, 0x58, 0x41, 0xdd, 0xf0, 0x14, 0x14, 0xce, 0x40, 0x33, 0x5e,
	0x24, 0x55, 0xc1, 0x52, 0x28, 0x41, 0xe8, 0xa4, 0xe1, 0xb0, 0x4a, 0x3a, 0x1a, 0x55, 0xb5, 0xd4,
	0xd2, 0x1f, 0x99, 0x51, 0xc4, 0x32, 0x85, 0x7a, 0x17, 0xd4, 0x84, 0xc8, 0xba, 0x04, 0x37, 0x43,
	0x39, 0x35, 0x28, 0xb9, 0xac, 0x07, 0x83, 0x4c, 0x40, 0xf0, 0xd8, 0x8e, 0x57, 0x1c, 0x33, 0x21,
	0xa4, 0x66, 0x9a, 0x4b, 0xa1, 0x3a, 0xf6, 0xe1, 0x0e, 0x9b, 0x16, 0x1c, 0x84, 0xee, 0x88, 0xb3,
	0x1d, 0xe2, 0x2d, 0x87, 0x22, 0x4b, 0x16, 0xf0, 0x8e, 0x35, 0x5c, 0xd6, 0x46, 0x70, 0xfe, 0xca,
	0x7b, 0x32, 0x01, 0x7d, 0xbb, 0x49, 0x9e, 0xd9, 0xe0, 0x88, 0xc3, 0x6a, 0x0e, 0xef, 0x97, 0xa0,
	0xb4, 0x7f, 0xe9, 0xdd, 0xb7, 0x1b, 0x26, 0x82, 0x95, 0x70, 0xea, 0x8c, 0x9c, 0xcb, 0xbb, 0xe3,
	0x93, 0x5f, 0xd4, 0x9d, 0xdf, 0xb3, 0xcc, 0x94, 0x95, 0xf0, 0xf4, 0xab, 0xeb, 0x05, 0x07, 0x8c,
	0x5e, 0x9b, 0x0b, 0xf8, 0xbf, 0x1d, 0xef, 0xc1, 0xe1, 0x28, 0xff, 0x05, 0x3a, 0x76, 0x3e, 0xf4,
	0xdf, 0x25, 0x83, 0xeb, 0x41, 0x83, 0xfe, 0xba, 0xe8, 0xc0, 0xf8, 0xf9, 0xf4, 0x27, 0xdd, 0x7f,
	0xdd, 0xa7, 0x1f, 0x7f, 0xbe, 0xb8, 0xcf, 0xfc, 0xeb, 0xb6, 0x98, 0x0f, 0x7b, 0xcc, 0x4d, 0xba,
	0x54, 0x5a, 0x96, 0x50, 0x2b, 0x7c, 0xd5, 0x35, 0xb5, 0xe7, 0xa5, 0xf0, 0xd5, 0xc7, 0xe0, 0xd1,
	0x9a, 0x9e, 0x6e, 0xe3, 0x3b, 0x54, 0x71, 0x85, 0x52, 0x59, 0x8e, 0x3f, 0xbb, 0xde, 0x45, 0x2a,
	0xcb, 0xa3, 0x6f, 0x1d, 0x9f, 0x0d, 0x5f, 0x71, 0xd6, 0x96, 0x36, 0x73, 0xde, 0xbc, 0xec, 0x4c,
	0x72, 0x59, 0x30, 0x91, 0x23, 0x59, 0xe7, 0x38, 0x07, 0xb1, 0xa9, 0x14, 0x6f, 0x63, 0x87, 0x7f,
	0xea, 0xe7, 0x16, 0x7c, 0x73, 0x4f, 0x26, 0x94, 0x7e, 0x77, 0x47, 0x13, 0x63, 0x48, 0x33, 0x85,
	0x0c, 0x6c, 0x51, 0x14, 0xa2, 0x2e, 0x58, 0xad, 0xad, 0x24, 0xa6, 0x99, 0x8a, 0x7b, 0x49, 0x1c,
	0x85, 0xb1, 0x95, 0xfc, 0x75, 0x2f, 0xcc, 0x77, 0x42, 0x68, 0xa6, 0x08, 0xe9, 0x45, 0x84, 0x44,
	0x21, 0x21, 0x56, 0xb6, 0xb8, 0xb3, 0xd9, 0x33, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x67, 0xa6,
	0xa3, 0x82, 0x7b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DetailPlacementViewServiceClient is the client API for DetailPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailPlacementViewServiceClient interface {
	// Returns the requested Detail Placement view in full detail.
	GetDetailPlacementView(ctx context.Context, in *GetDetailPlacementViewRequest, opts ...grpc.CallOption) (*resources.DetailPlacementView, error)
}

type detailPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailPlacementViewServiceClient(cc grpc.ClientConnInterface) DetailPlacementViewServiceClient {
	return &detailPlacementViewServiceClient{cc}
}

func (c *detailPlacementViewServiceClient) GetDetailPlacementView(ctx context.Context, in *GetDetailPlacementViewRequest, opts ...grpc.CallOption) (*resources.DetailPlacementView, error) {
	out := new(resources.DetailPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.DetailPlacementViewService/GetDetailPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailPlacementViewServiceServer is the server API for DetailPlacementViewService service.
type DetailPlacementViewServiceServer interface {
	// Returns the requested Detail Placement view in full detail.
	GetDetailPlacementView(context.Context, *GetDetailPlacementViewRequest) (*resources.DetailPlacementView, error)
}

// UnimplementedDetailPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailPlacementViewServiceServer struct {
}

func (*UnimplementedDetailPlacementViewServiceServer) GetDetailPlacementView(ctx context.Context, req *GetDetailPlacementViewRequest) (*resources.DetailPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailPlacementView not implemented")
}

func RegisterDetailPlacementViewServiceServer(s *grpc.Server, srv DetailPlacementViewServiceServer) {
	s.RegisterService(&_DetailPlacementViewService_serviceDesc, srv)
}

func _DetailPlacementViewService_GetDetailPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailPlacementViewServiceServer).GetDetailPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.DetailPlacementViewService/GetDetailPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailPlacementViewServiceServer).GetDetailPlacementView(ctx, req.(*GetDetailPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetailPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.DetailPlacementViewService",
	HandlerType: (*DetailPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailPlacementView",
			Handler:    _DetailPlacementViewService_GetDetailPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/detail_placement_view_service.proto",
}
