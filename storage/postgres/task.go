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

type TaskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) storage.TaskRepoI {
	return &TaskRepo{
		db: db,
	}
}

func (c *TaskRepo) Create(ctx context.Context, req *ct.CreateTask) (*ct.TaskPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.TaskPrimaryKey{Id: id}

	query := `INSERT INTO "Task" (
			"ID",
			"Label",
			"LessonID",
			"Deadline",
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				NOW()
			)`

	_, err := c.db.Exec(ctx, query, id, req.Label,req.LessonId,req.Deadline)
	if err != nil {
		log.Println("error while creating Task")
		return nil, err
	}

	return resp, err
}

func (c *TaskRepo) GetByID(ctx context.Context, req *ct.TaskPrimaryKey) (*ct.Task, error) {
	resp := &ct.Task{}
	query := `SELECT "ID",
					 "Label",
					 "LessonID",
					 "Deadline",
					 "Score",
					 "created_at",
					 "updated_at"
			FROM "Task"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var (updatedAt, createdAt, deadline  sql.NullTime
	     score int)
	if err := row.Scan(
		&resp.Id,
		&resp.Label,
		&resp.LessonId,
		&deadline,
		&score,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.Score=int32(score)
	resp.Deadline=helper.NullDateToString(deadline)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *TaskRepo) GetList(ctx context.Context, req *ct.GetListTaskRequest) (*ct.GetListTaskResponse, error) {
	resp := &ct.GetListTaskResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "LessonID" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "Label",
					 "LessonID",
					 "Deadline",
					 "Score",
					 "created_at",
					 "updated_at"
			FROM "Task"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Task := &ct.Task{}
		var createdAt, updatedAt,deadline  sql.NullTime
		if err := rows.Scan(
			&Task.Id,
			&Task.Label,
			&Task.LessonId,
			&deadline,
			&Task.Score,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Task.Deadline=helper.NullDateToString(deadline)
		Task.CreatedAt = helper.NullTimeStampToString(createdAt)
		Task.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Task = append(resp.Task, Task)
	}

	queryCount := `SELECT COUNT(*) FROM "Task" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TaskRepo) Update(ctx context.Context, req *ct.UpdateTask) (*ct.TSMessage, error) {
	resp := &ct.TSMessage{Message: "Task updated successfully"}
	query := `UPDATE "Task" SET
					 			 "Label"=$1,
								 "LessonID"=$2,
								 "Deadline"=$3,
								 "Score"=$4,
								 "updated_at"=NOW()
								 WHERE "ID"=$5 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Label,req.LessonId,req.Deadline,req.Score,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TaskRepo) Delete(ctx context.Context, req *ct.TaskPrimaryKey) (*ct.TSMessage, error) {
	resp := &ct.TSMessage{Message: "Task deleted successfully"}
	query := `UPDATE "Task" SET
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

func (c *TaskRepo) DoTask(ctx context.Context,req *ct.DoTaskReq) (*ct.TSMessage,error) {
	resp:=&ct.TSMessage{Message: "Task Completed Successfully"}
	query:=`UPDATE "Task" SET 
							"Label"=$1,
							"Status"=TRUE,
							updated_at=NOW()
							WHERE "ID"=$2 AND "deleted_at is null"`
	_,err:=c.db.Exec(ctx,query,req.Label,req.Id)
	if err != nil {
		return nil, err
	}

	return resp,nil
}