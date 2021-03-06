// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beeswax/base/eventid.proto

/*
Package base is a generated protocol buffer package.

It is generated from these files:
	beeswax/base/eventid.proto

It has these top-level messages:
	EventId
	AdGroupId
*/
package base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventId struct {
	Hostid           *uint32 `protobuf:"varint,1,opt,name=hostid" json:"hostid,omitempty"`
	Tid              *uint32 `protobuf:"varint,2,opt,name=tid" json:"tid,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EventId) Reset()                    { *m = EventId{} }
func (m *EventId) String() string            { return proto.CompactTextString(m) }
func (*EventId) ProtoMessage()               {}
func (*EventId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EventId) GetHostid() uint32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *EventId) GetTid() uint32 {
	if m != nil && m.Tid != nil {
		return *m.Tid
	}
	return 0
}

func (m *EventId) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

// Next Tag: 5
type AdGroupId struct {
	BuzzKey          *string `protobuf:"bytes,1,opt,name=buzz_key,json=buzzKey" json:"buzz_key,omitempty"`
	Accountid        *uint64 `protobuf:"varint,2,opt,name=accountid" json:"accountid,omitempty"`
	Campaignid       *uint64 `protobuf:"varint,3,opt,name=campaignid" json:"campaignid,omitempty"`
	Lineitemid       *uint64 `protobuf:"varint,4,opt,name=lineitemid" json:"lineitemid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdGroupId) Reset()                    { *m = AdGroupId{} }
func (m *AdGroupId) String() string            { return proto.CompactTextString(m) }
func (*AdGroupId) ProtoMessage()               {}
func (*AdGroupId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AdGroupId) GetBuzzKey() string {
	if m != nil && m.BuzzKey != nil {
		return *m.BuzzKey
	}
	return ""
}

func (m *AdGroupId) GetAccountid() uint64 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *AdGroupId) GetCampaignid() uint64 {
	if m != nil && m.Campaignid != nil {
		return *m.Campaignid
	}
	return 0
}

func (m *AdGroupId) GetLineitemid() uint64 {
	if m != nil && m.Lineitemid != nil {
		return *m.Lineitemid
	}
	return 0
}

func init() {
	proto.RegisterType((*EventId)(nil), "base.EventId")
	proto.RegisterType((*AdGroupId)(nil), "base.AdGroupId")
}

func init() { proto.RegisterFile("beeswax/base/eventid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x05, 0x60, 0xe2, 0x04, 0xc7, 0x5e, 0x10, 0x86, 0x2c, 0xa4, 0x8a, 0xc8, 0x30, 0xab, 0x59,
	0x75, 0x9e, 0x41, 0x41, 0x64, 0x70, 0x65, 0x5e, 0x40, 0xd2, 0xe4, 0xa2, 0x17, 0xcd, 0x0f, 0x4d,
	0xaa, 0xb6, 0x6b, 0x1f, 0x5c, 0x6e, 0xad, 0xd4, 0x5d, 0xce, 0x77, 0xe0, 0x84, 0x04, 0xae, 0x5a,
	0xc4, 0xfc, 0x69, 0xbe, 0x0e, 0xad, 0xc9, 0x78, 0xc0, 0x0f, 0x0c, 0x85, 0x5c, 0x93, 0xba, 0x58,
	0xa2, 0x92, 0x6c, 0xbb, 0x27, 0x58, 0xdf, 0x33, 0x1f, 0x9d, 0xba, 0x80, 0xd3, 0xd7, 0x98, 0x0b,
	0xb9, 0x5a, 0x6c, 0xc5, 0xfe, 0x5c, 0xcf, 0x49, 0x6d, 0x60, 0xc5, 0x78, 0x32, 0x21, 0x1f, 0xd5,
	0x35, 0x54, 0x85, 0x3c, 0xe6, 0x62, 0x7c, 0xaa, 0x57, 0x5b, 0xb1, 0x97, 0x7a, 0x81, 0xdd, 0xb7,
	0x80, 0xea, 0xd6, 0x3d, 0x74, 0xb1, 0x4f, 0x47, 0xa7, 0x2e, 0xe1, 0xac, 0xed, 0xc7, 0xf1, 0xf9,
	0x0d, 0x87, 0x69, 0xb7, 0xd2, 0x6b, 0xce, 0x8f, 0x38, 0xf0, 0x8c, 0xb1, 0x36, 0xf6, 0xe1, 0x6f,
	0x5e, 0xea, 0x05, 0xd4, 0x0d, 0x80, 0x35, 0x3e, 0x19, 0x7a, 0x09, 0xe4, 0xe6, 0x5b, 0xfe, 0x09,
	0xf7, 0xef, 0x14, 0x90, 0x0a, 0x7a, 0x72, 0xb5, 0xfc, 0xed, 0x17, 0xb9, 0x53, 0xb0, 0xb1, 0xd1,
	0x37, 0xf3, 0x0f, 0x34, 0xfc, 0xda, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xba, 0x8d, 0x7a,
	0x10, 0x01, 0x00, 0x00,
}
