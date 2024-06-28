package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/pkg/helper"
	"schedule_service/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) storage.EventRepoI {
	return &EventRepo{
		db: db,
	}
}

func (c *EventRepo) Create(ctx context.Context, req *ct.CreateEvent) (*ct.EventPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.EventPrimaryKey{Id: id}

	date, err := time.Parse("2006-01-02", req.Day)
	if err != nil {
		return nil, err
	}

	if date.Weekday() != time.Sunday {
		return nil, errors.New("wrong day: valid only Sunday")
	}

	querySelect := `SELECT 1 FROM "Event" WHERE "BranchID"=$1 AND "Day"=$2 AND "StartTime"=$3`
	var exists int
	_ = c.db.QueryRow(ctx, querySelect, req.BranchId, req.Day, req.Starttime).Scan(&exists)

	if exists == 1 {
	    return nil, errors.New("event already exists")
	}

	query := `INSERT INTO "Event" (
		"ID",
		"Topic",
		"StartTime",
		"Day",
		"BranchID",
		"created_at") VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			NOW()
		)`

	_, err = c.db.Exec(ctx, query, id, req.Topic, req.Starttime, req.Day, req.BranchId)
	if err != nil {
		log.Println("error while creating Event")
		return nil, err
	}

	return resp, nil
}

func (c *EventRepo) GetByID(ctx context.Context, req *ct.EventPrimaryKey) (*ct.Event, error) {
	resp := &ct.Event{}
	query := `SELECT "ID",
					 "Topic",
					 "StartTime",
					 "Day",
					 "BranchID",
					 "created_at",
					 "updated_at"
			FROM "Event"
			WHERE "ID"=$1 AND "deleted_at" is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var (updatedAt, createdAt,day sql.NullTime
	    starttime sql.NullString)
	if err := row.Scan(
		&resp.Id,
		&resp.Topic,
		&starttime,
		&day,
		&resp.BranchId,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.Day=helper.NullDateToString(day)
	resp.Starttime=helper.NullTimeToString(starttime)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *EventRepo) GetList(ctx context.Context, req *ct.GetListEventRequest) (*ct.GetListEventResponse, error) {
	resp := &ct.GetListEventResponse{}
	if req.Offset == 0 {
		req.Offset = 1
	}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND "BranchID" ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
					 "ID",
					 "Topic",
					 "StartTime",
					 "Day",
					 "BranchID",
					 "created_at",
					 "updated_at"
			FROM "Event"
        	WHERE "deleted_at" is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		Event := &ct.Event{}
		var updatedAt, createdAt,starttime,day sql.NullTime
		if err := rows.Scan(
			&Event.Id,
			&Event.Topic,
			&starttime,
			&day,
			&Event.BranchId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		Event.Day=helper.NullDateToString(day)
		Event.Starttime=helper.NullTimeStampToString(starttime)
		Event.CreatedAt = helper.NullTimeStampToString(createdAt)
		Event.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Event = append(resp.Event, Event)
	}

	queryCount := `SELECT COUNT(*) FROM "Event" WHERE "deleted_at" is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EventRepo) Update(ctx context.Context, req *ct.UpdateEventRequest) (*ct.EVMessage, error) {
	resp := &ct.EVMessage{Message: "Event updated successfully"}
	query := `UPDATE "Event" SET
					 			 "Topic"=$1,
					 			 "StartTime"=$2,
					 			 "Day"=$3,
					 			 "BranchID"=$4,
								 "updated_at"=NOW()
								 WHERE "ID"=$5 AND "deleted_at" is null`
	_, err := c.db.Exec(ctx, query, req.Topic,req.Starttime,req.Day,req.BranchId,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EventRepo) Delete(ctx context.Context, req *ct.EventPrimaryKey) (*ct.EVMessage, error) {
	resp := &ct.EVMessage{Message: "Event deleted successfully"}
	query := `UPDATE "Event" SET
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

func (c *EventRepo) RegisterEvent(ctx context.Context,req *ct.RegisterEv) (*ct.EVMessage,error) {
	resp:=&ct.EVMessage{Message: "Registered Successfully"}
	queryEvent:=`SELECT "StartTime","Day" FROM "Event" WHERE "ID"=$1`

	var (evTime sql.NullString
		evDay sql.NullTime)
	err:=c.db.QueryRow(ctx,queryEvent,req.EventId).Scan(&evTime,&evDay)
	if err != nil {
		return nil, err
	}

	EvTime:=helper.NullTimeToString(evTime)
	EvDay:=helper.NullDateToString(evDay)

	queryRegister:=`SELECT ev."StartTime",ev."Day" FROM "Register_Event" re
							      JOIN "Event" ev ON ev."ID"=re."EventID" WHERE re."StudentID"=$1`
	row,err:=c.db.Query(ctx,queryRegister,req.StudentId)
	if err != nil {
		return nil, err
	}
	for row.Next(){
		var (starttime sql.NullString
			day sql.NullTime)
		if err:=row.Scan(&starttime,&day);err!=nil{
			return nil,err
		}

		if helper.NullTimeToString(starttime)==EvTime && helper.NullDateToString(day)==EvDay {
			return nil,errors.New("student already registered for event at this time")
		}
	}

	queryInsert:=`INSERT INTO "Register_Event"(
									"EventID",
									"StudentID"
									) VALUES (
									 $1,
									 $2
									)`
	_,err=c.db.Exec(ctx,queryInsert,req.EventId,req.StudentId)
	if err != nil {
		return nil, err
	}

	return resp,nil	
}