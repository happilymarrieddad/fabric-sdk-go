/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/etcdraft/configuration.proto

package etcdraft

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "etcdraft".
type ConfigMetadata struct {
	Consenters           []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options              *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfigMetadata) Reset()         { *m = ConfigMetadata{} }
func (m *ConfigMetadata) String() string { return proto.CompactTextString(m) }
func (*ConfigMetadata) ProtoMessage()    {}
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f12d215c949b072, []int{0}
}

func (m *ConfigMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigMetadata.Unmarshal(m, b)
}
func (m *ConfigMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigMetadata.Marshal(b, m, deterministic)
}
func (m *ConfigMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigMetadata.Merge(m, src)
}
func (m *ConfigMetadata) XXX_Size() int {
	return xxx_messageInfo_ConfigMetadata.Size(m)
}
func (m *ConfigMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigMetadata proto.InternalMessageInfo

func (m *ConfigMetadata) GetConsenters() []*Consenter {
	if m != nil {
		return m.Consenters
	}
	return nil
}

func (m *ConfigMetadata) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ClientTlsCert        []byte   `protobuf:"bytes,3,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert        []byte   `protobuf:"bytes,4,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f12d215c949b072, []int{1}
}

func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consenter.Unmarshal(m, b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
}
func (m *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(m, src)
}
func (m *Consenter) XXX_Size() int {
	return xxx_messageInfo_Consenter.Size(m)
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func (m *Consenter) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Consenter) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Consenter) GetClientTlsCert() []byte {
	if m != nil {
		return m.ClientTlsCert
	}
	return nil
}

func (m *Consenter) GetServerTlsCert() []byte {
	if m != nil {
		return m.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the etcd/raft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	TickInterval      string `protobuf:"bytes,1,opt,name=tick_interval,json=tickInterval,proto3" json:"tick_interval,omitempty"`
	ElectionTick      uint32 `protobuf:"varint,2,opt,name=election_tick,json=electionTick,proto3" json:"election_tick,omitempty"`
	HeartbeatTick     uint32 `protobuf:"varint,3,opt,name=heartbeat_tick,json=heartbeatTick,proto3" json:"heartbeat_tick,omitempty"`
	MaxInflightBlocks uint32 `protobuf:"varint,4,opt,name=max_inflight_blocks,json=maxInflightBlocks,proto3" json:"max_inflight_blocks,omitempty"`
	// Take snapshot when cumulative data exceeds certain size in bytes.
	SnapshotIntervalSize uint32   `protobuf:"varint,5,opt,name=snapshot_interval_size,json=snapshotIntervalSize,proto3" json:"snapshot_interval_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f12d215c949b072, []int{2}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetTickInterval() string {
	if m != nil {
		return m.TickInterval
	}
	return ""
}

func (m *Options) GetElectionTick() uint32 {
	if m != nil {
		return m.ElectionTick
	}
	return 0
}

func (m *Options) GetHeartbeatTick() uint32 {
	if m != nil {
		return m.HeartbeatTick
	}
	return 0
}

func (m *Options) GetMaxInflightBlocks() uint32 {
	if m != nil {
		return m.MaxInflightBlocks
	}
	return 0
}

func (m *Options) GetSnapshotIntervalSize() uint32 {
	if m != nil {
		return m.SnapshotIntervalSize
	}
	return 0
}

func init() {
	proto.RegisterType((*ConfigMetadata)(nil), "etcdraft.ConfigMetadata")
	proto.RegisterType((*Consenter)(nil), "etcdraft.Consenter")
	proto.RegisterType((*Options)(nil), "etcdraft.Options")
}

func init() {
	proto.RegisterFile("orderer/etcdraft/configuration.proto", fileDescriptor_6f12d215c949b072)
}

var fileDescriptor_6f12d215c949b072 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x6b, 0xdc, 0x30,
	0x10, 0xc5, 0x71, 0x37, 0x6d, 0x1a, 0x65, 0x9d, 0x12, 0xa5, 0x14, 0x1f, 0xcd, 0xf6, 0x0f, 0x86,
	0x82, 0x0c, 0x49, 0xfb, 0x05, 0xb2, 0xa7, 0x1c, 0x4a, 0xc1, 0xcd, 0xa9, 0x17, 0x23, 0xcb, 0x63,
	0x5b, 0xac, 0xd6, 0x32, 0xa3, 0x49, 0x48, 0x73, 0xed, 0x17, 0xed, 0x47, 0x29, 0x96, 0x6c, 0xef,
	0x92, 0xdb, 0xf0, 0xde, 0xef, 0x8d, 0x1e, 0x68, 0xd8, 0x27, 0x8b, 0x35, 0x20, 0x60, 0x0e, 0xa4,
	0x6a, 0x94, 0x0d, 0xe5, 0xca, 0xf6, 0x8d, 0x6e, 0x1f, 0x50, 0x92, 0xb6, 0xbd, 0x18, 0xd0, 0x92,
	0xe5, 0x6f, 0x67, 0x77, 0x83, 0xec, 0x62, 0xeb, 0x81, 0x1f, 0x40, 0xb2, 0x96, 0x24, 0xf9, 0x0d,
	0x63, 0xca, 0xf6, 0x0e, 0x7a, 0x02, 0x74, 0x49, 0x94, 0xae, 0xb2, 0xf3, 0xeb, 0x2b, 0x31, 0x07,
	0xc4, 0x76, 0xf6, 0x8a, 0x23, 0x8c, 0x7f, 0x65, 0xa7, 0x76, 0x18, 0x1f, 0x70, 0xc9, 0xab, 0x34,
	0xca, 0xce, 0xaf, 0x2f, 0x0f, 0x89, 0x9f, 0xc1, 0x28, 0x66, 0x62, 0xf3, 0x37, 0x62, 0x67, 0xcb,
	0x1a, 0xce, 0xd9, 0x49, 0x67, 0x1d, 0x25, 0x51, 0x1a, 0x65, 0x67, 0x85, 0x9f, 0x47, 0x6d, 0xb0,
	0x48, 0x7e, 0x57, 0x5c, 0xf8, 0x99, 0x7f, 0x61, 0xef, 0x94, 0xd1, 0xd0, 0x53, 0x49, 0xc6, 0x95,
	0x0a, 0x90, 0x92, 0x55, 0x1a, 0x65, 0xeb, 0x22, 0x0e, 0xf2, 0xbd, 0x71, 0x5b, 0x08, 0x9c, 0x03,
	0x7c, 0x04, 0x3c, 0x70, 0x27, 0x81, 0x0b, 0xf2, 0xc4, 0x6d, 0xfe, 0x45, 0xec, 0x74, 0xaa, 0xc6,
	0x3f, 0xb2, 0x98, 0xb4, 0xda, 0x95, 0x7a, 0x6c, 0xf4, 0x28, 0xcd, 0x54, 0x66, 0x3d, 0x8a, 0x77,
	0x93, 0x36, 0x42, 0x60, 0x40, 0x8d, 0x89, 0x72, 0x34, 0xa6, 0x76, 0xeb, 0x59, 0xbc, 0xd7, 0x6a,
	0xc7, 0x3f, 0xb3, 0x8b, 0x0e, 0x24, 0x52, 0x05, 0x92, 0x02, 0xb5, 0xf2, 0x54, 0xbc, 0xa8, 0x1e,
	0x13, 0xec, 0x6a, 0x2f, 0x9f, 0x4a, 0xdd, 0x37, 0x46, 0xb7, 0x1d, 0x95, 0x95, 0xb1, 0x6a, 0xe7,
	0x7c, 0xd1, 0xb8, 0xb8, 0xdc, 0xcb, 0xa7, 0xbb, 0xc9, 0xb9, 0xf5, 0x06, 0xff, 0xc6, 0x3e, 0xb8,
	0x5e, 0x0e, 0xae, 0xb3, 0xb4, 0x94, 0x2c, 0x9d, 0x7e, 0x86, 0xe4, 0xb5, 0x8f, 0xbc, 0x9f, 0xdd,
	0xb9, 0xed, 0x2f, 0xfd, 0x0c, 0xb7, 0x2d, 0x13, 0x16, 0x5b, 0xd1, 0xfd, 0x19, 0x00, 0x0d, 0xd4,
	0x2d, 0xa0, 0x68, 0x64, 0x85, 0x5a, 0x85, 0x33, 0x70, 0x62, 0x3a, 0x96, 0xe5, 0xaf, 0x7e, 0x7f,
	0x6f, 0x35, 0x75, 0x0f, 0x95, 0x50, 0x76, 0x9f, 0x1f, 0xc5, 0xf2, 0x10, 0xcb, 0x43, 0x2c, 0x7f,
	0x79, 0x63, 0xd5, 0x1b, 0x6f, 0xdc, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x85, 0x77, 0x31,
	0x7e, 0x02, 0x00, 0x00,
}