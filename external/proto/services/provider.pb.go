// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/services/provider.proto

package services

import (
	context "context"
	commonproto "github.com/dysnix/ai-scale-proto/external/proto/commonproto"
	enums "github.com/dysnix/ai-scale-proto/external/proto/enums"
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

type ReqPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping *commonproto.Ping `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *ReqPing) Reset() {
	*x = ReqPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPing) ProtoMessage() {}

func (x *ReqPing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPing.ProtoReflect.Descriptor instead.
func (*ReqPing) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ReqPing) GetPing() *commonproto.Ping {
	if x != nil {
		return x.Ping
	}
	return nil
}

type ResPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong *commonproto.Pong `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ResPong) Reset() {
	*x = ResPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPong) ProtoMessage() {}

func (x *ResPong) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPong.ProtoReflect.Descriptor instead.
func (*ResPong) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ResPong) GetPong() *commonproto.Pong {
	if x != nil {
		return x.Pong
	}
	return nil
}

type ReqGetMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource     enums.ResourceType `protobuf:"varint,1,opt,name=resource,proto3,enum=enums.ResourceType" json:"resource,omitempty"`
	Name         string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace    string             `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Cluster      string             `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	StepDuration *duration.Duration `protobuf:"bytes,5,opt,name=stepDuration,proto3" json:"stepDuration,omitempty"`
	TimeWindow   *duration.Duration `protobuf:"bytes,6,opt,name=timeWindow,proto3" json:"timeWindow,omitempty"`
	// Types that are assignable to Query:
	//	*ReqGetMetrics_History
	Query isReqGetMetrics_Query `protobuf_oneof:"Query"`
}

func (x *ReqGetMetrics) Reset() {
	*x = ReqGetMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetMetrics) ProtoMessage() {}

func (x *ReqGetMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetMetrics.ProtoReflect.Descriptor instead.
func (*ReqGetMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGetMetrics) GetResource() enums.ResourceType {
	if x != nil {
		return x.Resource
	}
	return enums.ResourceType_Deployment
}

func (x *ReqGetMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqGetMetrics) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ReqGetMetrics) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ReqGetMetrics) GetStepDuration() *duration.Duration {
	if x != nil {
		return x.StepDuration
	}
	return nil
}

func (x *ReqGetMetrics) GetTimeWindow() *duration.Duration {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (m *ReqGetMetrics) GetQuery() isReqGetMetrics_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *ReqGetMetrics) GetHistory() *events.History {
	if x, ok := x.GetQuery().(*ReqGetMetrics_History); ok {
		return x.History
	}
	return nil
}

type isReqGetMetrics_Query interface {
	isReqGetMetrics_Query()
}

type ReqGetMetrics_History struct {
	History *events.History `protobuf:"bytes,7,opt,name=history,proto3,oneof"`
}

func (*ReqGetMetrics_History) isReqGetMetrics_Query() {}

type ResGetMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricValues []*commonproto.MetricValue `protobuf:"bytes,1,rep,name=metric_values,json=metricValues,proto3" json:"metric_values,omitempty"`
}

func (x *ResGetMetrics) Reset() {
	*x = ResGetMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetMetrics) ProtoMessage() {}

func (x *ResGetMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetMetrics.ProtoReflect.Descriptor instead.
func (*ResGetMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{3}
}

func (x *ResGetMetrics) GetMetricValues() []*commonproto.MetricValue {
	if x != nil {
		return x.MetricValues
	}
	return nil
}

var File_proto_services_provider_proto protoreflect.FileDescriptor

var file_proto_services_provider_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22,
	0xbc, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x2b, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x4e,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x3d, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0x83,
	0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6e, 0x67,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x73, 0x6e, 0x69, 0x78, 0x2f, 0x61, 0x69, 0x2d, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_provider_proto_rawDescOnce sync.Once
	file_proto_services_provider_proto_rawDescData = file_proto_services_provider_proto_rawDesc
)

func file_proto_services_provider_proto_rawDescGZIP() []byte {
	file_proto_services_provider_proto_rawDescOnce.Do(func() {
		file_proto_services_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_provider_proto_rawDescData)
	})
	return file_proto_services_provider_proto_rawDescData
}

var file_proto_services_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_services_provider_proto_goTypes = []interface{}{
	(*ReqPing)(nil),                 // 0: services.ReqPing
	(*ResPong)(nil),                 // 1: services.ResPong
	(*ReqGetMetrics)(nil),           // 2: services.ReqGetMetrics
	(*ResGetMetrics)(nil),           // 3: services.ResGetMetrics
	(*commonproto.Ping)(nil),        // 4: commonproto.Ping
	(*commonproto.Pong)(nil),        // 5: commonproto.Pong
	(enums.ResourceType)(0),         // 6: enums.ResourceType
	(*duration.Duration)(nil),       // 7: google.protobuf.Duration
	(*events.History)(nil),          // 8: events.History
	(*commonproto.MetricValue)(nil), // 9: commonproto.MetricValue
}
var file_proto_services_provider_proto_depIdxs = []int32{
	4, // 0: services.ReqPing.ping:type_name -> commonproto.Ping
	5, // 1: services.ResPong.pong:type_name -> commonproto.Pong
	6, // 2: services.ReqGetMetrics.resource:type_name -> enums.ResourceType
	7, // 3: services.ReqGetMetrics.stepDuration:type_name -> google.protobuf.Duration
	7, // 4: services.ReqGetMetrics.timeWindow:type_name -> google.protobuf.Duration
	8, // 5: services.ReqGetMetrics.history:type_name -> events.History
	9, // 6: services.ResGetMetrics.metric_values:type_name -> commonproto.MetricValue
	0, // 7: services.ProviderService.Ping:input_type -> services.ReqPing
	2, // 8: services.ProviderService.GetMetrics:input_type -> services.ReqGetMetrics
	1, // 9: services.ProviderService.Ping:output_type -> services.ResPong
	3, // 10: services.ProviderService.GetMetrics:output_type -> services.ResGetMetrics
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_services_provider_proto_init() }
func file_proto_services_provider_proto_init() {
	if File_proto_services_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPing); i {
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
		file_proto_services_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPong); i {
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
		file_proto_services_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetMetrics); i {
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
		file_proto_services_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetMetrics); i {
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
	file_proto_services_provider_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ReqGetMetrics_History)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_services_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_provider_proto_goTypes,
		DependencyIndexes: file_proto_services_provider_proto_depIdxs,
		MessageInfos:      file_proto_services_provider_proto_msgTypes,
	}.Build()
	File_proto_services_provider_proto = out.File
	file_proto_services_provider_proto_rawDesc = nil
	file_proto_services_provider_proto_goTypes = nil
	file_proto_services_provider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProviderServiceClient is the client API for ProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderServiceClient interface {
	Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error)
	GetMetrics(ctx context.Context, in *ReqGetMetrics, opts ...grpc.CallOption) (*ResGetMetrics, error)
}

type providerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceClient(cc grpc.ClientConnInterface) ProviderServiceClient {
	return &providerServiceClient{cc}
}

func (c *providerServiceClient) Ping(ctx context.Context, in *ReqPing, opts ...grpc.CallOption) (*ResPong, error) {
	out := new(ResPong)
	err := c.cc.Invoke(ctx, "/services.ProviderService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) GetMetrics(ctx context.Context, in *ReqGetMetrics, opts ...grpc.CallOption) (*ResGetMetrics, error) {
	out := new(ResGetMetrics)
	err := c.cc.Invoke(ctx, "/services.ProviderService/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServiceServer is the server API for ProviderService service.
type ProviderServiceServer interface {
	Ping(context.Context, *ReqPing) (*ResPong, error)
	GetMetrics(context.Context, *ReqGetMetrics) (*ResGetMetrics, error)
}

// UnimplementedProviderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServer struct {
}

func (*UnimplementedProviderServiceServer) Ping(context.Context, *ReqPing) (*ResPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedProviderServiceServer) GetMetrics(context.Context, *ReqGetMetrics) (*ResGetMetrics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func RegisterProviderServiceServer(s *grpc.Server, srv ProviderServiceServer) {
	s.RegisterService(&_ProviderService_serviceDesc, srv)
}

func _ProviderService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProviderService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).Ping(ctx, req.(*ReqPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProviderService/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).GetMetrics(ctx, req.(*ReqGetMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProviderService",
	HandlerType: (*ProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ProviderService_Ping_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _ProviderService_GetMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/provider.proto",
}
