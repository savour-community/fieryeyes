// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: scrapy.proto

package sav_scrapy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReturnCode int32

const (
	ReturnCode_SUCCESS ReturnCode = 0
	ReturnCode_ERROR   ReturnCode = 1
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	ReturnCode_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_scrapy_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_scrapy_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{0}
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Chain string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{0}
}

func (x *Chain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chain) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

type GiantWhale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Holder  string `protobuf:"bytes,3,opt,name=holder,proto3" json:"holder,omitempty"`
}

func (x *GiantWhale) Reset() {
	*x = GiantWhale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiantWhale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiantWhale) ProtoMessage() {}

func (x *GiantWhale) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiantWhale.ProtoReflect.Descriptor instead.
func (*GiantWhale) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{1}
}

func (x *GiantWhale) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GiantWhale) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GiantWhale) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

type SupportChainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *SupportChainReq) Reset() {
	*x = SupportChainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportChainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainReq) ProtoMessage() {}

func (x *SupportChainReq) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainReq.ProtoReflect.Descriptor instead.
func (*SupportChainReq) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{2}
}

func (x *SupportChainReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type SupportChainRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.keylocker.ReturnCode" json:"code,omitempty"`
	Msg    string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Chains []*Chain   `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (x *SupportChainRep) Reset() {
	*x = SupportChainRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportChainRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainRep) ProtoMessage() {}

func (x *SupportChainRep) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainRep.ProtoReflect.Descriptor instead.
func (*SupportChainRep) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{3}
}

func (x *SupportChainRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *SupportChainRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupportChainRep) GetChains() []*Chain {
	if x != nil {
		return x.Chains
	}
	return nil
}

type SetGiantWhaleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Address       string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Holder        string `protobuf:"bytes,4,opt,name=holder,proto3" json:"holder,omitempty"`
}

func (x *SetGiantWhaleReq) Reset() {
	*x = SetGiantWhaleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGiantWhaleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGiantWhaleReq) ProtoMessage() {}

func (x *SetGiantWhaleReq) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGiantWhaleReq.ProtoReflect.Descriptor instead.
func (*SetGiantWhaleReq) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{4}
}

func (x *SetGiantWhaleReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SetGiantWhaleReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SetGiantWhaleReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SetGiantWhaleReq) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

type SetGiantWhaleRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.keylocker.ReturnCode" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SetGiantWhaleRep) Reset() {
	*x = SetGiantWhaleRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGiantWhaleRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGiantWhaleRep) ProtoMessage() {}

func (x *SetGiantWhaleRep) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGiantWhaleRep.ProtoReflect.Descriptor instead.
func (*SetGiantWhaleRep) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{5}
}

func (x *SetGiantWhaleRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *SetGiantWhaleRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetGiantWhaleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Page          uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      uint32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetGiantWhaleReq) Reset() {
	*x = GetGiantWhaleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGiantWhaleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGiantWhaleReq) ProtoMessage() {}

func (x *GetGiantWhaleReq) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGiantWhaleReq.ProtoReflect.Descriptor instead.
func (*GetGiantWhaleReq) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{6}
}

func (x *GetGiantWhaleReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *GetGiantWhaleReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetGiantWhaleReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetGiantWhaleReq) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetGiantWhaleRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   ReturnCode    `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.keylocker.ReturnCode" json:"code,omitempty"`
	Msg    string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	GwList []*GiantWhale `protobuf:"bytes,3,rep,name=gw_list,json=gwList,proto3" json:"gw_list,omitempty"`
}

func (x *GetGiantWhaleRep) Reset() {
	*x = GetGiantWhaleRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGiantWhaleRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGiantWhaleRep) ProtoMessage() {}

func (x *GetGiantWhaleRep) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGiantWhaleRep.ProtoReflect.Descriptor instead.
func (*GetGiantWhaleRep) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{7}
}

func (x *GetGiantWhaleRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *GetGiantWhaleRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetGiantWhaleRep) GetGwList() []*GiantWhale {
	if x != nil {
		return x.GwList
	}
	return nil
}

type RemoveGiantWhaleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Address       uint32 `protobuf:"varint,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RemoveGiantWhaleReq) Reset() {
	*x = RemoveGiantWhaleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGiantWhaleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGiantWhaleReq) ProtoMessage() {}

func (x *RemoveGiantWhaleReq) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGiantWhaleReq.ProtoReflect.Descriptor instead.
func (*RemoveGiantWhaleReq) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveGiantWhaleReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *RemoveGiantWhaleReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *RemoveGiantWhaleReq) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type RemoveGiantWhaleRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=savourrpc.keylocker.ReturnCode" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RemoveGiantWhaleRep) Reset() {
	*x = RemoveGiantWhaleRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scrapy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGiantWhaleRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGiantWhaleRep) ProtoMessage() {}

func (x *RemoveGiantWhaleRep) ProtoReflect() protoreflect.Message {
	mi := &file_scrapy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGiantWhaleRep.ProtoReflect.Descriptor instead.
func (*RemoveGiantWhaleRep) Descriptor() ([]byte, []int) {
	return file_scrapy_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveGiantWhaleRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *RemoveGiantWhaleRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_scrapy_proto protoreflect.FileDescriptor

var file_scrapy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x22, 0x4e, 0x0a, 0x0a, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x22, 0x38, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8c, 0x01, 0x0a,
	0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x52, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x10,
	0x53, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22,
	0x59, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65,
	0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x93, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x07, 0x67, 0x77, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x06, 0x67, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x69, 0x61,
	0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x5c, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x69, 0x61, 0x6e, 0x74,
	0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72,
	0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a,
	0x24, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0xa0, 0x03, 0x0a, 0x11, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57,
	0x68, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0f, 0x67,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x24,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63,
	0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0d,
	0x73, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x12, 0x25, 0x2e,
	0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63,
	0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x69,
	0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5f, 0x0a,
	0x0d, 0x67, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x12, 0x25,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70,
	0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61,
	0x6c, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b,
	0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47,
	0x69, 0x61, 0x6e, 0x74, 0x57, 0x68, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x69, 0x61, 0x6e, 0x74, 0x57, 0x68,
	0x61, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x2e, 0x2e, 0x2f, 0x73,
	0x61, 0x76, 0x2d, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x61, 0x76, 0x2d, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_scrapy_proto_rawDescOnce sync.Once
	file_scrapy_proto_rawDescData = file_scrapy_proto_rawDesc
)

func file_scrapy_proto_rawDescGZIP() []byte {
	file_scrapy_proto_rawDescOnce.Do(func() {
		file_scrapy_proto_rawDescData = protoimpl.X.CompressGZIP(file_scrapy_proto_rawDescData)
	})
	return file_scrapy_proto_rawDescData
}

var file_scrapy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scrapy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_scrapy_proto_goTypes = []interface{}{
	(ReturnCode)(0),             // 0: savourrpc.keylocker.ReturnCode
	(*Chain)(nil),               // 1: savourrpc.keylocker.Chain
	(*GiantWhale)(nil),          // 2: savourrpc.keylocker.GiantWhale
	(*SupportChainReq)(nil),     // 3: savourrpc.keylocker.SupportChainReq
	(*SupportChainRep)(nil),     // 4: savourrpc.keylocker.SupportChainRep
	(*SetGiantWhaleReq)(nil),    // 5: savourrpc.keylocker.SetGiantWhaleReq
	(*SetGiantWhaleRep)(nil),    // 6: savourrpc.keylocker.SetGiantWhaleRep
	(*GetGiantWhaleReq)(nil),    // 7: savourrpc.keylocker.GetGiantWhaleReq
	(*GetGiantWhaleRep)(nil),    // 8: savourrpc.keylocker.GetGiantWhaleRep
	(*RemoveGiantWhaleReq)(nil), // 9: savourrpc.keylocker.RemoveGiantWhaleReq
	(*RemoveGiantWhaleRep)(nil), // 10: savourrpc.keylocker.RemoveGiantWhaleRep
}
var file_scrapy_proto_depIdxs = []int32{
	0,  // 0: savourrpc.keylocker.SupportChainRep.code:type_name -> savourrpc.keylocker.ReturnCode
	1,  // 1: savourrpc.keylocker.SupportChainRep.chains:type_name -> savourrpc.keylocker.Chain
	0,  // 2: savourrpc.keylocker.SetGiantWhaleRep.code:type_name -> savourrpc.keylocker.ReturnCode
	0,  // 3: savourrpc.keylocker.GetGiantWhaleRep.code:type_name -> savourrpc.keylocker.ReturnCode
	2,  // 4: savourrpc.keylocker.GetGiantWhaleRep.gw_list:type_name -> savourrpc.keylocker.GiantWhale
	0,  // 5: savourrpc.keylocker.RemoveGiantWhaleRep.code:type_name -> savourrpc.keylocker.ReturnCode
	3,  // 6: savourrpc.keylocker.GiantWhaleService.getSupportChain:input_type -> savourrpc.keylocker.SupportChainReq
	5,  // 7: savourrpc.keylocker.GiantWhaleService.setGiantWhale:input_type -> savourrpc.keylocker.SetGiantWhaleReq
	7,  // 8: savourrpc.keylocker.GiantWhaleService.getGiantWhale:input_type -> savourrpc.keylocker.GetGiantWhaleReq
	9,  // 9: savourrpc.keylocker.GiantWhaleService.removeGiantWhale:input_type -> savourrpc.keylocker.RemoveGiantWhaleReq
	4,  // 10: savourrpc.keylocker.GiantWhaleService.getSupportChain:output_type -> savourrpc.keylocker.SupportChainRep
	6,  // 11: savourrpc.keylocker.GiantWhaleService.setGiantWhale:output_type -> savourrpc.keylocker.SetGiantWhaleRep
	8,  // 12: savourrpc.keylocker.GiantWhaleService.getGiantWhale:output_type -> savourrpc.keylocker.GetGiantWhaleRep
	10, // 13: savourrpc.keylocker.GiantWhaleService.removeGiantWhale:output_type -> savourrpc.keylocker.RemoveGiantWhaleRep
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_scrapy_proto_init() }
func file_scrapy_proto_init() {
	if File_scrapy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scrapy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiantWhale); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportChainReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportChainRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGiantWhaleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGiantWhaleRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGiantWhaleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGiantWhaleRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGiantWhaleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scrapy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGiantWhaleRep); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scrapy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scrapy_proto_goTypes,
		DependencyIndexes: file_scrapy_proto_depIdxs,
		EnumInfos:         file_scrapy_proto_enumTypes,
		MessageInfos:      file_scrapy_proto_msgTypes,
	}.Build()
	File_scrapy_proto = out.File
	file_scrapy_proto_rawDesc = nil
	file_scrapy_proto_goTypes = nil
	file_scrapy_proto_depIdxs = nil
}
