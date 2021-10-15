// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protos/key_range.proto

package proto

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

type KeyRangeStatus int32

const (
	KeyRangeStatus_LOCKED    KeyRangeStatus = 0
	KeyRangeStatus_AVAILABLE KeyRangeStatus = 1
)

// Enum value maps for KeyRangeStatus.
var (
	KeyRangeStatus_name = map[int32]string{
		0: "LOCKED",
		1: "AVAILABLE",
	}
	KeyRangeStatus_value = map[string]int32{
		"LOCKED":    0,
		"AVAILABLE": 1,
	}
)

func (x KeyRangeStatus) Enum() *KeyRangeStatus {
	p := new(KeyRangeStatus)
	*p = x
	return p
}

func (x KeyRangeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyRangeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_key_range_proto_enumTypes[0].Descriptor()
}

func (KeyRangeStatus) Type() protoreflect.EnumType {
	return &file_protos_key_range_proto_enumTypes[0]
}

func (x KeyRangeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyRangeStatus.Descriptor instead.
func (KeyRangeStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{0}
}

type AddShardingColumnReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddShardingColumnReply) Reset() {
	*x = AddShardingColumnReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShardingColumnReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShardingColumnReply) ProtoMessage() {}

func (x *AddShardingColumnReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShardingColumnReply.ProtoReflect.Descriptor instead.
func (*AddShardingColumnReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{0}
}

type AddLocalTableReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddLocalTableReply) Reset() {
	*x = AddLocalTableReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLocalTableReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLocalTableReply) ProtoMessage() {}

func (x *AddLocalTableReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLocalTableReply.ProtoReflect.Descriptor instead.
func (*AddLocalTableReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{1}
}

type AddLocalTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tname string `protobuf:"bytes,1,opt,name=tname,proto3" json:"tname,omitempty"`
}

func (x *AddLocalTableRequest) Reset() {
	*x = AddLocalTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLocalTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLocalTableRequest) ProtoMessage() {}

func (x *AddLocalTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLocalTableRequest.ProtoReflect.Descriptor instead.
func (*AddLocalTableRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{2}
}

func (x *AddLocalTableRequest) GetTname() string {
	if x != nil {
		return x.Tname
	}
	return ""
}

type AddShardingColumnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Colname []string `protobuf:"bytes,1,rep,name=colname,proto3" json:"colname,omitempty"`
}

func (x *AddShardingColumnRequest) Reset() {
	*x = AddShardingColumnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShardingColumnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShardingColumnRequest) ProtoMessage() {}

func (x *AddShardingColumnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShardingColumnRequest.ProtoReflect.Descriptor instead.
func (*AddShardingColumnRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{3}
}

func (x *AddShardingColumnRequest) GetColname() []string {
	if x != nil {
		return x.Colname
	}
	return nil
}

// key range is mapped to shard
type KeyRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowerBound []byte `protobuf:"bytes,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound []byte `protobuf:"bytes,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	Krid       string `protobuf:"bytes,3,opt,name=krid,proto3" json:"krid,omitempty"`
	ShardId    string `protobuf:"bytes,4,opt,name=shardId,proto3" json:"shardId,omitempty"`
}

func (x *KeyRange) Reset() {
	*x = KeyRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRange) ProtoMessage() {}

func (x *KeyRange) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRange.ProtoReflect.Descriptor instead.
func (*KeyRange) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{4}
}

func (x *KeyRange) GetLowerBound() []byte {
	if x != nil {
		return x.LowerBound
	}
	return nil
}

func (x *KeyRange) GetUpperBound() []byte {
	if x != nil {
		return x.UpperBound
	}
	return nil
}

func (x *KeyRange) GetKrid() string {
	if x != nil {
		return x.Krid
	}
	return ""
}

func (x *KeyRange) GetShardId() string {
	if x != nil {
		return x.ShardId
	}
	return ""
}

type ListKeyRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListKeyRangeRequest) Reset() {
	*x = ListKeyRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyRangeRequest) ProtoMessage() {}

func (x *ListKeyRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyRangeRequest.ProtoReflect.Descriptor instead.
func (*ListKeyRangeRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{5}
}

type SplitKeyRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Krid  string `protobuf:"bytes,1,opt,name=krid,proto3" json:"krid,omitempty"`
	Bound []byte `protobuf:"bytes,2,opt,name=bound,proto3" json:"bound,omitempty"`
}

func (x *SplitKeyRangeRequest) Reset() {
	*x = SplitKeyRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitKeyRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitKeyRangeRequest) ProtoMessage() {}

func (x *SplitKeyRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitKeyRangeRequest.ProtoReflect.Descriptor instead.
func (*SplitKeyRangeRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{6}
}

func (x *SplitKeyRangeRequest) GetKrid() string {
	if x != nil {
		return x.Krid
	}
	return ""
}

func (x *SplitKeyRangeRequest) GetBound() []byte {
	if x != nil {
		return x.Bound
	}
	return nil
}

type SplitKeyRangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SplittedKrid string `protobuf:"bytes,1,opt,name=splitted_krid,json=splittedKrid,proto3" json:"splitted_krid,omitempty"`
}

func (x *SplitKeyRangeReply) Reset() {
	*x = SplitKeyRangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitKeyRangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitKeyRangeReply) ProtoMessage() {}

func (x *SplitKeyRangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitKeyRangeReply.ProtoReflect.Descriptor instead.
func (*SplitKeyRangeReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{7}
}

func (x *SplitKeyRangeReply) GetSplittedKrid() string {
	if x != nil {
		return x.SplittedKrid
	}
	return ""
}

type LockKeyRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Krid string `protobuf:"bytes,3,opt,name=krid,proto3" json:"krid,omitempty"`
}

func (x *LockKeyRangeRequest) Reset() {
	*x = LockKeyRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockKeyRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockKeyRangeRequest) ProtoMessage() {}

func (x *LockKeyRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockKeyRangeRequest.ProtoReflect.Descriptor instead.
func (*LockKeyRangeRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{8}
}

func (x *LockKeyRangeRequest) GetKrid() string {
	if x != nil {
		return x.Krid
	}
	return ""
}

type LockKeyRangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status KeyRangeStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.spqr.KeyRangeStatus" json:"status,omitempty"`
}

func (x *LockKeyRangeReply) Reset() {
	*x = LockKeyRangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockKeyRangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockKeyRangeReply) ProtoMessage() {}

func (x *LockKeyRangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockKeyRangeReply.ProtoReflect.Descriptor instead.
func (*LockKeyRangeReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{9}
}

func (x *LockKeyRangeReply) GetStatus() KeyRangeStatus {
	if x != nil {
		return x.Status
	}
	return KeyRangeStatus_LOCKED
}

type UnlockKeyRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Krid string `protobuf:"bytes,3,opt,name=krid,proto3" json:"krid,omitempty"`
}

func (x *UnlockKeyRangeRequest) Reset() {
	*x = UnlockKeyRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockKeyRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockKeyRangeRequest) ProtoMessage() {}

func (x *UnlockKeyRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockKeyRangeRequest.ProtoReflect.Descriptor instead.
func (*UnlockKeyRangeRequest) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{10}
}

func (x *UnlockKeyRangeRequest) GetKrid() string {
	if x != nil {
		return x.Krid
	}
	return ""
}

type UnlockKeyRangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status KeyRangeStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.spqr.KeyRangeStatus" json:"status,omitempty"`
}

func (x *UnlockKeyRangeReply) Reset() {
	*x = UnlockKeyRangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockKeyRangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockKeyRangeReply) ProtoMessage() {}

func (x *UnlockKeyRangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockKeyRangeReply.ProtoReflect.Descriptor instead.
func (*UnlockKeyRangeReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{11}
}

func (x *UnlockKeyRangeReply) GetStatus() KeyRangeStatus {
	if x != nil {
		return x.Status
	}
	return KeyRangeStatus_LOCKED
}

type KeyRangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyRanges []*KeyRange `protobuf:"bytes,1,rep,name=key_ranges,json=keyRanges,proto3" json:"key_ranges,omitempty"`
}

func (x *KeyRangeReply) Reset() {
	*x = KeyRangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_key_range_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRangeReply) ProtoMessage() {}

func (x *KeyRangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_key_range_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRangeReply.ProtoReflect.Descriptor instead.
func (*KeyRangeReply) Descriptor() ([]byte, []int) {
	return file_protos_key_range_proto_rawDescGZIP(), []int{12}
}

func (x *KeyRangeReply) GetKeyRanges() []*KeyRange {
	if x != nil {
		return x.KeyRanges
	}
	return nil
}

var File_protos_key_range_proto protoreflect.FileDescriptor

var file_protos_key_range_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x73, 0x70, 0x71, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x14,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x39,
	0x0a, 0x12, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x5f, 0x6b, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x4b, 0x72, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x6f, 0x63,
	0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x72, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2b,
	0x0a, 0x15, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x72, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x13, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x2a, 0x2b,
	0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x32, 0xa0, 0x04, 0x0a, 0x0f,
	0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x4c, 0x6f,
	0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x4c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73,
	0x70, 0x71, 0x72, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a,
	0x0d, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x53, 0x70, 0x6c,
	0x69, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x70, 0x71, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_key_range_proto_rawDescOnce sync.Once
	file_protos_key_range_proto_rawDescData = file_protos_key_range_proto_rawDesc
)

func file_protos_key_range_proto_rawDescGZIP() []byte {
	file_protos_key_range_proto_rawDescOnce.Do(func() {
		file_protos_key_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_key_range_proto_rawDescData)
	})
	return file_protos_key_range_proto_rawDescData
}

var file_protos_key_range_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_key_range_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protos_key_range_proto_goTypes = []interface{}{
	(KeyRangeStatus)(0),              // 0: yandex.spqr.KeyRangeStatus
	(*AddShardingColumnReply)(nil),   // 1: yandex.spqr.AddShardingColumnReply
	(*AddLocalTableReply)(nil),       // 2: yandex.spqr.AddLocalTableReply
	(*AddLocalTableRequest)(nil),     // 3: yandex.spqr.AddLocalTableRequest
	(*AddShardingColumnRequest)(nil), // 4: yandex.spqr.AddShardingColumnRequest
	(*KeyRange)(nil),                 // 5: yandex.spqr.KeyRange
	(*ListKeyRangeRequest)(nil),      // 6: yandex.spqr.ListKeyRangeRequest
	(*SplitKeyRangeRequest)(nil),     // 7: yandex.spqr.SplitKeyRangeRequest
	(*SplitKeyRangeReply)(nil),       // 8: yandex.spqr.SplitKeyRangeReply
	(*LockKeyRangeRequest)(nil),      // 9: yandex.spqr.LockKeyRangeRequest
	(*LockKeyRangeReply)(nil),        // 10: yandex.spqr.LockKeyRangeReply
	(*UnlockKeyRangeRequest)(nil),    // 11: yandex.spqr.UnlockKeyRangeRequest
	(*UnlockKeyRangeReply)(nil),      // 12: yandex.spqr.UnlockKeyRangeReply
	(*KeyRangeReply)(nil),            // 13: yandex.spqr.KeyRangeReply
}
var file_protos_key_range_proto_depIdxs = []int32{
	0,  // 0: yandex.spqr.LockKeyRangeReply.status:type_name -> yandex.spqr.KeyRangeStatus
	0,  // 1: yandex.spqr.UnlockKeyRangeReply.status:type_name -> yandex.spqr.KeyRangeStatus
	5,  // 2: yandex.spqr.KeyRangeReply.key_ranges:type_name -> yandex.spqr.KeyRange
	6,  // 3: yandex.spqr.KeyRangeService.ListKeyRange:input_type -> yandex.spqr.ListKeyRangeRequest
	9,  // 4: yandex.spqr.KeyRangeService.LockKeyRange:input_type -> yandex.spqr.LockKeyRangeRequest
	11, // 5: yandex.spqr.KeyRangeService.UnlockKeyRange:input_type -> yandex.spqr.UnlockKeyRangeRequest
	7,  // 6: yandex.spqr.KeyRangeService.SplitKeyRange:input_type -> yandex.spqr.SplitKeyRangeRequest
	4,  // 7: yandex.spqr.KeyRangeService.AddShardingColumn:input_type -> yandex.spqr.AddShardingColumnRequest
	3,  // 8: yandex.spqr.KeyRangeService.AddLocalTable:input_type -> yandex.spqr.AddLocalTableRequest
	13, // 9: yandex.spqr.KeyRangeService.ListKeyRange:output_type -> yandex.spqr.KeyRangeReply
	10, // 10: yandex.spqr.KeyRangeService.LockKeyRange:output_type -> yandex.spqr.LockKeyRangeReply
	12, // 11: yandex.spqr.KeyRangeService.UnlockKeyRange:output_type -> yandex.spqr.UnlockKeyRangeReply
	8,  // 12: yandex.spqr.KeyRangeService.SplitKeyRange:output_type -> yandex.spqr.SplitKeyRangeReply
	1,  // 13: yandex.spqr.KeyRangeService.AddShardingColumn:output_type -> yandex.spqr.AddShardingColumnReply
	2,  // 14: yandex.spqr.KeyRangeService.AddLocalTable:output_type -> yandex.spqr.AddLocalTableReply
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_key_range_proto_init() }
func file_protos_key_range_proto_init() {
	if File_protos_key_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_key_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShardingColumnReply); i {
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
		file_protos_key_range_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLocalTableReply); i {
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
		file_protos_key_range_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLocalTableRequest); i {
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
		file_protos_key_range_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShardingColumnRequest); i {
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
		file_protos_key_range_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRange); i {
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
		file_protos_key_range_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyRangeRequest); i {
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
		file_protos_key_range_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitKeyRangeRequest); i {
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
		file_protos_key_range_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitKeyRangeReply); i {
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
		file_protos_key_range_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockKeyRangeRequest); i {
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
		file_protos_key_range_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockKeyRangeReply); i {
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
		file_protos_key_range_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockKeyRangeRequest); i {
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
		file_protos_key_range_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockKeyRangeReply); i {
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
		file_protos_key_range_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRangeReply); i {
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
			RawDescriptor: file_protos_key_range_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_key_range_proto_goTypes,
		DependencyIndexes: file_protos_key_range_proto_depIdxs,
		EnumInfos:         file_protos_key_range_proto_enumTypes,
		MessageInfos:      file_protos_key_range_proto_msgTypes,
	}.Build()
	File_protos_key_range_proto = out.File
	file_protos_key_range_proto_rawDesc = nil
	file_protos_key_range_proto_goTypes = nil
	file_protos_key_range_proto_depIdxs = nil
}
