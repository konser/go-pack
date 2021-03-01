// Code generated by protoc-gen-go.
// source: protocol/slices.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SSPacketSlices struct {
	SeqNo            *int32 `protobuf:"varint,1,req" json:"SeqNo,omitempty"`
	TotalSize        *int32 `protobuf:"varint,2,req" json:"TotalSize,omitempty"`
	Offset           *int32 `protobuf:"varint,3,req" json:"Offset,omitempty"`
	PacketData       []byte `protobuf:"bytes,4,req" json:"PacketData,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SSPacketSlices) Reset()         { *m = SSPacketSlices{} }
func (m *SSPacketSlices) String() string { return proto.CompactTextString(m) }
func (*SSPacketSlices) ProtoMessage()    {}

func (m *SSPacketSlices) GetSeqNo() int32 {
	if m != nil && m.SeqNo != nil {
		return *m.SeqNo
	}
	return 0
}

func (m *SSPacketSlices) GetTotalSize() int32 {
	if m != nil && m.TotalSize != nil {
		return *m.TotalSize
	}
	return 0
}

func (m *SSPacketSlices) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *SSPacketSlices) GetPacketData() []byte {
	if m != nil {
		return m.PacketData
	}
	return nil
}

func init() {
}
