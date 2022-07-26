// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Crypto struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Upvote               int64    `protobuf:"varint,3,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote             int64    `protobuf:"varint,4,opt,name=downvote,proto3" json:"downvote,omitempty"`
	Totalscore           int64    `protobuf:"varint,5,opt,name=totalscore,proto3" json:"totalscore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crypto) Reset()         { *m = Crypto{} }
func (m *Crypto) String() string { return proto.CompactTextString(m) }
func (*Crypto) ProtoMessage()    {}
func (*Crypto) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{0}
}

func (m *Crypto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crypto.Unmarshal(m, b)
}
func (m *Crypto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crypto.Marshal(b, m, deterministic)
}
func (m *Crypto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crypto.Merge(m, src)
}
func (m *Crypto) XXX_Size() int {
	return xxx_messageInfo_Crypto.Size(m)
}
func (m *Crypto) XXX_DiscardUnknown() {
	xxx_messageInfo_Crypto.DiscardUnknown(m)
}

var xxx_messageInfo_Crypto proto.InternalMessageInfo

func (m *Crypto) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Crypto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Crypto) GetUpvote() int64 {
	if m != nil {
		return m.Upvote
	}
	return 0
}

func (m *Crypto) GetDownvote() int64 {
	if m != nil {
		return m.Downvote
	}
	return 0
}

func (m *Crypto) GetTotalscore() int64 {
	if m != nil {
		return m.Totalscore
	}
	return 0
}

type CreateCryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCryptoResponse) Reset()         { *m = CreateCryptoResponse{} }
func (m *CreateCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCryptoResponse) ProtoMessage()    {}
func (*CreateCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{1}
}

func (m *CreateCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCryptoResponse.Unmarshal(m, b)
}
func (m *CreateCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCryptoResponse.Marshal(b, m, deterministic)
}
func (m *CreateCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCryptoResponse.Merge(m, src)
}
func (m *CreateCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCryptoResponse.Size(m)
}
func (m *CreateCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCryptoResponse proto.InternalMessageInfo

func (m *CreateCryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

type GetByIdCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIdCryptoRequest) Reset()         { *m = GetByIdCryptoRequest{} }
func (m *GetByIdCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIdCryptoRequest) ProtoMessage()    {}
func (*GetByIdCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{2}
}

func (m *GetByIdCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIdCryptoRequest.Unmarshal(m, b)
}
func (m *GetByIdCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIdCryptoRequest.Marshal(b, m, deterministic)
}
func (m *GetByIdCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIdCryptoRequest.Merge(m, src)
}
func (m *GetByIdCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIdCryptoRequest.Size(m)
}
func (m *GetByIdCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIdCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIdCryptoRequest proto.InternalMessageInfo

func (m *GetByIdCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetByIdCryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIdCryptoResponse) Reset()         { *m = GetByIdCryptoResponse{} }
func (m *GetByIdCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*GetByIdCryptoResponse) ProtoMessage()    {}
func (*GetByIdCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{3}
}

func (m *GetByIdCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIdCryptoResponse.Unmarshal(m, b)
}
func (m *GetByIdCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIdCryptoResponse.Marshal(b, m, deterministic)
}
func (m *GetByIdCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIdCryptoResponse.Merge(m, src)
}
func (m *GetByIdCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_GetByIdCryptoResponse.Size(m)
}
func (m *GetByIdCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIdCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIdCryptoResponse proto.InternalMessageInfo

func (m *GetByIdCryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

type GetAllCryptoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCryptoRequest) Reset()         { *m = GetAllCryptoRequest{} }
func (m *GetAllCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllCryptoRequest) ProtoMessage()    {}
func (*GetAllCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{4}
}

func (m *GetAllCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCryptoRequest.Unmarshal(m, b)
}
func (m *GetAllCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCryptoRequest.Marshal(b, m, deterministic)
}
func (m *GetAllCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCryptoRequest.Merge(m, src)
}
func (m *GetAllCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllCryptoRequest.Size(m)
}
func (m *GetAllCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCryptoRequest proto.InternalMessageInfo

type GetAllCryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllCryptoResponse) Reset()         { *m = GetAllCryptoResponse{} }
func (m *GetAllCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllCryptoResponse) ProtoMessage()    {}
func (*GetAllCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{5}
}

func (m *GetAllCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllCryptoResponse.Unmarshal(m, b)
}
func (m *GetAllCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllCryptoResponse.Marshal(b, m, deterministic)
}
func (m *GetAllCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllCryptoResponse.Merge(m, src)
}
func (m *GetAllCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllCryptoResponse.Size(m)
}
func (m *GetAllCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllCryptoResponse proto.InternalMessageInfo

func (m *GetAllCryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

type UpVoteCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpVoteCryptoRequest) Reset()         { *m = UpVoteCryptoRequest{} }
func (m *UpVoteCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*UpVoteCryptoRequest) ProtoMessage()    {}
func (*UpVoteCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{6}
}

func (m *UpVoteCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpVoteCryptoRequest.Unmarshal(m, b)
}
func (m *UpVoteCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpVoteCryptoRequest.Marshal(b, m, deterministic)
}
func (m *UpVoteCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpVoteCryptoRequest.Merge(m, src)
}
func (m *UpVoteCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_UpVoteCryptoRequest.Size(m)
}
func (m *UpVoteCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpVoteCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpVoteCryptoRequest proto.InternalMessageInfo

func (m *UpVoteCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpVoteCryptoResponse struct {
	Upvote               bool     `protobuf:"varint,1,opt,name=upvote,proto3" json:"upvote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpVoteCryptoResponse) Reset()         { *m = UpVoteCryptoResponse{} }
func (m *UpVoteCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*UpVoteCryptoResponse) ProtoMessage()    {}
func (*UpVoteCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{7}
}

func (m *UpVoteCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpVoteCryptoResponse.Unmarshal(m, b)
}
func (m *UpVoteCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpVoteCryptoResponse.Marshal(b, m, deterministic)
}
func (m *UpVoteCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpVoteCryptoResponse.Merge(m, src)
}
func (m *UpVoteCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_UpVoteCryptoResponse.Size(m)
}
func (m *UpVoteCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpVoteCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpVoteCryptoResponse proto.InternalMessageInfo

func (m *UpVoteCryptoResponse) GetUpvote() bool {
	if m != nil {
		return m.Upvote
	}
	return false
}

type DownVoteCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownVoteCryptoRequest) Reset()         { *m = DownVoteCryptoRequest{} }
func (m *DownVoteCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*DownVoteCryptoRequest) ProtoMessage()    {}
func (*DownVoteCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{8}
}

func (m *DownVoteCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownVoteCryptoRequest.Unmarshal(m, b)
}
func (m *DownVoteCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownVoteCryptoRequest.Marshal(b, m, deterministic)
}
func (m *DownVoteCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownVoteCryptoRequest.Merge(m, src)
}
func (m *DownVoteCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_DownVoteCryptoRequest.Size(m)
}
func (m *DownVoteCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownVoteCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownVoteCryptoRequest proto.InternalMessageInfo

func (m *DownVoteCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DownVoteCryptoResponse struct {
	Downvote             bool     `protobuf:"varint,1,opt,name=downvote,proto3" json:"downvote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownVoteCryptoResponse) Reset()         { *m = DownVoteCryptoResponse{} }
func (m *DownVoteCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*DownVoteCryptoResponse) ProtoMessage()    {}
func (*DownVoteCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{9}
}

func (m *DownVoteCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownVoteCryptoResponse.Unmarshal(m, b)
}
func (m *DownVoteCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownVoteCryptoResponse.Marshal(b, m, deterministic)
}
func (m *DownVoteCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownVoteCryptoResponse.Merge(m, src)
}
func (m *DownVoteCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_DownVoteCryptoResponse.Size(m)
}
func (m *DownVoteCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownVoteCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownVoteCryptoResponse proto.InternalMessageInfo

func (m *DownVoteCryptoResponse) GetDownvote() bool {
	if m != nil {
		return m.Downvote
	}
	return false
}

type DeleteCryptoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCryptoRequest) Reset()         { *m = DeleteCryptoRequest{} }
func (m *DeleteCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCryptoRequest) ProtoMessage()    {}
func (*DeleteCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{10}
}

func (m *DeleteCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCryptoRequest.Unmarshal(m, b)
}
func (m *DeleteCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCryptoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCryptoRequest.Merge(m, src)
}
func (m *DeleteCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCryptoRequest.Size(m)
}
func (m *DeleteCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCryptoRequest proto.InternalMessageInfo

func (m *DeleteCryptoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteCryptoResponse struct {
	Delete               bool     `protobuf:"varint,1,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCryptoResponse) Reset()         { *m = DeleteCryptoResponse{} }
func (m *DeleteCryptoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCryptoResponse) ProtoMessage()    {}
func (*DeleteCryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_527278fb02d03321, []int{11}
}

func (m *DeleteCryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCryptoResponse.Unmarshal(m, b)
}
func (m *DeleteCryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCryptoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCryptoResponse.Merge(m, src)
}
func (m *DeleteCryptoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCryptoResponse.Size(m)
}
func (m *DeleteCryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCryptoResponse proto.InternalMessageInfo

func (m *DeleteCryptoResponse) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func init() {
	proto.RegisterType((*Crypto)(nil), "crypto.Crypto")
	proto.RegisterType((*CreateCryptoResponse)(nil), "crypto.CreateCryptoResponse")
	proto.RegisterType((*GetByIdCryptoRequest)(nil), "crypto.GetByIdCryptoRequest")
	proto.RegisterType((*GetByIdCryptoResponse)(nil), "crypto.GetByIdCryptoResponse")
	proto.RegisterType((*GetAllCryptoRequest)(nil), "crypto.GetAllCryptoRequest")
	proto.RegisterType((*GetAllCryptoResponse)(nil), "crypto.GetAllCryptoResponse")
	proto.RegisterType((*UpVoteCryptoRequest)(nil), "crypto.UpVoteCryptoRequest")
	proto.RegisterType((*UpVoteCryptoResponse)(nil), "crypto.UpVoteCryptoResponse")
	proto.RegisterType((*DownVoteCryptoRequest)(nil), "crypto.DownVoteCryptoRequest")
	proto.RegisterType((*DownVoteCryptoResponse)(nil), "crypto.DownVoteCryptoResponse")
	proto.RegisterType((*DeleteCryptoRequest)(nil), "crypto.DeleteCryptoRequest")
	proto.RegisterType((*DeleteCryptoResponse)(nil), "crypto.DeleteCryptoResponse")
}

func init() { proto.RegisterFile("crypto.proto", fileDescriptor_527278fb02d03321) }

var fileDescriptor_527278fb02d03321 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x35, 0xd1, 0x17, 0x7c, 0xd7, 0x3c, 0x1f, 0x5c, 0x3f, 0x08, 0x79, 0x3e, 0x91, 0x40, 0xad,
	0x2b, 0x29, 0xb6, 0xeb, 0x16, 0x3f, 0x40, 0x8a, 0x54, 0x5a, 0xa1, 0x5d, 0x74, 0xa7, 0x66, 0x16,
	0x42, 0x9a, 0x49, 0x93, 0xb1, 0xe2, 0xae, 0x3f, 0xb4, 0x3f, 0xa6, 0x38, 0xf9, 0x98, 0x24, 0x1d,
	0x69, 0xdd, 0xcd, 0xbd, 0xf7, 0xcc, 0x99, 0x73, 0x73, 0x0e, 0x01, 0x7d, 0xed, 0xef, 0x3d, 0x46,
	0xfb, 0x9e, 0x4f, 0x19, 0x45, 0x2d, 0xac, 0xac, 0x77, 0x05, 0xb4, 0x31, 0x3f, 0x62, 0x15, 0xd4,
	0x8d, 0x6d, 0x28, 0x1d, 0xa5, 0xf7, 0x7b, 0xa1, 0x6e, 0x6c, 0x44, 0x28, 0xb9, 0xcb, 0x17, 0x62,
	0xa8, 0xbc, 0xc3, 0xcf, 0xd8, 0x04, 0x6d, 0xeb, 0xbd, 0x51, 0x46, 0x8c, 0x62, 0x47, 0xe9, 0x15,
	0x17, 0x51, 0x85, 0x26, 0x94, 0x6d, 0xba, 0x73, 0xf9, 0xa4, 0xc4, 0x27, 0x49, 0x8d, 0x6d, 0x00,
	0x46, 0xd9, 0xd2, 0x09, 0xd6, 0xd4, 0x27, 0xc6, 0x2f, 0x3e, 0x4d, 0x75, 0xac, 0x6b, 0xa8, 0x8f,
	0x7d, 0xb2, 0x64, 0x24, 0xd4, 0xb1, 0x20, 0x81, 0x47, 0xdd, 0x80, 0x60, 0x17, 0x22, 0x91, 0x5c,
	0x53, 0x65, 0x50, 0xed, 0x47, 0x1b, 0x44, 0xb8, 0x78, 0x85, 0x2e, 0xd4, 0xa7, 0x84, 0x8d, 0xf6,
	0xb7, 0x76, 0x4c, 0xf0, 0xba, 0x25, 0x01, 0xcb, 0xef, 0x63, 0xdd, 0x40, 0x23, 0x87, 0x3b, 0xf1,
	0xa1, 0x06, 0xd4, 0xa6, 0x84, 0x0d, 0x1d, 0x27, 0xf3, 0xce, 0x41, 0x7f, 0xb6, 0x7d, 0x22, 0xed,
	0x19, 0xd4, 0x1e, 0xbd, 0x27, 0x2a, 0xf6, 0x97, 0xcb, 0xef, 0x43, 0x3d, 0x0b, 0x8b, 0x9e, 0x11,
	0x96, 0x1c, 0xb0, 0xe5, 0xd8, 0x12, 0xeb, 0x1c, 0x1a, 0x13, 0xba, 0x73, 0xbf, 0x27, 0xbe, 0x82,
	0x66, 0x1e, 0x18, 0x51, 0xa7, 0x5d, 0x0d, 0xc9, 0x93, 0xfa, 0xa0, 0x7a, 0x42, 0x1c, 0xf2, 0x03,
	0xd5, 0x59, 0x98, 0x50, 0x6d, 0xf3, 0x7e, 0xac, 0x3a, 0xac, 0x06, 0x1f, 0x45, 0xa8, 0x84, 0xd0,
	0x7b, 0x9e, 0xd3, 0x21, 0xfc, 0x0d, 0xc3, 0x31, 0x27, 0xbb, 0x38, 0xa7, 0xd9, 0xef, 0x68, 0xb6,
	0x44, 0xfd, 0x35, 0x45, 0x56, 0x01, 0xe7, 0xf0, 0x27, 0xe3, 0x3b, 0x26, 0x17, 0x64, 0xb1, 0x31,
	0xff, 0x1f, 0x99, 0x26, 0x7c, 0x77, 0xa0, 0xa7, 0xfd, 0xc6, 0x7f, 0xa9, 0x0b, 0xf9, 0x70, 0x98,
	0x2d, 0xf9, 0x30, 0x26, 0xbb, 0x50, 0x70, 0x06, 0x7a, 0xda, 0x57, 0x41, 0x27, 0x09, 0x85, 0xa0,
	0x93, 0x45, 0xc1, 0x2a, 0xe0, 0x03, 0x54, 0xb3, 0x5e, 0x62, 0xb2, 0x8e, 0x34, 0x0c, 0x66, 0xfb,
	0xd8, 0x38, 0xa1, 0x9c, 0x81, 0x9e, 0x76, 0x50, 0xe8, 0x93, 0xd8, 0x2f, 0xf4, 0xc9, 0x4c, 0xb7,
	0x0a, 0xa3, 0xd2, 0xb3, 0xea, 0xad, 0x56, 0x1a, 0xff, 0x07, 0x5d, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0xd4, 0x79, 0x60, 0x93, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CryptoProtoClient is the client API for CryptoProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoProtoClient interface {
	CreateNewCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CreateCryptoResponse, error)
	GetByIdCrypto(ctx context.Context, in *GetByIdCryptoRequest, opts ...grpc.CallOption) (*GetByIdCryptoResponse, error)
	GetAllCrypto(ctx context.Context, in *GetAllCryptoRequest, opts ...grpc.CallOption) (CryptoProto_GetAllCryptoClient, error)
	UpVoteCrypto(ctx context.Context, in *UpVoteCryptoRequest, opts ...grpc.CallOption) (*UpVoteCryptoResponse, error)
	DownVoteCrypto(ctx context.Context, in *DownVoteCryptoRequest, opts ...grpc.CallOption) (*DownVoteCryptoResponse, error)
	DeleteCrypto(ctx context.Context, in *DeleteCryptoRequest, opts ...grpc.CallOption) (*DeleteCryptoResponse, error)
}

type cryptoProtoClient struct {
	cc *grpc.ClientConn
}

func NewCryptoProtoClient(cc *grpc.ClientConn) CryptoProtoClient {
	return &cryptoProtoClient{cc}
}

func (c *cryptoProtoClient) CreateNewCrypto(ctx context.Context, in *Crypto, opts ...grpc.CallOption) (*CreateCryptoResponse, error) {
	out := new(CreateCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoProto/CreateNewCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoProtoClient) GetByIdCrypto(ctx context.Context, in *GetByIdCryptoRequest, opts ...grpc.CallOption) (*GetByIdCryptoResponse, error) {
	out := new(GetByIdCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoProto/GetByIdCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoProtoClient) GetAllCrypto(ctx context.Context, in *GetAllCryptoRequest, opts ...grpc.CallOption) (CryptoProto_GetAllCryptoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CryptoProto_serviceDesc.Streams[0], "/crypto.CryptoProto/GetAllCrypto", opts...)
	if err != nil {
		return nil, err
	}
	x := &cryptoProtoGetAllCryptoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CryptoProto_GetAllCryptoClient interface {
	Recv() (*GetAllCryptoResponse, error)
	grpc.ClientStream
}

type cryptoProtoGetAllCryptoClient struct {
	grpc.ClientStream
}

func (x *cryptoProtoGetAllCryptoClient) Recv() (*GetAllCryptoResponse, error) {
	m := new(GetAllCryptoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cryptoProtoClient) UpVoteCrypto(ctx context.Context, in *UpVoteCryptoRequest, opts ...grpc.CallOption) (*UpVoteCryptoResponse, error) {
	out := new(UpVoteCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoProto/UpVoteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoProtoClient) DownVoteCrypto(ctx context.Context, in *DownVoteCryptoRequest, opts ...grpc.CallOption) (*DownVoteCryptoResponse, error) {
	out := new(DownVoteCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoProto/DownVoteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoProtoClient) DeleteCrypto(ctx context.Context, in *DeleteCryptoRequest, opts ...grpc.CallOption) (*DeleteCryptoResponse, error) {
	out := new(DeleteCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoProto/DeleteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoProtoServer is the server API for CryptoProto service.
type CryptoProtoServer interface {
	CreateNewCrypto(context.Context, *Crypto) (*CreateCryptoResponse, error)
	GetByIdCrypto(context.Context, *GetByIdCryptoRequest) (*GetByIdCryptoResponse, error)
	GetAllCrypto(*GetAllCryptoRequest, CryptoProto_GetAllCryptoServer) error
	UpVoteCrypto(context.Context, *UpVoteCryptoRequest) (*UpVoteCryptoResponse, error)
	DownVoteCrypto(context.Context, *DownVoteCryptoRequest) (*DownVoteCryptoResponse, error)
	DeleteCrypto(context.Context, *DeleteCryptoRequest) (*DeleteCryptoResponse, error)
}

// UnimplementedCryptoProtoServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoProtoServer struct {
}

func (*UnimplementedCryptoProtoServer) CreateNewCrypto(ctx context.Context, req *Crypto) (*CreateCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCrypto not implemented")
}
func (*UnimplementedCryptoProtoServer) GetByIdCrypto(ctx context.Context, req *GetByIdCryptoRequest) (*GetByIdCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdCrypto not implemented")
}
func (*UnimplementedCryptoProtoServer) GetAllCrypto(req *GetAllCryptoRequest, srv CryptoProto_GetAllCryptoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllCrypto not implemented")
}
func (*UnimplementedCryptoProtoServer) UpVoteCrypto(ctx context.Context, req *UpVoteCryptoRequest) (*UpVoteCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpVoteCrypto not implemented")
}
func (*UnimplementedCryptoProtoServer) DownVoteCrypto(ctx context.Context, req *DownVoteCryptoRequest) (*DownVoteCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownVoteCrypto not implemented")
}
func (*UnimplementedCryptoProtoServer) DeleteCrypto(ctx context.Context, req *DeleteCryptoRequest) (*DeleteCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCrypto not implemented")
}

func RegisterCryptoProtoServer(s *grpc.Server, srv CryptoProtoServer) {
	s.RegisterService(&_CryptoProto_serviceDesc, srv)
}

func _CryptoProto_CreateNewCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crypto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoProtoServer).CreateNewCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoProto/CreateNewCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoProtoServer).CreateNewCrypto(ctx, req.(*Crypto))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoProto_GetByIdCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoProtoServer).GetByIdCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoProto/GetByIdCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoProtoServer).GetByIdCrypto(ctx, req.(*GetByIdCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoProto_GetAllCrypto_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllCryptoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CryptoProtoServer).GetAllCrypto(m, &cryptoProtoGetAllCryptoServer{stream})
}

type CryptoProto_GetAllCryptoServer interface {
	Send(*GetAllCryptoResponse) error
	grpc.ServerStream
}

type cryptoProtoGetAllCryptoServer struct {
	grpc.ServerStream
}

func (x *cryptoProtoGetAllCryptoServer) Send(m *GetAllCryptoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CryptoProto_UpVoteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpVoteCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoProtoServer).UpVoteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoProto/UpVoteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoProtoServer).UpVoteCrypto(ctx, req.(*UpVoteCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoProto_DownVoteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownVoteCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoProtoServer).DownVoteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoProto/DownVoteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoProtoServer).DownVoteCrypto(ctx, req.(*DownVoteCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoProto_DeleteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoProtoServer).DeleteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoProto/DeleteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoProtoServer).DeleteCrypto(ctx, req.(*DeleteCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoProto",
	HandlerType: (*CryptoProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewCrypto",
			Handler:    _CryptoProto_CreateNewCrypto_Handler,
		},
		{
			MethodName: "GetByIdCrypto",
			Handler:    _CryptoProto_GetByIdCrypto_Handler,
		},
		{
			MethodName: "UpVoteCrypto",
			Handler:    _CryptoProto_UpVoteCrypto_Handler,
		},
		{
			MethodName: "DownVoteCrypto",
			Handler:    _CryptoProto_DownVoteCrypto_Handler,
		},
		{
			MethodName: "DeleteCrypto",
			Handler:    _CryptoProto_DeleteCrypto_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllCrypto",
			Handler:       _CryptoProto_GetAllCrypto_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crypto.proto",
}
