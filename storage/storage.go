package storage

import (
	"context"
	ct "schedule_service/genproto/genproto/schedule_service"
)

type StorageI interface {
	CloseDB()
	Student() StudentRepoI
	Group() GroupRepoI
}

type StudentRepoI interface {
	Create(ctx context.Context, req *ct.CreateStudent) (resp *ct.StudentPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.StudentPrimaryKey) (resp *ct.Student, err error)
	GetList(ctx context.Context, req *ct.GetListStudentRequest) (resp *ct.GetListStudentResponse, err error)
	Update(ctx context.Context, req *ct.UpdateStudentRequest) (resp *ct.STMessage, err error)
	Delete(ctx context.Context, req *ct.StudentPrimaryKey) (resp *ct.STMessage, err error)
	GetByGmail(ctx context.Context, req *ct.StudentGmail) (*ct.StudentPrimaryKey, error)
}

type GroupRepoI interface {
	Create(ctx context.Context, req *ct.CreateGroup) (resp *ct.GroupPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.GroupPrimaryKey) (resp *ct.Group, err error)
	GetList(ctx context.Context, req *ct.GetListGroupRequest) (resp *ct.GetListGroupResponse, err error)
	Update(ctx context.Context, req *ct.UpdateGroupRequest) (resp *ct.GRMessage, err error)
	Delete(ctx context.Context, req *ct.GroupPrimaryKey) (resp *ct.GRMessage, err error)
}