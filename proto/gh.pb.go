// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gh.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type EmploymentType int32

const (
	EmploymentType_WORK       EmploymentType = 0
	EmploymentType_OPENSOURCE EmploymentType = 1
	EmploymentType_HOBBY      EmploymentType = 2
)

var EmploymentType_name = map[int32]string{
	0: "WORK",
	1: "OPENSOURCE",
	2: "HOBBY",
}

var EmploymentType_value = map[string]int32{
	"WORK":       0,
	"OPENSOURCE": 1,
	"HOBBY":      2,
}

func (x EmploymentType) String() string {
	return proto.EnumName(EmploymentType_name, int32(x))
}

func (EmploymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22afb19a307f14c9, []int{0}
}

type OwnProfileRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnProfileRequest) Reset()         { *m = OwnProfileRequest{} }
func (m *OwnProfileRequest) String() string { return proto.CompactTextString(m) }
func (*OwnProfileRequest) ProtoMessage()    {}
func (*OwnProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22afb19a307f14c9, []int{0}
}

func (m *OwnProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnProfileRequest.Unmarshal(m, b)
}
func (m *OwnProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnProfileRequest.Marshal(b, m, deterministic)
}
func (m *OwnProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnProfileRequest.Merge(m, src)
}
func (m *OwnProfileRequest) XXX_Size() int {
	return xxx_messageInfo_OwnProfileRequest.Size(m)
}
func (m *OwnProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnProfileRequest proto.InternalMessageInfo

func (m *OwnProfileRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OwnProfileResponse struct {
	ContributionsDispersion float32        `protobuf:"fixed32,1,opt,name=contributionsDispersion,proto3" json:"contributionsDispersion,omitempty"`
	Type                    EmploymentType `protobuf:"varint,2,opt,name=type,proto3,enum=pb.EmploymentType" json:"type,omitempty"`
	Languages               []string       `protobuf:"bytes,3,rep,name=languages,proto3" json:"languages,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}       `json:"-"`
	XXX_unrecognized        []byte         `json:"-"`
	XXX_sizecache           int32          `json:"-"`
}

func (m *OwnProfileResponse) Reset()         { *m = OwnProfileResponse{} }
func (m *OwnProfileResponse) String() string { return proto.CompactTextString(m) }
func (*OwnProfileResponse) ProtoMessage()    {}
func (*OwnProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22afb19a307f14c9, []int{1}
}

func (m *OwnProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnProfileResponse.Unmarshal(m, b)
}
func (m *OwnProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnProfileResponse.Marshal(b, m, deterministic)
}
func (m *OwnProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnProfileResponse.Merge(m, src)
}
func (m *OwnProfileResponse) XXX_Size() int {
	return xxx_messageInfo_OwnProfileResponse.Size(m)
}
func (m *OwnProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OwnProfileResponse proto.InternalMessageInfo

func (m *OwnProfileResponse) GetContributionsDispersion() float32 {
	if m != nil {
		return m.ContributionsDispersion
	}
	return 0
}

func (m *OwnProfileResponse) GetType() EmploymentType {
	if m != nil {
		return m.Type
	}
	return EmploymentType_WORK
}

func (m *OwnProfileResponse) GetLanguages() []string {
	if m != nil {
		return m.Languages
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.EmploymentType", EmploymentType_name, EmploymentType_value)
	proto.RegisterType((*OwnProfileRequest)(nil), "pb.OwnProfileRequest")
	proto.RegisterType((*OwnProfileResponse)(nil), "pb.OwnProfileResponse")
}

func init() {
	proto.RegisterFile("gh.proto", fileDescriptor_22afb19a307f14c9)
}

var fileDescriptor_22afb19a307f14c9 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x9b, 0xb4, 0x7d, 0x68, 0xe6, 0x10, 0xda, 0xe1, 0x51, 0x83, 0x78, 0x08, 0x39, 0x48,
	0xf4, 0x10, 0xa1, 0x22, 0x78, 0x6d, 0x34, 0x58, 0x11, 0x4c, 0xd9, 0xfa, 0x82, 0xde, 0x92, 0x32,
	0x8d, 0x8b, 0xe9, 0xee, 0x9a, 0xdd, 0x28, 0xf9, 0x1e, 0x7e, 0x60, 0xb1, 0x39, 0xd4, 0x17, 0x3c,
	0xce, 0xff, 0x3f, 0x3f, 0x98, 0xf9, 0xc1, 0xa0, 0x78, 0x8a, 0x54, 0x25, 0x8d, 0x44, 0x5b, 0xe5,
	0xc1, 0x01, 0x8c, 0xd2, 0x37, 0x31, 0xab, 0xe4, 0x92, 0x97, 0xc4, 0xe8, 0xa5, 0x26, 0x6d, 0xf0,
	0x3f, 0xf4, 0x8d, 0x7c, 0x26, 0xe1, 0x59, 0xbe, 0x15, 0x3a, 0xac, 0x1d, 0x82, 0x77, 0x0b, 0xf0,
	0xeb, 0xae, 0x56, 0x52, 0x68, 0xc2, 0x53, 0xd8, 0x59, 0x48, 0x61, 0x2a, 0x9e, 0xd7, 0x86, 0x4b,
	0xa1, 0xcf, 0xb9, 0x56, 0x54, 0x69, 0x2e, 0x5b, 0xdc, 0x66, 0x7f, 0xd5, 0xb8, 0x0f, 0x3d, 0xd3,
	0x28, 0xf2, 0x6c, 0xdf, 0x0a, 0xdd, 0x31, 0x46, 0x2a, 0x8f, 0x92, 0x95, 0x2a, 0x65, 0xb3, 0x22,
	0x61, 0x6e, 0x1a, 0x45, 0x6c, 0xdd, 0xe3, 0x1e, 0x38, 0x65, 0x26, 0x8a, 0x3a, 0x2b, 0x48, 0x7b,
	0x5d, 0xbf, 0x1b, 0x3a, 0x6c, 0x13, 0x1c, 0x9e, 0x80, 0xfb, 0x9d, 0xc2, 0x01, 0xf4, 0xee, 0x53,
	0x76, 0x35, 0xec, 0xa0, 0x0b, 0x90, 0xce, 0x92, 0xeb, 0x79, 0x7a, 0xcb, 0xce, 0x92, 0xa1, 0x85,
	0x0e, 0xf4, 0xa7, 0x69, 0x1c, 0x3f, 0x0c, 0xed, 0xf1, 0x1d, 0x8c, 0x2e, 0xa6, 0x13, 0x91, 0x95,
	0x8d, 0xe6, 0x7a, 0x4e, 0xd5, 0x2b, 0x5f, 0x10, 0x4e, 0xc0, 0xdd, 0x7c, 0x78, 0x29, 0x96, 0x12,
	0xb7, 0x3e, 0xaf, 0xfa, 0x65, 0x68, 0x77, 0xfb, 0x67, 0xdc, 0xca, 0x08, 0x3a, 0x71, 0xff, 0xb1,
	0x7b, 0xa4, 0xf2, 0xfc, 0xdf, 0x5a, 0xf1, 0xf1, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xd0,
	0xa0, 0xc0, 0x6e, 0x01, 0x00, 0x00,
}
