// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
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

type ValueType int32

const (
	ValueType_STRING  ValueType = 0
	ValueType_BOOL    ValueType = 1
	ValueType_INT64   ValueType = 2
	ValueType_FLOAT64 ValueType = 3
	ValueType_BINARY  ValueType = 4
)

var ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
	3: "FLOAT64",
	4: "BINARY",
}

var ValueType_value = map[string]int32{
	"STRING":  0,
	"BOOL":    1,
	"INT64":   2,
	"FLOAT64": 3,
	"BINARY":  4,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

type SpanRefType int32

const (
	SpanRefType_CHILD_OF     SpanRefType = 0
	SpanRefType_FOLLOWS_FROM SpanRefType = 1
)

var SpanRefType_name = map[int32]string{
	0: "CHILD_OF",
	1: "FOLLOWS_FROM",
}

var SpanRefType_value = map[string]int32{
	"CHILD_OF":     0,
	"FOLLOWS_FROM": 1,
}

func (x SpanRefType) String() string {
	return proto.EnumName(SpanRefType_name, int32(x))
}

func (SpanRefType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

type KeyValue struct {
	Key                  string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	VType                ValueType `protobuf:"varint,2,opt,name=v_type,json=vType,proto3,enum=jaeger.api_v2.ValueType" json:"v_type,omitempty"`
	VStr                 string    `protobuf:"bytes,3,opt,name=v_str,json=vStr,proto3" json:"v_str,omitempty"`
	VBool                bool      `protobuf:"varint,4,opt,name=v_bool,json=vBool,proto3" json:"v_bool,omitempty"`
	VInt64               int64     `protobuf:"varint,5,opt,name=v_int64,json=vInt64,proto3" json:"v_int64,omitempty"`
	VFloat64             float64   `protobuf:"fixed64,6,opt,name=v_float64,json=vFloat64,proto3" json:"v_float64,omitempty"`
	VBinary              []byte    `protobuf:"bytes,7,opt,name=v_binary,json=vBinary,proto3" json:"v_binary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetVType() ValueType {
	if m != nil {
		return m.VType
	}
	return ValueType_STRING
}

func (m *KeyValue) GetVStr() string {
	if m != nil {
		return m.VStr
	}
	return ""
}

func (m *KeyValue) GetVBool() bool {
	if m != nil {
		return m.VBool
	}
	return false
}

func (m *KeyValue) GetVInt64() int64 {
	if m != nil {
		return m.VInt64
	}
	return 0
}

func (m *KeyValue) GetVFloat64() float64 {
	if m != nil {
		return m.VFloat64
	}
	return 0
}

func (m *KeyValue) GetVBinary() []byte {
	if m != nil {
		return m.VBinary
	}
	return nil
}

type Log struct {
	Timestamp            *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Fields               []*KeyValue      `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Log) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

type SpanRef struct {
	TraceId              []byte      `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId               []byte      `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	RefType              SpanRefType `protobuf:"varint,3,opt,name=ref_type,json=refType,proto3,enum=jaeger.api_v2.SpanRefType" json:"ref_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SpanRef) Reset()         { *m = SpanRef{} }
func (m *SpanRef) String() string { return proto.CompactTextString(m) }
func (*SpanRef) ProtoMessage()    {}
func (*SpanRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *SpanRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanRef.Unmarshal(m, b)
}
func (m *SpanRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanRef.Marshal(b, m, deterministic)
}
func (m *SpanRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanRef.Merge(m, src)
}
func (m *SpanRef) XXX_Size() int {
	return xxx_messageInfo_SpanRef.Size(m)
}
func (m *SpanRef) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanRef.DiscardUnknown(m)
}

var xxx_messageInfo_SpanRef proto.InternalMessageInfo

func (m *SpanRef) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *SpanRef) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *SpanRef) GetRefType() SpanRefType {
	if m != nil {
		return m.RefType
	}
	return SpanRefType_CHILD_OF
}

type Process struct {
	ServiceName          string      `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Tags                 []*KeyValue `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

func (m *Process) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Process) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Span struct {
	TraceId              []byte           `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId               []byte           `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	OperationName        string           `protobuf:"bytes,3,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	References           []*SpanRef       `protobuf:"bytes,4,rep,name=references,proto3" json:"references,omitempty"`
	Flags                uint32           `protobuf:"varint,5,opt,name=flags,proto3" json:"flags,omitempty"`
	StartTime            *types.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration             *types.Duration  `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Tags                 []*KeyValue      `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*Log           `protobuf:"bytes,9,rep,name=logs,proto3" json:"logs,omitempty"`
	Process              *Process         `protobuf:"bytes,10,opt,name=process,proto3" json:"process,omitempty"`
	ProcessId            string           `protobuf:"bytes,11,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	Warnings             []string         `protobuf:"bytes,12,rep,name=warnings,proto3" json:"warnings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Span) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Span) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Span) GetReferences() []*SpanRef {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *Span) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Span) GetStartTime() *types.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Span) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Span) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Span) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *Span) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *Span) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

type Trace struct {
	Spans                []*Span                 `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	ProcessMap           []*Trace_ProcessMapping `protobuf:"bytes,2,rep,name=process_map,json=processMap,proto3" json:"process_map,omitempty"`
	Warnings             []string                `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *Trace) GetProcessMap() []*Trace_ProcessMapping {
	if m != nil {
		return m.ProcessMap
	}
	return nil
}

func (m *Trace) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

type Trace_ProcessMapping struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	Process              *Process `protobuf:"bytes,2,opt,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace_ProcessMapping) Reset()         { *m = Trace_ProcessMapping{} }
func (m *Trace_ProcessMapping) String() string { return proto.CompactTextString(m) }
func (*Trace_ProcessMapping) ProtoMessage()    {}
func (*Trace_ProcessMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5, 0}
}

func (m *Trace_ProcessMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace_ProcessMapping.Unmarshal(m, b)
}
func (m *Trace_ProcessMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace_ProcessMapping.Marshal(b, m, deterministic)
}
func (m *Trace_ProcessMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace_ProcessMapping.Merge(m, src)
}
func (m *Trace_ProcessMapping) XXX_Size() int {
	return xxx_messageInfo_Trace_ProcessMapping.Size(m)
}
func (m *Trace_ProcessMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace_ProcessMapping.DiscardUnknown(m)
}

var xxx_messageInfo_Trace_ProcessMapping proto.InternalMessageInfo

func (m *Trace_ProcessMapping) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *Trace_ProcessMapping) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type Batch struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	Process              *Process `protobuf:"bytes,2,opt,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}

func (m *Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Batch.Unmarshal(m, b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return xxx_messageInfo_Batch.Size(m)
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *Batch) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

type DependencyLink struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Child                string   `protobuf:"bytes,2,opt,name=child,proto3" json:"child,omitempty"`
	CallCount            uint64   `protobuf:"varint,3,opt,name=call_count,json=callCount,proto3" json:"call_count,omitempty"`
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DependencyLink) Reset()         { *m = DependencyLink{} }
func (m *DependencyLink) String() string { return proto.CompactTextString(m) }
func (*DependencyLink) ProtoMessage()    {}
func (*DependencyLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{7}
}

func (m *DependencyLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DependencyLink.Unmarshal(m, b)
}
func (m *DependencyLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DependencyLink.Marshal(b, m, deterministic)
}
func (m *DependencyLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DependencyLink.Merge(m, src)
}
func (m *DependencyLink) XXX_Size() int {
	return xxx_messageInfo_DependencyLink.Size(m)
}
func (m *DependencyLink) XXX_DiscardUnknown() {
	xxx_messageInfo_DependencyLink.DiscardUnknown(m)
}

var xxx_messageInfo_DependencyLink proto.InternalMessageInfo

func (m *DependencyLink) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *DependencyLink) GetChild() string {
	if m != nil {
		return m.Child
	}
	return ""
}

func (m *DependencyLink) GetCallCount() uint64 {
	if m != nil {
		return m.CallCount
	}
	return 0
}

func (m *DependencyLink) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func init() {
	proto.RegisterEnum("jaeger.api_v2.ValueType", ValueType_name, ValueType_value)
	proto.RegisterEnum("jaeger.api_v2.SpanRefType", SpanRefType_name, SpanRefType_value)
	proto.RegisterType((*KeyValue)(nil), "jaeger.api_v2.KeyValue")
	proto.RegisterType((*Log)(nil), "jaeger.api_v2.Log")
	proto.RegisterType((*SpanRef)(nil), "jaeger.api_v2.SpanRef")
	proto.RegisterType((*Process)(nil), "jaeger.api_v2.Process")
	proto.RegisterType((*Span)(nil), "jaeger.api_v2.Span")
	proto.RegisterType((*Trace)(nil), "jaeger.api_v2.Trace")
	proto.RegisterType((*Trace_ProcessMapping)(nil), "jaeger.api_v2.Trace.ProcessMapping")
	proto.RegisterType((*Batch)(nil), "jaeger.api_v2.Batch")
	proto.RegisterType((*DependencyLink)(nil), "jaeger.api_v2.DependencyLink")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0xee, 0x24, 0x76, 0x6c, 0xbf, 0x64, 0x57, 0xd1, 0xb4, 0x74, 0xdd, 0x20, 0x36, 0x21, 0x15,
	0x52, 0xa8, 0x4a, 0xb6, 0x5d, 0xda, 0x3d, 0x20, 0x24, 0x54, 0xef, 0x12, 0x30, 0x64, 0x37, 0x68,
	0x76, 0x05, 0x82, 0x8b, 0x35, 0xeb, 0x4c, 0x5c, 0xb7, 0x8e, 0xc7, 0xb2, 0x1d, 0xa3, 0xdc, 0xf8,
	0x09, 0x88, 0x13, 0x47, 0xb8, 0xf2, 0x2b, 0x38, 0xf6, 0xc8, 0x81, 0x13, 0x12, 0x0b, 0x5a, 0x2e,
	0xfd, 0x19, 0x68, 0xc6, 0xe3, 0x6c, 0x37, 0x54, 0xb0, 0xbd, 0x70, 0xca, 0xbc, 0x99, 0xef, 0x7b,
	0xf3, 0xde, 0xf7, 0xbe, 0x89, 0xa1, 0x39, 0xe7, 0x53, 0x16, 0x0d, 0x93, 0x94, 0xe7, 0x1c, 0x6f,
	0x3c, 0xa1, 0x2c, 0x60, 0xe9, 0x90, 0x26, 0xa1, 0x57, 0xec, 0x76, 0x6e, 0x04, 0x3c, 0xe0, 0xf2,
	0x64, 0x47, 0xac, 0x4a, 0x50, 0xa7, 0x1b, 0x70, 0x1e, 0x44, 0x6c, 0x47, 0x46, 0xa7, 0x8b, 0xd9,
	0x4e, 0x1e, 0xce, 0x59, 0x96, 0xd3, 0x79, 0xa2, 0x00, 0xdb, 0xeb, 0x80, 0xe9, 0x22, 0xa5, 0x79,
	0xc8, 0xe3, 0xf2, 0xbc, 0xff, 0x2b, 0x02, 0xf3, 0x53, 0xb6, 0xfc, 0x9c, 0x46, 0x0b, 0x86, 0xdb,
	0x50, 0x7f, 0xca, 0x96, 0x36, 0xea, 0xa1, 0x81, 0x45, 0xc4, 0x12, 0xef, 0x40, 0xa3, 0xf0, 0xf2,
	0x65, 0xc2, 0xec, 0x5a, 0x0f, 0x0d, 0x36, 0x77, 0xed, 0xe1, 0xa5, 0xaa, 0x86, 0x92, 0x77, 0xb2,
	0x4c, 0x18, 0xd1, 0x0b, 0xf1, 0x83, 0xaf, 0x83, 0x5e, 0x78, 0x59, 0x9e, 0xda, 0x75, 0x99, 0x44,
	0x2b, 0x8e, 0xf3, 0x14, 0xbf, 0x26, 0xb2, 0x9c, 0x72, 0x1e, 0xd9, 0x5a, 0x0f, 0x0d, 0x4c, 0xa2,
	0x17, 0x0e, 0xe7, 0x11, 0xde, 0x02, 0xa3, 0xf0, 0xc2, 0x38, 0xdf, 0x7b, 0x60, 0xeb, 0x3d, 0x34,
	0xa8, 0x93, 0x46, 0xe1, 0x8a, 0x08, 0xbf, 0x0e, 0x56, 0xe1, 0xcd, 0x22, 0x4e, 0xc5, 0x51, 0xa3,
	0x87, 0x06, 0x88, 0x98, 0xc5, 0xa8, 0x8c, 0xf1, 0x2d, 0x30, 0x0b, 0xef, 0x34, 0x8c, 0x69, 0xba,
	0xb4, 0x8d, 0x1e, 0x1a, 0xb4, 0x88, 0x51, 0x38, 0x32, 0x7c, 0xcf, 0x7c, 0xfe, 0x43, 0x17, 0x3d,
	0xff, 0xb1, 0x8b, 0xfa, 0xdf, 0x20, 0xa8, 0x8f, 0x79, 0x80, 0x1d, 0xb0, 0x56, 0x8a, 0xc8, 0xbe,
	0x9a, 0xbb, 0x9d, 0x61, 0x29, 0xc9, 0xb0, 0x92, 0x64, 0x78, 0x52, 0x21, 0x1c, 0xf3, 0xd9, 0x59,
	0xf7, 0xda, 0xb7, 0x7f, 0x74, 0x11, 0xb9, 0xa0, 0xe1, 0x87, 0xd0, 0x98, 0x85, 0x2c, 0x9a, 0x66,
	0x76, 0xad, 0x57, 0x1f, 0x34, 0x77, 0xb7, 0xd6, 0x34, 0xa8, 0xe4, 0x73, 0x34, 0xc1, 0x26, 0x0a,
	0xdc, 0xff, 0x09, 0x81, 0x71, 0x9c, 0xd0, 0x98, 0xb0, 0x19, 0x7e, 0x08, 0x66, 0x9e, 0x52, 0x9f,
	0x79, 0xe1, 0x54, 0x56, 0xd1, 0x72, 0x3a, 0x02, 0xfb, 0xdb, 0x59, 0xd7, 0x38, 0x11, 0xfb, 0xee,
	0xc1, 0xf9, 0xc5, 0x92, 0x18, 0x12, 0xeb, 0x4e, 0xf1, 0x7d, 0x30, 0xb2, 0x84, 0xc6, 0x82, 0x55,
	0x93, 0x2c, 0x5b, 0xb1, 0x1a, 0x22, 0xb1, 0x24, 0xa9, 0x15, 0x69, 0x08, 0xa0, 0x3b, 0x15, 0x37,
	0xa5, 0x6c, 0x56, 0x8e, 0xac, 0x2e, 0x47, 0xd6, 0x59, 0x2b, 0x57, 0xd5, 0x24, 0x87, 0x66, 0xa4,
	0xe5, 0xa2, 0xef, 0x81, 0xf1, 0x59, 0xca, 0x7d, 0x96, 0x65, 0xf8, 0x4d, 0x68, 0x65, 0x2c, 0x2d,
	0x42, 0x9f, 0x79, 0x31, 0x9d, 0x33, 0xe5, 0x86, 0xa6, 0xda, 0x3b, 0xa2, 0x73, 0x86, 0xef, 0x83,
	0x96, 0xd3, 0xe0, 0x8a, 0x7a, 0x48, 0x68, 0xff, 0x77, 0x0d, 0x34, 0x71, 0xf3, 0xff, 0x28, 0xc5,
	0x5b, 0xb0, 0xc9, 0x13, 0x56, 0xba, 0xbd, 0x6c, 0xa5, 0xf4, 0xe4, 0xc6, 0x6a, 0x57, 0x36, 0xf3,
	0x3e, 0x40, 0xca, 0x66, 0x2c, 0x65, 0xb1, 0xcf, 0x32, 0x5b, 0x93, 0x2d, 0xdd, 0x7c, 0xb9, 0x66,
	0xaa, 0xa3, 0x17, 0xf0, 0xf8, 0x36, 0xe8, 0xb3, 0x48, 0x68, 0x21, 0x1c, 0xbc, 0xe1, 0x6c, 0xa8,
	0xaa, 0xf4, 0x91, 0xd8, 0x24, 0xe5, 0x19, 0xde, 0x07, 0xc8, 0x72, 0x9a, 0xe6, 0x9e, 0x30, 0x95,
	0x34, 0xf4, 0x95, 0x6d, 0x28, 0x79, 0xe2, 0x04, 0x7f, 0x00, 0x66, 0xf5, 0x76, 0xa5, 0xef, 0x9b,
	0xbb, 0xb7, 0xfe, 0x91, 0xe2, 0x40, 0x01, 0xca, 0x0c, 0xdf, 0x8b, 0x0c, 0x2b, 0xd2, 0x6a, 0x6a,
	0xe6, 0x95, 0xa7, 0x86, 0xef, 0x82, 0x16, 0xf1, 0x20, 0xb3, 0x2d, 0x49, 0xc1, 0x6b, 0x94, 0x31,
	0x0f, 0x2a, 0xb4, 0x40, 0xe1, 0x7b, 0x60, 0x24, 0xa5, 0x89, 0x6c, 0x90, 0x05, 0xae, 0xcb, 0xa8,
	0x2c, 0x46, 0x2a, 0x18, 0xbe, 0x0b, 0xa0, 0x96, 0x62, 0xb0, 0x4d, 0x31, 0x1e, 0x67, 0xe3, 0xfc,
	0xac, 0x6b, 0x29, 0xa4, 0x7b, 0x40, 0x2c, 0x05, 0x70, 0xa7, 0xb8, 0x03, 0xe6, 0xd7, 0x34, 0x8d,
	0xc3, 0x38, 0xc8, 0xec, 0x56, 0xaf, 0x3e, 0xb0, 0xc8, 0x2a, 0xee, 0x7f, 0x57, 0x03, 0x5d, 0x9a,
	0x06, 0xbf, 0x0d, 0xba, 0x30, 0x40, 0x66, 0x23, 0x59, 0xf4, 0xf5, 0x97, 0x8d, 0xb2, 0x44, 0xe0,
	0x4f, 0xa0, 0x59, 0x5d, 0x3f, 0xa7, 0x89, 0xb2, 0xf3, 0xed, 0x35, 0x82, 0xcc, 0x5a, 0x95, 0x7e,
	0x48, 0x93, 0x24, 0x8c, 0xab, 0xb6, 0xab, 0xe2, 0x0f, 0x69, 0x72, 0xa9, 0xb8, 0xfa, 0xe5, 0xe2,
	0x3a, 0x05, 0x6c, 0x5e, 0xe6, 0xaf, 0x35, 0x8e, 0xfe, 0xa3, 0xf1, 0xbd, 0x0b, 0x61, 0x6b, 0xff,
	0x26, 0xac, 0x2a, 0xab, 0x02, 0xf7, 0x9f, 0x80, 0xee, 0xd0, 0xdc, 0x7f, 0xfc, 0x2a, 0x9a, 0xbc,
	0xd2, 0x5d, 0xe8, 0xe2, 0xae, 0x05, 0x6c, 0x1e, 0xb0, 0x84, 0xc5, 0x53, 0x16, 0xfb, 0xcb, 0x71,
	0x18, 0x3f, 0xc5, 0x37, 0xa1, 0x91, 0xd0, 0x94, 0xc5, 0xb9, 0xfa, 0x0b, 0x51, 0x11, 0xbe, 0x01,
	0xba, 0xff, 0x38, 0x8c, 0xca, 0x87, 0x6c, 0x91, 0x32, 0xc0, 0x6f, 0x00, 0xf8, 0x34, 0x8a, 0x3c,
	0x9f, 0x2f, 0xe2, 0x5c, 0xbe, 0x54, 0x8d, 0x58, 0x62, 0x67, 0x5f, 0x6c, 0x88, 0x64, 0x19, 0x5f,
	0xa4, 0x3e, 0x93, 0x9f, 0x10, 0x8b, 0xa8, 0xe8, 0xce, 0x87, 0x60, 0xad, 0xbe, 0x41, 0x18, 0xa0,
	0x71, 0x7c, 0x42, 0xdc, 0xa3, 0x8f, 0xda, 0xd7, 0xb0, 0x09, 0x9a, 0x33, 0x99, 0x8c, 0xdb, 0x08,
	0x5b, 0xa0, 0xbb, 0x47, 0x27, 0x7b, 0x0f, 0xda, 0x35, 0xdc, 0x04, 0x63, 0x34, 0x9e, 0x3c, 0x12,
	0x41, 0x5d, 0xa0, 0x1d, 0xf7, 0xe8, 0x11, 0xf9, 0xb2, 0xad, 0xdd, 0x79, 0x07, 0x9a, 0x2f, 0xfc,
	0x2f, 0xe2, 0x16, 0x98, 0xfb, 0x1f, 0xbb, 0xe3, 0x03, 0x6f, 0x32, 0x6a, 0x5f, 0xc3, 0x6d, 0x68,
	0x8d, 0x26, 0xe3, 0xf1, 0xe4, 0x8b, 0x63, 0x6f, 0x44, 0x26, 0x87, 0x6d, 0xe4, 0xdc, 0x7b, 0x76,
	0xbe, 0x8d, 0x7e, 0x39, 0xdf, 0x46, 0x7f, 0x9e, 0x6f, 0xa3, 0x9f, 0xff, 0xda, 0x46, 0xb0, 0x15,
	0x72, 0xa5, 0x93, 0xf8, 0xc7, 0x0a, 0xe3, 0x40, 0xc9, 0xf5, 0x95, 0x2e, 0xbf, 0xe9, 0xa7, 0x0d,
	0xf9, 0x46, 0xdf, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x18, 0x22, 0x2c, 0x8e, 0xe3, 0x07, 0x00,
	0x00,
}
