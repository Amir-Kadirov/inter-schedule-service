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

type JournalRepo struct {
	db *pgxpool.Pool
}

func NewJournalRepo(db *pgxpool.Pool) storage.JournalRepoI {
	return &JournalRepo{
		db: db,
	}
}

func (c *JournalRepo) Create(ctx context.Context, req *ct.CreateJournal) (*ct.JournalPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.JournalPrimaryKey{Id: id}

	query := `INSERT INTO "Journal" (
			"ID",
			"From",
			"To",
			"created_at") VALUES (
				$1,
				$2,
				$3,
				NOW()
			)`

	_, err := c.db.Exec(ctx, query, id, req.From,req.To)
	if err != nil {
		log.Println("error while creating Journal")
		return nil, err
	}

	return resp, err
}

func (c *JournalRepo) GetByID(ctx context.Context, req *ct.JournalPrimaryKey) (*ct.Journal, error) {
	resp := &ct.Journal{}
	query := `SELECT "ID",
					 "From",
					 "To",
					 "created_at",
					 "updated_at"
			FROM "Journal"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, createdAt,from,to sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&from,
		&to,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.From=helper.NullDateToString(from)
	resp.From=helper.NullDateToString(to)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *JournalRepo) GetList(ctx context.Context, req *ct.GetListJournalRequest) (*ct.GetListJournalResponse, error) {
	resp := &ct.GetListJournalResponse{}
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
					 "From",
					 "To",
					 "created_at",
					 "updated_at"
			FROM "Journal"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Journal := &ct.Journal{}
		var updatedAt, createdAt,from,to sql.NullTime
		if err := rows.Scan(
			&Journal.Id,
			&from,
			&to,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Journal.From=helper.NullDateToString(from)
		Journal.To=helper.NullDateToString(to)
		Journal.CreatedAt = helper.NullTimeStampToString(createdAt)
		Journal.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Journal = append(resp.Journal, Journal)
	}

	queryCount := `SELECT COUNT(*) FROM "Journal" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *JournalRepo) Update(ctx context.Context, req *ct.UpdateJournalRequest) (*ct.JRMessage, error) {
	resp := &ct.JRMessage{Message: "Journal updated successfully"}
	query := `UPDATE "Journal" SET
					 			 "From"=$1,
								 "To"=$2,
								 "updated_at"=NOW()
								 WHERE "ID"=$3 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.From,req.To,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *JournalRepo) Delete(ctx context.Context, req *ct.JournalPrimaryKey) (*ct.JRMessage, error) {
	resp := &ct.JRMessage{Message: "Journal deleted successfully"}
	query := `UPDATE "Journal" SET
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