package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "schedule_service/genproto/genproto/user_service"
	"schedule_service/pkg/hash"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AdminRepo struct {
	db *pgxpool.Pool
}

func NewAdminRepo(db *pgxpool.Pool) storage.AdminRepoI {
	return &AdminRepo{
		db: db,
	}
}

func (c *AdminRepo) Create(ctx context.Context, req *ct.CreateAdmin) (*ct.AdminPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.AdminPrimaryKey{Id: id}

	query := `INSERT INTO "Administration" (
			"ID",
			"FullName",
			"Phone",
			"Password",
			"Email",
			"Salary",
			"IeltsScore",
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
				NOW()
			)`
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = c.db.Exec(ctx, query, id, req.Fullname, req.Phone, hashedPassword, req.Email,
		req.Salary, req.Ieltsscore, req.Branchid)
	if err != nil {
		log.Println("error while creating admin")
		return nil, err
	}

	return resp, err
}

func (c *AdminRepo) GetByID(ctx context.Context, req *ct.AdminPrimaryKey) (*ct.Admin, error) {
	resp := &ct.Admin{}
	query := `SELECT "ID",
					 "FullName",
					 "Phone",
					 "Email",
					 "Salary",
					 "IeltsScore",
					 "BranchID",
					 "created_at",
					"updated_at"
			FROM "Administration"
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
		&resp.Branchid,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *AdminRepo) GetList(ctx context.Context, req *ct.GetListAdminRequest) (*ct.GetListAdminResponse, error) {
	resp := &ct.GetListAdminResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

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
					 "BranchID",
					 "created_at",
					 "updated_at"
			FROM "Administration"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Admin := &ct.Admin{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Admin.Id,
			&Admin.Fullname,
			&Admin.Phone,
			&Admin.Email,
			&Admin.Salary,
			&Admin.Ieltsscore,
			&Admin.Branchid,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Admin.CreatedAt = helper.NullTimeStampToString(createdAt)
		Admin.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Admin = append(resp.Admin, Admin)
	}

	queryCount := `SELECT COUNT(*) FROM "Administration" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *AdminRepo) Update(ctx context.Context, req *ct.UpdateAdminRequest) (*ct.ADMessage, error) {
	resp := &ct.ADMessage{Message: "Admin updated successfully"}
	query := `UPDATE "Administration" SET
								 "FullName"=$1,
								 "Phone"=$2,
								 "Email"=$3,
								 "Salary"=$4,
								 "IeltsScore"=$5,
								 "BranchID"=$6,
								 "updated_at"=NOW()
								 WHERE "ID"=$7 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Fullname, req.Phone, req.Email, req.Salary, req.Ieltsscore,
		req.Branchid, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *AdminRepo) Delete(ctx context.Context, req *ct.AdminPrimaryKey) (*ct.ADMessage, error) {
	resp := &ct.ADMessage{Message: "Admin deleted successfully"}
	query := `UPDATE "Administration" SET
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

func (c *AdminRepo) GetByGmail(ctx context.Context, req *ct.AdminGmail) (*ct.AdminPrimaryKey, error) {
	query := `SELECT "ID" FROM "Administration" WHERE "Email"=$1`
	var id string
	err := c.db.QueryRow(ctx, query, req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.AdminPrimaryKey{Id: id}, nil
}
