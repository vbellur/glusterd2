// Code generated by protoc-gen-go.
// source: peer-svc.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	peer-svc.proto

It has these top-level messages:
	RPCPeerAddReq
	RPCPeerAddResp
*/
package services

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RPCPeerAddReq struct {
	Name             *string  `protobuf:"bytes,1,req" json:"Name,omitempty"`
	Addresses        []string `protobuf:"bytes,2,rep" json:"Addresses,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RPCPeerAddReq) Reset()         { *m = RPCPeerAddReq{} }
func (m *RPCPeerAddReq) String() string { return proto.CompactTextString(m) }
func (*RPCPeerAddReq) ProtoMessage()    {}

func (m *RPCPeerAddReq) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RPCPeerAddReq) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type RPCPeerAddResp struct {
	OpRet            *int32  `protobuf:"varint,1,req" json:"OpRet,omitempty"`
	OpError          *string `protobuf:"bytes,2,req" json:"OpError,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RPCPeerAddResp) Reset()         { *m = RPCPeerAddResp{} }
func (m *RPCPeerAddResp) String() string { return proto.CompactTextString(m) }
func (*RPCPeerAddResp) ProtoMessage()    {}

func (m *RPCPeerAddResp) GetOpRet() int32 {
	if m != nil && m.OpRet != nil {
		return *m.OpRet
	}
	return 0
}

func (m *RPCPeerAddResp) GetOpError() string {
	if m != nil && m.OpError != nil {
		return *m.OpError
	}
	return ""
}
