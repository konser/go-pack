// Code generated by protoc-gen-go.
// source: protocol/packetid.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type MmoPacketID int32

const (
	MmoPacketID_PACKET_GB_CUR_LOAD     MmoPacketID = 1000
	MmoPacketID_PACKET_GB_STATE_SWITCH MmoPacketID = 1001
	MmoPacketID_PACKET_SC_GATEINFO     MmoPacketID = 1002
)

var MmoPacketID_name = map[int32]string{
	1000: "PACKET_GB_CUR_LOAD",
	1001: "PACKET_GB_STATE_SWITCH",
	1002: "PACKET_SC_GATEINFO",
}
var MmoPacketID_value = map[string]int32{
	"PACKET_GB_CUR_LOAD":     1000,
	"PACKET_GB_STATE_SWITCH": 1001,
	"PACKET_SC_GATEINFO":     1002,
}

func (x MmoPacketID) Enum() *MmoPacketID {
	p := new(MmoPacketID)
	*p = x
	return p
}
func (x MmoPacketID) String() string {
	return proto.EnumName(MmoPacketID_name, int32(x))
}
func (x MmoPacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *MmoPacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MmoPacketID_value, data, "MmoPacketID")
	if err != nil {
		return err
	}
	*x = MmoPacketID(value)
	return nil
}

func init() {
	proto.RegisterEnum("protocol.MmoPacketID", MmoPacketID_name, MmoPacketID_value)
}