// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: lesson.proto

package schedule_service

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

type LessonPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LessonPrimaryKey) Reset() {
	*x = LessonPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LessonPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LessonPrimaryKey) ProtoMessage() {}

func (x *LessonPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LessonPrimaryKey.ProtoReflect.Descriptor instead.
func (*LessonPrimaryKey) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{0}
}

func (x *LessonPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LSMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LSMessage) Reset() {
	*x = LSMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LSMessage) ProtoMessage() {}

func (x *LSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LSMessage.ProtoReflect.Descriptor instead.
func (*LSMessage) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{1}
}

func (x *LSMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateLesson) Reset() {
	*x = CreateLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLesson) ProtoMessage() {}

func (x *CreateLesson) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLesson.ProtoReflect.Descriptor instead.
func (*CreateLesson) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{3}
}

func (x *Lesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lesson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lesson) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Lesson) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Lesson) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type UpdateLessonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateLessonRequest) Reset() {
	*x = UpdateLessonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLessonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLessonRequest) ProtoMessage() {}

func (x *UpdateLessonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLessonRequest.ProtoReflect.Descriptor instead.
func (*UpdateLessonRequest) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLessonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateLessonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListLessonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListLessonRequest) Reset() {
	*x = GetListLessonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListLessonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListLessonRequest) ProtoMessage() {}

func (x *GetListLessonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListLessonRequest.ProtoReflect.Descriptor instead.
func (*GetListLessonRequest) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{5}
}

func (x *GetListLessonRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListLessonRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListLessonRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListLessonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Lesson []*Lesson `protobuf:"bytes,2,rep,name=Lesson,proto3" json:"Lesson,omitempty"`
}

func (x *GetListLessonResponse) Reset() {
	*x = GetListLessonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListLessonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListLessonResponse) ProtoMessage() {}

func (x *GetListLessonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListLessonResponse.ProtoReflect.Descriptor instead.
func (*GetListLessonResponse) Descriptor() ([]byte, []int) {
	return file_lesson_proto_rawDescGZIP(), []int{6}
}

func (x *GetListLessonResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListLessonResponse) GetLesson() []*Lesson {
	if x != nil {
		return x.Lesson
	}
	return nil
}

var File_lesson_proto protoreflect.FileDescriptor

var file_lesson_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x22, 0x0a, 0x10, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x09, 0x4c, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x89, 0x01, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x5f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x32, 0xa5, 0x03, 0x0a, 0x0d, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x1a, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1b,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a,
	0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_lesson_proto_rawDescOnce sync.Once
	file_lesson_proto_rawDescData = file_lesson_proto_rawDesc
)

func file_lesson_proto_rawDescGZIP() []byte {
	file_lesson_proto_rawDescOnce.Do(func() {
		file_lesson_proto_rawDescData = protoimpl.X.CompressGZIP(file_lesson_proto_rawDescData)
	})
	return file_lesson_proto_rawDescData
}

var file_lesson_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lesson_proto_goTypes = []interface{}{
	(*LessonPrimaryKey)(nil),      // 0: schedule_service.LessonPrimaryKey
	(*LSMessage)(nil),             // 1: schedule_service.LSMessage
	(*CreateLesson)(nil),          // 2: schedule_service.CreateLesson
	(*Lesson)(nil),                // 3: schedule_service.Lesson
	(*UpdateLessonRequest)(nil),   // 4: schedule_service.UpdateLessonRequest
	(*GetListLessonRequest)(nil),  // 5: schedule_service.GetListLessonRequest
	(*GetListLessonResponse)(nil), // 6: schedule_service.GetListLessonResponse
}
var file_lesson_proto_depIdxs = []int32{
	3, // 0: schedule_service.GetListLessonResponse.Lesson:type_name -> schedule_service.Lesson
	2, // 1: schedule_service.LessonService.Create:input_type -> schedule_service.CreateLesson
	0, // 2: schedule_service.LessonService.GetByID:input_type -> schedule_service.LessonPrimaryKey
	5, // 3: schedule_service.LessonService.GetList:input_type -> schedule_service.GetListLessonRequest
	4, // 4: schedule_service.LessonService.Update:input_type -> schedule_service.UpdateLessonRequest
	0, // 5: schedule_service.LessonService.Delete:input_type -> schedule_service.LessonPrimaryKey
	0, // 6: schedule_service.LessonService.Create:output_type -> schedule_service.LessonPrimaryKey
	3, // 7: schedule_service.LessonService.GetByID:output_type -> schedule_service.Lesson
	6, // 8: schedule_service.LessonService.GetList:output_type -> schedule_service.GetListLessonResponse
	1, // 9: schedule_service.LessonService.Update:output_type -> schedule_service.LSMessage
	1, // 10: schedule_service.LessonService.Delete:output_type -> schedule_service.LSMessage
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_lesson_proto_init() }
func file_lesson_proto_init() {
	if File_lesson_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lesson_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LessonPrimaryKey); i {
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
		file_lesson_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LSMessage); i {
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
		file_lesson_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLesson); i {
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
		file_lesson_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
		file_lesson_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLessonRequest); i {
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
		file_lesson_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListLessonRequest); i {
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
		file_lesson_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListLessonResponse); i {
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
			RawDescriptor: file_lesson_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lesson_proto_goTypes,
		DependencyIndexes: file_lesson_proto_depIdxs,
		MessageInfos:      file_lesson_proto_msgTypes,
	}.Build()
	File_lesson_proto = out.File
	file_lesson_proto_rawDesc = nil
	file_lesson_proto_goTypes = nil
	file_lesson_proto_depIdxs = nil
}
