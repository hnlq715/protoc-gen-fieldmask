// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata.proto

package testdata // import "github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	A                    *Test_TestNested `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	CustomName           *Test_TestNested `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C                    Test_TestNested  `protobuf:"bytes,3,opt,name=c" json:"c"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdata_a2f9de05b7077ca1, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetA() *Test_TestNested {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Test) GetCustomName() *Test_TestNested {
	if m != nil {
		return m.CustomName
	}
	return nil
}

func (m *Test) GetC() Test_TestNested {
	if m != nil {
		return m.C
	}
	return Test_TestNested{}
}

type Test_TestNested struct {
	A                    *Test_TestNested_TestNestedNested `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B                    []byte                            `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    *time.Duration                    `protobuf:"bytes,3,opt,name=c,stdduration" json:"c,omitempty"`
	D                    *time.Time                        `protobuf:"bytes,4,opt,name=d,stdtime" json:"d,omitempty"`
	E                    *CustomType                       `protobuf:"bytes,5,opt,name=e,proto3,customtype=CustomType" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Test_TestNested) Reset()         { *m = Test_TestNested{} }
func (m *Test_TestNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested) ProtoMessage()    {}
func (*Test_TestNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdata_a2f9de05b7077ca1, []int{0, 0}
}
func (m *Test_TestNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested.Unmarshal(m, b)
}
func (m *Test_TestNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested.Marshal(b, m, deterministic)
}
func (dst *Test_TestNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested.Merge(dst, src)
}
func (m *Test_TestNested) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested.Size(m)
}
func (m *Test_TestNested) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested proto.InternalMessageInfo

func (m *Test_TestNested) GetA() *Test_TestNested_TestNestedNested {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Test_TestNested) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Test_TestNested) GetC() *time.Duration {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Test_TestNested) GetD() *time.Time {
	if m != nil {
		return m.D
	}
	return nil
}

type Test_TestNested_TestNestedNested struct {
	A                    int32            `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int64            `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    [][]byte         `protobuf:"bytes,3,rep,name=c" json:"c,omitempty"`
	D                    map[int32]uint32 `protobuf:"bytes,4,rep,name=d" json:"d,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Test_TestNested_TestNestedNested) Reset()         { *m = Test_TestNested_TestNestedNested{} }
func (m *Test_TestNested_TestNestedNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested_TestNestedNested) ProtoMessage()    {}
func (*Test_TestNested_TestNestedNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdata_a2f9de05b7077ca1, []int{0, 0, 0}
}
func (m *Test_TestNested_TestNestedNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Unmarshal(m, b)
}
func (m *Test_TestNested_TestNestedNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Marshal(b, m, deterministic)
}
func (dst *Test_TestNested_TestNestedNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested_TestNestedNested.Merge(dst, src)
}
func (m *Test_TestNested_TestNestedNested) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Size(m)
}
func (m *Test_TestNested_TestNestedNested) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested_TestNestedNested.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested_TestNestedNested proto.InternalMessageInfo

func (m *Test_TestNested_TestNestedNested) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetC() [][]byte {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetD() map[int32]uint32 {
	if m != nil {
		return m.D
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "testdata.Test")
	proto.RegisterType((*Test_TestNested)(nil), "testdata.Test.TestNested")
	proto.RegisterType((*Test_TestNested_TestNestedNested)(nil), "testdata.Test.TestNested.TestNestedNested")
	proto.RegisterMapType((map[int32]uint32)(nil), "testdata.Test.TestNested.TestNestedNested.DEntry")
}

func init() { proto.RegisterFile("testdata.proto", fileDescriptor_testdata_a2f9de05b7077ca1) }

var fileDescriptor_testdata_a2f9de05b7077ca1 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0x37, 0xe9, 0x45, 0xe7, 0xd6, 0x52, 0x06, 0x17, 0x31, 0x88, 0xb9, 0xb8, 0xb1,
	0x08, 0x49, 0xf1, 0x2a, 0x58, 0x44, 0x14, 0x62, 0x5d, 0xb8, 0xe9, 0x22, 0x64, 0xe5, 0x6e, 0x92,
	0x39, 0x4d, 0x43, 0x9b, 0x4c, 0xc9, 0x4c, 0x84, 0x3e, 0x81, 0x5b, 0x97, 0x6e, 0x7d, 0x06, 0x5f,
	0xc2, 0x67, 0x70, 0x51, 0xc1, 0x27, 0x91, 0xcc, 0x34, 0x4d, 0x6d, 0x91, 0xde, 0x4d, 0x38, 0x7f,
	0xfe, 0xff, 0x30, 0xdf, 0xfc, 0x0c, 0x19, 0x28, 0x94, 0x8a, 0x33, 0xc5, 0x82, 0x75, 0x25, 0x94,
	0xa0, 0x77, 0x5b, 0xed, 0xfa, 0x59, 0xae, 0x16, 0x75, 0x12, 0xa4, 0xa2, 0x18, 0x67, 0x22, 0x13,
	0x63, 0x1d, 0x48, 0xea, 0xb9, 0x56, 0x5a, 0xe8, 0xc9, 0x2c, 0xba, 0x8f, 0x33, 0x21, 0xb2, 0x15,
	0x76, 0x29, 0x5e, 0x57, 0x4c, 0xe5, 0xa2, 0xdc, 0xf9, 0xde, 0xb1, 0xaf, 0xf2, 0x02, 0xa5, 0x62,
	0xc5, 0xda, 0x04, 0x9e, 0x7c, 0xb7, 0x89, 0x1d, 0xa3, 0x54, 0xf4, 0x29, 0x01, 0xe6, 0xc0, 0x35,
	0x8c, 0xae, 0x6e, 0x1e, 0x06, 0x7b, 0xbc, 0xc6, 0xd2, 0x9f, 0x19, 0x4a, 0x85, 0x3c, 0x02, 0x46,
	0x5f, 0x11, 0x48, 0x9c, 0x8b, 0x33, 0xc1, 0x70, 0xf0, 0x67, 0xeb, 0x91, 0xf7, 0xb5, 0x54, 0xa2,
	0x98, 0xb1, 0x02, 0x23, 0x48, 0xa8, 0x4f, 0x20, 0x75, 0xac, 0x73, 0x8b, 0xf6, 0xcf, 0xad, 0x77,
	0x27, 0x82, 0xd4, 0xfd, 0x62, 0x11, 0xd2, 0xfd, 0xa7, 0x93, 0x8e, 0xef, 0xd9, 0x7f, 0xb7, 0x0f,
	0xc6, 0x0e, 0xb8, 0xdf, 0x02, 0xf7, 0x4f, 0x28, 0x4c, 0x3b, 0x41, 0xdb, 0x4e, 0x30, 0xdd, 0xb5,
	0x17, 0xda, 0xdf, 0x7e, 0x7b, 0x10, 0x41, 0x4a, 0x03, 0x02, 0xdc, 0xb1, 0x75, 0xdc, 0x3d, 0x89,
	0xc7, 0x6d, 0x99, 0xa1, 0xfd, 0x55, 0xe7, 0x39, 0x7d, 0x44, 0x00, 0x9d, 0xde, 0x35, 0x8c, 0xee,
	0x85, 0x83, 0x5f, 0xfb, 0x0a, 0xe2, 0xcd, 0x1a, 0x23, 0x40, 0xf7, 0x07, 0x90, 0xe1, 0x31, 0x62,
	0xc3, 0x67, 0x6e, 0xd6, 0xfb, 0x87, 0x76, 0xd8, 0xd0, 0xf6, 0x0d, 0xad, 0xd5, 0xb0, 0xa7, 0xf4,
	0x9d, 0x81, 0xb1, 0x46, 0x57, 0x37, 0xcf, 0x6f, 0xdf, 0x41, 0x30, 0xfd, 0x50, 0xaa, 0x6a, 0x13,
	0x01, 0x77, 0x5f, 0x92, 0x4b, 0x23, 0xe8, 0x90, 0x58, 0x4b, 0xdc, 0xec, 0x8e, 0x6d, 0x46, 0xfa,
	0x80, 0xf4, 0x3e, 0xb3, 0x55, 0x8d, 0xfa, 0xf0, 0xfb, 0x91, 0x11, 0xaf, 0x2f, 0x26, 0x10, 0xbe,
	0xfd, 0xf4, 0xe6, 0xe0, 0x55, 0xc6, 0x0b, 0x8c, 0x17, 0x79, 0x99, 0xc9, 0x8f, 0x25, 0xaf, 0xa5,
	0xaa, 0x72, 0x94, 0xe6, 0x79, 0xa5, 0x7e, 0x86, 0xa5, 0x3f, 0xcf, 0x71, 0xc5, 0x0b, 0x26, 0x97,
	0xe3, 0x16, 0x30, 0xb9, 0xd4, 0xf6, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x0b, 0xfa,
	0x1f, 0xf6, 0x02, 0x00, 0x00,
}
