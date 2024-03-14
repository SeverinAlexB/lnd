// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: peersrpc/peers.proto

package peersrpc

import (
	lnrpc "github.com/lightningnetwork/lnd/lnrpc"
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

// UpdateAction is used to determine the kind of action we are referring to.
type UpdateAction int32

const (
	// ADD indicates this is an "insertion" kind of action.
	UpdateAction_ADD UpdateAction = 0
	// REMOVE indicates this is a "deletion" kind of action.
	UpdateAction_REMOVE UpdateAction = 1
)

// Enum value maps for UpdateAction.
var (
	UpdateAction_name = map[int32]string{
		0: "ADD",
		1: "REMOVE",
	}
	UpdateAction_value = map[string]int32{
		"ADD":    0,
		"REMOVE": 1,
	}
)

func (x UpdateAction) Enum() *UpdateAction {
	p := new(UpdateAction)
	*p = x
	return p
}

func (x UpdateAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateAction) Descriptor() protoreflect.EnumDescriptor {
	return file_peersrpc_peers_proto_enumTypes[0].Descriptor()
}

func (UpdateAction) Type() protoreflect.EnumType {
	return &file_peersrpc_peers_proto_enumTypes[0]
}

func (x UpdateAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateAction.Descriptor instead.
func (UpdateAction) EnumDescriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{0}
}

type FeatureSet int32

const (
	// SET_INIT identifies features that should be sent in an Init message to
	// a remote peer.
	FeatureSet_SET_INIT FeatureSet = 0
	// SET_LEGACY_GLOBAL identifies features that should be set in the legacy
	// GlobalFeatures field of an Init message, which maintains backwards
	// compatibility with nodes that haven't implemented flat features.
	FeatureSet_SET_LEGACY_GLOBAL FeatureSet = 1
	// SET_NODE_ANN identifies features that should be advertised on node
	// announcements.
	FeatureSet_SET_NODE_ANN FeatureSet = 2
	// SET_INVOICE identifies features that should be advertised on invoices
	// generated by the daemon.
	FeatureSet_SET_INVOICE FeatureSet = 3
	// SET_INVOICE_AMP identifies the features that should be advertised on
	// AMP invoices generated by the daemon.
	FeatureSet_SET_INVOICE_AMP FeatureSet = 4
)

// Enum value maps for FeatureSet.
var (
	FeatureSet_name = map[int32]string{
		0: "SET_INIT",
		1: "SET_LEGACY_GLOBAL",
		2: "SET_NODE_ANN",
		3: "SET_INVOICE",
		4: "SET_INVOICE_AMP",
	}
	FeatureSet_value = map[string]int32{
		"SET_INIT":          0,
		"SET_LEGACY_GLOBAL": 1,
		"SET_NODE_ANN":      2,
		"SET_INVOICE":       3,
		"SET_INVOICE_AMP":   4,
	}
)

func (x FeatureSet) Enum() *FeatureSet {
	p := new(FeatureSet)
	*p = x
	return p
}

func (x FeatureSet) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureSet) Descriptor() protoreflect.EnumDescriptor {
	return file_peersrpc_peers_proto_enumTypes[1].Descriptor()
}

func (FeatureSet) Type() protoreflect.EnumType {
	return &file_peersrpc_peers_proto_enumTypes[1]
}

func (x FeatureSet) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureSet.Descriptor instead.
func (FeatureSet) EnumDescriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{1}
}

type UpdateAddressAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Determines the kind of action.
	Action UpdateAction `protobuf:"varint,1,opt,name=action,proto3,enum=peersrpc.UpdateAction" json:"action,omitempty"`
	// The address used to apply the update action.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UpdateAddressAction) Reset() {
	*x = UpdateAddressAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peersrpc_peers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAddressAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddressAction) ProtoMessage() {}

func (x *UpdateAddressAction) ProtoReflect() protoreflect.Message {
	mi := &file_peersrpc_peers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddressAction.ProtoReflect.Descriptor instead.
func (*UpdateAddressAction) Descriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAddressAction) GetAction() UpdateAction {
	if x != nil {
		return x.Action
	}
	return UpdateAction_ADD
}

func (x *UpdateAddressAction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UpdateFeatureAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Determines the kind of action.
	Action UpdateAction `protobuf:"varint,1,opt,name=action,proto3,enum=peersrpc.UpdateAction" json:"action,omitempty"`
	// The feature bit used to apply the update action.
	FeatureBit lnrpc.FeatureBit `protobuf:"varint,2,opt,name=feature_bit,json=featureBit,proto3,enum=lnrpc.FeatureBit" json:"feature_bit,omitempty"`
}

func (x *UpdateFeatureAction) Reset() {
	*x = UpdateFeatureAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peersrpc_peers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeatureAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeatureAction) ProtoMessage() {}

func (x *UpdateFeatureAction) ProtoReflect() protoreflect.Message {
	mi := &file_peersrpc_peers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeatureAction.ProtoReflect.Descriptor instead.
func (*UpdateFeatureAction) Descriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateFeatureAction) GetAction() UpdateAction {
	if x != nil {
		return x.Action
	}
	return UpdateAction_ADD
}

func (x *UpdateFeatureAction) GetFeatureBit() lnrpc.FeatureBit {
	if x != nil {
		return x.FeatureBit
	}
	return lnrpc.FeatureBit(0)
}

type NodeAnnouncementUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of changes for the features that the node supports.
	FeatureUpdates []*UpdateFeatureAction `protobuf:"bytes,1,rep,name=feature_updates,json=featureUpdates,proto3" json:"feature_updates,omitempty"`
	// Color is the node's color in hex code format.
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// Alias or nick name of the node.
	Alias string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	// Set of changes for the node's known addresses.
	AddressUpdates []*UpdateAddressAction `protobuf:"bytes,4,rep,name=address_updates,json=addressUpdates,proto3" json:"address_updates,omitempty"`
}

func (x *NodeAnnouncementUpdateRequest) Reset() {
	*x = NodeAnnouncementUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peersrpc_peers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAnnouncementUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAnnouncementUpdateRequest) ProtoMessage() {}

func (x *NodeAnnouncementUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peersrpc_peers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAnnouncementUpdateRequest.ProtoReflect.Descriptor instead.
func (*NodeAnnouncementUpdateRequest) Descriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{2}
}

func (x *NodeAnnouncementUpdateRequest) GetFeatureUpdates() []*UpdateFeatureAction {
	if x != nil {
		return x.FeatureUpdates
	}
	return nil
}

func (x *NodeAnnouncementUpdateRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *NodeAnnouncementUpdateRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *NodeAnnouncementUpdateRequest) GetAddressUpdates() []*UpdateAddressAction {
	if x != nil {
		return x.AddressUpdates
	}
	return nil
}

type NodeAnnouncementUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ops []*lnrpc.Op `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty"`
}

func (x *NodeAnnouncementUpdateResponse) Reset() {
	*x = NodeAnnouncementUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peersrpc_peers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAnnouncementUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAnnouncementUpdateResponse) ProtoMessage() {}

func (x *NodeAnnouncementUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peersrpc_peers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAnnouncementUpdateResponse.ProtoReflect.Descriptor instead.
func (*NodeAnnouncementUpdateResponse) Descriptor() ([]byte, []int) {
	return file_peersrpc_peers_proto_rawDescGZIP(), []int{3}
}

func (x *NodeAnnouncementUpdateResponse) GetOps() []*lnrpc.Op {
	if x != nil {
		return x.Ops
	}
	return nil
}

var File_peersrpc_peers_proto protoreflect.FileDescriptor

var file_peersrpc_peers_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x65, 0x72, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x65, 0x65, 0x72, 0x73, 0x72, 0x70, 0x63,
	0x1a, 0x0f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x79, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x6c, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x69,
	0x74, 0x52, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x69, 0x74, 0x22, 0xdb, 0x01,
	0x0a, 0x1d, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x46, 0x0a, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x12, 0x46, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x1e, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x03, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6c, 0x6e, 0x72,
	0x70, 0x63, 0x2e, 0x4f, 0x70, 0x52, 0x03, 0x6f, 0x70, 0x73, 0x2a, 0x23, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x2a,
	0x69, 0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x45, 0x54, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x45, 0x54, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x41,
	0x4e, 0x4e, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f, 0x49, 0x4e, 0x56,
	0x4f, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x4d, 0x50, 0x10, 0x04, 0x32, 0x74, 0x0a, 0x05, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x12, 0x6b, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x72, 0x70,
	0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x6c, 0x6e, 0x64, 0x2f, 0x6c, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peersrpc_peers_proto_rawDescOnce sync.Once
	file_peersrpc_peers_proto_rawDescData = file_peersrpc_peers_proto_rawDesc
)

func file_peersrpc_peers_proto_rawDescGZIP() []byte {
	file_peersrpc_peers_proto_rawDescOnce.Do(func() {
		file_peersrpc_peers_proto_rawDescData = protoimpl.X.CompressGZIP(file_peersrpc_peers_proto_rawDescData)
	})
	return file_peersrpc_peers_proto_rawDescData
}

var file_peersrpc_peers_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_peersrpc_peers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_peersrpc_peers_proto_goTypes = []interface{}{
	(UpdateAction)(0),                      // 0: peersrpc.UpdateAction
	(FeatureSet)(0),                        // 1: peersrpc.FeatureSet
	(*UpdateAddressAction)(nil),            // 2: peersrpc.UpdateAddressAction
	(*UpdateFeatureAction)(nil),            // 3: peersrpc.UpdateFeatureAction
	(*NodeAnnouncementUpdateRequest)(nil),  // 4: peersrpc.NodeAnnouncementUpdateRequest
	(*NodeAnnouncementUpdateResponse)(nil), // 5: peersrpc.NodeAnnouncementUpdateResponse
	(lnrpc.FeatureBit)(0),                  // 6: lnrpc.FeatureBit
	(*lnrpc.Op)(nil),                       // 7: lnrpc.Op
}
var file_peersrpc_peers_proto_depIdxs = []int32{
	0, // 0: peersrpc.UpdateAddressAction.action:type_name -> peersrpc.UpdateAction
	0, // 1: peersrpc.UpdateFeatureAction.action:type_name -> peersrpc.UpdateAction
	6, // 2: peersrpc.UpdateFeatureAction.feature_bit:type_name -> lnrpc.FeatureBit
	3, // 3: peersrpc.NodeAnnouncementUpdateRequest.feature_updates:type_name -> peersrpc.UpdateFeatureAction
	2, // 4: peersrpc.NodeAnnouncementUpdateRequest.address_updates:type_name -> peersrpc.UpdateAddressAction
	7, // 5: peersrpc.NodeAnnouncementUpdateResponse.ops:type_name -> lnrpc.Op
	4, // 6: peersrpc.Peers.UpdateNodeAnnouncement:input_type -> peersrpc.NodeAnnouncementUpdateRequest
	5, // 7: peersrpc.Peers.UpdateNodeAnnouncement:output_type -> peersrpc.NodeAnnouncementUpdateResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_peersrpc_peers_proto_init() }
func file_peersrpc_peers_proto_init() {
	if File_peersrpc_peers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peersrpc_peers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAddressAction); i {
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
		file_peersrpc_peers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeatureAction); i {
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
		file_peersrpc_peers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAnnouncementUpdateRequest); i {
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
		file_peersrpc_peers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAnnouncementUpdateResponse); i {
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
			RawDescriptor: file_peersrpc_peers_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peersrpc_peers_proto_goTypes,
		DependencyIndexes: file_peersrpc_peers_proto_depIdxs,
		EnumInfos:         file_peersrpc_peers_proto_enumTypes,
		MessageInfos:      file_peersrpc_peers_proto_msgTypes,
	}.Build()
	File_peersrpc_peers_proto = out.File
	file_peersrpc_peers_proto_rawDesc = nil
	file_peersrpc_peers_proto_goTypes = nil
	file_peersrpc_peers_proto_depIdxs = nil
}
