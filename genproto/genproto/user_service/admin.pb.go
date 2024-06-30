// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: admin.proto

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

type AdminPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminPrimaryKey) Reset() {
	*x = AdminPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminPrimaryKey) ProtoMessage() {}

func (x *AdminPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminPrimaryKey.ProtoReflect.Descriptor instead.
func (*AdminPrimaryKey) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AdminPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AdminGmailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail    string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AdminGmailRes) Reset() {
	*x = AdminGmailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGmailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGmailRes) ProtoMessage() {}

func (x *AdminGmailRes) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGmailRes.ProtoReflect.Descriptor instead.
func (*AdminGmailRes) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminGmailRes) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *AdminGmailRes) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminGmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *AdminGmail) Reset() {
	*x = AdminGmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGmail) ProtoMessage() {}

func (x *AdminGmail) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGmail.ProtoReflect.Descriptor instead.
func (*AdminGmail) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminGmail) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type ADMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ADMessage) Reset() {
	*x = ADMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADMessage) ProtoMessage() {}

func (x *ADMessage) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADMessage.ProtoReflect.Descriptor instead.
func (*ADMessage) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *ADMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname   string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary     int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Branchid   string `protobuf:"bytes,6,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateAdmin) Reset() {
	*x = CreateAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdmin) ProtoMessage() {}

func (x *CreateAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdmin.ProtoReflect.Descriptor instead.
func (*CreateAdmin) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAdmin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateAdmin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateAdmin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAdmin) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateAdmin) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *CreateAdmin) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *CreateAdmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname   string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary     int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Branchid   string `protobuf:"bytes,6,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Id         string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt  string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{5}
}

func (x *Admin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Admin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Admin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Admin) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Admin) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *Admin) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *Admin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Admin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Admin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Admin) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type UpdateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname   string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salary     int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	Branchid   string `protobuf:"bytes,6,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Id         string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateAdminRequest) Reset() {
	*x = UpdateAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminRequest) ProtoMessage() {}

func (x *UpdateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAdminRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateAdminRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateAdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateAdminRequest) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateAdminRequest) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *UpdateAdminRequest) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *UpdateAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateAdminRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListAdminRequest) Reset() {
	*x = GetListAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAdminRequest) ProtoMessage() {}

func (x *GetListAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAdminRequest.ProtoReflect.Descriptor instead.
func (*GetListAdminRequest) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{7}
}

func (x *GetListAdminRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListAdminRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListAdminRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Admin []*Admin `protobuf:"bytes,2,rep,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetListAdminResponse) Reset() {
	*x = GetListAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAdminResponse) ProtoMessage() {}

func (x *GetListAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAdminResponse.ProtoReflect.Descriptor instead.
func (*GetListAdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{8}
}

func (x *GetListAdminResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListAdminResponse) GetAdmin() []*Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

type AdminRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname    string `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone       string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary      int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Ieltsscore  int32  `protobuf:"varint,5,opt,name=ieltsscore,proto3" json:"ieltsscore,omitempty"`
	BranchId    string `protobuf:"bytes,6,opt,name=branchId,proto3" json:"branchId,omitempty"`
	MonhtWorked int32  `protobuf:"varint,7,opt,name=monhtWorked,proto3" json:"monhtWorked,omitempty"`
	TotalSum    int32  `protobuf:"varint,8,opt,name=totalSum,proto3" json:"totalSum,omitempty"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *AdminRep) Reset() {
	*x = AdminRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRep) ProtoMessage() {}

func (x *AdminRep) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRep.ProtoReflect.Descriptor instead.
func (*AdminRep) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{9}
}

func (x *AdminRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminRep) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *AdminRep) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AdminRep) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *AdminRep) GetIeltsscore() int32 {
	if x != nil {
		return x.Ieltsscore
	}
	return 0
}

func (x *AdminRep) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *AdminRep) GetMonhtWorked() int32 {
	if x != nil {
		return x.MonhtWorked
	}
	return 0
}

func (x *AdminRep) GetTotalSum() int32 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *AdminRep) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AdminRep) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AdminRep) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type GetRepAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Admin []*AdminRep `protobuf:"bytes,2,rep,name=admin,proto3" json:"admin,omitempty"`
}

func (x *GetRepAdminResponse) Reset() {
	*x = GetRepAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepAdminResponse) ProtoMessage() {}

func (x *GetRepAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepAdminResponse.ProtoReflect.Descriptor instead.
func (*GetRepAdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{10}
}

func (x *GetRepAdminResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetRepAdminResponse) GetAdmin() []*AdminRep {
	if x != nil {
		return x.Admin
	}
	return nil
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41,
	0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x09, 0x41, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc5, 0x01, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0xac, 0x02, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22,
	0x57, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xbb, 0x02, 0x0a, 0x08, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x6f, 0x6e, 0x68, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x68, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x32, 0x92, 0x04, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x44, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x6d, 0x61, 0x69, 0x6c,
	0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_admin_proto_goTypes = []interface{}{
	(*AdminPrimaryKey)(nil),      // 0: user_service.AdminPrimaryKey
	(*AdminGmailRes)(nil),        // 1: user_service.AdminGmailRes
	(*AdminGmail)(nil),           // 2: user_service.AdminGmail
	(*ADMessage)(nil),            // 3: user_service.ADMessage
	(*CreateAdmin)(nil),          // 4: user_service.CreateAdmin
	(*Admin)(nil),                // 5: user_service.Admin
	(*UpdateAdminRequest)(nil),   // 6: user_service.UpdateAdminRequest
	(*GetListAdminRequest)(nil),  // 7: user_service.GetListAdminRequest
	(*GetListAdminResponse)(nil), // 8: user_service.GetListAdminResponse
	(*AdminRep)(nil),             // 9: user_service.AdminRep
	(*GetRepAdminResponse)(nil),  // 10: user_service.GetRepAdminResponse
}
var file_admin_proto_depIdxs = []int32{
	5,  // 0: user_service.GetListAdminResponse.admin:type_name -> user_service.Admin
	9,  // 1: user_service.GetRepAdminResponse.admin:type_name -> user_service.AdminRep
	4,  // 2: user_service.AdminService.Create:input_type -> user_service.CreateAdmin
	0,  // 3: user_service.AdminService.GetByID:input_type -> user_service.AdminPrimaryKey
	7,  // 4: user_service.AdminService.GetList:input_type -> user_service.GetListAdminRequest
	6,  // 5: user_service.AdminService.Update:input_type -> user_service.UpdateAdminRequest
	0,  // 6: user_service.AdminService.Delete:input_type -> user_service.AdminPrimaryKey
	2,  // 7: user_service.AdminService.GetByGmail:input_type -> user_service.AdminGmail
	7,  // 8: user_service.AdminService.AdminReport:input_type -> user_service.GetListAdminRequest
	0,  // 9: user_service.AdminService.Create:output_type -> user_service.AdminPrimaryKey
	5,  // 10: user_service.AdminService.GetByID:output_type -> user_service.Admin
	8,  // 11: user_service.AdminService.GetList:output_type -> user_service.GetListAdminResponse
	3,  // 12: user_service.AdminService.Update:output_type -> user_service.ADMessage
	3,  // 13: user_service.AdminService.Delete:output_type -> user_service.ADMessage
	1,  // 14: user_service.AdminService.GetByGmail:output_type -> user_service.AdminGmailRes
	10, // 15: user_service.AdminService.AdminReport:output_type -> user_service.GetRepAdminResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminPrimaryKey); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGmailRes); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGmail); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADMessage); i {
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
		file_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdmin); i {
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
		file_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminRequest); i {
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
		file_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAdminRequest); i {
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
		file_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAdminResponse); i {
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
		file_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRep); i {
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
		file_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepAdminResponse); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
