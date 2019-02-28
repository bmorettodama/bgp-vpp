// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bgp.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type PeerConf_RemovePrivateAs int32

const (
	PeerConf_NONE    PeerConf_RemovePrivateAs = 0
	PeerConf_ALL     PeerConf_RemovePrivateAs = 1
	PeerConf_REPLACE PeerConf_RemovePrivateAs = 2
)

var PeerConf_RemovePrivateAs_name = map[int32]string{
	0: "NONE",
	1: "ALL",
	2: "REPLACE",
}

var PeerConf_RemovePrivateAs_value = map[string]int32{
	"NONE":    0,
	"ALL":     1,
	"REPLACE": 2,
}

func (x PeerConf_RemovePrivateAs) String() string {
	return proto.EnumName(PeerConf_RemovePrivateAs_name, int32(x))
}

func (PeerConf_RemovePrivateAs) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2, 0}
}

//BGP configuration
type BgpConf struct {
	Global               *Global     `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Peers                []*PeerConf `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BgpConf) Reset()         { *m = BgpConf{} }
func (m *BgpConf) String() string { return proto.CompactTextString(m) }
func (*BgpConf) ProtoMessage()    {}
func (*BgpConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{0}
}
func (m *BgpConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConf.Unmarshal(m, b)
}
func (m *BgpConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConf.Marshal(b, m, deterministic)
}
func (m *BgpConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConf.Merge(m, src)
}
func (m *BgpConf) XXX_Size() int {
	return xxx_messageInfo_BgpConf.Size(m)
}
func (m *BgpConf) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConf.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConf proto.InternalMessageInfo

func (m *BgpConf) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpConf) GetPeers() []*PeerConf {
	if m != nil {
		return m.Peers
	}
	return nil
}

// global configuration
type Global struct {
	As                   uint32   `protobuf:"varint,1,opt,name=as,proto3" json:"as,omitempty"`
	RouterId             string   `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	ListenPort           int32    `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	ListenAddresses      []string `protobuf:"bytes,4,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listen_addresses,omitempty"`
	Families             []uint32 `protobuf:"varint,5,rep,packed,name=families,proto3" json:"families,omitempty"`
	UseMultiplePaths     bool     `protobuf:"varint,6,opt,name=use_multiple_paths,json=useMultiplePaths,proto3" json:"use_multiple_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{1}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (m *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(m, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *Global) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *Global) GetListenPort() int32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *Global) GetListenAddresses() []string {
	if m != nil {
		return m.ListenAddresses
	}
	return nil
}

func (m *Global) GetFamilies() []uint32 {
	if m != nil {
		return m.Families
	}
	return nil
}

func (m *Global) GetUseMultiplePaths() bool {
	if m != nil {
		return m.UseMultiplePaths
	}
	return false
}

// neighbor configuration
type PeerConf struct {
	AuthPassword         string                   `protobuf:"bytes,1,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	Description          string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	LocalAs              uint32                   `protobuf:"varint,3,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	NeighborAddress      string                   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	PeerAs               uint32                   `protobuf:"varint,5,opt,name=peer_as,json=peerAs,proto3" json:"peer_as,omitempty"`
	PeerGroup            string                   `protobuf:"bytes,6,opt,name=peer_group,json=peerGroup,proto3" json:"peer_group,omitempty"`
	PeerType             uint32                   `protobuf:"varint,7,opt,name=peer_type,json=peerType,proto3" json:"peer_type,omitempty"`
	RemovePrivateAs      PeerConf_RemovePrivateAs `protobuf:"varint,8,opt,name=remove_private_as,json=removePrivateAs,proto3,enum=model.PeerConf_RemovePrivateAs" json:"remove_private_as,omitempty"`
	RouteFlapDamping     bool                     `protobuf:"varint,9,opt,name=route_flap_damping,json=routeFlapDamping,proto3" json:"route_flap_damping,omitempty"`
	SendCommunity        uint32                   `protobuf:"varint,10,opt,name=send_community,json=sendCommunity,proto3" json:"send_community,omitempty"`
	NeighborInterface    string                   `protobuf:"bytes,11,opt,name=neighbor_interface,json=neighborInterface,proto3" json:"neighbor_interface,omitempty"`
	Vrf                  string                   `protobuf:"bytes,12,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AllowOwnAs           uint32                   `protobuf:"varint,13,opt,name=allow_own_as,json=allowOwnAs,proto3" json:"allow_own_as,omitempty"`
	ReplacePeerAs        bool                     `protobuf:"varint,14,opt,name=replace_peer_as,json=replacePeerAs,proto3" json:"replace_peer_as,omitempty"`
	AdminDown            bool                     `protobuf:"varint,15,opt,name=admin_down,json=adminDown,proto3" json:"admin_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PeerConf) Reset()         { *m = PeerConf{} }
func (m *PeerConf) String() string { return proto.CompactTextString(m) }
func (*PeerConf) ProtoMessage()    {}
func (*PeerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e12fe45eec524a6, []int{2}
}
func (m *PeerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerConf.Unmarshal(m, b)
}
func (m *PeerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerConf.Marshal(b, m, deterministic)
}
func (m *PeerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerConf.Merge(m, src)
}
func (m *PeerConf) XXX_Size() int {
	return xxx_messageInfo_PeerConf.Size(m)
}
func (m *PeerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerConf.DiscardUnknown(m)
}

var xxx_messageInfo_PeerConf proto.InternalMessageInfo

func (m *PeerConf) GetAuthPassword() string {
	if m != nil {
		return m.AuthPassword
	}
	return ""
}

func (m *PeerConf) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PeerConf) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *PeerConf) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *PeerConf) GetPeerAs() uint32 {
	if m != nil {
		return m.PeerAs
	}
	return 0
}

func (m *PeerConf) GetPeerGroup() string {
	if m != nil {
		return m.PeerGroup
	}
	return ""
}

func (m *PeerConf) GetPeerType() uint32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *PeerConf) GetRemovePrivateAs() PeerConf_RemovePrivateAs {
	if m != nil {
		return m.RemovePrivateAs
	}
	return PeerConf_NONE
}

func (m *PeerConf) GetRouteFlapDamping() bool {
	if m != nil {
		return m.RouteFlapDamping
	}
	return false
}

func (m *PeerConf) GetSendCommunity() uint32 {
	if m != nil {
		return m.SendCommunity
	}
	return 0
}

func (m *PeerConf) GetNeighborInterface() string {
	if m != nil {
		return m.NeighborInterface
	}
	return ""
}

func (m *PeerConf) GetVrf() string {
	if m != nil {
		return m.Vrf
	}
	return ""
}

func (m *PeerConf) GetAllowOwnAs() uint32 {
	if m != nil {
		return m.AllowOwnAs
	}
	return 0
}

func (m *PeerConf) GetReplacePeerAs() bool {
	if m != nil {
		return m.ReplacePeerAs
	}
	return false
}

func (m *PeerConf) GetAdminDown() bool {
	if m != nil {
		return m.AdminDown
	}
	return false
}

func init() {
	proto.RegisterEnum("model.PeerConf_RemovePrivateAs", PeerConf_RemovePrivateAs_name, PeerConf_RemovePrivateAs_value)
	proto.RegisterType((*BgpConf)(nil), "model.BgpConf")
	proto.RegisterType((*Global)(nil), "model.Global")
	proto.RegisterType((*PeerConf)(nil), "model.PeerConf")
}

func init() { proto.RegisterFile("bgp.proto", fileDescriptor_2e12fe45eec524a6) }

var fileDescriptor_2e12fe45eec524a6 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0xff, 0x26, 0x5f, 0x97, 0xb6, 0xf3, 0x0d, 0x06, 0x84, 0x16, 0x15, 0x0d, 0x05,
	0x09, 0x2a, 0x31, 0x9e, 0x20, 0x6c, 0x63, 0x9a, 0x18, 0x5b, 0x64, 0x21, 0x71, 0x19, 0xb9, 0x8d,
	0xdb, 0x45, 0x72, 0x6c, 0xcb, 0x76, 0x56, 0xed, 0x19, 0x79, 0x14, 0x5e, 0x02, 0xd9, 0x49, 0x26,
	0xd8, 0x5d, 0xfd, 0x3b, 0x9f, 0x4f, 0x7d, 0xce, 0xa7, 0x40, 0xb4, 0xd9, 0xab, 0xb5, 0xd2, 0xd2,
	0x4a, 0x34, 0xae, 0x65, 0xc9, 0xf8, 0xea, 0x17, 0x4c, 0xbf, 0xee, 0xd5, 0xb9, 0x14, 0x3b, 0x74,
	0x0a, 0x93, 0x3d, 0x97, 0x1b, 0xca, 0x71, 0x90, 0x04, 0xe9, 0xec, 0x2c, 0x5e, 0xfb, 0x91, 0xf5,
	0x95, 0x87, 0xa4, 0x13, 0xd1, 0x29, 0x8c, 0x15, 0x63, 0xda, 0xe0, 0x41, 0x32, 0x4c, 0x67, 0x67,
	0x8b, 0x6e, 0x2a, 0x67, 0x4c, 0x3b, 0x1b, 0xd2, 0xaa, 0xab, 0xdf, 0x01, 0x4c, 0xda, 0x9b, 0x68,
	0x0e, 0x03, 0x6a, 0xbc, 0x69, 0x4c, 0x06, 0xd4, 0xa0, 0x37, 0x10, 0x69, 0xd9, 0x58, 0xa6, 0x8b,
	0xaa, 0xc4, 0x83, 0x24, 0x48, 0x23, 0x12, 0xb6, 0xe0, 0xba, 0x44, 0x27, 0x30, 0xe3, 0x95, 0xb1,
	0x4c, 0x14, 0x4a, 0x6a, 0x8b, 0x87, 0x49, 0x90, 0x8e, 0x09, 0xb4, 0x28, 0x97, 0xda, 0xa2, 0x0f,
	0xb0, 0xec, 0x06, 0x68, 0x59, 0x6a, 0x66, 0x0c, 0x33, 0x78, 0x94, 0x0c, 0xd3, 0x88, 0x2c, 0x5a,
	0x9e, 0xf5, 0x18, 0xbd, 0x86, 0x70, 0x47, 0xeb, 0x8a, 0x57, 0xcc, 0xe0, 0x71, 0x32, 0x4c, 0x63,
	0xf2, 0x74, 0x46, 0x1f, 0x01, 0x35, 0x86, 0x15, 0x75, 0xc3, 0x6d, 0xa5, 0x38, 0x2b, 0x14, 0xb5,
	0xf7, 0x06, 0x4f, 0x92, 0x20, 0x0d, 0xc9, 0xb2, 0x31, 0xec, 0x47, 0x27, 0xe4, 0x8e, 0xaf, 0xfe,
	0x8c, 0x20, 0xec, 0x13, 0xa2, 0x77, 0x10, 0xd3, 0xc6, 0xde, 0x17, 0x8a, 0x1a, 0x73, 0x90, 0xba,
	0xf4, 0xd1, 0x22, 0x72, 0xe4, 0x60, 0xde, 0x31, 0x94, 0xc0, 0xac, 0x64, 0x66, 0xab, 0x2b, 0x65,
	0x2b, 0x29, 0xba, 0x98, 0xff, 0x22, 0xf4, 0x0a, 0x42, 0x2e, 0xb7, 0x94, 0x17, 0xd4, 0xf8, 0x98,
	0x31, 0x99, 0xfa, 0x73, 0x66, 0x5c, 0x46, 0xc1, 0xaa, 0xfd, 0xfd, 0x46, 0xea, 0x3e, 0x25, 0x1e,
	0x79, 0x87, 0x45, 0xcf, 0xbb, 0x94, 0xe8, 0x25, 0x4c, 0x5d, 0xe1, 0xce, 0x64, 0xec, 0x4d, 0x26,
	0xee, 0x98, 0x19, 0xf4, 0x16, 0xc0, 0x0b, 0x7b, 0x2d, 0x1b, 0xe5, 0x83, 0x45, 0x24, 0x72, 0xe4,
	0xca, 0x01, 0xb7, 0x04, 0x2f, 0xdb, 0x47, 0xc5, 0xf0, 0xd4, 0xdf, 0x0c, 0x1d, 0xf8, 0xf9, 0xa8,
	0x18, 0xfa, 0x0e, 0xc7, 0x9a, 0xd5, 0xf2, 0x81, 0x15, 0x4a, 0x57, 0x0f, 0xd4, 0x32, 0x67, 0x1f,
	0x26, 0x41, 0x3a, 0x3f, 0x3b, 0x79, 0xb6, 0xef, 0x35, 0xf1, 0x83, 0x79, 0x3b, 0x97, 0x19, 0xb2,
	0xd0, 0xff, 0x03, 0xd7, 0xb4, 0xdf, 0x6e, 0xb1, 0xe3, 0x54, 0x15, 0x25, 0xad, 0x55, 0x25, 0xf6,
	0x38, 0x6a, 0x9b, 0xf6, 0xca, 0x37, 0x4e, 0xd5, 0x45, 0xcb, 0xd1, 0x29, 0xcc, 0x0d, 0x13, 0x65,
	0xb1, 0x95, 0x75, 0xdd, 0x88, 0xca, 0x3e, 0x62, 0xf0, 0x8f, 0x8b, 0x1d, 0x3d, 0xef, 0x21, 0xfa,
	0x04, 0xe8, 0xa9, 0xa1, 0x4a, 0x58, 0xa6, 0x77, 0x74, 0xcb, 0xf0, 0xcc, 0xa7, 0x3c, 0xee, 0x95,
	0xeb, 0x5e, 0x40, 0x4b, 0x18, 0x3e, 0xe8, 0x1d, 0x3e, 0xf2, 0xba, 0xfb, 0x89, 0x12, 0x38, 0xa2,
	0x9c, 0xcb, 0x43, 0x21, 0x0f, 0xc2, 0xa5, 0x8b, 0xfd, 0xbf, 0x80, 0x67, 0x77, 0x07, 0x91, 0x19,
	0xf4, 0x1e, 0x16, 0x9a, 0x29, 0x4e, 0xb7, 0xac, 0xe8, 0x1b, 0x9e, 0xfb, 0x47, 0xc7, 0x1d, 0xce,
	0x9f, 0x8a, 0xa6, 0x65, 0x5d, 0x89, 0xa2, 0x94, 0x07, 0x81, 0x17, 0x7e, 0x24, 0xf2, 0xe4, 0x42,
	0x1e, 0xc4, 0xea, 0x33, 0x2c, 0x9e, 0x55, 0x84, 0x42, 0x18, 0xdd, 0xde, 0xdd, 0x5e, 0x2e, 0x5f,
	0xa0, 0x29, 0x0c, 0xb3, 0x9b, 0x9b, 0x65, 0x80, 0x66, 0x30, 0x25, 0x97, 0xf9, 0x4d, 0x76, 0x7e,
	0xb9, 0x1c, 0x6c, 0x26, 0xfe, 0x13, 0xfd, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x96, 0xb7, 0xeb,
	0xa3, 0xaf, 0x03, 0x00, 0x00,
}
