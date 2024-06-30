package postgres

import (
	"context"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AttendanceRepo struct {
	db *pgxpool.Pool
}

func NewAttendanceRepo(db *pgxpool.Pool) storage.AttendanceRepoI {
	return &AttendanceRepo{
		db: db,
	}
}

func (c *AttendanceRepo) Create(ctx context.Context, req *ct.CreateAttendance) (*ct.ATMessage, error) {
	id := uuid.NewString()
	resp := &ct.ATMessage{Message: "Attendance created successfully"}
	query := `INSERT INTO "Attendance" (
			"StudentID",
			"LessonID",
			"Status") VALUES (
				$1,
				$2,
				$3,
			)`

	_, err := c.db.Exec(ctx, query, id, req.StudentId,req.LessonId,req.Status)
	if err != nil {
		log.Println("error while creating Attendance")

		return nil, err
	}

	return resp, err
}

func (c *AttendanceRepo) GetByID(ctx context.Context, req *ct.AttendancePrimaryKey) (*ct.Attendance, error) {
	resp := &ct.Attendance{}
	query := `SELECT "StudentID",
					 "Status",
					 "LessonID"
			FROM "Attendance"
			WHERE "StudentID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	if err := row.Scan(
		&resp.StudentId,
		&resp.Status,
		&resp.LessonId); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *AttendanceRepo) GetList(ctx context.Context, req *ct.GetListAttendanceRequest) (*ct.GetListAttendanceResponse, error) {
	resp := &ct.GetListAttendanceResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "LessonID" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "StudentID",
					 "Status",
					 "LessonID"
			FROM "Attendance"
        	WHERE TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Attendance := &ct.Attendance{}
		if err := rows.Scan(
			&Attendance.StudentId,
			&Attendance.Status,
			&Attendance.LessonId,
			); err != nil {
			return nil, err
		}

		resp.Attendance = append(resp.Attendance, Attendance)
	}

	queryCount := `SELECT COUNT(*) FROM "Attendance" WHERE TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}