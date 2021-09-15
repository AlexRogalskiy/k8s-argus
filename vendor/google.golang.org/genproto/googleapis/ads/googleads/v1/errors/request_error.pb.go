// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/request_error.proto

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

// Enum describing possible request errors.
type RequestErrorEnum_RequestError int32

const (
	// Enum unspecified.
	RequestErrorEnum_UNSPECIFIED RequestErrorEnum_RequestError = 0
	// The received error code is not known in this version.
	RequestErrorEnum_UNKNOWN RequestErrorEnum_RequestError = 1
	// Resource name is required for this request.
	RequestErrorEnum_RESOURCE_NAME_MISSING RequestErrorEnum_RequestError = 3
	// Resource name provided is malformed.
	RequestErrorEnum_RESOURCE_NAME_MALFORMED RequestErrorEnum_RequestError = 4
	// Resource name provided is malformed.
	RequestErrorEnum_BAD_RESOURCE_ID RequestErrorEnum_RequestError = 17
	// Customer ID is invalid.
	RequestErrorEnum_INVALID_CUSTOMER_ID RequestErrorEnum_RequestError = 16
	// Mutate operation should have either create, update, or remove specified.
	RequestErrorEnum_OPERATION_REQUIRED RequestErrorEnum_RequestError = 5
	// Requested resource not found.
	RequestErrorEnum_RESOURCE_NOT_FOUND RequestErrorEnum_RequestError = 6
	// Next page token specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_TOKEN RequestErrorEnum_RequestError = 7
	// Next page token specified in user request has expired.
	RequestErrorEnum_EXPIRED_PAGE_TOKEN RequestErrorEnum_RequestError = 8
	// Page size specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_SIZE RequestErrorEnum_RequestError = 22
	// Required field is missing.
	RequestErrorEnum_REQUIRED_FIELD_MISSING RequestErrorEnum_RequestError = 9
	// The field cannot be modified because it's immutable. It's also possible
	// that the field can be modified using 'create' operation but not 'update'.
	RequestErrorEnum_IMMUTABLE_FIELD RequestErrorEnum_RequestError = 11
	// Received too many entries in request.
	RequestErrorEnum_TOO_MANY_MUTATE_OPERATIONS RequestErrorEnum_RequestError = 13
	// Request cannot be executed by a manager account.
	RequestErrorEnum_CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT RequestErrorEnum_RequestError = 14
	// Mutate request was attempting to modify a readonly field.
	// For instance, Budget fields can be requested for Ad Group,
	// but are read-only for adGroups:mutate.
	RequestErrorEnum_CANNOT_MODIFY_FOREIGN_FIELD RequestErrorEnum_RequestError = 15
	// Enum value is not permitted.
	RequestErrorEnum_INVALID_ENUM_VALUE RequestErrorEnum_RequestError = 18
	// The developer-token parameter is required for all requests.
	RequestErrorEnum_DEVELOPER_TOKEN_PARAMETER_MISSING RequestErrorEnum_RequestError = 19
	// The login-customer-id parameter is required for this request.
	RequestErrorEnum_LOGIN_CUSTOMER_ID_PARAMETER_MISSING RequestErrorEnum_RequestError = 20
	// page_token is set in the validate only request
	RequestErrorEnum_VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN RequestErrorEnum_RequestError = 21
)

var RequestErrorEnum_RequestError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NAME_MISSING",
	4:  "RESOURCE_NAME_MALFORMED",
	17: "BAD_RESOURCE_ID",
	16: "INVALID_CUSTOMER_ID",
	5:  "OPERATION_REQUIRED",
	6:  "RESOURCE_NOT_FOUND",
	7:  "INVALID_PAGE_TOKEN",
	8:  "EXPIRED_PAGE_TOKEN",
	22: "INVALID_PAGE_SIZE",
	9:  "REQUIRED_FIELD_MISSING",
	11: "IMMUTABLE_FIELD",
	13: "TOO_MANY_MUTATE_OPERATIONS",
	14: "CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT",
	15: "CANNOT_MODIFY_FOREIGN_FIELD",
	18: "INVALID_ENUM_VALUE",
	19: "DEVELOPER_TOKEN_PARAMETER_MISSING",
	20: "LOGIN_CUSTOMER_ID_PARAMETER_MISSING",
	21: "VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN",
}

var RequestErrorEnum_RequestError_value = map[string]int32{
	"UNSPECIFIED":                           0,
	"UNKNOWN":                               1,
	"RESOURCE_NAME_MISSING":                 3,
	"RESOURCE_NAME_MALFORMED":               4,
	"BAD_RESOURCE_ID":                       17,
	"INVALID_CUSTOMER_ID":                   16,
	"OPERATION_REQUIRED":                    5,
	"RESOURCE_NOT_FOUND":                    6,
	"INVALID_PAGE_TOKEN":                    7,
	"EXPIRED_PAGE_TOKEN":                    8,
	"INVALID_PAGE_SIZE":                     22,
	"REQUIRED_FIELD_MISSING":                9,
	"IMMUTABLE_FIELD":                       11,
	"TOO_MANY_MUTATE_OPERATIONS":            13,
	"CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT": 14,
	"CANNOT_MODIFY_FOREIGN_FIELD":           15,
	"INVALID_ENUM_VALUE":                    18,
	"DEVELOPER_TOKEN_PARAMETER_MISSING":     19,
	"LOGIN_CUSTOMER_ID_PARAMETER_MISSING":   20,
	"VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN":  21,
}

func (x RequestErrorEnum_RequestError) String() string {
	return proto.EnumName(RequestErrorEnum_RequestError_name, int32(x))
}

func (RequestErrorEnum_RequestError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fef0616bab723331, []int{0, 0}
}

// Container for enum describing possible request errors.
type RequestErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestErrorEnum) Reset()         { *m = RequestErrorEnum{} }
func (m *RequestErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RequestErrorEnum) ProtoMessage()    {}
func (*RequestErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fef0616bab723331, []int{0}
}

func (m *RequestErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestErrorEnum.Unmarshal(m, b)
}
func (m *RequestErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestErrorEnum.Marshal(b, m, deterministic)
}
func (m *RequestErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestErrorEnum.Merge(m, src)
}
func (m *RequestErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RequestErrorEnum.Size(m)
}
func (m *RequestErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RequestErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.RequestErrorEnum_RequestError", RequestErrorEnum_RequestError_name, RequestErrorEnum_RequestError_value)
	proto.RegisterType((*RequestErrorEnum)(nil), "google.ads.googleads.v1.errors.RequestErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/request_error.proto", fileDescriptor_fef0616bab723331)
}

var fileDescriptor_fef0616bab723331 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x66, 0x3f, 0x6c, 0xe0, 0x01, 0xf3, 0x3c, 0xb6, 0x89, 0x0d, 0x0d, 0x51, 0x98, 0x80, 0x9b,
	0x54, 0x85, 0xbb, 0x70, 0xe5, 0xc6, 0xa7, 0xc1, 0x5a, 0x62, 0x07, 0x27, 0x29, 0xeb, 0x54, 0xc9,
	0x2a, 0xb4, 0xaa, 0x26, 0x6d, 0xc9, 0x48, 0xba, 0x3d, 0x10, 0x97, 0x3c, 0x01, 0xcf, 0x80, 0xc4,
	0x8b, 0x20, 0x1e, 0x02, 0x39, 0x6e, 0x43, 0x40, 0xc0, 0x55, 0x4e, 0xbe, 0xf3, 0x7d, 0xe7, 0x3b,
	0xe7, 0xe8, 0x18, 0xbd, 0x9c, 0xe6, 0xf9, 0xf4, 0x7c, 0xd2, 0x1e, 0x8d, 0xcb, 0xb6, 0x0d, 0x4d,
	0x74, 0xdd, 0x69, 0x4f, 0x8a, 0x22, 0x2f, 0xca, 0x76, 0x31, 0xf9, 0x78, 0x35, 0x29, 0x67, 0xba,
	0xfa, 0x75, 0x2e, 0x8b, 0x7c, 0x96, 0x93, 0x43, 0x4b, 0x74, 0x46, 0xe3, 0xd2, 0xa9, 0x35, 0xce,
	0x75, 0xc7, 0xb1, 0x9a, 0xfd, 0x87, 0x8b, 0x9a, 0x97, 0x67, 0xed, 0x51, 0x96, 0xe5, 0xb3, 0xd1,
	0xec, 0x2c, 0xcf, 0x4a, 0xab, 0x6e, 0x7d, 0x5b, 0x45, 0x58, 0xd9, 0xaa, 0x60, 0xf8, 0x90, 0x5d,
	0x5d, 0xb4, 0xbe, 0xac, 0xa2, 0x3b, 0x4d, 0x90, 0x6c, 0xa2, 0x8d, 0x54, 0xc4, 0x11, 0x78, 0xbc,
	0xc7, 0x81, 0xe1, 0x1b, 0x64, 0x03, 0xad, 0xa7, 0xe2, 0x58, 0xc8, 0x77, 0x02, 0x2f, 0x91, 0x07,
	0x68, 0x47, 0x41, 0x2c, 0x53, 0xe5, 0x81, 0x16, 0x34, 0x04, 0x1d, 0xf2, 0x38, 0xe6, 0xc2, 0xc7,
	0x2b, 0xe4, 0x00, 0xed, 0xfd, 0x91, 0xa2, 0x41, 0x4f, 0xaa, 0x10, 0x18, 0x5e, 0x25, 0xdb, 0x68,
	0xb3, 0x4b, 0x99, 0xae, 0x09, 0x9c, 0xe1, 0x2d, 0xb2, 0x87, 0xb6, 0xb9, 0xe8, 0xd3, 0x80, 0x33,
	0xed, 0xa5, 0x71, 0x22, 0x43, 0x50, 0x26, 0x81, 0xc9, 0x2e, 0x22, 0x32, 0x02, 0x45, 0x13, 0x2e,
	0x85, 0x56, 0xf0, 0x36, 0xe5, 0x0a, 0x18, 0xbe, 0x69, 0xf0, 0x5f, 0x16, 0x32, 0xd1, 0x3d, 0x99,
	0x0a, 0x86, 0xd7, 0x0c, 0xbe, 0x28, 0x14, 0x51, 0x1f, 0x74, 0x22, 0x8f, 0x41, 0xe0, 0x75, 0x83,
	0xc3, 0x49, 0x64, 0xc4, 0x4d, 0xfc, 0x16, 0xd9, 0x41, 0x5b, 0xbf, 0xf1, 0x63, 0x7e, 0x0a, 0x78,
	0x97, 0xec, 0xa3, 0xdd, 0x85, 0x99, 0xee, 0x71, 0x08, 0x58, 0x3d, 0xdd, 0x6d, 0x33, 0x00, 0x0f,
	0xc3, 0x34, 0xa1, 0xdd, 0x00, 0x6c, 0x12, 0x6f, 0x90, 0x43, 0xb4, 0x9f, 0x48, 0xa9, 0x43, 0x2a,
	0x06, 0xda, 0xe4, 0x12, 0xd0, 0x75, 0xdf, 0x31, 0xbe, 0x4b, 0x5e, 0xa0, 0x23, 0x8f, 0x0a, 0xd3,
	0x69, 0x17, 0x34, 0x9c, 0x80, 0x97, 0x26, 0xc0, 0x74, 0x77, 0x60, 0x14, 0xd4, 0x07, 0xa5, 0xa9,
	0xe7, 0xc9, 0x54, 0x24, 0xf8, 0x1e, 0x79, 0x84, 0x0e, 0xe6, 0xd4, 0x50, 0x32, 0xde, 0x1b, 0xe8,
	0x9e, 0x54, 0xc0, 0x7d, 0x31, 0xf7, 0xda, 0x6c, 0xce, 0x08, 0x22, 0x0d, 0x75, 0x9f, 0x06, 0x29,
	0x60, 0x42, 0x8e, 0xd0, 0x63, 0x06, 0x7d, 0x08, 0x8c, 0xb1, 0x1d, 0x50, 0x47, 0x54, 0xd1, 0x10,
	0x12, 0x50, 0x75, 0xff, 0xdb, 0xe4, 0x19, 0x7a, 0x12, 0x48, 0x9f, 0x8b, 0xe6, 0xa6, 0xff, 0x42,
	0xbc, 0x4f, 0x9e, 0xa3, 0xa7, 0x95, 0x4b, 0x35, 0x8c, 0x08, 0x06, 0xd5, 0xfe, 0x21, 0x4e, 0xf4,
	0x1b, 0x1a, 0x37, 0xb7, 0xb8, 0xd3, 0xfd, 0xb1, 0x84, 0x5a, 0x1f, 0xf2, 0x0b, 0xe7, 0xff, 0x47,
	0xd9, 0xdd, 0x6a, 0x9e, 0x57, 0x64, 0x2e, 0x31, 0x5a, 0x3a, 0x65, 0x73, 0xd1, 0x34, 0x3f, 0x1f,
	0x65, 0x53, 0x27, 0x2f, 0xa6, 0xed, 0xe9, 0x24, 0xab, 0xee, 0x74, 0xf1, 0x1a, 0x2e, 0xcf, 0xca,
	0x7f, 0x3d, 0x8e, 0xd7, 0xf6, 0xf3, 0x69, 0x79, 0xc5, 0xa7, 0xf4, 0xf3, 0xf2, 0xa1, 0x6f, 0x8b,
	0xd1, 0x71, 0xe9, 0xd8, 0xd0, 0x44, 0xfd, 0x8e, 0x53, 0x59, 0x96, 0x5f, 0x17, 0x84, 0x21, 0x1d,
	0x97, 0xc3, 0x9a, 0x30, 0xec, 0x77, 0x86, 0x96, 0xf0, 0x7d, 0xb9, 0x65, 0x51, 0xd7, 0xa5, 0xe3,
	0xd2, 0x75, 0x6b, 0x8a, 0xeb, 0xf6, 0x3b, 0xae, 0x6b, 0x49, 0xef, 0xd7, 0xaa, 0xee, 0x5e, 0xfd,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x0c, 0x84, 0xce, 0xb9, 0x03, 0x00, 0x00,
}
