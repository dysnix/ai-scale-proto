// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/services/auth.proto

package services

import (
	context "context"
	events "github.com/dysnix/ai-scale-proto/external/proto/events"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReqCreateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *events.Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *ReqCreateClient) Reset() {
	*x = ReqCreateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCreateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCreateClient) ProtoMessage() {}

func (x *ReqCreateClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCreateClient.ProtoReflect.Descriptor instead.
func (*ReqCreateClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ReqCreateClient) GetClient() *events.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type ResCreateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResCreateClient) Reset() {
	*x = ResCreateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResCreateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResCreateClient) ProtoMessage() {}

func (x *ResCreateClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResCreateClient.ProtoReflect.Descriptor instead.
func (*ResCreateClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{1}
}

type ReqUpdateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *events.Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *ReqUpdateClient) Reset() {
	*x = ReqUpdateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateClient) ProtoMessage() {}

func (x *ReqUpdateClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateClient.ProtoReflect.Descriptor instead.
func (*ReqUpdateClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ReqUpdateClient) GetClient() *events.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type ResUpdateClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResUpdateClient) Reset() {
	*x = ResUpdateClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResUpdateClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResUpdateClient) ProtoMessage() {}

func (x *ResUpdateClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResUpdateClient.ProtoReflect.Descriptor instead.
func (*ResUpdateClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{3}
}

type ReqDeleteClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//	*ReqDeleteClient_Name
	//	*ReqDeleteClient_ClusterId
	Query isReqDeleteClient_Query `protobuf_oneof:"Query"`
}

func (x *ReqDeleteClient) Reset() {
	*x = ReqDeleteClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDeleteClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDeleteClient) ProtoMessage() {}

func (x *ReqDeleteClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDeleteClient.ProtoReflect.Descriptor instead.
func (*ReqDeleteClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{4}
}

func (m *ReqDeleteClient) GetQuery() isReqDeleteClient_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *ReqDeleteClient) GetName() string {
	if x, ok := x.GetQuery().(*ReqDeleteClient_Name); ok {
		return x.Name
	}
	return ""
}

func (x *ReqDeleteClient) GetClusterId() string {
	if x, ok := x.GetQuery().(*ReqDeleteClient_ClusterId); ok {
		return x.ClusterId
	}
	return ""
}

type isReqDeleteClient_Query interface {
	isReqDeleteClient_Query()
}

type ReqDeleteClient_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type ReqDeleteClient_ClusterId struct {
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3,oneof"`
}

func (*ReqDeleteClient_Name) isReqDeleteClient_Query() {}

func (*ReqDeleteClient_ClusterId) isReqDeleteClient_Query() {}

type ResDeleteClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResDeleteClient) Reset() {
	*x = ResDeleteClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResDeleteClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResDeleteClient) ProtoMessage() {}

func (x *ResDeleteClient) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResDeleteClient.ProtoReflect.Descriptor instead.
func (*ResDeleteClient) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{5}
}

type ReqGenerateToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//	*ReqGenerateToken_Name
	//	*ReqGenerateToken_ClusterId
	Query       isReqGenerateToken_Query `protobuf_oneof:"Query"`
	ExtDuration *duration.Duration       `protobuf:"bytes,3,opt,name=ext_duration,json=extDuration,proto3" json:"ext_duration,omitempty"`
}

func (x *ReqGenerateToken) Reset() {
	*x = ReqGenerateToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGenerateToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGenerateToken) ProtoMessage() {}

func (x *ReqGenerateToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGenerateToken.ProtoReflect.Descriptor instead.
func (*ReqGenerateToken) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{6}
}

func (m *ReqGenerateToken) GetQuery() isReqGenerateToken_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *ReqGenerateToken) GetName() string {
	if x, ok := x.GetQuery().(*ReqGenerateToken_Name); ok {
		return x.Name
	}
	return ""
}

func (x *ReqGenerateToken) GetClusterId() string {
	if x, ok := x.GetQuery().(*ReqGenerateToken_ClusterId); ok {
		return x.ClusterId
	}
	return ""
}

func (x *ReqGenerateToken) GetExtDuration() *duration.Duration {
	if x != nil {
		return x.ExtDuration
	}
	return nil
}

type isReqGenerateToken_Query interface {
	isReqGenerateToken_Query()
}

type ReqGenerateToken_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type ReqGenerateToken_ClusterId struct {
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3,oneof"`
}

func (*ReqGenerateToken_Name) isReqGenerateToken_Query() {}

func (*ReqGenerateToken_ClusterId) isReqGenerateToken_Query() {}

type ResGenerateToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResGenerateToken) Reset() {
	*x = ResGenerateToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGenerateToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGenerateToken) ProtoMessage() {}

func (x *ResGenerateToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGenerateToken.ProtoReflect.Descriptor instead.
func (*ResGenerateToken) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ResGenerateToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ReqRefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Query:
	//	*ReqRefreshToken_Name
	//	*ReqRefreshToken_ClusterId
	Query       isReqRefreshToken_Query `protobuf_oneof:"Query"`
	Token       string                  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" validate:"jwt"` // @gotags: validate:"jwt"
	ExtDuration *duration.Duration      `protobuf:"bytes,4,opt,name=ext_duration,json=extDuration,proto3" json:"ext_duration,omitempty"`
}

func (x *ReqRefreshToken) Reset() {
	*x = ReqRefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRefreshToken) ProtoMessage() {}

func (x *ReqRefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRefreshToken.ProtoReflect.Descriptor instead.
func (*ReqRefreshToken) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{8}
}

func (m *ReqRefreshToken) GetQuery() isReqRefreshToken_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *ReqRefreshToken) GetName() string {
	if x, ok := x.GetQuery().(*ReqRefreshToken_Name); ok {
		return x.Name
	}
	return ""
}

func (x *ReqRefreshToken) GetClusterId() string {
	if x, ok := x.GetQuery().(*ReqRefreshToken_ClusterId); ok {
		return x.ClusterId
	}
	return ""
}

func (x *ReqRefreshToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ReqRefreshToken) GetExtDuration() *duration.Duration {
	if x != nil {
		return x.ExtDuration
	}
	return nil
}

type isReqRefreshToken_Query interface {
	isReqRefreshToken_Query()
}

type ReqRefreshToken_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type ReqRefreshToken_ClusterId struct {
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3,oneof"`
}

func (*ReqRefreshToken_Name) isReqRefreshToken_Query() {}

func (*ReqRefreshToken_ClusterId) isReqRefreshToken_Query() {}

type ResRefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResRefreshToken) Reset() {
	*x = ResRefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResRefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResRefreshToken) ProtoMessage() {}

func (x *ResRefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResRefreshToken.ProtoReflect.Descriptor instead.
func (*ResRefreshToken) Descriptor() ([]byte, []int) {
	return file_proto_services_auth_proto_rawDescGZIP(), []int{9}
}

func (x *ResRefreshToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_services_auth_proto protoreflect.FileDescriptor

var file_proto_services_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x0f, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x0f,
	0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x0f, 0x52, 0x65,
	0x71, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x11, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x90, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0c,
	0x65, 0x78, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65,
	0x78, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa5, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x71, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3c,
	0x0a, 0x0c, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x65, 0x78, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xf8,
	0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x73, 0x6e, 0x69, 0x78, 0x2f, 0x61,
	0x69, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_auth_proto_rawDescOnce sync.Once
	file_proto_services_auth_proto_rawDescData = file_proto_services_auth_proto_rawDesc
)

func file_proto_services_auth_proto_rawDescGZIP() []byte {
	file_proto_services_auth_proto_rawDescOnce.Do(func() {
		file_proto_services_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_auth_proto_rawDescData)
	})
	return file_proto_services_auth_proto_rawDescData
}

var file_proto_services_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_services_auth_proto_goTypes = []interface{}{
	(*ReqCreateClient)(nil),   // 0: services.ReqCreateClient
	(*ResCreateClient)(nil),   // 1: services.ResCreateClient
	(*ReqUpdateClient)(nil),   // 2: services.ReqUpdateClient
	(*ResUpdateClient)(nil),   // 3: services.ResUpdateClient
	(*ReqDeleteClient)(nil),   // 4: services.ReqDeleteClient
	(*ResDeleteClient)(nil),   // 5: services.ResDeleteClient
	(*ReqGenerateToken)(nil),  // 6: services.ReqGenerateToken
	(*ResGenerateToken)(nil),  // 7: services.ResGenerateToken
	(*ReqRefreshToken)(nil),   // 8: services.ReqRefreshToken
	(*ResRefreshToken)(nil),   // 9: services.ResRefreshToken
	(*events.Client)(nil),     // 10: events.Client
	(*duration.Duration)(nil), // 11: google.protobuf.Duration
}
var file_proto_services_auth_proto_depIdxs = []int32{
	10, // 0: services.ReqCreateClient.client:type_name -> events.Client
	10, // 1: services.ReqUpdateClient.client:type_name -> events.Client
	11, // 2: services.ReqGenerateToken.ext_duration:type_name -> google.protobuf.Duration
	11, // 3: services.ReqRefreshToken.ext_duration:type_name -> google.protobuf.Duration
	0,  // 4: services.AuthService.CreateClient:input_type -> services.ReqCreateClient
	2,  // 5: services.AuthService.UpdateClient:input_type -> services.ReqUpdateClient
	4,  // 6: services.AuthService.DeleteClient:input_type -> services.ReqDeleteClient
	6,  // 7: services.AuthService.GenerateToken:input_type -> services.ReqGenerateToken
	8,  // 8: services.AuthService.RefreshToken:input_type -> services.ReqRefreshToken
	1,  // 9: services.AuthService.CreateClient:output_type -> services.ResCreateClient
	3,  // 10: services.AuthService.UpdateClient:output_type -> services.ResUpdateClient
	5,  // 11: services.AuthService.DeleteClient:output_type -> services.ResDeleteClient
	7,  // 12: services.AuthService.GenerateToken:output_type -> services.ResGenerateToken
	9,  // 13: services.AuthService.RefreshToken:output_type -> services.ResRefreshToken
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_services_auth_proto_init() }
func file_proto_services_auth_proto_init() {
	if File_proto_services_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCreateClient); i {
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
		file_proto_services_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResCreateClient); i {
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
		file_proto_services_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdateClient); i {
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
		file_proto_services_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResUpdateClient); i {
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
		file_proto_services_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDeleteClient); i {
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
		file_proto_services_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResDeleteClient); i {
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
		file_proto_services_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGenerateToken); i {
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
		file_proto_services_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGenerateToken); i {
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
		file_proto_services_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRefreshToken); i {
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
		file_proto_services_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResRefreshToken); i {
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
	file_proto_services_auth_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ReqDeleteClient_Name)(nil),
		(*ReqDeleteClient_ClusterId)(nil),
	}
	file_proto_services_auth_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ReqGenerateToken_Name)(nil),
		(*ReqGenerateToken_ClusterId)(nil),
	}
	file_proto_services_auth_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ReqRefreshToken_Name)(nil),
		(*ReqRefreshToken_ClusterId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_services_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_auth_proto_goTypes,
		DependencyIndexes: file_proto_services_auth_proto_depIdxs,
		MessageInfos:      file_proto_services_auth_proto_msgTypes,
	}.Build()
	File_proto_services_auth_proto = out.File
	file_proto_services_auth_proto_rawDesc = nil
	file_proto_services_auth_proto_goTypes = nil
	file_proto_services_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateClient(ctx context.Context, in *ReqCreateClient, opts ...grpc.CallOption) (*ResCreateClient, error)
	UpdateClient(ctx context.Context, in *ReqUpdateClient, opts ...grpc.CallOption) (*ResUpdateClient, error)
	DeleteClient(ctx context.Context, in *ReqDeleteClient, opts ...grpc.CallOption) (*ResDeleteClient, error)
	GenerateToken(ctx context.Context, in *ReqGenerateToken, opts ...grpc.CallOption) (*ResGenerateToken, error)
	RefreshToken(ctx context.Context, in *ReqRefreshToken, opts ...grpc.CallOption) (*ResRefreshToken, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateClient(ctx context.Context, in *ReqCreateClient, opts ...grpc.CallOption) (*ResCreateClient, error) {
	out := new(ResCreateClient)
	err := c.cc.Invoke(ctx, "/services.AuthService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateClient(ctx context.Context, in *ReqUpdateClient, opts ...grpc.CallOption) (*ResUpdateClient, error) {
	out := new(ResUpdateClient)
	err := c.cc.Invoke(ctx, "/services.AuthService/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteClient(ctx context.Context, in *ReqDeleteClient, opts ...grpc.CallOption) (*ResDeleteClient, error) {
	out := new(ResDeleteClient)
	err := c.cc.Invoke(ctx, "/services.AuthService/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateToken(ctx context.Context, in *ReqGenerateToken, opts ...grpc.CallOption) (*ResGenerateToken, error) {
	out := new(ResGenerateToken)
	err := c.cc.Invoke(ctx, "/services.AuthService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *ReqRefreshToken, opts ...grpc.CallOption) (*ResRefreshToken, error) {
	out := new(ResRefreshToken)
	err := c.cc.Invoke(ctx, "/services.AuthService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	CreateClient(context.Context, *ReqCreateClient) (*ResCreateClient, error)
	UpdateClient(context.Context, *ReqUpdateClient) (*ResUpdateClient, error)
	DeleteClient(context.Context, *ReqDeleteClient) (*ResDeleteClient, error)
	GenerateToken(context.Context, *ReqGenerateToken) (*ResGenerateToken, error)
	RefreshToken(context.Context, *ReqRefreshToken) (*ResRefreshToken, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) CreateClient(context.Context, *ReqCreateClient) (*ResCreateClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (*UnimplementedAuthServiceServer) UpdateClient(context.Context, *ReqUpdateClient) (*ResUpdateClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (*UnimplementedAuthServiceServer) DeleteClient(context.Context, *ReqDeleteClient) (*ResDeleteClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (*UnimplementedAuthServiceServer) GenerateToken(context.Context, *ReqGenerateToken) (*ResGenerateToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (*UnimplementedAuthServiceServer) RefreshToken(context.Context, *ReqRefreshToken) (*ResRefreshToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.AuthService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateClient(ctx, req.(*ReqCreateClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.AuthService/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateClient(ctx, req.(*ReqUpdateClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.AuthService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteClient(ctx, req.(*ReqDeleteClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGenerateToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.AuthService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateToken(ctx, req.(*ReqGenerateToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqRefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*ReqRefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _AuthService_CreateClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _AuthService_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _AuthService_DeleteClient_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _AuthService_GenerateToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/auth.proto",
}
