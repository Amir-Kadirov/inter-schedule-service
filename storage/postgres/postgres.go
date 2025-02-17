package postgres

import (
	"context"
	"fmt"
	"log"
	"schedule_service/config"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db      *pgxpool.Pool
	services client.ServiceManagerI
	student storage.StudentRepoI
	group storage.GroupRepoI
	lesson storage.LessonRepoI
	task storage.TaskRepoI
	journal storage.JournalRepoI
	schedule storage.ScheduleRepoI
	groupmany storage.GroupManyRepoI
	event storage.EventRepoI
	attendence	storage.AttendanceRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Student() storage.StudentRepoI {
	if s.student == nil {
		s.student = NewStudentRepo(s.db)
	}

	return s.student
}

func (s *Store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db,s.services)
	}

	return s.group
}

func (s *Store) Lesson() storage.LessonRepoI {
	if s.lesson == nil {
		s.lesson = NewLessonRepo(s.db)
	}

	return s.lesson
}


func (s *Store) Task() storage.TaskRepoI {
	if s.task == nil {
		s.task = NewTaskRepo(s.db)
	}

	return s.task
}

func (s *Store) Journal() storage.JournalRepoI {
	if s.journal == nil {
		s.journal = NewJournalRepo(s.db)
	}

	return s.journal
}


func (s *Store) Schedule() storage.ScheduleRepoI {
	if s.schedule == nil {
		s.schedule = NewScheduleRepo(s.db)
	}

	return s.schedule
}

func (s *Store) GroupMany() storage.GroupManyRepoI {
	if s.groupmany == nil {
		s.groupmany = NewGroupManyRepo(s.db,s.services)
	}

	return s.groupmany
}

func (s *Store) Event() storage.EventRepoI {
	if s.event==nil {
		s.event=NewEventRepo(s.db)
	}

	return s.event
}

func (s *Store) Attendance() storage.AttendanceRepoI {
	if s.attendence==nil {
		s.attendence=NewAttendanceRepo(s.db)
	}

	return s.attendence
}

func (s Store) GetDB() *pgxpool.Pool {
	return s.db
}