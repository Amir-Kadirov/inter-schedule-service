package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupManyRepo struct {
	db *pgxpool.Pool
}

func NewGroupManyRepo(db *pgxpool.Pool) storage.GroupManyRepoI {
	return &GroupManyRepo{
		db: db,
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

	_, err := c.db.Exec(ctx, query, req.GroupId,req.ScheduleId,req.JournalId)
	if err != nil {
		log.Println("error while creating GroupMany")
		return nil, err
	}

	return resp, err
}

func (c *GroupManyRepo) ScheduleM(ctx context.Context,req *ct.Empty) (*ct.ScheduleMonth,error) {
	resp:=&ct.ScheduleMonth{}
	query:=`SELECT 
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
				    gr."SupportTeacherID";
				`
	var starttime,endtime sql.NullTime
	err:=c.db.QueryRow(ctx,query).Scan(
		&resp.GroupID,
		&resp.GroupType,
		&starttime,
		&endtime,
		&resp.BranchID,
		&resp.TeacherID,
		&resp.SupportTeacherID,
		&resp.StudentCount)
	if err != nil {
		return nil, err
	}
	resp.StartTime=helper.NullTimeStampToString(starttime)
	resp.EndTime=helper.NullTimeStampToString(endtime)

	return resp,nil
}