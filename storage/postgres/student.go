package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/pkg/hash"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepoI {
	return &StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(ctx context.Context, req *ct.CreateStudent) (*ct.StudentPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.StudentPrimaryKey{Id: id}

	query := `INSERT INTO "Student" (
			"ID",
			"FullName",
			"Phone",
			"Email",
			"Password",
			"GroupID",
			"BranchID",
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				NOW()
			)`
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = c.db.Exec(ctx, query, id, req.Fullname, req.Phone, req.Email, hashedPassword,
		req.GroupId, req.Branchid)
	if err != nil {
		log.Println("error while creating Student")
		return nil, err
	}

	return resp, err
}

func (c *StudentRepo) GetByID(ctx context.Context, req *ct.StudentPrimaryKey) (*ct.Student, error) {
	resp := &ct.Student{}
	query := `SELECT "ID",
					"FullName",
					"Phone",
					"Email",
					"GroupID",
					"PaidSum",
					"CourseCount",
					"TotalSum",
					"BranchID",
					"created_at",
					"updated_at"
			FROM "Student"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Email,
		&resp.GroupId,
		&resp.PaidSum,
		&resp.CourseCount,
		&resp.TotalSum,
		&resp.Branchid,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *StudentRepo) GetList(ctx context.Context, req *ct.GetListStudentRequest) (*ct.GetListStudentResponse, error) {
	resp := &ct.GetListStudentResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "GroupID" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					"ID",
					"FullName",
					"Phone",
					"Email",
					"GroupID",
					"PaidSum",
					"CourseCount",
					"TotalSum",
					"BranchID",
					"created_at",
					"updated_at"
			FROM "Student"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Student := &ct.Student{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Student.Id,
			&Student.Fullname,
			&Student.Phone,
			&Student.Email,
			&Student.GroupId,
			&Student.PaidSum,
			&Student.CourseCount,
			&Student.TotalSum,
			&Student.Branchid,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Student.CreatedAt = helper.NullTimeStampToString(createdAt)
		Student.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Student = append(resp.Student, Student)
	}

	queryCount := `SELECT COUNT(*) FROM "Student" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *StudentRepo) Update(ctx context.Context, req *ct.UpdateStudentRequest) (*ct.STMessage, error) {
	resp := &ct.STMessage{Message: "Student updated successfully"}
	query := `UPDATE "Student" SET
								 "FullName"=$1,
								 "Phone"=$2,
								 "Email"=$3,
								 "GroupID"=$4,
								 "PaidSum"=$5,
								 "CourseCount"=$6,
								 "TotalSum"="TotalSum"+$7,
								 "BranchID"=$8,
								 "updated_at"=NOW()
								 WHERE "ID"=$9 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Fullname, req.Phone, req.Email, req.GroupId, req.PaidSum,
		req.CourseCount, req.PaidSum,req.Branchid,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *StudentRepo) Delete(ctx context.Context, req *ct.StudentPrimaryKey) (*ct.STMessage, error) {
	resp := &ct.STMessage{Message: "Student deleted successfully"}
	query := `UPDATE "Student" SET
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

func (c *StudentRepo) GetByGmail(ctx context.Context, req *ct.StudentGmail) (*ct.StudentPrimaryKey, error) {
	query := `SELECT "ID" FROM "Student" WHERE "Email"=$1`
	var id string
	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.StudentPrimaryKey{Id: id}, nil
}
