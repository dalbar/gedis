// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: gedis.proto

package grpc

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

type InstanceKeyBroadcast_Operation int32

const (
	InstanceKeyBroadcast_WRITE  InstanceKeyBroadcast_Operation = 0
	InstanceKeyBroadcast_DELETE InstanceKeyBroadcast_Operation = 1
)

// Enum value maps for InstanceKeyBroadcast_Operation.
var (
	InstanceKeyBroadcast_Operation_name = map[int32]string{
		0: "WRITE",
		1: "DELETE",
	}
	InstanceKeyBroadcast_Operation_value = map[string]int32{
		"WRITE":  0,
		"DELETE": 1,
	}
)

func (x InstanceKeyBroadcast_Operation) Enum() *InstanceKeyBroadcast_Operation {
	p := new(InstanceKeyBroadcast_Operation)
	*p = x
	return p
}

func (x InstanceKeyBroadcast_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceKeyBroadcast_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_gedis_proto_enumTypes[0].Descriptor()
}

func (InstanceKeyBroadcast_Operation) Type() protoreflect.EnumType {
	return &file_gedis_proto_enumTypes[0]
}

func (x InstanceKeyBroadcast_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceKeyBroadcast_Operation.Descriptor instead.
func (InstanceKeyBroadcast_Operation) EnumDescriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{3, 0}
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{0}
}

func (x *Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type InstanceKeyBroadcast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string                         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     string                         `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Instance  string                         `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	Operation InstanceKeyBroadcast_Operation `protobuf:"varint,4,opt,name=operation,proto3,enum=InstanceKeyBroadcast_Operation" json:"operation,omitempty"`
}

func (x *InstanceKeyBroadcast) Reset() {
	*x = InstanceKeyBroadcast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceKeyBroadcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceKeyBroadcast) ProtoMessage() {}

func (x *InstanceKeyBroadcast) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceKeyBroadcast.ProtoReflect.Descriptor instead.
func (*InstanceKeyBroadcast) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceKeyBroadcast) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *InstanceKeyBroadcast) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *InstanceKeyBroadcast) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *InstanceKeyBroadcast) GetOperation() InstanceKeyBroadcast_Operation {
	if x != nil {
		return x.Operation
	}
	return InstanceKeyBroadcast_WRITE
}

type InstanceKeyBroadCastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance string `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *InstanceKeyBroadCastReply) Reset() {
	*x = InstanceKeyBroadCastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceKeyBroadCastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceKeyBroadCastReply) ProtoMessage() {}

func (x *InstanceKeyBroadCastReply) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceKeyBroadCastReply.ProtoReflect.Descriptor instead.
func (*InstanceKeyBroadCastReply) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{4}
}

func (x *InstanceKeyBroadCastReply) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*InstanceKeyBroadCastReply `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatResponse) GetInstances() []*InstanceKeyBroadCastReply {
	if x != nil {
		return x.Instances
	}
	return nil
}

type DatabaseDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*Pair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *DatabaseDump) Reset() {
	*x = DatabaseDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseDump) ProtoMessage() {}

func (x *DatabaseDump) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseDump.ProtoReflect.Descriptor instead.
func (*DatabaseDump) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{6}
}

func (x *DatabaseDump) GetPairs() []*Pair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// Empty can be extended in the future
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gedis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gedis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gedis_proto_rawDescGZIP(), []int{7}
}

var File_gedis_proto protoreflect.FileDescriptor

var file_gedis_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a,
	0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x21, 0x0a,
	0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0xbd, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x09,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01,
	0x22, 0x37, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x81,
	0x02, 0x0a, 0x05, 0x47, 0x65, 0x64, 0x69, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0b, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0b, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x04, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x3b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x1a, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x12,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4b, 0x65, 0x79, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gedis_proto_rawDescOnce sync.Once
	file_gedis_proto_rawDescData = file_gedis_proto_rawDesc
)

func file_gedis_proto_rawDescGZIP() []byte {
	file_gedis_proto_rawDescOnce.Do(func() {
		file_gedis_proto_rawDescData = protoimpl.X.CompressGZIP(file_gedis_proto_rawDescData)
	})
	return file_gedis_proto_rawDescData
}

var file_gedis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gedis_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gedis_proto_goTypes = []interface{}{
	(InstanceKeyBroadcast_Operation)(0), // 0: InstanceKeyBroadcast.Operation
	(*Pair)(nil),                        // 1: Pair
	(*KeyRequest)(nil),                  // 2: KeyRequest
	(*ErrorResponse)(nil),               // 3: ErrorResponse
	(*InstanceKeyBroadcast)(nil),        // 4: InstanceKeyBroadcast
	(*InstanceKeyBroadCastReply)(nil),   // 5: InstanceKeyBroadCastReply
	(*HeartbeatResponse)(nil),           // 6: HeartbeatResponse
	(*DatabaseDump)(nil),                // 7: DatabaseDump
	(*Empty)(nil),                       // 8: Empty
}
var file_gedis_proto_depIdxs = []int32{
	0, // 0: InstanceKeyBroadcast.operation:type_name -> InstanceKeyBroadcast.Operation
	5, // 1: HeartbeatResponse.instances:type_name -> InstanceKeyBroadCastReply
	1, // 2: DatabaseDump.pairs:type_name -> Pair
	2, // 3: Gedis.GetValue:input_type -> KeyRequest
	1, // 4: Gedis.PutPair:input_type -> Pair
	2, // 5: Gedis.DeleteValue:input_type -> KeyRequest
	8, // 6: Gedis.Dump:input_type -> Empty
	5, // 7: Gedis.Heartbeat:input_type -> InstanceKeyBroadCastReply
	4, // 8: Gedis.Broadcast:input_type -> InstanceKeyBroadcast
	1, // 9: Gedis.GetValue:output_type -> Pair
	8, // 10: Gedis.PutPair:output_type -> Empty
	8, // 11: Gedis.DeleteValue:output_type -> Empty
	7, // 12: Gedis.Dump:output_type -> DatabaseDump
	6, // 13: Gedis.Heartbeat:output_type -> HeartbeatResponse
	5, // 14: Gedis.Broadcast:output_type -> InstanceKeyBroadCastReply
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gedis_proto_init() }
func file_gedis_proto_init() {
	if File_gedis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gedis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_gedis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
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
		file_gedis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_gedis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceKeyBroadcast); i {
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
		file_gedis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceKeyBroadCastReply); i {
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
		file_gedis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_gedis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseDump); i {
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
		file_gedis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_gedis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gedis_proto_goTypes,
		DependencyIndexes: file_gedis_proto_depIdxs,
		EnumInfos:         file_gedis_proto_enumTypes,
		MessageInfos:      file_gedis_proto_msgTypes,
	}.Build()
	File_gedis_proto = out.File
	file_gedis_proto_rawDesc = nil
	file_gedis_proto_goTypes = nil
	file_gedis_proto_depIdxs = nil
}
