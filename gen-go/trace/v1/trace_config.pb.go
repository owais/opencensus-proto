// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opencensus/proto/trace/v1/trace_config.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// How spans should be sampled:
// - Always off
// - Always on
// - Always follow the parent Span's decision (off if no parent).
type ConstantSampler_ConstantDecision int32

const (
	ConstantSampler_ALWAYS_OFF    ConstantSampler_ConstantDecision = 0
	ConstantSampler_ALWAYS_ON     ConstantSampler_ConstantDecision = 1
	ConstantSampler_ALWAYS_PARENT ConstantSampler_ConstantDecision = 2
)

var ConstantSampler_ConstantDecision_name = map[int32]string{
	0: "ALWAYS_OFF",
	1: "ALWAYS_ON",
	2: "ALWAYS_PARENT",
}

var ConstantSampler_ConstantDecision_value = map[string]int32{
	"ALWAYS_OFF":    0,
	"ALWAYS_ON":     1,
	"ALWAYS_PARENT": 2,
}

func (x ConstantSampler_ConstantDecision) String() string {
	return proto.EnumName(ConstantSampler_ConstantDecision_name, int32(x))
}

func (ConstantSampler_ConstantDecision) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2, 0}
}

// Global configuration of the trace service. All fields must be specified, or
// the default (zero) values will be used for each type.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfAnnotations int64 `protobuf:"varint,5,opt,name=max_number_of_annotations,json=maxNumberOfAnnotations,proto3" json:"max_number_of_annotations,omitempty"`
	// The global default max number of message events per span.
	MaxNumberOfMessageEvents int64 `protobuf:"varint,6,opt,name=max_number_of_message_events,json=maxNumberOfMessageEvents,proto3" json:"max_number_of_message_events,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks int64 `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,1,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof"`
}
type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,2,opt,name=constant_sampler,json=constantSampler,proto3,oneof"`
}
type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof"`
}

func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler()  {}
func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler()     {}
func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAnnotations() int64 {
	if m != nil {
		return m.MaxNumberOfAnnotations
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfMessageEvents() int64 {
	if m != nil {
		return m.MaxNumberOfMessageEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TraceConfig_OneofMarshaler, _TraceConfig_OneofUnmarshaler, _TraceConfig_OneofSizer, []interface{}{
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

func _TraceConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TraceConfig)
	// sampler
	switch x := m.Sampler.(type) {
	case *TraceConfig_ProbabilitySampler:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProbabilitySampler); err != nil {
			return err
		}
	case *TraceConfig_ConstantSampler:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConstantSampler); err != nil {
			return err
		}
	case *TraceConfig_RateLimitingSampler:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RateLimitingSampler); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TraceConfig.Sampler has unexpected type %T", x)
	}
	return nil
}

func _TraceConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TraceConfig)
	switch tag {
	case 1: // sampler.probability_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProbabilitySampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_ProbabilitySampler{msg}
		return true, err
	case 2: // sampler.constant_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConstantSampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_ConstantSampler{msg}
		return true, err
	case 3: // sampler.rate_limiting_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RateLimitingSampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_RateLimitingSampler{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TraceConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TraceConfig)
	// sampler
	switch x := m.Sampler.(type) {
	case *TraceConfig_ProbabilitySampler:
		s := proto.Size(x.ProbabilitySampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TraceConfig_ConstantSampler:
		s := proto.Size(x.ConstantSampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TraceConfig_RateLimitingSampler:
		s := proto.Size(x.RateLimitingSampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability float64 `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{1}
}
func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(m, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return m.Size()
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that always makes a constant decision on span sampling.
type ConstantSampler struct {
	Decision ConstantSampler_ConstantDecision `protobuf:"varint,1,opt,name=decision,proto3,enum=opencensus.proto.trace.v1.ConstantSampler_ConstantDecision" json:"decision,omitempty"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{2}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(m, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return m.Size()
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() ConstantSampler_ConstantDecision {
	if m != nil {
		return m.Decision
	}
	return ConstantSampler_ALWAYS_OFF
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps int64 `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_5359209b41ff50c5, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(m, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterEnum("opencensus.proto.trace.v1.ConstantSampler_ConstantDecision", ConstantSampler_ConstantDecision_name, ConstantSampler_ConstantDecision_value)
	proto.RegisterType((*TraceConfig)(nil), "opencensus.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ProbabilitySampler)(nil), "opencensus.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*ConstantSampler)(nil), "opencensus.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opencensus.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opencensus/proto/trace/v1/trace_config.proto", fileDescriptor_5359209b41ff50c5)
}

var fileDescriptor_5359209b41ff50c5 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x76, 0x6c, 0xec, 0x9b, 0xb6, 0x05, 0x57, 0x43, 0xa9, 0x34, 0x45, 0x53, 0x2f,
	0x4c, 0x88, 0xa6, 0x74, 0x1c, 0x10, 0x42, 0x42, 0x6a, 0xbb, 0x55, 0x1c, 0x4a, 0x5b, 0x65, 0x13,
	0x15, 0x5c, 0x82, 0x93, 0xb9, 0xc1, 0xa2, 0xb1, 0x43, 0xec, 0x56, 0xe3, 0x2d, 0x78, 0x00, 0x9e,
	0x80, 0x13, 0x8f, 0xc1, 0x71, 0x47, 0x8e, 0xa8, 0xbd, 0xf1, 0x14, 0xa8, 0x4e, 0x69, 0xd2, 0x76,
	0x9b, 0xb8, 0xf9, 0xfb, 0xff, 0xbf, 0xdf, 0xdf, 0x49, 0xfc, 0x39, 0xf0, 0x44, 0x44, 0x94, 0xfb,
	0x94, 0xcb, 0x91, 0xac, 0x46, 0xb1, 0x50, 0xa2, 0xaa, 0x62, 0xe2, 0xd3, 0xea, 0xb8, 0x96, 0x2c,
	0x5c, 0x5f, 0xf0, 0x01, 0x0b, 0x6c, 0xed, 0xe1, 0x52, 0xda, 0x9d, 0x28, 0xb6, 0x6e, 0xb2, 0xc7,
	0xb5, 0xf2, 0xb7, 0x0d, 0xd8, 0xb9, 0x98, 0x15, 0x4d, 0x0d, 0xe0, 0x0f, 0x50, 0x8c, 0x62, 0xe1,
	0x11, 0x8f, 0x0d, 0x99, 0xfa, 0xe2, 0x4a, 0x12, 0x46, 0x43, 0x1a, 0x9b, 0xe8, 0x08, 0x1d, 0xef,
	0x9c, 0x54, 0xec, 0x5b, 0x83, 0xec, 0x5e, 0x4a, 0x9d, 0x27, 0xd0, 0xeb, 0x9c, 0x83, 0xa3, 0x35,
	0x15, 0xf7, 0xc1, 0xf0, 0x05, 0x97, 0x8a, 0x70, 0xb5, 0x88, 0xcf, 0xeb, 0xf8, 0xc7, 0x77, 0xc4,
	0x37, 0xe7, 0x48, 0x9a, 0xbd, 0xef, 0x2f, 0x4b, 0xf8, 0x12, 0x0e, 0x62, 0xa2, 0xa8, 0x3b, 0x64,
	0x21, 0x53, 0x8c, 0x07, 0x8b, 0xf4, 0x82, 0x4e, 0xb7, 0xef, 0x48, 0x77, 0x88, 0xa2, 0xed, 0x39,
	0x96, 0xee, 0x50, 0x8c, 0xd7, 0x65, 0xfc, 0x1c, 0xcc, 0x90, 0x5c, 0xb9, 0x7c, 0x14, 0x7a, 0x34,
	0x76, 0xc5, 0xc0, 0x25, 0x4a, 0xc5, 0xcc, 0x1b, 0x29, 0x2a, 0xcd, 0x8d, 0x23, 0x74, 0x5c, 0x70,
	0x0e, 0x42, 0x72, 0xd5, 0xd1, 0x76, 0x77, 0x50, 0x5f, 0x98, 0xf8, 0x05, 0x94, 0x56, 0x40, 0xce,
	0x85, 0x22, 0x8a, 0x09, 0x2e, 0xcd, 0x7b, 0x9a, 0x7c, 0x98, 0x25, 0x53, 0x17, 0xbf, 0x82, 0xc3,
	0x65, 0x34, 0xa4, 0x52, 0x92, 0x80, 0xba, 0x74, 0x4c, 0xb9, 0x92, 0xe6, 0xa6, 0xa6, 0xcd, 0x0c,
	0xfd, 0x26, 0x69, 0x38, 0xd3, 0x3e, 0xae, 0x40, 0x71, 0x99, 0x1f, 0x32, 0xfe, 0x49, 0x9a, 0x5b,
	0x1a, 0x33, 0x32, 0x58, 0x7b, 0xa6, 0x37, 0xb6, 0x61, 0x6b, 0xfe, 0xe9, 0xca, 0x2d, 0xc0, 0xeb,
	0x07, 0x8b, 0x9f, 0x42, 0x51, 0x37, 0x30, 0x1e, 0x64, 0x5c, 0x3d, 0x24, 0xc8, 0xb9, 0xc9, 0x2a,
	0xff, 0x40, 0xb0, 0xbf, 0x72, 0x84, 0xb8, 0x0f, 0xf7, 0x2f, 0xa9, 0xcf, 0x24, 0x13, 0x5c, 0xa3,
	0x7b, 0x27, 0x2f, 0xff, 0x7f, 0x00, 0x16, 0xf5, 0xe9, 0x3c, 0xc2, 0x59, 0x84, 0x95, 0x4f, 0xc1,
	0x58, 0x75, 0xf1, 0x1e, 0x40, 0xbd, 0xdd, 0xaf, 0xbf, 0x3b, 0x77, 0xbb, 0xad, 0x96, 0x91, 0xc3,
	0xbb, 0xb0, 0xfd, 0xaf, 0xee, 0x18, 0x08, 0x3f, 0x80, 0xdd, 0x79, 0xd9, 0xab, 0x3b, 0x67, 0x9d,
	0x0b, 0x23, 0x5f, 0x7e, 0x04, 0xc5, 0x1b, 0xc6, 0x02, 0x1b, 0x50, 0xf8, 0x1c, 0x49, 0xfd, 0xc0,
	0x05, 0x67, 0xb6, 0x6c, 0x7c, 0x47, 0x3f, 0x27, 0x16, 0xba, 0x9e, 0x58, 0xe8, 0xf7, 0xc4, 0x42,
	0x5f, 0xa7, 0x56, 0xee, 0x7a, 0x6a, 0xe5, 0x7e, 0x4d, 0xad, 0x1c, 0x1c, 0x32, 0x71, 0xfb, 0x2b,
	0x35, 0x8c, 0xcc, 0xc5, 0xeb, 0xcd, 0xac, 0x1e, 0x7a, 0xdf, 0x08, 0x98, 0xfa, 0x38, 0xf2, 0x6c,
	0x5f, 0x84, 0xd5, 0x84, 0xaa, 0x30, 0x2e, 0x55, 0x3c, 0x0a, 0x29, 0x4f, 0xc6, 0xa1, 0x9a, 0x06,
	0x56, 0x92, 0xab, 0x1f, 0x50, 0x5e, 0x09, 0xd2, 0x3f, 0xc0, 0x9f, 0x7c, 0xa9, 0x1b, 0x51, 0xde,
	0x4c, 0xf6, 0xd4, 0xc1, 0xb6, 0xde, 0xc9, 0x7e, 0x5b, 0xf3, 0x36, 0x35, 0xf2, 0xec, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x75, 0xe7, 0x3a, 0x4c, 0x41, 0x04, 0x00, 0x00,
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sampler != nil {
		nn1, err1 := m.Sampler.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += nn1
	}
	if m.MaxNumberOfAttributes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAttributes))
	}
	if m.MaxNumberOfAnnotations != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfAnnotations))
	}
	if m.MaxNumberOfMessageEvents != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfMessageEvents))
	}
	if m.MaxNumberOfLinks != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.MaxNumberOfLinks))
	}
	return i, nil
}

func (m *TraceConfig_ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ProbabilitySampler != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.ProbabilitySampler.Size()))
		n2, err2 := m.ProbabilitySampler.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}
func (m *TraceConfig_ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ConstantSampler != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.ConstantSampler.Size()))
		n3, err3 := m.ConstantSampler.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	return i, nil
}
func (m *TraceConfig_RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.RateLimitingSampler != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.RateLimitingSampler.Size()))
		n4, err4 := m.RateLimitingSampler.MarshalTo(dAtA[i:])
		if err4 != nil {
			return 0, err4
		}
		i += n4
	}
	return i, nil
}
func (m *ProbabilitySampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbabilitySampler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SamplingProbability))))
		i += 8
	}
	return i, nil
}

func (m *ConstantSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConstantSampler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Decision != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Decision))
	}
	return i, nil
}

func (m *RateLimitingSampler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitingSampler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Qps != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTraceConfig(dAtA, i, uint64(m.Qps))
	}
	return i, nil
}

func encodeVarintTraceConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sampler != nil {
		n += m.Sampler.Size()
	}
	if m.MaxNumberOfAttributes != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAttributes))
	}
	if m.MaxNumberOfAnnotations != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfAnnotations))
	}
	if m.MaxNumberOfMessageEvents != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfMessageEvents))
	}
	if m.MaxNumberOfLinks != 0 {
		n += 1 + sovTraceConfig(uint64(m.MaxNumberOfLinks))
	}
	return n
}

func (m *TraceConfig_ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProbabilitySampler != nil {
		l = m.ProbabilitySampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConstantSampler != nil {
		l = m.ConstantSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *TraceConfig_RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RateLimitingSampler != nil {
		l = m.RateLimitingSampler.Size()
		n += 1 + l + sovTraceConfig(uint64(l))
	}
	return n
}
func (m *ProbabilitySampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SamplingProbability != 0 {
		n += 9
	}
	return n
}

func (m *ConstantSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Decision != 0 {
		n += 1 + sovTraceConfig(uint64(m.Decision))
	}
	return n
}

func (m *RateLimitingSampler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Qps != 0 {
		n += 1 + sovTraceConfig(uint64(m.Qps))
	}
	return n
}

func sovTraceConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTraceConfig(x uint64) (n int) {
	return sovTraceConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProbabilitySampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ProbabilitySampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ProbabilitySampler{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConstantSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_ConstantSampler{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitingSampler", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTraceConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RateLimitingSampler{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sampler = &TraceConfig_RateLimitingSampler{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAttributes", wireType)
			}
			m.MaxNumberOfAttributes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAttributes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfAnnotations", wireType)
			}
			m.MaxNumberOfAnnotations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfAnnotations |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfMessageEvents", wireType)
			}
			m.MaxNumberOfMessageEvents = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfMessageEvents |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumberOfLinks", wireType)
			}
			m.MaxNumberOfLinks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNumberOfLinks |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProbabilitySampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProbabilitySampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProbabilitySampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingProbability", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SamplingProbability = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConstantSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConstantSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConstantSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decision", wireType)
			}
			m.Decision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decision |= ConstantSampler_ConstantDecision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitingSampler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimitingSampler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitingSampler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qps", wireType)
			}
			m.Qps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Qps |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraceConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTraceConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTraceConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraceConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTraceConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTraceConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTraceConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTraceConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTraceConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTraceConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTraceConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraceConfig   = fmt.Errorf("proto: integer overflow")
)
