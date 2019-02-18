// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata.proto

package testdata

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
	time "time"
)

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Test struct {
	A          *Test_TestNested `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	C          Test_TestNested  `protobuf:"bytes,2,opt,name=c,proto3" json:"c"`
	CustomName *Test_TestNested `protobuf:"bytes,3,opt,name=b,proto3" json:"b,omitempty"`
	// Types that are valid to be assigned to TestOneof:
	//	*Test_E
	//	*Test_D
	//	*Test_F
	TestOneof            isTest_TestOneof      `protobuf_oneof:"testOneof"`
	G                    *Empty                `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H                    *wrappers.StringValue `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

type isTest_TestOneof interface {
	isTest_TestOneof()
}

type Test_E struct {
	E uint32 `protobuf:"varint,4,opt,name=e,proto3,oneof"`
}
type Test_D struct {
	D int32 `protobuf:"varint,5,opt,name=d,proto3,oneof"`
}
type Test_F struct {
	F []byte `protobuf:"bytes,6,opt,name=f,proto3,oneof"`
}

func (*Test_E) isTest_TestOneof() {}
func (*Test_D) isTest_TestOneof() {}
func (*Test_F) isTest_TestOneof() {}

func (m *Test) GetTestOneof() isTest_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *Test) GetA() *Test_TestNested {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Test) GetC() Test_TestNested {
	if m != nil {
		return m.C
	}
	return Test_TestNested{}
}

func (m *Test) GetCustomName() *Test_TestNested {
	if m != nil {
		return m.CustomName
	}
	return nil
}

func (m *Test) GetE() uint32 {
	if x, ok := m.GetTestOneof().(*Test_E); ok {
		return x.E
	}
	return 0
}

func (m *Test) GetD() int32 {
	if x, ok := m.GetTestOneof().(*Test_D); ok {
		return x.D
	}
	return 0
}

func (m *Test) GetF() []byte {
	if x, ok := m.GetTestOneof().(*Test_F); ok {
		return x.F
	}
	return nil
}

func (m *Test) GetG() *Empty {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *Test) GetH() *wrappers.StringValue {
	if m != nil {
		return m.H
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_OneofMarshaler, _Test_OneofUnmarshaler, _Test_OneofSizer, []interface{}{
		(*Test_E)(nil),
		(*Test_D)(nil),
		(*Test_F)(nil),
	}
}

func _Test_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test)
	// testOneof
	switch x := m.TestOneof.(type) {
	case *Test_E:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.E))
	case *Test_D:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.D))
	case *Test_F:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.F)
	case nil:
	default:
		return fmt.Errorf("Test.TestOneof has unexpected type %T", x)
	}
	return nil
}

func _Test_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test)
	switch tag {
	case 4: // testOneof.e
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneof = &Test_E{uint32(x)}
		return true, err
	case 5: // testOneof.d
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneof = &Test_D{int32(x)}
		return true, err
	case 6: // testOneof.f
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TestOneof = &Test_F{x}
		return true, err
	default:
		return false, nil
	}
}

func _Test_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test)
	// testOneof
	switch x := m.TestOneof.(type) {
	case *Test_E:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.E))
	case *Test_D:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.D))
	case *Test_F:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.F)))
		n += len(x.F)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Test_TestNested struct {
	A                    *Test_TestNested_TestNestedNested `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    []byte                            `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    *time.Duration                    `protobuf:"bytes,3,opt,name=c,proto3,stdduration" json:"c,omitempty"`
	D                    *time.Time                        `protobuf:"bytes,4,opt,name=d,proto3,stdtime" json:"d,omitempty"`
	E                    *CustomType                       `protobuf:"bytes,5,opt,name=e,proto3,customtype=CustomType" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Test_TestNested) Reset()         { *m = Test_TestNested{} }
func (m *Test_TestNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested) ProtoMessage()    {}
func (*Test_TestNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0}
}
func (m *Test_TestNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested.Unmarshal(m, b)
}
func (m *Test_TestNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested.Marshal(b, m, deterministic)
}
func (m *Test_TestNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested.Merge(m, src)
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
	A int32            `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64            `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	C [][]byte         `protobuf:"bytes,3,rep,name=c,proto3" json:"c,omitempty"`
	D map[int32]uint32 `protobuf:"bytes,4,rep,name=d,proto3" json:"d,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to TestNestedNestedOneOf:
	//	*Test_TestNested_TestNestedNested_E
	//	*Test_TestNested_TestNestedNested_F
	//	*Test_TestNested_TestNestedNested_G
	TestNestedNestedOneOf isTest_TestNested_TestNestedNested_TestNestedNestedOneOf `protobuf_oneof:"testNestedNestedOneOf"`
	XXX_NoUnkeyedLiteral  struct{}                                                 `json:"-"`
	XXX_unrecognized      []byte                                                   `json:"-"`
	XXX_sizecache         int32                                                    `json:"-"`
}

func (m *Test_TestNested_TestNestedNested) Reset()         { *m = Test_TestNested_TestNestedNested{} }
func (m *Test_TestNested_TestNestedNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested_TestNestedNested) ProtoMessage()    {}
func (*Test_TestNested_TestNestedNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0, 0}
}
func (m *Test_TestNested_TestNestedNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Unmarshal(m, b)
}
func (m *Test_TestNested_TestNestedNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Marshal(b, m, deterministic)
}
func (m *Test_TestNested_TestNestedNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested_TestNestedNested.Merge(m, src)
}
func (m *Test_TestNested_TestNestedNested) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Size(m)
}
func (m *Test_TestNested_TestNestedNested) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested_TestNestedNested.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested_TestNestedNested proto.InternalMessageInfo

type isTest_TestNested_TestNestedNested_TestNestedNestedOneOf interface {
	isTest_TestNested_TestNestedNested_TestNestedNestedOneOf()
}

type Test_TestNested_TestNestedNested_E struct {
	E *Empty `protobuf:"bytes,5,opt,name=e,proto3,oneof"`
}
type Test_TestNested_TestNestedNested_F struct {
	F uint32 `protobuf:"varint,6,opt,name=f,proto3,oneof"`
}
type Test_TestNested_TestNestedNested_G struct {
	G *wrappers.UInt64Value `protobuf:"bytes,7,opt,name=g,proto3,oneof"`
}

func (*Test_TestNested_TestNestedNested_E) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}
func (*Test_TestNested_TestNestedNested_F) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}
func (*Test_TestNested_TestNestedNested_G) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}

func (m *Test_TestNested_TestNestedNested) GetTestNestedNestedOneOf() isTest_TestNested_TestNestedNested_TestNestedNestedOneOf {
	if m != nil {
		return m.TestNestedNestedOneOf
	}
	return nil
}

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

func (m *Test_TestNested_TestNestedNested) GetE() *Empty {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_E); ok {
		return x.E
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetF() uint32 {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_F); ok {
		return x.F
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetG() *wrappers.UInt64Value {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_G); ok {
		return x.G
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test_TestNested_TestNestedNested) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_TestNested_TestNestedNested_OneofMarshaler, _Test_TestNested_TestNestedNested_OneofUnmarshaler, _Test_TestNested_TestNestedNested_OneofSizer, []interface{}{
		(*Test_TestNested_TestNestedNested_E)(nil),
		(*Test_TestNested_TestNestedNested_F)(nil),
		(*Test_TestNested_TestNestedNested_G)(nil),
	}
}

func _Test_TestNested_TestNestedNested_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test_TestNested_TestNestedNested)
	// testNestedNestedOneOf
	switch x := m.TestNestedNestedOneOf.(type) {
	case *Test_TestNested_TestNestedNested_E:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.E); err != nil {
			return err
		}
	case *Test_TestNested_TestNestedNested_F:
		_ = b.EncodeVarint(6<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.F))
	case *Test_TestNested_TestNestedNested_G:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.G); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Test_TestNested_TestNestedNested.TestNestedNestedOneOf has unexpected type %T", x)
	}
	return nil
}

func _Test_TestNested_TestNestedNested_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test_TestNested_TestNestedNested)
	switch tag {
	case 5: // testNestedNestedOneOf.e
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_E{msg}
		return true, err
	case 6: // testNestedNestedOneOf.f
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_F{uint32(x)}
		return true, err
	case 7: // testNestedNestedOneOf.g
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.UInt64Value)
		err := b.DecodeMessage(msg)
		m.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_G{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Test_TestNested_TestNestedNested_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test_TestNested_TestNestedNested)
	// testNestedNestedOneOf
	switch x := m.TestNestedNestedOneOf.(type) {
	case *Test_TestNested_TestNestedNested_E:
		s := proto.Size(x.E)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_TestNested_TestNestedNested_F:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.F))
	case *Test_TestNested_TestNestedNested_G:
		s := proto.Size(x.G)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Empty)(nil), "testdata.Empty")
	proto.RegisterType((*Test)(nil), "testdata.Test")
	proto.RegisterType((*Test_TestNested)(nil), "testdata.Test.TestNested")
	proto.RegisterType((*Test_TestNested_TestNestedNested)(nil), "testdata.Test.TestNested.TestNestedNested")
	proto.RegisterMapType((map[int32]uint32)(nil), "testdata.Test.TestNested.TestNestedNested.DEntry")
}

func init() { proto.RegisterFile("testdata.proto", fileDescriptor_40c4782d007dfce9) }

var fileDescriptor_40c4782d007dfce9 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0xee, 0x24, 0x9b, 0xb4, 0x9d, 0x7c, 0x18, 0x06, 0xc5, 0x75, 0xa9, 0x26, 0xf4, 0x62, 0x28,
	0x66, 0x83, 0xb5, 0x68, 0x11, 0x51, 0x58, 0x5b, 0xb0, 0x97, 0x06, 0xd6, 0xe8, 0xc1, 0xdb, 0x26,
	0xfb, 0xee, 0x64, 0x69, 0xf6, 0x83, 0x9d, 0x59, 0x25, 0xff, 0xc2, 0xa3, 0x57, 0xcf, 0xfe, 0x11,
	0x7f, 0x83, 0x87, 0x0a, 0xfe, 0x11, 0xe5, 0x9d, 0xc9, 0x66, 0x6b, 0x42, 0x69, 0x2f, 0x61, 0x1e,
	0x9e, 0x27, 0x3b, 0xcf, 0x07, 0x43, 0xdb, 0x12, 0x84, 0xf4, 0x3d, 0xe9, 0xd9, 0x69, 0x96, 0xc8,
	0x84, 0xed, 0x14, 0xd8, 0x1a, 0xf0, 0x50, 0xce, 0xf2, 0x89, 0x3d, 0x4d, 0xa2, 0x21, 0x4f, 0x78,
	0x32, 0x54, 0x82, 0x49, 0x1e, 0x28, 0xa4, 0x80, 0x3a, 0xe9, 0x3f, 0x5a, 0x8f, 0x78, 0x92, 0xf0,
	0x39, 0x94, 0x2a, 0x3f, 0xcf, 0x3c, 0x19, 0x26, 0xf1, 0x92, 0xef, 0xae, 0xf3, 0x32, 0x8c, 0x40,
	0x48, 0x2f, 0x4a, 0xaf, 0xfb, 0xc0, 0x97, 0xcc, 0x4b, 0x53, 0xc8, 0x84, 0xe6, 0xf7, 0xb7, 0x69,
	0xed, 0x34, 0x4a, 0xe5, 0x62, 0xff, 0x7b, 0x9d, 0x1a, 0x63, 0x10, 0x92, 0x3d, 0xa6, 0xc4, 0x33,
	0x49, 0x8f, 0xf4, 0x1b, 0x87, 0x0f, 0xec, 0x55, 0x0e, 0xa4, 0xd4, 0xcf, 0x39, 0x08, 0x09, 0xbe,
	0x4b, 0x3c, 0x36, 0xa0, 0x64, 0x6a, 0x56, 0x6e, 0x10, 0x3a, 0xc6, 0xcf, 0xcb, 0xee, 0x96, 0x4b,
	0xa6, 0xec, 0x05, 0x25, 0x13, 0xb3, 0x7a, 0x93, 0xbc, 0xfd, 0xe7, 0xb2, 0x4b, 0xdf, 0xe6, 0x42,
	0x26, 0xd1, 0xb9, 0x17, 0x81, 0x4b, 0x26, 0xac, 0x4d, 0x09, 0x98, 0x46, 0x8f, 0xf4, 0x5b, 0xef,
	0xb6, 0x5c, 0x02, 0x88, 0x7d, 0xb3, 0xd6, 0x23, 0xfd, 0x1a, 0x62, 0x1f, 0x71, 0x60, 0xd6, 0x7b,
	0xa4, 0xdf, 0x44, 0x1c, 0xb0, 0x87, 0x94, 0x70, 0x73, 0x5b, 0x5d, 0x74, 0xa7, 0xbc, 0x48, 0xa5,
	0x74, 0x09, 0x67, 0x07, 0x94, 0xcc, 0xcc, 0x1d, 0x45, 0xef, 0xd9, 0xba, 0x1d, 0xbb, 0x68, 0xc7,
	0x7e, 0x2f, 0xb3, 0x30, 0xe6, 0x1f, 0xbd, 0x79, 0x0e, 0x2e, 0x99, 0x59, 0x7f, 0xab, 0x94, 0x96,
	0xe6, 0xd8, 0x71, 0x59, 0xcd, 0xc1, 0xb5, 0x11, 0xae, 0x1c, 0xcb, 0xae, 0x9a, 0x18, 0x1e, 0xbb,
	0x6a, 0x62, 0x22, 0xd5, 0x5c, 0x51, 0xc5, 0xba, 0x85, 0x93, 0xe5, 0xc2, 0x8e, 0xf1, 0xed, 0x77,
	0x97, 0x60, 0x73, 0x36, 0x06, 0x36, 0x94, 0xdc, 0xda, 0x90, 0x8f, 0x8b, 0xc1, 0x1d, 0xe3, 0xab,
	0xd2, 0xfb, 0x6c, 0x0f, 0x0b, 0xc3, 0x82, 0x76, 0x9d, 0xf6, 0xaf, 0x55, 0x9d, 0xe3, 0x45, 0x0a,
	0x2e, 0x01, 0xeb, 0x47, 0x85, 0x76, 0xd6, 0x2d, 0xa2, 0x3f, 0x9d, 0xac, 0xf6, 0x9f, 0xdb, 0x0e,
	0xba, 0x6d, 0x6a, 0xb7, 0x55, 0xf4, 0x3e, 0x65, 0x6f, 0xb4, 0x99, 0x6a, 0xbf, 0x71, 0xf8, 0xf4,
	0xf6, 0x1d, 0xd8, 0x27, 0xa7, 0xb1, 0xcc, 0x16, 0xe8, 0xae, 0x5b, 0xb8, 0xdb, 0x9c, 0x67, 0xb5,
	0xaf, 0xde, 0xb3, 0xa5, 0xf7, 0x7c, 0x52, 0xee, 0xb9, 0x39, 0xd8, 0x87, 0xb3, 0x58, 0x3e, 0x3f,
	0x52, 0x83, 0xa1, 0x9a, 0x5b, 0x47, 0xb4, 0xae, 0xef, 0x62, 0x1d, 0x5a, 0xbd, 0x80, 0xc5, 0x32,
	0x15, 0x1e, 0xd9, 0x5d, 0x5a, 0xfb, 0x8c, 0x4a, 0x95, 0xad, 0xe5, 0x6a, 0xf0, 0xb2, 0x72, 0x4c,
	0x9c, 0xfb, 0xf4, 0x9e, 0x5c, 0xb3, 0x3c, 0x8a, 0x61, 0x14, 0x38, 0x0d, 0xba, 0x8b, 0xc4, 0x28,
	0x86, 0x24, 0x70, 0x5e, 0x7f, 0x7a, 0x75, 0xe5, 0xf9, 0x8e, 0x67, 0x30, 0x9e, 0x85, 0x31, 0x17,
	0x67, 0xb1, 0x9f, 0x0b, 0x99, 0x85, 0x20, 0xf4, 0x33, 0x9b, 0x0e, 0x38, 0xc4, 0x83, 0x20, 0x84,
	0xb9, 0x1f, 0x79, 0xe2, 0x62, 0x58, 0x84, 0x9c, 0xd4, 0x15, 0xfd, 0xec, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x93, 0x23, 0xb8, 0xf0, 0x1f, 0x04, 0x00, 0x00,
}
