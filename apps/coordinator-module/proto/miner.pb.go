// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: miner.proto

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

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_miner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_miner_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string          `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Bio         string          `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Following   int32           `protobuf:"varint,4,opt,name=following,proto3" json:"following,omitempty"`
	Followers   int32           `protobuf:"varint,5,opt,name=followers,proto3" json:"followers,omitempty"`
	Likes       int32           `protobuf:"varint,6,opt,name=likes,proto3" json:"likes,omitempty"`
	Url         string          `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Videos      []*VideoPreview `protobuf:"bytes,8,rep,name=videos,proto3" json:"videos,omitempty"`
	Timestamp   int64           `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_miner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_miner_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Account) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Account) GetFollowing() int32 {
	if x != nil {
		return x.Following
	}
	return 0
}

func (x *Account) GetFollowers() int32 {
	if x != nil {
		return x.Followers
	}
	return 0
}

func (x *Account) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Account) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Account) GetVideos() []*VideoPreview {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *Account) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type VideoPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Id       string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Views    int32  `protobuf:"varint,2,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *VideoPreview) Reset() {
	*x = VideoPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPreview) ProtoMessage() {}

func (x *VideoPreview) ProtoReflect() protoreflect.Message {
	mi := &file_miner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPreview.ProtoReflect.Descriptor instead.
func (*VideoPreview) Descriptor() ([]byte, []int) {
	return file_miner_proto_rawDescGZIP(), []int{2}
}

func (x *VideoPreview) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoPreview) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VideoPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VideoPreview) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

type VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *VideoRequest) Reset() {
	*x = VideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRequest) ProtoMessage() {}

func (x *VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_miner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRequest.ProtoReflect.Descriptor instead.
func (*VideoRequest) Descriptor() ([]byte, []int) {
	return file_miner_proto_rawDescGZIP(), []int{3}
}

func (x *VideoRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type VideoDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url            string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	VideoUrl       string `protobuf:"bytes,2,opt,name=videoUrl,proto3" json:"videoUrl,omitempty"`
	VideoTimestamp string `protobuf:"bytes,3,opt,name=videoTimestamp,proto3" json:"videoTimestamp,omitempty"`
	ThumbnailUrl   string `protobuf:"bytes,4,opt,name=thumbnailUrl,proto3" json:"thumbnailUrl,omitempty"`
	Views          int32  `protobuf:"varint,5,opt,name=views,proto3" json:"views,omitempty"`
	Likes          int32  `protobuf:"varint,6,opt,name=likes,proto3" json:"likes,omitempty"`
	Comments       int32  `protobuf:"varint,7,opt,name=comments,proto3" json:"comments,omitempty"`
	Shares         int32  `protobuf:"varint,8,opt,name=shares,proto3" json:"shares,omitempty"`
	AudioName      string `protobuf:"bytes,9,opt,name=audioName,proto3" json:"audioName,omitempty"`
	Description    string `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Timestamp      int64  `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *VideoDetails) Reset() {
	*x = VideoDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_miner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoDetails) ProtoMessage() {}

func (x *VideoDetails) ProtoReflect() protoreflect.Message {
	mi := &file_miner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoDetails.ProtoReflect.Descriptor instead.
func (*VideoDetails) Descriptor() ([]byte, []int) {
	return file_miner_proto_rawDescGZIP(), []int{4}
}

func (x *VideoDetails) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoDetails) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoDetails) GetVideoTimestamp() string {
	if x != nil {
		return x.VideoTimestamp
	}
	return ""
}

func (x *VideoDetails) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *VideoDetails) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *VideoDetails) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *VideoDetails) GetComments() int32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *VideoDetails) GetShares() int32 {
	if x != nil {
		return x.Shares
	}
	return 0
}

func (x *VideoDetails) GetAudioName() string {
	if x != nil {
		return x.AudioName
	}
	return ""
}

func (x *VideoDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VideoDetails) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_miner_proto protoreflect.FileDescriptor

var file_miner_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x69, 0x6e, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x62, 0x0a,
	0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x22, 0x20, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0xc6, 0x02, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x7d, 0x0a, 0x05,
	0x4d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x13, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_miner_proto_rawDescOnce sync.Once
	file_miner_proto_rawDescData = file_miner_proto_rawDesc
)

func file_miner_proto_rawDescGZIP() []byte {
	file_miner_proto_rawDescOnce.Do(func() {
		file_miner_proto_rawDescData = protoimpl.X.CompressGZIP(file_miner_proto_rawDescData)
	})
	return file_miner_proto_rawDescData
}

var file_miner_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_miner_proto_goTypes = []interface{}{
	(*AccountRequest)(nil), // 0: miner.AccountRequest
	(*Account)(nil),        // 1: miner.Account
	(*VideoPreview)(nil),   // 2: miner.VideoPreview
	(*VideoRequest)(nil),   // 3: miner.VideoRequest
	(*VideoDetails)(nil),   // 4: miner.VideoDetails
}
var file_miner_proto_depIdxs = []int32{
	2, // 0: miner.Account.videos:type_name -> miner.VideoPreview
	0, // 1: miner.Miner.GetAccount:input_type -> miner.AccountRequest
	3, // 2: miner.Miner.GetVideoDetails:input_type -> miner.VideoRequest
	1, // 3: miner.Miner.GetAccount:output_type -> miner.Account
	4, // 4: miner.Miner.GetVideoDetails:output_type -> miner.VideoDetails
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_miner_proto_init() }
func file_miner_proto_init() {
	if File_miner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_miner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequest); i {
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
		file_miner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_miner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPreview); i {
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
		file_miner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoRequest); i {
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
		file_miner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoDetails); i {
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
			RawDescriptor: file_miner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_miner_proto_goTypes,
		DependencyIndexes: file_miner_proto_depIdxs,
		MessageInfos:      file_miner_proto_msgTypes,
	}.Build()
	File_miner_proto = out.File
	file_miner_proto_rawDesc = nil
	file_miner_proto_goTypes = nil
	file_miner_proto_depIdxs = nil
}
