package postgres

import (
	"context"
	"log"
	ct "schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type GroupManyRepo struct {
	db       *pgxpool.Pool
	services client.ServiceManagerI
}

func NewGroupManyRepo(db *pgxpool.Pool,service client.ServiceManagerI) storage.GroupManyRepoI {
	return &GroupManyRepo{
		db: db,
		services: service,
	}
}

func (c *GroupManyRepo) Create(ctx context.Context, req *ct.CreateGroupMany) (*ct.GMMessage, error) {
	// id := uuid.NewString()
	resp := &ct.GMMessage{Message: "Created successfully"}

	query := `INSERT INTO "Group_Many" (
			"GroupID",
			"ScheduleID",
			"JournalID"
			) VALUES (
				$1,
				$2,
				$3
			)`

	_, err := c.db.Exec(ctx, query, req.GroupId, req.ScheduleId, req.JournalId)
	if err != nil {
		log.Println("error while creating GroupMany")
		return nil, err
	}

	return resp, err
}