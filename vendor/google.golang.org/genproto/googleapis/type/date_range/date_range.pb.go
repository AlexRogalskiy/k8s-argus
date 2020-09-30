// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/actions/type/date_range.proto

package date_range

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	date "google.golang.org/genproto/googleapis/type/date"
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

// Represents a range based on whole or partial calendar dates, e.g. the
// duration of a hotel reservation or the Common Era. This can represent:
//
// * A range between full dates, e.g. the duration of a hotel reservation
// * A range between years, e.g. a historical era
// * A range between year/month dates, e.g. the duration of a job on a resume
// * A range beginning in a year, e.g. the Common Era
// * A range ending on a specific date, e.g. the period of time before an event
//
// While [google.type.Date][google.type.Date] allows zero years, DateRange does not. Year must
// always be non-zero.
//
// End cannot be chronologically before start. For example, if start has year
// 2000, end cannot have year 1999.
//
// When both set, start and end must have exactly the same precision. That is,
// they must have the same fields populated with non-zero values. For example,
// if start specifies only year and month, then end must also specify only year
// and month (but not day).
//
// The date range is inclusive. That is, the dates specified by start and end
// are part of the date range. For example, the date January 1, 2000 falls
// within any date with start or end equal to January 1, 2000. When determining
// whether a date is inside a date range, the date should only be compared to
// start and end when those values are set.
//
// When a date and date range are specified to different degrees of precision,
// the rules for evaluating whether that date is inside the date range are as
// follows:
//
//  * When comparing the date to the start of the date range, unspecified months
//    should be replaced with 1, and unspecified days should be replaced with 1.
//    For example, the year 2000 is within the date range with start equal to
//    January 1, 2000 and no end. And the date January 1, 2000 is within the
//    date range with start equal to the year 2000 and no end.
//
//  * When comparing the date to the end of the date range, unspecified months
//    should be replaced with 12, and unspecified days should be replaced with
//    the last valid day for the month/year. For example, the year 2000 is
//    within the date range with start equal to January 1, 1999 and end equal to
//    December 31, 2000. And the date December 31, 2001 is within the date range
//    with start equal to the year 2000 and end equal to the year 2001.
//
// The semantics of start and end are the same as those of [google.type.Date][google.type.Date],
// except that year must always be non-zero in DateRange.
type DateRange struct {
	// Date at which the date range begins. If unset, the date range has no
	// beginning bound.
	Start *date.Date `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// Date at which the date range ends. If unset, the date range has no ending
	// bound.
	End                  *date.Date `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1397d007038b766, []int{0}
}

func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (m *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(m, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStart() *date.Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *DateRange) GetEnd() *date.Date {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.actions.type.DateRange")
}

func init() {
	proto.RegisterFile("google/actions/type/date_range.proto", fileDescriptor_d1397d007038b766)
}

var fileDescriptor_d1397d007038b766 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0xd5,
	0x4f, 0x49, 0x2c, 0x49, 0x8d, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x86, 0xa8, 0xd2, 0x83, 0xaa, 0xd2, 0x03, 0xa9, 0x92, 0x12, 0x83, 0x6a, 0x85, 0x6b,
	0x81, 0x28, 0x56, 0x8a, 0xe4, 0xe2, 0x74, 0x49, 0x2c, 0x49, 0x0d, 0x02, 0xe9, 0x17, 0x52, 0xe7,
	0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x2a, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd4, 0x83,
	0x9a, 0x04, 0xd2, 0xa4, 0x07, 0x56, 0x06, 0x91, 0x17, 0x52, 0xe6, 0x62, 0x4e, 0xcd, 0x4b, 0x91,
	0x60, 0xc2, 0xa5, 0x0c, 0x24, 0xeb, 0x54, 0xca, 0x25, 0x9e, 0x9c, 0x9f, 0xab, 0x87, 0xc5, 0x35,
	0x4e, 0x7c, 0x70, 0x3b, 0x03, 0x40, 0xae, 0x08, 0x60, 0x8c, 0x72, 0x80, 0x2a, 0x4b, 0xcf, 0xcf,
	0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0xbb, 0x51, 0x1f, 0x22,
	0x95, 0x58, 0x90, 0x89, 0xe1, 0x63, 0x6b, 0x04, 0x73, 0x11, 0x13, 0xab, 0xa3, 0xbf, 0x7b, 0x48,
	0x40, 0x12, 0x1b, 0x58, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x22, 0x31, 0x5b, 0x2d,
	0x01, 0x00, 0x00,
}
