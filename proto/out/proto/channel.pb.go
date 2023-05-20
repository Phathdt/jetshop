// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/channel.proto

package jetshop_proto

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

type HermesChannelCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelCode  string `protobuf:"bytes,1,opt,name=channel_code,json=channelCode,proto3" json:"channel_code,omitempty"`
	PlatformCode string `protobuf:"bytes,2,opt,name=platform_code,json=platformCode,proto3" json:"platform_code,omitempty"`
	IsEnabled    bool   `protobuf:"varint,3,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	SellerId     string `protobuf:"bytes,4,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
}

func (x *HermesChannelCredential) Reset() {
	*x = HermesChannelCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HermesChannelCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HermesChannelCredential) ProtoMessage() {}

func (x *HermesChannelCredential) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HermesChannelCredential.ProtoReflect.Descriptor instead.
func (*HermesChannelCredential) Descriptor() ([]byte, []int) {
	return file_proto_channel_proto_rawDescGZIP(), []int{0}
}

func (x *HermesChannelCredential) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

func (x *HermesChannelCredential) GetPlatformCode() string {
	if x != nil {
		return x.PlatformCode
	}
	return ""
}

func (x *HermesChannelCredential) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *HermesChannelCredential) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

type ChannelListHermesCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnabled bool `protobuf:"varint,1,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
}

func (x *ChannelListHermesCredentialRequest) Reset() {
	*x = ChannelListHermesCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelListHermesCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelListHermesCredentialRequest) ProtoMessage() {}

func (x *ChannelListHermesCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelListHermesCredentialRequest.ProtoReflect.Descriptor instead.
func (*ChannelListHermesCredentialRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelListHermesCredentialRequest) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

type ChannelListHermesCredentialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creds []*HermesChannelCredential `protobuf:"bytes,1,rep,name=creds,proto3" json:"creds,omitempty"`
}

func (x *ChannelListHermesCredentialResponse) Reset() {
	*x = ChannelListHermesCredentialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelListHermesCredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelListHermesCredentialResponse) ProtoMessage() {}

func (x *ChannelListHermesCredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelListHermesCredentialResponse.ProtoReflect.Descriptor instead.
func (*ChannelListHermesCredentialResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelListHermesCredentialResponse) GetCreds() []*HermesChannelCredential {
	if x != nil {
		return x.Creds
	}
	return nil
}

type ChannelGetHermesCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelCode string `protobuf:"bytes,1,opt,name=channel_code,json=channelCode,proto3" json:"channel_code,omitempty"`
}

func (x *ChannelGetHermesCredentialRequest) Reset() {
	*x = ChannelGetHermesCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelGetHermesCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGetHermesCredentialRequest) ProtoMessage() {}

func (x *ChannelGetHermesCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGetHermesCredentialRequest.ProtoReflect.Descriptor instead.
func (*ChannelGetHermesCredentialRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelGetHermesCredentialRequest) GetChannelCode() string {
	if x != nil {
		return x.ChannelCode
	}
	return ""
}

type ChannelGetHermesCredentialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cred *HermesChannelCredential `protobuf:"bytes,1,opt,name=cred,proto3" json:"cred,omitempty"`
}

func (x *ChannelGetHermesCredentialResponse) Reset() {
	*x = ChannelGetHermesCredentialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelGetHermesCredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelGetHermesCredentialResponse) ProtoMessage() {}

func (x *ChannelGetHermesCredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelGetHermesCredentialResponse.ProtoReflect.Descriptor instead.
func (*ChannelGetHermesCredentialResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelGetHermesCredentialResponse) GetCred() *HermesChannelCredential {
	if x != nil {
		return x.Cred
	}
	return nil
}

var File_proto_channel_proto protoreflect.FileDescriptor

var file_proto_channel_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x17, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x22,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65,
	0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x5b, 0x0a, 0x23, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x63, 0x72, 0x65, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x05, 0x63, 0x72, 0x65, 0x64, 0x73, 0x22, 0x46,
	0x0a, 0x21, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6d,
	0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x58, 0x0a, 0x22, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x63, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x04, 0x63, 0x72, 0x65, 0x64,
	0x32, 0xfd, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6d,
	0x65, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x6a, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_channel_proto_rawDescOnce sync.Once
	file_proto_channel_proto_rawDescData = file_proto_channel_proto_rawDesc
)

func file_proto_channel_proto_rawDescGZIP() []byte {
	file_proto_channel_proto_rawDescOnce.Do(func() {
		file_proto_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_channel_proto_rawDescData)
	})
	return file_proto_channel_proto_rawDescData
}

var file_proto_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_channel_proto_goTypes = []interface{}{
	(*HermesChannelCredential)(nil),             // 0: proto.HermesChannelCredential
	(*ChannelListHermesCredentialRequest)(nil),  // 1: proto.ChannelListHermesCredentialRequest
	(*ChannelListHermesCredentialResponse)(nil), // 2: proto.ChannelListHermesCredentialResponse
	(*ChannelGetHermesCredentialRequest)(nil),   // 3: proto.ChannelGetHermesCredentialRequest
	(*ChannelGetHermesCredentialResponse)(nil),  // 4: proto.ChannelGetHermesCredentialResponse
}
var file_proto_channel_proto_depIdxs = []int32{
	0, // 0: proto.ChannelListHermesCredentialResponse.creds:type_name -> proto.HermesChannelCredential
	0, // 1: proto.ChannelGetHermesCredentialResponse.cred:type_name -> proto.HermesChannelCredential
	1, // 2: proto.ChannelService.ListHermesChannelCredential:input_type -> proto.ChannelListHermesCredentialRequest
	3, // 3: proto.ChannelService.GetHermesChannelCredential:input_type -> proto.ChannelGetHermesCredentialRequest
	2, // 4: proto.ChannelService.ListHermesChannelCredential:output_type -> proto.ChannelListHermesCredentialResponse
	4, // 5: proto.ChannelService.GetHermesChannelCredential:output_type -> proto.ChannelGetHermesCredentialResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_channel_proto_init() }
func file_proto_channel_proto_init() {
	if File_proto_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HermesChannelCredential); i {
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
		file_proto_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelListHermesCredentialRequest); i {
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
		file_proto_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelListHermesCredentialResponse); i {
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
		file_proto_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelGetHermesCredentialRequest); i {
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
		file_proto_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelGetHermesCredentialResponse); i {
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
			RawDescriptor: file_proto_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_channel_proto_goTypes,
		DependencyIndexes: file_proto_channel_proto_depIdxs,
		MessageInfos:      file_proto_channel_proto_msgTypes,
	}.Build()
	File_proto_channel_proto = out.File
	file_proto_channel_proto_rawDesc = nil
	file_proto_channel_proto_goTypes = nil
	file_proto_channel_proto_depIdxs = nil
}
