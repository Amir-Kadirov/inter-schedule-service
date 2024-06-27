package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ScheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) storage.ScheduleRepoI {
	return &ScheduleRepo{
		db: db,
	}
}

func (c *ScheduleRepo) Create(ctx context.Context, req *ct.CreateSchedule) (*ct.SchedulePrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.SchedulePrimaryKey{Id: id}

	query := `INSERT INTO "Schedule" (
			"ID",
			"EndTime",
			"StartTime",
			"LessonID",
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				NOW()
			)`

	_, err := c.db.Exec(ctx, query, id, req.EndTime,req.StartTime,req.LessonId)
	if err != nil {
		log.Println("error while creating Schedule")
		return nil, err
	}

	return resp, err
}

func (c *ScheduleRepo) GetByID(ctx context.Context, req *ct.SchedulePrimaryKey) (*ct.Schedule, error) {
	resp := &ct.Schedule{}
	query := `SELECT "ID",
					 "EndTime",
					 "StartTime",
					 "LessonID",
					 "created_at",
					 "updated_at"
			FROM "Schedule"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt,endTime,startTime sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&endTime,
		&startTime,
		&resp.LessonId,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.EndTime=helper.NullTimeStampToString(endTime)
	resp.StartTime=helper.NullTimeStampToString(startTime)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *ScheduleRepo) GetList(ctx context.Context, req *ct.GetListScheduleRequest) (*ct.GetListScheduleResponse, error) {
	resp := &ct.GetListScheduleResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "From" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "EndTime",
					 "StartTime",
					 "LessonID",
					 "created_at",
					 "updated_at"
			FROM "Schedule"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Schedule := &ct.Schedule{}
		var updatedAt, createdAt,endTime,startTime sql.NullTime
		if err := rows.Scan(
			&Schedule.Id,
			&endTime,
			&startTime,
			&Schedule.LessonId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Schedule.EndTime=helper.NullTimeStampToString(endTime)
		Schedule.StartTime=helper.NullTimeStampToString(startTime)
		Schedule.CreatedAt = helper.NullTimeStampToString(createdAt)
		Schedule.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Schedule = append(resp.Schedule, Schedule)
	}

	queryCount := `SELECT COUNT(*) FROM "Schedule" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleRepo) Update(ctx context.Context, req *ct.UpdateScheduleRequest) (*ct.SHMessage, error) {
	resp := &ct.SHMessage{Message: "Schedule updated successfully"}
	query := `UPDATE "Schedule" SET
					 			 "EndTime"=$1,
								 "StartTime"=$2,
								 "LessonID"=$3,
								 "updated_at"=NOW()
								 WHERE "ID"=$4 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.EndTime,req.StartTime,req.LessonId,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleRepo) Delete(ctx context.Context, req *ct.SchedulePrimaryKey) (*ct.SHMessage, error) {
	resp := &ct.SHMessage{Message: "Schedule deleted successfully"}
	query := `UPDATE "Schedule" SET
							 "deleted_at"=NOW()
							 WHERE "ID"=$1 AND "deleted_at" is null RETURNING "created_at"`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err = helper.DeleteChecker(createdAt); err != nil {
		return resp, nil
	}

	return resp, nil
}