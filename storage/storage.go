package storage

import (
	"context"
	ct "schedule_service/genproto/genproto/schedule_service"

	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageI interface {
	CloseDB()
	Student() StudentRepoI
	Group() GroupRepoI
	Lesson() LessonRepoI
	Task() TaskRepoI
	Journal() JournalRepoI
	Schedule() ScheduleRepoI
	GroupMany() GroupManyRepoI
	Event() EventRepoI
	Attendance() AttendanceRepoI
	GetDB() *pgxpool.Pool
}

type StudentRepoI interface {
	Create(ctx context.Context, req *ct.CreateStudent) (resp *ct.StudentPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.StudentPrimaryKey) (resp *ct.Student, err error)
	GetList(ctx context.Context, req *ct.GetListStudentRequest) (resp *ct.GetListStudentResponse, err error)
	Update(ctx context.Context, req *ct.UpdateStudentRequest) (resp *ct.STMessage, err error)
	Delete(ctx context.Context, req *ct.StudentPrimaryKey) (resp *ct.STMessage, err error)
	GetByGmail(ctx context.Context, req *ct.StudentGmail) (*ct.StudentGmailRes, error)
	StudentReport(ctx context.Context, req *ct.GetListStudentRequest) (*ct.GetRepStudent, error)
}

type GroupRepoI interface {
	Create(ctx context.Context, req *ct.CreateGroup) (resp *ct.GroupPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.GroupPrimaryKey) (resp *ct.Group, err error)
	GetList(ctx context.Context, req *ct.GetListGroupRequest) (resp *ct.GetListGroupResponse, err error)
	Update(ctx context.Context, req *ct.UpdateGroupRequest) (resp *ct.GRMessage, err error)
	Delete(ctx context.Context, req *ct.GroupPrimaryKey) (resp *ct.GRMessage, err error)
}

type LessonRepoI interface {
	Create(ctx context.Context, req *ct.CreateLesson) (resp *ct.LessonPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.LessonPrimaryKey) (resp *ct.Lesson, err error)
	GetList(ctx context.Context, req *ct.GetListLessonRequest) (resp *ct.GetListLessonResponse, err error)
	Update(ctx context.Context, req *ct.UpdateLessonRequest) (resp *ct.LSMessage, err error)
	Delete(ctx context.Context, req *ct.LessonPrimaryKey) (resp *ct.LSMessage, err error)
}

type TaskRepoI interface {
	Create(ctx context.Context, req *ct.CreateTask) (resp *ct.TaskPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.TaskPrimaryKey) (resp *ct.Task, err error)
	GetList(ctx context.Context, req *ct.GetListTaskRequest) (resp *ct.GetListTaskResponse, err error)
	Update(ctx context.Context, req *ct.UpdateTask) (resp *ct.TSMessage, err error)
	Delete(ctx context.Context, req *ct.TaskPrimaryKey) (resp *ct.TSMessage, err error)
	DoTask(ctx context.Context,req *ct.DoTaskReq) (*ct.TSMessage,error)
}

type JournalRepoI interface {
	Create(ctx context.Context, req *ct.CreateJournal) (resp *ct.JournalPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.JournalPrimaryKey) (resp *ct.Journal, err error)
	GetList(ctx context.Context, req *ct.GetListJournalRequest) (resp *ct.GetListJournalResponse, err error)
	Update(ctx context.Context, req *ct.UpdateJournalRequest) (resp *ct.JRMessage, err error)
	Delete(ctx context.Context, req *ct.JournalPrimaryKey) (resp *ct.JRMessage, err error)
}

type ScheduleRepoI interface {
	Create(ctx context.Context, req *ct.CreateSchedule) (resp *ct.SchedulePrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.SchedulePrimaryKey) (resp *ct.Schedule, err error)
	GetList(ctx context.Context, req *ct.GetListScheduleRequest) (resp *ct.GetListScheduleResponse, err error)
	Update(ctx context.Context, req *ct.UpdateScheduleRequest) (resp *ct.SHMessage, err error)
	Delete(ctx context.Context, req *ct.SchedulePrimaryKey) (resp *ct.SHMessage, err error)
}

type GroupManyRepoI interface {
	Create(ctx context.Context, req *ct.CreateGroupMany) (resp *ct.GMMessage, err error)
}

type EventRepoI interface {
	Create(ctx context.Context, req *ct.CreateEvent) (resp *ct.EventPrimaryKey, err error)
	GetByID(ctx context.Context, req *ct.EventPrimaryKey) (resp *ct.Event, err error)
	GetList(ctx context.Context, req *ct.GetListEventRequest) (resp *ct.GetListEventResponse, err error)
	Update(ctx context.Context, req *ct.UpdateEventRequest) (resp *ct.EVMessage, err error)
	Delete(ctx context.Context, req *ct.EventPrimaryKey) (resp *ct.EVMessage, err error)
	RegisterEvent(ctx context.Context,req *ct.RegisterEv) (*ct.EVMessage,error)
}

type AttendanceRepoI interface {
	Create(ctx context.Context, req *ct.CreateAttendance) (*ct.ATMessage, error)
	GetByID(ctx context.Context, req *ct.AttendancePrimaryKey) (*ct.Attendance, error)
	GetList(ctx context.Context, req *ct.GetListAttendanceRequest) (*ct.GetListAttendanceResponse, error)
}