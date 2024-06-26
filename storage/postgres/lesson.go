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

type LessonRepo struct {
	db *pgxpool.Pool
}

func NewLessonRepo(db *pgxpool.Pool) storage.LessonRepoI {
	return &LessonRepo{
		db: db,
	}
}

func (c *LessonRepo) Create(ctx context.Context, req *ct.CreateLesson) (*ct.LessonPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.LessonPrimaryKey{Id: id}

	query := `INSERT INTO "Lesson" (
			"ID",
			"Name",
			"created_at") VALUES (
				$1,
				$2,
				NOW()
			)`

	_, err := c.db.Exec(ctx, query, id, req.Name)
	if err != nil {
		log.Println("error while creating Lesson")
		return nil, err
	}

	return resp, err
}

func (c *LessonRepo) GetByID(ctx context.Context, req *ct.LessonPrimaryKey) (*ct.Lesson, error) {
	resp := &ct.Lesson{}
	query := `SELECT "ID",
					 "Name",
					 "created_at",
					 "updated_at"
			FROM "Lesson"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Name,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *LessonRepo) GetList(ctx context.Context, req *ct.GetListLessonRequest) (*ct.GetListLessonResponse, error) {
	resp := &ct.GetListLessonResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "Name" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "Name",
					 "created_at",
					 "updated_at"
			FROM "Lesson"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Lesson := &ct.Lesson{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Lesson.Id,
			&Lesson.Name,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Lesson.CreatedAt = helper.NullTimeStampToString(createdAt)
		Lesson.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Lesson = append(resp.Lesson, Lesson)
	}

	queryCount := `SELECT COUNT(*) FROM "Lesson" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *LessonRepo) Update(ctx context.Context, req *ct.UpdateLessonRequest) (*ct.LSMessage, error) {
	resp := &ct.LSMessage{Message: "Lesson updated successfully"}
	query := `UPDATE "Lesson" SET
					 			 "Name"=$1,
								 "updated_at"=NOW()
								 WHERE "ID"=$2 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Name,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *LessonRepo) Delete(ctx context.Context, req *ct.LessonPrimaryKey) (*ct.LSMessage, error) {
	resp := &ct.LSMessage{Message: "Lesson deleted successfully"}
	query := `UPDATE "Lesson" SET
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