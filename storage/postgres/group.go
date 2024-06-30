package postgres

import (
	"context"
	"database/sql"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupRepo struct {
	db *pgxpool.Pool
	services client.ServiceManagerI
}

func NewGroupRepo(db *pgxpool.Pool,service client.ServiceManagerI) storage.GroupRepoI {
	return &GroupRepo{
		db: db,
		services: service,
	}
}

func (c *GroupRepo) Create(ctx context.Context, req *ct.CreateGroup) (*ct.GroupPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.GroupPrimaryKey{Id: id}


	query := `INSERT INTO "Group" (
			"ID",
			"TeacherID",
			"SupportTeacherID",
			"BranchID",
			"Type",	
			"created_at") VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				NOW()
			)`

	_, err := c.db.Exec(ctx, query, id, req.TeacherId, req.SupportTeacherId, req.Branchid, req.Type)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *GroupRepo) GetByID(ctx context.Context, req *ct.GroupPrimaryKey) (*ct.Group, error) {
	resp := &ct.Group{}
	query := `SELECT "ID",
					 "TeacherID",
					 "SupportTeacherID",
					 "BranchID",
					 "Type",
					 "created_at",
					 "updated_at"
			FROM "Group"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.TeacherId,
		&resp.SupportTeacherId,
		&resp.Branchid,
		&resp.Type,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *GroupRepo) GetList(ctx context.Context, req *ct.GetListGroupRequest) (*ct.GetListGroupResponse, error) {
	resp := &ct.GetListGroupResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "Type" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "TeacherID",
					 "SupportTeacherID",
					 "BranchID",
					 "Type",
					 "created_at",
					 "updated_at"
			FROM "Group"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Group := &ct.Group{}
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&Group.Id,
			&Group.TeacherId,
			&Group.SupportTeacherId,
			&Group.Branchid,
			&Group.Type,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Group.CreatedAt = helper.NullTimeStampToString(createdAt)
		Group.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Group = append(resp.Group, Group)
	}

	queryCount := `SELECT COUNT(*) FROM "Group" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *GroupRepo) Update(ctx context.Context, req *ct.UpdateGroupRequest) (*ct.GRMessage, error) {
	resp := &ct.GRMessage{Message: "Group updated successfully"}
	query := `UPDATE "Group" SET
					 			 "TeacherID"=$1,
					 			 "SupportTeacherID"=$2,
					 			 "BranchID"=$3,
					 			 "Type"=$4,
								 "updated_at"=NOW()
								 WHERE "ID"=$5 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.TeacherId, req.SupportTeacherId, req.Branchid, req.Type)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *GroupRepo) Delete(ctx context.Context, req *ct.GroupPrimaryKey) (*ct.GRMessage, error) {
	resp := &ct.GRMessage{Message: "Group deleted successfully"}
	query := `UPDATE "Group" SET
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