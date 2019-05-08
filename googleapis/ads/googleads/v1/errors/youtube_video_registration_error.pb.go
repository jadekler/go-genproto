// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/youtube_video_registration_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing YouTube video registration errors.
type YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError int32

const (
	// Enum unspecified.
	YoutubeVideoRegistrationErrorEnum_UNSPECIFIED YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 0
	// The received error code is not known in this version.
	YoutubeVideoRegistrationErrorEnum_UNKNOWN YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 1
	// Video to be registered wasn't found.
	YoutubeVideoRegistrationErrorEnum_VIDEO_NOT_FOUND YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 2
	// Video to be registered is not accessible (e.g. private).
	YoutubeVideoRegistrationErrorEnum_VIDEO_NOT_ACCESSIBLE YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 3
)

var YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "VIDEO_NOT_FOUND",
	3: "VIDEO_NOT_ACCESSIBLE",
}
var YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"VIDEO_NOT_FOUND":      2,
	"VIDEO_NOT_ACCESSIBLE": 3,
}

func (x YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError) String() string {
	return proto.EnumName(YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name, int32(x))
}
func (YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_youtube_video_registration_error_b51a63b007d526e2, []int{0, 0}
}

// Container for enum describing YouTube video registration errors.
type YoutubeVideoRegistrationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YoutubeVideoRegistrationErrorEnum) Reset()         { *m = YoutubeVideoRegistrationErrorEnum{} }
func (m *YoutubeVideoRegistrationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*YoutubeVideoRegistrationErrorEnum) ProtoMessage()    {}
func (*YoutubeVideoRegistrationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_youtube_video_registration_error_b51a63b007d526e2, []int{0}
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Unmarshal(m, b)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *YoutubeVideoRegistrationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Merge(dst, src)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Size(m)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_YoutubeVideoRegistrationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*YoutubeVideoRegistrationErrorEnum)(nil), "google.ads.googleads.v1.errors.YoutubeVideoRegistrationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError", YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name, YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/youtube_video_registration_error.proto", fileDescriptor_youtube_video_registration_error_b51a63b007d526e2)
}

var fileDescriptor_youtube_video_registration_error_b51a63b007d526e2 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x4d, 0x0a, 0x0a, 0xd3, 0x45, 0x43, 0x74, 0x21, 0xa2, 0x05, 0x73, 0x80, 0x09, 0xc1,
	0xdd, 0xb8, 0x4a, 0x9b, 0x69, 0x09, 0x4a, 0x52, 0xac, 0x8d, 0x28, 0x81, 0x90, 0x3a, 0xc3, 0x10,
	0x68, 0xe7, 0x95, 0x99, 0xb4, 0xe0, 0x31, 0xbc, 0x82, 0x4b, 0x8f, 0xe2, 0x51, 0xbc, 0x83, 0x20,
	0xc9, 0xd8, 0xea, 0xc6, 0xae, 0xf2, 0xf3, 0xf8, 0xde, 0xff, 0xbf, 0x3f, 0x83, 0xa8, 0x00, 0x10,
	0x0b, 0xee, 0x97, 0x4c, 0xfb, 0x46, 0x36, 0x6a, 0x13, 0xf8, 0x5c, 0x29, 0x50, 0xda, 0x7f, 0x81,
	0x75, 0xbd, 0x9e, 0xf3, 0x62, 0x53, 0x31, 0x0e, 0x85, 0xe2, 0xa2, 0xd2, 0xb5, 0x2a, 0xeb, 0x0a,
	0x64, 0xd1, 0x12, 0x78, 0xa5, 0xa0, 0x06, 0xb7, 0x6f, 0x76, 0x71, 0xc9, 0x34, 0xde, 0xd9, 0xe0,
	0x4d, 0x80, 0x8d, 0xcd, 0xd9, 0xf9, 0x36, 0x66, 0x55, 0xf9, 0xa5, 0x94, 0x50, 0xb7, 0x16, 0xda,
	0x6c, 0x7b, 0xaf, 0x16, 0xba, 0x7c, 0x34, 0x41, 0x59, 0x93, 0x73, 0xf7, 0x27, 0x86, 0x36, 0x06,
	0x54, 0xae, 0x97, 0xde, 0x02, 0x5d, 0xec, 0x85, 0xdc, 0x1e, 0xea, 0xce, 0x92, 0xe9, 0x84, 0x0e,
	0xe3, 0x51, 0x4c, 0x23, 0xe7, 0xc0, 0xed, 0xa2, 0xa3, 0x59, 0x72, 0x93, 0xa4, 0x0f, 0x89, 0x63,
	0xb9, 0xc7, 0xa8, 0x97, 0xc5, 0x11, 0x4d, 0x8b, 0x24, 0xbd, 0x2f, 0x46, 0xe9, 0x2c, 0x89, 0x1c,
	0xdb, 0x3d, 0x45, 0x27, 0xbf, 0xc3, 0x70, 0x38, 0xa4, 0xd3, 0x69, 0x3c, 0xb8, 0xa5, 0x4e, 0x67,
	0xf0, 0x65, 0x21, 0xef, 0x19, 0x96, 0x78, 0x7f, 0xb1, 0x81, 0xb7, 0xf7, 0xa4, 0x49, 0x53, 0x6f,
	0x62, 0x3d, 0x45, 0x3f, 0x2e, 0x02, 0x16, 0xa5, 0x14, 0x18, 0x94, 0xf0, 0x05, 0x97, 0x6d, 0xf9,
	0xed, 0x5f, 0x5f, 0x55, 0xfa, 0xbf, 0x47, 0xb8, 0x36, 0x9f, 0x37, 0xbb, 0x33, 0x0e, 0xc3, 0x77,
	0xbb, 0x3f, 0x36, 0x66, 0x21, 0xd3, 0xd8, 0xc8, 0x46, 0x65, 0x01, 0x6e, 0x23, 0xf5, 0xc7, 0x16,
	0xc8, 0x43, 0xa6, 0xf3, 0x1d, 0x90, 0x67, 0x41, 0x6e, 0x80, 0x4f, 0xdb, 0x33, 0x53, 0x42, 0x42,
	0xa6, 0x09, 0xd9, 0x21, 0x84, 0x64, 0x01, 0x21, 0x06, 0x9a, 0x1f, 0xb6, 0xd7, 0x5d, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x20, 0xae, 0x04, 0xe3, 0x21, 0x02, 0x00, 0x00,
}
