// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: banners-rotator.proto

package api

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

type BannerOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId   int64 `protobuf:"varint,1,opt,name=slotId,proto3" json:"slotId,omitempty"`
	BannerId int64 `protobuf:"varint,2,opt,name=bannerId,proto3" json:"bannerId,omitempty"`
}

func (x *BannerOperationRequest) Reset() {
	*x = BannerOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banners_rotator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerOperationRequest) ProtoMessage() {}

func (x *BannerOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banners_rotator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerOperationRequest.ProtoReflect.Descriptor instead.
func (*BannerOperationRequest) Descriptor() ([]byte, []int) {
	return file_banners_rotator_proto_rawDescGZIP(), []int{0}
}

func (x *BannerOperationRequest) GetSlotId() int64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *BannerOperationRequest) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

type BannerOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BannerOperationResponse) Reset() {
	*x = BannerOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banners_rotator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerOperationResponse) ProtoMessage() {}

func (x *BannerOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banners_rotator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerOperationResponse.ProtoReflect.Descriptor instead.
func (*BannerOperationResponse) Descriptor() ([]byte, []int) {
	return file_banners_rotator_proto_rawDescGZIP(), []int{1}
}

func (x *BannerOperationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CountBannerClickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req        *BannerOperationRequest `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	DemGroupId int64                   `protobuf:"varint,2,opt,name=demGroupId,proto3" json:"demGroupId,omitempty"`
}

func (x *CountBannerClickRequest) Reset() {
	*x = CountBannerClickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banners_rotator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountBannerClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountBannerClickRequest) ProtoMessage() {}

func (x *CountBannerClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banners_rotator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountBannerClickRequest.ProtoReflect.Descriptor instead.
func (*CountBannerClickRequest) Descriptor() ([]byte, []int) {
	return file_banners_rotator_proto_rawDescGZIP(), []int{2}
}

func (x *CountBannerClickRequest) GetReq() *BannerOperationRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

func (x *CountBannerClickRequest) GetDemGroupId() int64 {
	if x != nil {
		return x.DemGroupId
	}
	return 0
}

type SelectBannerForShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId     int64 `protobuf:"varint,1,opt,name=slotId,proto3" json:"slotId,omitempty"`
	DemGroupId int64 `protobuf:"varint,2,opt,name=demGroupId,proto3" json:"demGroupId,omitempty"`
}

func (x *SelectBannerForShowRequest) Reset() {
	*x = SelectBannerForShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banners_rotator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectBannerForShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectBannerForShowRequest) ProtoMessage() {}

func (x *SelectBannerForShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banners_rotator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectBannerForShowRequest.ProtoReflect.Descriptor instead.
func (*SelectBannerForShowRequest) Descriptor() ([]byte, []int) {
	return file_banners_rotator_proto_rawDescGZIP(), []int{3}
}

func (x *SelectBannerForShowRequest) GetSlotId() int64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *SelectBannerForShowRequest) GetDemGroupId() int64 {
	if x != nil {
		return x.DemGroupId
	}
	return 0
}

type SelectBannerForShowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int64 `protobuf:"varint,1,opt,name=bannerId,proto3" json:"bannerId,omitempty"`
}

func (x *SelectBannerForShowResponse) Reset() {
	*x = SelectBannerForShowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banners_rotator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectBannerForShowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectBannerForShowResponse) ProtoMessage() {}

func (x *SelectBannerForShowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banners_rotator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectBannerForShowResponse.ProtoReflect.Descriptor instead.
func (*SelectBannerForShowResponse) Descriptor() ([]byte, []int) {
	return file_banners_rotator_proto_rawDescGZIP(), []int{4}
}

func (x *SelectBannerForShowResponse) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

var File_banners_rotator_proto protoreflect.FileDescriptor

var file_banners_rotator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c,
	0x0a, 0x16, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x17,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x6a, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x03,
	0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x64, 0x65, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x1a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6c, 0x6f,
	0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6d, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x1b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x32, 0xeb,
	0x02, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x50, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c,
	0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f,
	0x72, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x68, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banners_rotator_proto_rawDescOnce sync.Once
	file_banners_rotator_proto_rawDescData = file_banners_rotator_proto_rawDesc
)

func file_banners_rotator_proto_rawDescGZIP() []byte {
	file_banners_rotator_proto_rawDescOnce.Do(func() {
		file_banners_rotator_proto_rawDescData = protoimpl.X.CompressGZIP(file_banners_rotator_proto_rawDescData)
	})
	return file_banners_rotator_proto_rawDescData
}

var file_banners_rotator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_banners_rotator_proto_goTypes = []interface{}{
	(*BannerOperationRequest)(nil),      // 0: proto.BannerOperationRequest
	(*BannerOperationResponse)(nil),     // 1: proto.BannerOperationResponse
	(*CountBannerClickRequest)(nil),     // 2: proto.CountBannerClickRequest
	(*SelectBannerForShowRequest)(nil),  // 3: proto.SelectBannerForShowRequest
	(*SelectBannerForShowResponse)(nil), // 4: proto.SelectBannerForShowResponse
}
var file_banners_rotator_proto_depIdxs = []int32{
	0, // 0: proto.CountBannerClickRequest.req:type_name -> proto.BannerOperationRequest
	0, // 1: proto.BannersRotator.AddBannerToSlot:input_type -> proto.BannerOperationRequest
	0, // 2: proto.BannersRotator.RemoveBannerFromSlot:input_type -> proto.BannerOperationRequest
	2, // 3: proto.BannersRotator.CountBannerClick:input_type -> proto.CountBannerClickRequest
	3, // 4: proto.BannersRotator.SelectBannerForShow:input_type -> proto.SelectBannerForShowRequest
	1, // 5: proto.BannersRotator.AddBannerToSlot:output_type -> proto.BannerOperationResponse
	1, // 6: proto.BannersRotator.RemoveBannerFromSlot:output_type -> proto.BannerOperationResponse
	1, // 7: proto.BannersRotator.CountBannerClick:output_type -> proto.BannerOperationResponse
	4, // 8: proto.BannersRotator.SelectBannerForShow:output_type -> proto.SelectBannerForShowResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_banners_rotator_proto_init() }
func file_banners_rotator_proto_init() {
	if File_banners_rotator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banners_rotator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerOperationRequest); i {
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
		file_banners_rotator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerOperationResponse); i {
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
		file_banners_rotator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountBannerClickRequest); i {
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
		file_banners_rotator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectBannerForShowRequest); i {
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
		file_banners_rotator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectBannerForShowResponse); i {
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
			RawDescriptor: file_banners_rotator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banners_rotator_proto_goTypes,
		DependencyIndexes: file_banners_rotator_proto_depIdxs,
		MessageInfos:      file_banners_rotator_proto_msgTypes,
	}.Build()
	File_banners_rotator_proto = out.File
	file_banners_rotator_proto_rawDesc = nil
	file_banners_rotator_proto_goTypes = nil
	file_banners_rotator_proto_depIdxs = nil
}