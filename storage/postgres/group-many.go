package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/genproto/genproto/user_service"
	"schedule_service/grpc/client"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupManyRepo struct {
	db       *pgxpool.Pool
	services client.ServiceManagerI
}

func NewGroupManyRepo(db *pgxpool.Pool,service client.ServiceManagerI) storage.GroupManyRepoI {
	return &GroupManyRepo{
		db: db,
		services: service,
	}
}

func (c *GroupManyRepo) Create(ctx context.Context, req *ct.CreateGroupMany) (*ct.GMMessage, error) {
	// id := uuid.NewString()
	resp := &ct.GMMessage{Message: "Created successfully"}

	query := `INSERT INTO "Group_Many" (
			"GroupID",
			"ScheduleID",
			"JournalID"
			) VALUES (
				$1,
				$2,
				$3
			)`

	_, err := c.db.Exec(ctx, query, req.GroupId, req.ScheduleId, req.JournalId)
	if err != nil {
		log.Println("error while creating GroupMany")
		return nil, err
	}

	return resp, err
}

func (c *GroupManyRepo) ScheduleM(ctx context.Context, req *ct.Empty) (*ct.ScheduleMonth, error) {
	resp := &ct.ScheduleMonth{}

	query := `SELECT 
				    gr."ID",
				    gr."Type",
				    sch."StartTime",
				    sch."EndTime",
				    gr."BranchID",
				    gr."TeacherID",
				    gr."SupportTeacherID",
				    COUNT(gt."StudentID")
				FROM 
				    "Group_Many" gm
				    JOIN "Group" gr ON gr."ID" = gm."GroupID"
				    JOIN "Schedule" sch ON sch."ID" = gm."ScheduleID"
				    JOIN "Group_Table" gt ON gt."GroupID" = gm."GroupID"
				GROUP BY 
				    gr."ID",
				    gr."Type",
				    sch."StartTime",
				    sch."EndTime",
				    gr."BranchID",
				    gr."TeacherID",
				    gr."SupportTeacherID"
				`
	row, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var starttime, endtime sql.NullTime
		scheduleMonth := &ct.SchMonth{}
		if err := row.Scan(
			&scheduleMonth.GroupID,
			&scheduleMonth.GroupType,
			&starttime,
			&endtime,
			&scheduleMonth.BranchName,
			&scheduleMonth.TeacherName,
			&scheduleMonth.SupportTeacherName,
			&scheduleMonth.StudentCount); err != nil {
			return nil, err
		}
		scheduleMonth.StartTime=helper.NullTimeStampToString(starttime)
		scheduleMonth.EndTime=helper.NullTimeStampToString(endtime)
		branch,err:=c.services.BranchService().GetByID(ctx,&user_service.BranchPrimaryKey{Id: scheduleMonth.BranchName})
		if err != nil {
			return nil, err
		}
		teacher,err:=c.services.TeacherService().GetByID(ctx,&user_service.TeacherPrimaryKey{Id: scheduleMonth.TeacherName})
		if err != nil {
			return nil, err
		}
		supportteacher,err:=c.services.SupportTeacherService().GetByID(ctx,&user_service.SupportTeacherPrimaryKey{Id: scheduleMonth.SupportTeacherName})
		if err != nil {
			return nil, err
		}
		scheduleMonth.BranchName=branch.Address
		scheduleMonth.TeacherName=teacher.Fullname
		scheduleMonth.SupportTeacherName=supportteacher.Fullname

		resp.SMonth=append(resp.SMonth, scheduleMonth)
	}

	return resp, nil
}
