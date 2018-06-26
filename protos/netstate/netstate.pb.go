// Code generated by protoc-gen-go.
// source: netstate.proto
// DO NOT EDIT!

/*
Package netstate is a generated protocol buffer package.

It is generated from these files:
	netstate.proto

It has these top-level messages:
	RedundancyScheme
	EncryptionScheme
	RemotePiece
	RemoteSegment
	Pointer
	PutRequest
	GetRequest
	ListRequest
	PutResponse
	GetResponse
	ListResponse
	DeleteRequest
	DeleteResponse
*/
package netstate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_RS RedundancyScheme_SchemeType = 0
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "RS",
}
var RedundancyScheme_SchemeType_value = map[string]int32{
	"RS": 0,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}
func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type EncryptionScheme_EncryptionType int32

const (
	EncryptionScheme_AESGCM    EncryptionScheme_EncryptionType = 0
	EncryptionScheme_SECRETBOX EncryptionScheme_EncryptionType = 1
)

var EncryptionScheme_EncryptionType_name = map[int32]string{
	0: "AESGCM",
	1: "SECRETBOX",
}
var EncryptionScheme_EncryptionType_value = map[string]int32{
	"AESGCM":    0,
	"SECRETBOX": 1,
}

func (x EncryptionScheme_EncryptionType) String() string {
	return proto.EnumName(EncryptionScheme_EncryptionType_name, int32(x))
}
func (EncryptionScheme_EncryptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}
var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}
func (Pointer_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,enum=netstate.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq           int64 `protobuf:"varint,2,opt,name=min_req,json=minReq" json:"min_req,omitempty"`
	Total            int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	RepairThreshold  int64 `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold" json:"repair_threshold,omitempty"`
	SuccessThreshold int64 `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold" json:"success_threshold,omitempty"`
}

func (m *RedundancyScheme) Reset()                    { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string            { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()               {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_RS
}

func (m *RedundancyScheme) GetMinReq() int64 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int64 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int64 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

type EncryptionScheme struct {
	Type                   EncryptionScheme_EncryptionType `protobuf:"varint,1,opt,name=type,enum=netstate.EncryptionScheme_EncryptionType" json:"type,omitempty"`
	EncryptedEncryptionKey []byte                          `protobuf:"bytes,2,opt,name=encrypted_encryption_key,json=encryptedEncryptionKey,proto3" json:"encrypted_encryption_key,omitempty"`
	EncryptedStartingNonce []byte                          `protobuf:"bytes,3,opt,name=encrypted_starting_nonce,json=encryptedStartingNonce,proto3" json:"encrypted_starting_nonce,omitempty"`
}

func (m *EncryptionScheme) Reset()                    { *m = EncryptionScheme{} }
func (m *EncryptionScheme) String() string            { return proto.CompactTextString(m) }
func (*EncryptionScheme) ProtoMessage()               {}
func (*EncryptionScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EncryptionScheme) GetType() EncryptionScheme_EncryptionType {
	if m != nil {
		return m.Type
	}
	return EncryptionScheme_AESGCM
}

func (m *EncryptionScheme) GetEncryptedEncryptionKey() []byte {
	if m != nil {
		return m.EncryptedEncryptionKey
	}
	return nil
}

func (m *EncryptionScheme) GetEncryptedStartingNonce() []byte {
	if m != nil {
		return m.EncryptedStartingNonce
	}
	return nil
}

type RemotePiece struct {
	PieceNum int64  `protobuf:"varint,1,opt,name=piece_num,json=pieceNum" json:"piece_num,omitempty"`
	NodeId   string `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *RemotePiece) Reset()                    { *m = RemotePiece{} }
func (m *RemotePiece) String() string            { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()               {}
func (*RemotePiece) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemotePiece) GetPieceNum() int64 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

func (m *RemotePiece) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type RemoteSegment struct {
	Redundancy   *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy" json:"redundancy,omitempty"`
	PieceId      string            `protobuf:"bytes,2,opt,name=piece_id,json=pieceId" json:"piece_id,omitempty"`
	RemotePieces []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces" json:"remote_pieces,omitempty"`
	MerkleRoot   []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
}

func (m *RemoteSegment) Reset()                    { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string            { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()               {}
func (*RemoteSegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetPieceId() string {
	if m != nil {
		return m.PieceId
	}
	return ""
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type Pointer struct {
	Type           Pointer_DataType           `protobuf:"varint,1,opt,name=type,enum=netstate.Pointer_DataType" json:"type,omitempty"`
	InlineSegment  []byte                     `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote         *RemoteSegment             `protobuf:"bytes,4,opt,name=remote" json:"remote,omitempty"`
	Size           int64                      `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	CreationDate   *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	ExpirationDate *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`
	Metadata       []byte                     `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Pointer) Reset()                    { *m = Pointer{} }
func (m *Pointer) String() string            { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()               {}
func (*Pointer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Pointer) GetCreationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Pointer) GetExpirationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpirationDate
	}
	return nil
}

func (m *Pointer) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// PutRequest is a request message for the Put rpc call
type PutRequest struct {
	Path    []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer *Pointer `protobuf:"bytes,2,opt,name=pointer" json:"pointer,omitempty"`
	APIKey  []byte   `protobuf:"bytes,3,opt,name=APIKey,json=aPIKey,proto3" json:"APIKey,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PutRequest) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *PutRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,json=aPIKey,proto3" json:"APIKey,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// ListRequest is a request message for the List rpc call
type ListRequest struct {
	StartingPathKey []byte `protobuf:"bytes,1,opt,name=starting_path_key,json=startingPathKey,proto3" json:"starting_path_key,omitempty"`
	Limit           int64  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	APIKey          []byte `protobuf:"bytes,3,opt,name=APIKey,json=aPIKey,proto3" json:"APIKey,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetStartingPathKey() []byte {
	if m != nil {
		return m.StartingPathKey
	}
	return nil
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// PutResponse is a response message for the Put rpc call
type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Pointer []byte `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetResponse) GetPointer() []byte {
	if m != nil {
		return m.Pointer
	}
	return nil
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Paths [][]byte `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListResponse) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

type DeleteRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,json=aPIKey,proto3" json:"APIKey,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *DeleteRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// DeleteResponse is a response message for the Delete rpc call
type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*RedundancyScheme)(nil), "netstate.RedundancyScheme")
	proto.RegisterType((*EncryptionScheme)(nil), "netstate.EncryptionScheme")
	proto.RegisterType((*RemotePiece)(nil), "netstate.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "netstate.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "netstate.Pointer")
	proto.RegisterType((*PutRequest)(nil), "netstate.PutRequest")
	proto.RegisterType((*GetRequest)(nil), "netstate.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "netstate.ListRequest")
	proto.RegisterType((*PutResponse)(nil), "netstate.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "netstate.GetResponse")
	proto.RegisterType((*ListResponse)(nil), "netstate.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "netstate.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "netstate.DeleteResponse")
	proto.RegisterEnum("netstate.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("netstate.EncryptionScheme_EncryptionType", EncryptionScheme_EncryptionType_name, EncryptionScheme_EncryptionType_value)
	proto.RegisterEnum("netstate.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetState service

type NetStateClient interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type netStateClient struct {
	cc *grpc.ClientConn
}

func NewNetStateClient(cc *grpc.ClientConn) NetStateClient {
	return &netStateClient{cc}
}

func (c *netStateClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetState service

type NetStateServer interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterNetStateServer(s *grpc.Server, srv NetStateServer) {
	s.RegisterService(&_NetState_serviceDesc, srv)
}

func _NetState_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetState_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netstate.NetState",
	HandlerType: (*NetStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _NetState_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetState_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetState_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetState_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netstate.proto",
}

func init() { proto.RegisterFile("netstate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x2d, 0x9b, 0x92, 0x87, 0x92, 0x4c, 0x2f, 0x14, 0x87, 0x55, 0x0f, 0x31, 0x88, 0x06,
	0x75, 0x6a, 0x40, 0x06, 0x54, 0x14, 0x48, 0x13, 0x14, 0x45, 0x2a, 0x0b, 0x86, 0x90, 0x44, 0x11,
	0x56, 0x3a, 0xf4, 0x46, 0x6c, 0xc4, 0xa9, 0x44, 0x44, 0xdc, 0xa5, 0xb9, 0x2b, 0xa0, 0xea, 0xeb,
	0xf5, 0x5d, 0x7a, 0x68, 0x6f, 0x7d, 0x82, 0x82, 0xbb, 0xfc, 0x93, 0x02, 0xb7, 0x40, 0x4f, 0xe2,
	0xcc, 0x7e, 0xdf, 0xfc, 0x7c, 0x33, 0x23, 0xe8, 0x72, 0x54, 0x52, 0x31, 0x85, 0x83, 0x24, 0x15,
	0x4a, 0x90, 0x56, 0x61, 0xf7, 0x9f, 0xad, 0x84, 0x58, 0x6d, 0xf0, 0x56, 0xfb, 0x3f, 0x6e, 0x7f,
	0xb9, 0x55, 0x51, 0x8c, 0x52, 0xb1, 0x38, 0x31, 0x50, 0xff, 0x4f, 0x0b, 0x5c, 0x8a, 0xe1, 0x96,
	0x87, 0x8c, 0x2f, 0x77, 0xf3, 0xe5, 0x1a, 0x63, 0x24, 0xdf, 0xc3, 0x89, 0xda, 0x25, 0xe8, 0x59,
	0x57, 0xd6, 0x75, 0x77, 0xf8, 0x7c, 0x50, 0x86, 0x3f, 0x44, 0x0e, 0xcc, 0xcf, 0x62, 0x97, 0x20,
	0xd5, 0x14, 0xf2, 0x14, 0x9a, 0x71, 0xc4, 0x83, 0x14, 0x1f, 0xbc, 0xe3, 0x2b, 0xeb, 0xba, 0x41,
	0xed, 0x38, 0xe2, 0x14, 0x1f, 0x48, 0x0f, 0x4e, 0x95, 0x50, 0x6c, 0xe3, 0x35, 0xb4, 0xdb, 0x18,
	0xe4, 0x05, 0xb8, 0x29, 0x26, 0x2c, 0x4a, 0x03, 0xb5, 0x4e, 0x51, 0xae, 0xc5, 0x26, 0xf4, 0x4e,
	0x34, 0xe0, 0xdc, 0xf8, 0x17, 0x85, 0x9b, 0xdc, 0xc0, 0x85, 0xdc, 0x2e, 0x97, 0x28, 0x65, 0x0d,
	0x7b, 0xaa, 0xb1, 0x6e, 0xfe, 0x50, 0x82, 0xfd, 0x1e, 0x40, 0x55, 0x1a, 0xb1, 0xe1, 0x98, 0xce,
	0xdd, 0x23, 0xff, 0x6f, 0x0b, 0xdc, 0x31, 0x5f, 0xa6, 0xbb, 0x44, 0x45, 0x82, 0xe7, 0xcd, 0xfe,
	0xb0, 0xd7, 0xec, 0x8b, 0xaa, 0xd9, 0x43, 0x64, 0xcd, 0x51, 0x6b, 0xf8, 0x25, 0x78, 0x68, 0xfc,
	0x18, 0x06, 0x58, 0x22, 0x82, 0x4f, 0xb8, 0xd3, 0x0a, 0xb4, 0xe9, 0x65, 0xf9, 0x5e, 0x05, 0x78,
	0x8b, 0xbb, 0x7d, 0xa6, 0x54, 0x2c, 0x55, 0x11, 0x5f, 0x05, 0x5c, 0xf0, 0x25, 0x6a, 0x91, 0xea,
	0xcc, 0x79, 0xfe, 0x3c, 0xcd, 0x5e, 0xfd, 0x1b, 0xe8, 0xee, 0xd7, 0x42, 0x00, 0xec, 0x37, 0xe3,
	0xf9, 0xfd, 0xe8, 0xbd, 0x7b, 0x44, 0x3a, 0x70, 0x36, 0x1f, 0x8f, 0xe8, 0x78, 0xf1, 0xd3, 0x87,
	0x9f, 0x5d, 0xcb, 0x1f, 0x81, 0x43, 0x31, 0x16, 0x0a, 0x67, 0x11, 0x2e, 0x91, 0x7c, 0x09, 0x67,
	0x49, 0xf6, 0x11, 0xf0, 0x6d, 0xac, 0x7b, 0x6e, 0xd0, 0x96, 0x76, 0x4c, 0xb7, 0x71, 0x36, 0x3d,
	0x2e, 0x42, 0x0c, 0xa2, 0x50, 0xd7, 0x7e, 0x46, 0xed, 0xcc, 0x9c, 0x84, 0xfe, 0xef, 0x16, 0x74,
	0x4c, 0x94, 0x39, 0xae, 0x62, 0xe4, 0x8a, 0xbc, 0x02, 0x48, 0xcb, 0x6d, 0xd0, 0x81, 0x9c, 0x61,
	0xff, 0xf1, 0x4d, 0xa1, 0x35, 0x34, 0xf9, 0x02, 0x4c, 0xca, 0x2a, 0x4f, 0x53, 0xdb, 0x93, 0x90,
	0xbc, 0x82, 0x4e, 0xaa, 0xf3, 0x04, 0xda, 0x23, 0xbd, 0xc6, 0x55, 0xe3, 0xda, 0x19, 0x3e, 0xa9,
	0x47, 0x2e, 0x9b, 0xa1, 0xed, 0xb4, 0x32, 0x24, 0x79, 0x06, 0x4e, 0x8c, 0xe9, 0xa7, 0x0d, 0x06,
	0xa9, 0x10, 0x4a, 0xef, 0x51, 0x9b, 0x82, 0x71, 0x51, 0x21, 0x94, 0xff, 0xd7, 0x31, 0x34, 0x67,
	0x22, 0xe2, 0x0a, 0x53, 0x32, 0xd8, 0x1b, 0x7b, 0xad, 0xf2, 0x1c, 0x30, 0xb8, 0x63, 0x8a, 0xd5,
	0xe6, 0xfc, 0x1c, 0xba, 0x11, 0xdf, 0x44, 0x1c, 0x03, 0x69, 0x14, 0xc8, 0x67, 0xd4, 0x31, 0xde,
	0x42, 0x96, 0x5b, 0xb0, 0x4d, 0x4d, 0x3a, 0xbd, 0x33, 0x7c, 0x7a, 0x58, 0x78, 0x0e, 0xa4, 0x39,
	0x8c, 0x10, 0x38, 0x91, 0xd1, 0x6f, 0x98, 0x6f, 0xb2, 0xfe, 0x26, 0x3f, 0x42, 0x67, 0x99, 0x22,
	0xd3, 0x7b, 0x14, 0x32, 0x85, 0x9e, 0x9d, 0xcb, 0x6b, 0xae, 0x79, 0x50, 0x5c, 0xf3, 0x60, 0x51,
	0x5c, 0x33, 0x6d, 0x17, 0x84, 0x3b, 0xa6, 0x90, 0x8c, 0xe0, 0x1c, 0x7f, 0x4d, 0xa2, 0xb4, 0x16,
	0xa2, 0xf9, 0x9f, 0x21, 0xba, 0x15, 0x45, 0x07, 0xe9, 0x43, 0x2b, 0x46, 0xc5, 0x42, 0xa6, 0x98,
	0xd7, 0xd2, 0xbd, 0x96, 0xb6, 0xef, 0x43, 0xab, 0xd0, 0x27, 0xdb, 0xbd, 0xc9, 0xf4, 0xdd, 0x64,
	0x3a, 0x76, 0x8f, 0xb2, 0x6f, 0x3a, 0x7e, 0xff, 0x61, 0x31, 0x76, 0x2d, 0x1f, 0x01, 0x66, 0x5b,
	0x45, 0xf1, 0x61, 0x8b, 0x52, 0x65, 0x7d, 0x26, 0x4c, 0xad, 0xb5, 0xde, 0x6d, 0xaa, 0xbf, 0xc9,
	0x0d, 0x34, 0x13, 0xa3, 0xb6, 0x5e, 0x03, 0x67, 0x78, 0xf1, 0xd9, 0x18, 0x68, 0x81, 0x20, 0x97,
	0x60, 0xbf, 0x99, 0x4d, 0xde, 0xe2, 0x2e, 0x17, 0xde, 0x66, 0xda, 0xf2, 0x5f, 0x02, 0xdc, 0xe3,
	0xbf, 0xa6, 0xa9, 0x98, 0xc7, 0x7b, 0xcc, 0x15, 0x38, 0xef, 0x22, 0x59, 0x52, 0xbf, 0x81, 0x8b,
	0xf2, 0x0a, 0x33, 0x9e, 0x3e, 0x61, 0x13, 0xe7, 0xbc, 0x78, 0x98, 0x31, 0xb5, 0xce, 0x6e, 0xb7,
	0x07, 0xa7, 0x9b, 0x28, 0x8e, 0x54, 0xfe, 0x27, 0x67, 0x8c, 0x47, 0x4b, 0xec, 0x80, 0xa3, 0x95,
	0x90, 0x89, 0xe0, 0x12, 0xfd, 0xaf, 0xc1, 0xd1, 0x15, 0x1b, 0x93, 0x78, 0x95, 0x0a, 0x26, 0x5b,
	0x61, 0xfa, 0x5f, 0x41, 0xdb, 0x14, 0x98, 0x23, 0x7b, 0x70, 0x9a, 0x15, 0x26, 0x3d, 0xeb, 0xaa,
	0x71, 0xdd, 0xa6, 0xc6, 0xf0, 0x5f, 0x43, 0xe7, 0x0e, 0x37, 0xa8, 0xf0, 0xff, 0x68, 0xe0, 0x42,
	0xb7, 0x20, 0x9b, 0x24, 0xc3, 0x3f, 0x2c, 0x68, 0x4d, 0x51, 0xcd, 0xb3, 0x29, 0x90, 0x21, 0x34,
	0x66, 0x5b, 0x45, 0x7a, 0xb5, 0xb9, 0x94, 0x23, 0xed, 0x3f, 0x39, 0xf0, 0xe6, 0x55, 0x0e, 0xa1,
	0x71, 0x8f, 0x7b, 0x9c, 0x6a, 0x3e, 0x75, 0x4e, 0x5d, 0x83, 0xef, 0xe0, 0x24, 0xeb, 0x94, 0xd4,
	0x9e, 0x6b, 0xa3, 0xe9, 0x5f, 0x1e, 0xba, 0x73, 0xda, 0x6b, 0xb0, 0x4d, 0xf5, 0xa4, 0x76, 0x67,
	0x7b, 0x62, 0xf4, 0xbd, 0xcf, 0x1f, 0x0c, 0xf9, 0xa3, 0xad, 0x6f, 0xe0, 0xdb, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0x95, 0x0c, 0x06, 0x3e, 0x07, 0x00, 0x00,
}
