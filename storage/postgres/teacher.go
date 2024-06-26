package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	ct "schedule_service/genproto/genproto/user_service"
	"schedule_service/pkg/hash"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepoI {
	return &TeacherRepo{
		db: db,
	}
}

func (c *TeacherRepo) Create(ctx context.Context, req *ct.CreateTeacher) (*ct.TeacherPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.TeacherPrimaryKey{Id: id}

	query := `INSERT INTO "Teacher" (
			"ID",
			"FullName",
			"Phone",
			"Password",
			"Email",
			"Salary",
			"IeltsScore",
			"IeltsAttemptsCount",
			"SupportTeacherID",	
			"BranchID",
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				$9,
				$10,
				NOW()
			)`
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = c.db.Exec(ctx, query, id, req.Fullname, req.Phone, hashedPassword, req.Email,
		req.Salary, req.Ieltsscore, req.Ieltsattemptscount, req.Supportteacherid, req.Branchid)
	if err != nil {
		log.Println("error while creating teacher")
		return nil, err
	}

	return resp, err
}

func (c *TeacherRepo) GetByID(ctx context.Context, req *ct.TeacherPrimaryKey) (*ct.Teacher, error) {
	resp := &ct.Teacher{}
	query := `SELECT "ID",
					"FullName",
					"Phone",
					"Email",
					"Salary",
					"IeltsScore",
					"IeltsAttemptsCount",
					"SupportTeacherID",	
					"BranchID",
					"created_at",
					"updated_at"
			FROM "Teacher"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Email,
		&resp.Salary,
		&resp.Ieltsscore,
		&resp.Ieltsattemptscount,
		&resp.Supportteacherid,
		&resp.Branchid,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *TeacherRepo) GetList(ctx context.Context, req *ct.GetListTeacherRequest) (*ct.GetListTeacherResponse, error) {
	resp := &ct.GetListTeacherResponse{}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "IeltsScore" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT
					"ID",
					"FullName",
					"Phone",
					"Email",
					"Salary",
					"IeltsScore",
					"IeltsAttemptsCount",
					"SupportTeacherID",	
					"BranchID",
					"created_at",
					"updated_at"
			FROM "Teacher"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Teacher := &ct.Teacher{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Teacher.Id,
			&Teacher.Fullname,
			&Teacher.Phone,
			&Teacher.Email,
			&Teacher.Salary,
			&Teacher.Ieltsscore,
			&Teacher.Ieltsattemptscount,
			&Teacher.Supportteacherid,
			&Teacher.Branchid,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Teacher.CreatedAt = helper.NullTimeStampToString(createdAt)
		Teacher.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Teacher = append(resp.Teacher, Teacher)
	}

	queryCount := `SELECT COUNT(*) FROM "Teacher" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TeacherRepo) Update(ctx context.Context, req *ct.UpdateTeacherRequest) (*ct.Message, error) {
	resp := &ct.Message{Message: "Teacher updated successfully"}
	query := `UPDATE "Teacher" SET "FullName"=$1,
								 "Phone"=$2,
								 "Email"=$3,
								 "Salary"=$4,
								 "IeltsScore"=$5,
								 "IeltsAttemptsCount"=$6,
								 "SupportTeacherID"=$7,
								 "BranchID"=$8,
								 "updated_at"=NOW()
								 WHERE "ID"=$9 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Fullname, req.Phone, req.Email, req.Salary, req.Ieltsscore, req.Ieltsattemptscount,
		req.Supportteacherid, req.Branchid, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TeacherRepo) Delete(ctx context.Context, req *ct.TeacherPrimaryKey) (*ct.Message, error) {
	resp := &ct.Message{Message: "Teacher deleted successfully"}
	query := `UPDATE "Teacher" SET
							 "deleted_at"=NOW()
							 WHERE "ID"=$1 AND "deleted_at" is null RETURNING "created_at"`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("teacher is not exits")
	} else if err != nil {
		return nil, err
	}

	if err = helper.DeleteChecker(createdAt); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *TeacherRepo) GetByGmail(ctx context.Context, req *ct.TeacherGmail) (*ct.TeacherPrimaryKey, error) {
	query := `SELECT "ID" FROM "Teacher" WHERE "Email"=$1`
	var id string
	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.TeacherPrimaryKey{Id: id}, nil
}
