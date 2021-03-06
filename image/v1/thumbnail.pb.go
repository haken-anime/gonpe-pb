// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gonpe/image/v1/thumbnail.proto

package imagev1

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

// A thumbnail image information.
type Thumbnail struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width                uint64   `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thumbnail) Reset()         { *m = Thumbnail{} }
func (m *Thumbnail) String() string { return proto.CompactTextString(m) }
func (*Thumbnail) ProtoMessage()    {}
func (*Thumbnail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5f02178c4a651, []int{0}
}

func (m *Thumbnail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thumbnail.Unmarshal(m, b)
}
func (m *Thumbnail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thumbnail.Marshal(b, m, deterministic)
}
func (m *Thumbnail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thumbnail.Merge(m, src)
}
func (m *Thumbnail) XXX_Size() int {
	return xxx_messageInfo_Thumbnail.Size(m)
}
func (m *Thumbnail) XXX_DiscardUnknown() {
	xxx_messageInfo_Thumbnail.DiscardUnknown(m)
}

var xxx_messageInfo_Thumbnail proto.InternalMessageInfo

func (m *Thumbnail) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Thumbnail) GetWidth() uint64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Thumbnail) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Thumbnail)(nil), "gonpe.image.v1.Thumbnail")
}

func init() { proto.RegisterFile("gonpe/image/v1/thumbnail.proto", fileDescriptor_c4f5f02178c4a651) }

var fileDescriptor_c4f5f02178c4a651 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x2b,
	0x48, 0xd5, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xc9, 0x28, 0xcd, 0x4d,
	0xca, 0x4b, 0xcc, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0xcb, 0xeb, 0x81,
	0xe5, 0xf5, 0xca, 0x0c, 0x95, 0xbc, 0xb9, 0x38, 0x43, 0x60, 0x4a, 0x84, 0x04, 0xb8, 0x98, 0x4b,
	0x8b, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xf2,
	0xcc, 0x94, 0x92, 0x0c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b,
	0x2d, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x19, 0x2c, 0x0c, 0xe5, 0x39, 0x85, 0x73, 0x09,
	0x25, 0xe7, 0xe7, 0xea, 0xa1, 0x5a, 0xe1, 0xc4, 0x07, 0xb7, 0x20, 0x00, 0xe4, 0x84, 0x00, 0xc6,
	0x28, 0x76, 0xb0, 0x5c, 0x99, 0xe1, 0x22, 0x26, 0x66, 0x77, 0xcf, 0x88, 0x55, 0x4c, 0x7c, 0xee,
	0x60, 0x1d, 0x9e, 0x60, 0x1d, 0x61, 0x86, 0xa7, 0xa0, 0x02, 0x31, 0x60, 0x81, 0x98, 0x30, 0xc3,
	0x24, 0x36, 0xb0, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x6b, 0x5e, 0x4f, 0xde,
	0x00, 0x00, 0x00,
}
