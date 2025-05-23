// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: api/review/v1/review.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        int64                  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID       int64                  `protobuf:"varint,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	StoreID       int64                  `protobuf:"varint,10,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Score         int32                  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	ServiceScore  int32                  `protobuf:"varint,4,opt,name=serviceScore,proto3" json:"serviceScore,omitempty"`
	ExpressScore  int32                  `protobuf:"varint,5,opt,name=expressScore,proto3" json:"expressScore,omitempty"`
	Content       string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	PicInfo       string                 `protobuf:"bytes,7,opt,name=picInfo,proto3" json:"picInfo,omitempty"`
	VideoInfo     string                 `protobuf:"bytes,8,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	Anonymous     bool                   `protobuf:"varint,9,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReviewRequest) Reset() {
	*x = CreateReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewRequest) ProtoMessage() {}

func (x *CreateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReviewRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateReviewRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *CreateReviewRequest) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *CreateReviewRequest) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CreateReviewRequest) GetServiceScore() int32 {
	if x != nil {
		return x.ServiceScore
	}
	return 0
}

func (x *CreateReviewRequest) GetExpressScore() int32 {
	if x != nil {
		return x.ExpressScore
	}
	return 0
}

func (x *CreateReviewRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateReviewRequest) GetPicInfo() string {
	if x != nil {
		return x.PicInfo
	}
	return ""
}

func (x *CreateReviewRequest) GetVideoInfo() string {
	if x != nil {
		return x.VideoInfo
	}
	return ""
}

func (x *CreateReviewRequest) GetAnonymous() bool {
	if x != nil {
		return x.Anonymous
	}
	return false
}

type CreateReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReviewReply) Reset() {
	*x = CreateReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewReply) ProtoMessage() {}

func (x *CreateReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewReply.ProtoReflect.Descriptor instead.
func (*CreateReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReviewReply) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

type UpdateReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReviewRequest) Reset() {
	*x = UpdateReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewRequest) ProtoMessage() {}

func (x *UpdateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{2}
}

type UpdateReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReviewReply) Reset() {
	*x = UpdateReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewReply) ProtoMessage() {}

func (x *UpdateReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewReply.ProtoReflect.Descriptor instead.
func (*UpdateReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{3}
}

type DeleteReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteReviewRequest) Reset() {
	*x = DeleteReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewRequest) ProtoMessage() {}

func (x *DeleteReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewRequest.ProtoReflect.Descriptor instead.
func (*DeleteReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{4}
}

type DeleteReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteReviewReply) Reset() {
	*x = DeleteReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewReply) ProtoMessage() {}

func (x *DeleteReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewReply.ProtoReflect.Descriptor instead.
func (*DeleteReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{5}
}

type GetReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewRequest) Reset() {
	*x = GetReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewRequest) ProtoMessage() {}

func (x *GetReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewRequest.ProtoReflect.Descriptor instead.
func (*GetReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{6}
}

func (x *GetReviewRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

type GetReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,11,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	UserID        int64                  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID       int64                  `protobuf:"varint,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	StoreID       int64                  `protobuf:"varint,10,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Score         int32                  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	ServiceScore  int32                  `protobuf:"varint,4,opt,name=serviceScore,proto3" json:"serviceScore,omitempty"`
	ExpressScore  int32                  `protobuf:"varint,5,opt,name=expressScore,proto3" json:"expressScore,omitempty"`
	Content       string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	PicInfo       string                 `protobuf:"bytes,7,opt,name=picInfo,proto3" json:"picInfo,omitempty"`
	VideoInfo     string                 `protobuf:"bytes,8,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	Anonymous     bool                   `protobuf:"varint,9,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewReply) Reset() {
	*x = GetReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewReply) ProtoMessage() {}

func (x *GetReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewReply.ProtoReflect.Descriptor instead.
func (*GetReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{7}
}

func (x *GetReviewReply) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *GetReviewReply) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetReviewReply) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *GetReviewReply) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *GetReviewReply) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GetReviewReply) GetServiceScore() int32 {
	if x != nil {
		return x.ServiceScore
	}
	return 0
}

func (x *GetReviewReply) GetExpressScore() int32 {
	if x != nil {
		return x.ExpressScore
	}
	return 0
}

func (x *GetReviewReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetReviewReply) GetPicInfo() string {
	if x != nil {
		return x.PicInfo
	}
	return ""
}

func (x *GetReviewReply) GetVideoInfo() string {
	if x != nil {
		return x.VideoInfo
	}
	return ""
}

func (x *GetReviewReply) GetAnonymous() bool {
	if x != nil {
		return x.Anonymous
	}
	return false
}

type ListReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListReviewRequest) Reset() {
	*x = ListReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewRequest) ProtoMessage() {}

func (x *ListReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewRequest.ProtoReflect.Descriptor instead.
func (*ListReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{8}
}

type ListReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListReviewReply) Reset() {
	*x = ListReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReviewReply) ProtoMessage() {}

func (x *ListReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReviewReply.ProtoReflect.Descriptor instead.
func (*ListReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{9}
}

type ReplyReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	StoreID       int64                  `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PicInfo       string                 `protobuf:"bytes,4,opt,name=picInfo,proto3" json:"picInfo,omitempty"`
	VideoInfo     string                 `protobuf:"bytes,5,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyReviewRequest) Reset() {
	*x = ReplyReviewRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyReviewRequest) ProtoMessage() {}

func (x *ReplyReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyReviewRequest.ProtoReflect.Descriptor instead.
func (*ReplyReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{10}
}

func (x *ReplyReviewRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *ReplyReviewRequest) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *ReplyReviewRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReplyReviewRequest) GetPicInfo() string {
	if x != nil {
		return x.PicInfo
	}
	return ""
}

func (x *ReplyReviewRequest) GetVideoInfo() string {
	if x != nil {
		return x.VideoInfo
	}
	return ""
}

type ReplyReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyID       int64                  `protobuf:"varint,1,opt,name=replyID,proto3" json:"replyID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyReviewReply) Reset() {
	*x = ReplyReviewReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyReviewReply) ProtoMessage() {}

func (x *ReplyReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyReviewReply.ProtoReflect.Descriptor instead.
func (*ReplyReviewReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{11}
}

func (x *ReplyReviewReply) GetReplyID() int64 {
	if x != nil {
		return x.ReplyID
	}
	return 0
}

type CreateAppealRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	StoreID       int64                  `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Reason        string                 `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	PicInfo       string                 `protobuf:"bytes,5,opt,name=picInfo,proto3" json:"picInfo,omitempty"`
	VideoInfo     string                 `protobuf:"bytes,6,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAppealRequest) Reset() {
	*x = CreateAppealRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAppealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppealRequest) ProtoMessage() {}

func (x *CreateAppealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppealRequest.ProtoReflect.Descriptor instead.
func (*CreateAppealRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{12}
}

func (x *CreateAppealRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *CreateAppealRequest) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *CreateAppealRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CreateAppealRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateAppealRequest) GetPicInfo() string {
	if x != nil {
		return x.PicInfo
	}
	return ""
}

func (x *CreateAppealRequest) GetVideoInfo() string {
	if x != nil {
		return x.VideoInfo
	}
	return ""
}

type CreateAppealReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppealID      int64                  `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAppealReply) Reset() {
	*x = CreateAppealReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAppealReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppealReply) ProtoMessage() {}

func (x *CreateAppealReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppealReply.ProtoReflect.Descriptor instead.
func (*CreateAppealReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{13}
}

func (x *CreateAppealReply) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

type OperateAppealRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppealID      int64                  `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	StoreID       int64                  `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID,omitempty"`
	ReviewID      int64                  `protobuf:"varint,3,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	Status        int32                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	OpRemark      string                 `protobuf:"bytes,5,opt,name=opRemark,proto3" json:"opRemark,omitempty"`
	OpUser        string                 `protobuf:"bytes,6,opt,name=opUser,proto3" json:"opUser,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperateAppealRequest) Reset() {
	*x = OperateAppealRequest{}
	mi := &file_api_review_v1_review_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperateAppealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateAppealRequest) ProtoMessage() {}

func (x *OperateAppealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateAppealRequest.ProtoReflect.Descriptor instead.
func (*OperateAppealRequest) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{14}
}

func (x *OperateAppealRequest) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

func (x *OperateAppealRequest) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *OperateAppealRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *OperateAppealRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OperateAppealRequest) GetOpRemark() string {
	if x != nil {
		return x.OpRemark
	}
	return ""
}

func (x *OperateAppealRequest) GetOpUser() string {
	if x != nil {
		return x.OpUser
	}
	return ""
}

type OperateAppealReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppealID      int64                  `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperateAppealReply) Reset() {
	*x = OperateAppealReply{}
	mi := &file_api_review_v1_review_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperateAppealReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateAppealReply) ProtoMessage() {}

func (x *OperateAppealReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_review_v1_review_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateAppealReply.ProtoReflect.Descriptor instead.
func (*OperateAppealReply) Descriptor() ([]byte, []int) {
	return file_api_review_v1_review_proto_rawDescGZIP(), []int{15}
}

func (x *OperateAppealReply) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

var File_api_review_v1_review_proto protoreflect.FileDescriptor

const file_api_review_v1_review_proto_rawDesc = "" +
	"\n" +
	"\x1aapi/review/v1/review.proto\x12\rapi.review.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"\x86\x03\n" +
	"\x13CreateReviewRequest\x12\x1f\n" +
	"\x06userID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x06userID\x12!\n" +
	"\aorderID\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\aorderID\x12!\n" +
	"\astoreID\x18\n" +
	" \x01(\x03B\a\xfaB\x04\"\x02 \x00R\astoreID\x12%\n" +
	"\x05score\x18\x03 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\x05score\x123\n" +
	"\fserviceScore\x18\x04 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\fserviceScore\x123\n" +
	"\fexpressScore\x18\x05 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\fexpressScore\x12!\n" +
	"\acontent\x18\x06 \x01(\tB\a\xfaB\x04r\x02\x10\n" +
	"R\acontent\x12\x18\n" +
	"\apicInfo\x18\a \x01(\tR\apicInfo\x12\x1c\n" +
	"\tvideoInfo\x18\b \x01(\tR\tvideoInfo\x12\x1c\n" +
	"\tanonymous\x18\t \x01(\bR\tanonymous\"/\n" +
	"\x11CreateReviewReply\x12\x1a\n" +
	"\breviewID\x18\x01 \x01(\x03R\breviewID\"\x15\n" +
	"\x13UpdateReviewRequest\"\x13\n" +
	"\x11UpdateReviewReply\"\x15\n" +
	"\x13DeleteReviewRequest\"\x13\n" +
	"\x11DeleteReviewReply\"7\n" +
	"\x10GetReviewRequest\x12#\n" +
	"\breviewID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\breviewID\"\x9d\x03\n" +
	"\x0eGetReviewReply\x12\x1a\n" +
	"\breviewID\x18\v \x01(\x03R\breviewID\x12\x1f\n" +
	"\x06userID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\x06userID\x12!\n" +
	"\aorderID\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\aorderID\x12!\n" +
	"\astoreID\x18\n" +
	" \x01(\x03B\a\xfaB\x04\"\x02 \x00R\astoreID\x12%\n" +
	"\x05score\x18\x03 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\x05score\x123\n" +
	"\fserviceScore\x18\x04 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\fserviceScore\x123\n" +
	"\fexpressScore\x18\x05 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\fexpressScore\x12!\n" +
	"\acontent\x18\x06 \x01(\tB\a\xfaB\x04r\x02\x10\n" +
	"R\acontent\x12\x18\n" +
	"\apicInfo\x18\a \x01(\tR\apicInfo\x12\x1c\n" +
	"\tvideoInfo\x18\b \x01(\tR\tvideoInfo\x12\x1c\n" +
	"\tanonymous\x18\t \x01(\bR\tanonymous\"\x13\n" +
	"\x11ListReviewRequest\"\x11\n" +
	"\x0fListReviewReply\"\xb7\x01\n" +
	"\x12ReplyReviewRequest\x12#\n" +
	"\breviewID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\breviewID\x12!\n" +
	"\astoreID\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\astoreID\x12!\n" +
	"\acontent\x18\x03 \x01(\tB\a\xfaB\x04r\x02\x10\x04R\acontent\x12\x18\n" +
	"\apicInfo\x18\x04 \x01(\tR\apicInfo\x12\x1c\n" +
	"\tvideoInfo\x18\x05 \x01(\tR\tvideoInfo\",\n" +
	"\x10ReplyReviewReply\x12\x18\n" +
	"\areplyID\x18\x01 \x01(\x03R\areplyID\"\xc7\x01\n" +
	"\x13CreateAppealRequest\x12#\n" +
	"\breviewID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\breviewID\x12!\n" +
	"\astoreID\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\astoreID\x12\x16\n" +
	"\x06reason\x18\x03 \x01(\tR\x06reason\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12\x18\n" +
	"\apicInfo\x18\x05 \x01(\tR\apicInfo\x12\x1c\n" +
	"\tvideoInfo\x18\x06 \x01(\tR\tvideoInfo\"/\n" +
	"\x11CreateAppealReply\x12\x1a\n" +
	"\bappealID\x18\x01 \x01(\x03R\bappealID\"\xe0\x01\n" +
	"\x14OperateAppealRequest\x12#\n" +
	"\bappealID\x18\x01 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\bappealID\x12!\n" +
	"\astoreID\x18\x02 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\astoreID\x12#\n" +
	"\breviewID\x18\x03 \x01(\x03B\a\xfaB\x04\"\x02 \x00R\breviewID\x12'\n" +
	"\x06status\x18\x04 \x01(\x05B\x0f\xfaB\f\x1a\n" +
	"0\x010\x020\x030\x040\x05R\x06status\x12\x1a\n" +
	"\bopRemark\x18\x05 \x01(\tR\bopRemark\x12\x16\n" +
	"\x06opUser\x18\x06 \x01(\tR\x06opUser\"0\n" +
	"\x12OperateAppealReply\x12\x1a\n" +
	"\bappealID\x18\x01 \x01(\x03R\bappealID2\x8f\x06\n" +
	"\x06Review\x12k\n" +
	"\fCreateReview\x12\".api.review.v1.CreateReviewRequest\x1a .api.review.v1.CreateReviewReply\"\x15\x82\xd3\xe4\x93\x02\x0f:\x01*\"\n" +
	"/v1/review\x12T\n" +
	"\fUpdateReview\x12\".api.review.v1.UpdateReviewRequest\x1a .api.review.v1.UpdateReviewReply\x12T\n" +
	"\fDeleteReview\x12\".api.review.v1.DeleteReviewRequest\x1a .api.review.v1.DeleteReviewReply\x12}\n" +
	"\x0eGetReviewByRID\x12\x1f.api.review.v1.GetReviewRequest\x1a\x1d.api.review.v1.GetReviewReply\"+\x82\xd3\xe4\x93\x02%Z\x17\x12\x15/v1/review/{reviewID}\x12\n" +
	"/v1/review\x12N\n" +
	"\n" +
	"ListReview\x12 .api.review.v1.ListReviewRequest\x1a\x1e.api.review.v1.ListReviewReply\x12n\n" +
	"\vReplyReview\x12!.api.review.v1.ReplyReviewRequest\x1a\x1f.api.review.v1.ReplyReviewReply\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/v1/review/reply\x12T\n" +
	"\fCreateAppeal\x12\".api.review.v1.CreateAppealRequest\x1a .api.review.v1.CreateAppealReply\x12W\n" +
	"\rOperateAppeal\x12#.api.review.v1.OperateAppealRequest\x1a!.api.review.v1.OperateAppealReplyB2\n" +
	"\rapi.review.v1P\x01Z\x1freview-service/api/review/v1;v1b\x06proto3"

var (
	file_api_review_v1_review_proto_rawDescOnce sync.Once
	file_api_review_v1_review_proto_rawDescData []byte
)

func file_api_review_v1_review_proto_rawDescGZIP() []byte {
	file_api_review_v1_review_proto_rawDescOnce.Do(func() {
		file_api_review_v1_review_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_review_v1_review_proto_rawDesc), len(file_api_review_v1_review_proto_rawDesc)))
	})
	return file_api_review_v1_review_proto_rawDescData
}

var file_api_review_v1_review_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_review_v1_review_proto_goTypes = []any{
	(*CreateReviewRequest)(nil),  // 0: api.review.v1.CreateReviewRequest
	(*CreateReviewReply)(nil),    // 1: api.review.v1.CreateReviewReply
	(*UpdateReviewRequest)(nil),  // 2: api.review.v1.UpdateReviewRequest
	(*UpdateReviewReply)(nil),    // 3: api.review.v1.UpdateReviewReply
	(*DeleteReviewRequest)(nil),  // 4: api.review.v1.DeleteReviewRequest
	(*DeleteReviewReply)(nil),    // 5: api.review.v1.DeleteReviewReply
	(*GetReviewRequest)(nil),     // 6: api.review.v1.GetReviewRequest
	(*GetReviewReply)(nil),       // 7: api.review.v1.GetReviewReply
	(*ListReviewRequest)(nil),    // 8: api.review.v1.ListReviewRequest
	(*ListReviewReply)(nil),      // 9: api.review.v1.ListReviewReply
	(*ReplyReviewRequest)(nil),   // 10: api.review.v1.ReplyReviewRequest
	(*ReplyReviewReply)(nil),     // 11: api.review.v1.ReplyReviewReply
	(*CreateAppealRequest)(nil),  // 12: api.review.v1.CreateAppealRequest
	(*CreateAppealReply)(nil),    // 13: api.review.v1.CreateAppealReply
	(*OperateAppealRequest)(nil), // 14: api.review.v1.OperateAppealRequest
	(*OperateAppealReply)(nil),   // 15: api.review.v1.OperateAppealReply
}
var file_api_review_v1_review_proto_depIdxs = []int32{
	0,  // 0: api.review.v1.Review.CreateReview:input_type -> api.review.v1.CreateReviewRequest
	2,  // 1: api.review.v1.Review.UpdateReview:input_type -> api.review.v1.UpdateReviewRequest
	4,  // 2: api.review.v1.Review.DeleteReview:input_type -> api.review.v1.DeleteReviewRequest
	6,  // 3: api.review.v1.Review.GetReviewByRID:input_type -> api.review.v1.GetReviewRequest
	8,  // 4: api.review.v1.Review.ListReview:input_type -> api.review.v1.ListReviewRequest
	10, // 5: api.review.v1.Review.ReplyReview:input_type -> api.review.v1.ReplyReviewRequest
	12, // 6: api.review.v1.Review.CreateAppeal:input_type -> api.review.v1.CreateAppealRequest
	14, // 7: api.review.v1.Review.OperateAppeal:input_type -> api.review.v1.OperateAppealRequest
	1,  // 8: api.review.v1.Review.CreateReview:output_type -> api.review.v1.CreateReviewReply
	3,  // 9: api.review.v1.Review.UpdateReview:output_type -> api.review.v1.UpdateReviewReply
	5,  // 10: api.review.v1.Review.DeleteReview:output_type -> api.review.v1.DeleteReviewReply
	7,  // 11: api.review.v1.Review.GetReviewByRID:output_type -> api.review.v1.GetReviewReply
	9,  // 12: api.review.v1.Review.ListReview:output_type -> api.review.v1.ListReviewReply
	11, // 13: api.review.v1.Review.ReplyReview:output_type -> api.review.v1.ReplyReviewReply
	13, // 14: api.review.v1.Review.CreateAppeal:output_type -> api.review.v1.CreateAppealReply
	15, // 15: api.review.v1.Review.OperateAppeal:output_type -> api.review.v1.OperateAppealReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_review_v1_review_proto_init() }
func file_api_review_v1_review_proto_init() {
	if File_api_review_v1_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_review_v1_review_proto_rawDesc), len(file_api_review_v1_review_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_review_v1_review_proto_goTypes,
		DependencyIndexes: file_api_review_v1_review_proto_depIdxs,
		MessageInfos:      file_api_review_v1_review_proto_msgTypes,
	}.Build()
	File_api_review_v1_review_proto = out.File
	file_api_review_v1_review_proto_goTypes = nil
	file_api_review_v1_review_proto_depIdxs = nil
}
