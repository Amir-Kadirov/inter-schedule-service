// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: supportteacher.proto

package user_service

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

type SupportTeacherPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SupportTeacherPrimaryKey) Reset() {
	*x = SupportTeacherPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportTeacherPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportTeacherPrimaryKey) ProtoMessage() {}

func (x *SupportTeacherPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportTeacherPrimaryKey.ProtoReflect.Descriptor instead.
func (*SupportTeacherPrimaryKey) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{0}
}

func (x *SupportTeacherPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupportTeacherGmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *SupportTeacherGmail) Reset() {
	*x = SupportTeacherGmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportTeacherGmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportTeacherGmail) ProtoMessage() {}

func (x *SupportTeacherGmail) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportTeacherGmail.ProtoReflect.Descriptor instead.
func (*SupportTeacherGmail) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{1}
}

func (x *SupportTeacherGmail) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type STMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *STMessage) Reset() {
	*x = STMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STMessage) ProtoMessage() {}

func (x *STMessage) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STMessage.ProtoReflect.Descriptor instead.
func (*STMessage) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{2}
}

func (x *STMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateSupportTeacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname           string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone              string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password           string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary             int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore         int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Ieltsattemptscount int32  `protobuf:"varint,6,opt,name=ieltsattemptscount,proto3" json:"ieltsattemptscount,omitempty"`
	Branchid           string `protobuf:"bytes,8,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email              string `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateSupportTeacher) Reset() {
	*x = CreateSupportTeacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupportTeacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupportTeacher) ProtoMessage() {}

func (x *CreateSupportTeacher) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupportTeacher.ProtoReflect.Descriptor instead.
func (*CreateSupportTeacher) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSupportTeacher) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateSupportTeacher) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateSupportTeacher) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateSupportTeacher) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateSupportTeacher) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *CreateSupportTeacher) GetIeltsattemptscount() int32 {
	if x != nil {
		return x.Ieltsattemptscount
	}
	return 0
}

func (x *CreateSupportTeacher) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *CreateSupportTeacher) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SupportTeacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname           string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone              string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password           string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary             int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore         int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Ieltsattemptscount int32  `protobuf:"varint,6,opt,name=ieltsattemptscount,proto3" json:"ieltsattemptscount,omitempty"`
	Branchid           string `protobuf:"bytes,8,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email              string `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Id                 string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt          string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt          string `protobuf:"bytes,13,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *SupportTeacher) Reset() {
	*x = SupportTeacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportTeacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportTeacher) ProtoMessage() {}

func (x *SupportTeacher) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportTeacher.ProtoReflect.Descriptor instead.
func (*SupportTeacher) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{4}
}

func (x *SupportTeacher) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *SupportTeacher) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SupportTeacher) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SupportTeacher) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *SupportTeacher) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *SupportTeacher) GetIeltsattemptscount() int32 {
	if x != nil {
		return x.Ieltsattemptscount
	}
	return 0
}

func (x *SupportTeacher) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *SupportTeacher) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SupportTeacher) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupportTeacher) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SupportTeacher) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SupportTeacher) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type UpdateSupportTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname           string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone              string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password           string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary             int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore         int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Ieltsattemptscount int32  `protobuf:"varint,6,opt,name=ieltsattemptscount,proto3" json:"ieltsattemptscount,omitempty"`
	Branchid           string `protobuf:"bytes,7,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email              string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Id                 string `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateSupportTeacherRequest) Reset() {
	*x = UpdateSupportTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupportTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupportTeacherRequest) ProtoMessage() {}

func (x *UpdateSupportTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupportTeacherRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupportTeacherRequest) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSupportTeacherRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateSupportTeacherRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateSupportTeacherRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateSupportTeacherRequest) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateSupportTeacherRequest) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *UpdateSupportTeacherRequest) GetIeltsattemptscount() int32 {
	if x != nil {
		return x.Ieltsattemptscount
	}
	return 0
}

func (x *UpdateSupportTeacherRequest) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *UpdateSupportTeacherRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateSupportTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListSupportTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListSupportTeacherRequest) Reset() {
	*x = GetListSupportTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSupportTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSupportTeacherRequest) ProtoMessage() {}

func (x *GetListSupportTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSupportTeacherRequest.ProtoReflect.Descriptor instead.
func (*GetListSupportTeacherRequest) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{6}
}

func (x *GetListSupportTeacherRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListSupportTeacherRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListSupportTeacherRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListSupportTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count          int64             `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Supportteacher []*SupportTeacher `protobuf:"bytes,2,rep,name=supportteacher,proto3" json:"supportteacher,omitempty"`
}

func (x *GetListSupportTeacherResponse) Reset() {
	*x = GetListSupportTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supportteacher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSupportTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSupportTeacherResponse) ProtoMessage() {}

func (x *GetListSupportTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supportteacher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSupportTeacherResponse.ProtoReflect.Descriptor instead.
func (*GetListSupportTeacherResponse) Descriptor() ([]byte, []int) {
	return file_supportteacher_proto_rawDescGZIP(), []int{7}
}

func (x *GetListSupportTeacherResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListSupportTeacherResponse) GetSupportteacher() []*SupportTeacher {
	if x != nil {
		return x.Supportteacher
	}
	return nil
}

var File_supportteacher_proto protoreflect.FileDescriptor

var file_supportteacher_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2b, 0x0a, 0x13, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a,
	0x09, 0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe5, 0x02, 0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e,
	0x0a, 0x12, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x69, 0x65, 0x6c, 0x74,
	0x73, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x95, 0x02,
	0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x7b, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x32, 0xa0, 0x04, 0x0a, 0x15, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x26, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69,
	0x6c, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supportteacher_proto_rawDescOnce sync.Once
	file_supportteacher_proto_rawDescData = file_supportteacher_proto_rawDesc
)

func file_supportteacher_proto_rawDescGZIP() []byte {
	file_supportteacher_proto_rawDescOnce.Do(func() {
		file_supportteacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_supportteacher_proto_rawDescData)
	})
	return file_supportteacher_proto_rawDescData
}

var file_supportteacher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_supportteacher_proto_goTypes = []interface{}{
	(*SupportTeacherPrimaryKey)(nil),      // 0: user_service.SupportTeacherPrimaryKey
	(*SupportTeacherGmail)(nil),           // 1: user_service.SupportTeacherGmail
	(*STMessage)(nil),                     // 2: user_service.STMessage
	(*CreateSupportTeacher)(nil),          // 3: user_service.CreateSupportTeacher
	(*SupportTeacher)(nil),                // 4: user_service.SupportTeacher
	(*UpdateSupportTeacherRequest)(nil),   // 5: user_service.UpdateSupportTeacherRequest
	(*GetListSupportTeacherRequest)(nil),  // 6: user_service.GetListSupportTeacherRequest
	(*GetListSupportTeacherResponse)(nil), // 7: user_service.GetListSupportTeacherResponse
}
var file_supportteacher_proto_depIdxs = []int32{
	4, // 0: user_service.GetListSupportTeacherResponse.supportteacher:type_name -> user_service.SupportTeacher
	3, // 1: user_service.SupportTeacherService.Create:input_type -> user_service.CreateSupportTeacher
	0, // 2: user_service.SupportTeacherService.GetByID:input_type -> user_service.SupportTeacherPrimaryKey
	6, // 3: user_service.SupportTeacherService.GetList:input_type -> user_service.GetListSupportTeacherRequest
	5, // 4: user_service.SupportTeacherService.Update:input_type -> user_service.UpdateSupportTeacherRequest
	0, // 5: user_service.SupportTeacherService.Delete:input_type -> user_service.SupportTeacherPrimaryKey
	1, // 6: user_service.SupportTeacherService.GetByGmail:input_type -> user_service.SupportTeacherGmail
	0, // 7: user_service.SupportTeacherService.Create:output_type -> user_service.SupportTeacherPrimaryKey
	4, // 8: user_service.SupportTeacherService.GetByID:output_type -> user_service.SupportTeacher
	7, // 9: user_service.SupportTeacherService.GetList:output_type -> user_service.GetListSupportTeacherResponse
	2, // 10: user_service.SupportTeacherService.Update:output_type -> user_service.STMessage
	2, // 11: user_service.SupportTeacherService.Delete:output_type -> user_service.STMessage
	0, // 12: user_service.SupportTeacherService.GetByGmail:output_type -> user_service.SupportTeacherPrimaryKey
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supportteacher_proto_init() }
func file_supportteacher_proto_init() {
	if File_supportteacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supportteacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportTeacherPrimaryKey); i {
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
		file_supportteacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportTeacherGmail); i {
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
		file_supportteacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STMessage); i {
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
		file_supportteacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupportTeacher); i {
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
		file_supportteacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportTeacher); i {
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
		file_supportteacher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupportTeacherRequest); i {
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
		file_supportteacher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSupportTeacherRequest); i {
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
		file_supportteacher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSupportTeacherResponse); i {
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
			RawDescriptor: file_supportteacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supportteacher_proto_goTypes,
		DependencyIndexes: file_supportteacher_proto_depIdxs,
		MessageInfos:      file_supportteacher_proto_msgTypes,
	}.Build()
	File_supportteacher_proto = out.File
	file_supportteacher_proto_rawDesc = nil
	file_supportteacher_proto_goTypes = nil
	file_supportteacher_proto_depIdxs = nil
}
