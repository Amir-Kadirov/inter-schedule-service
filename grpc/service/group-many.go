package service

import (
	"context"
	"database/sql"
	"errors"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/genproto/genproto/user_service"
	"schedule_service/grpc/client"
	"schedule_service/pkg/helper"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type GroupManyService struct {
	schedule_service.UnimplementedGroupManyServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewGroupManyService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupManyService {
	return &GroupManyService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (c *GroupManyService) Create(ctx context.Context, req *schedule_service.CreateGroupMany) (*schedule_service.GMMessage, error) {

	c.log.Info("---CreateGroupMany--->>>", logger.Any("req", req))

	resp, err := c.strg.GroupMany().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateGroupMany--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *GroupManyService) ScheduleM(ctx context.Context, req *schedule_service.Empty) (*schedule_service.ScheduleMonth, error) {

	c.log.Info("---ScheduleGroupMany--->>>", logger.Any("req", req))

	db:=c.strg.GetDB()

	resp := &schedule_service.ScheduleMonth{}

	query := `SELECT 
						gr."ID",
						gr."Type",
						sch."StartTime",
						sch."EndTime",
						gr."BranchID",
						gr."TeacherID",
						gr."SupportTeacherID",
						COUNT(gt."StudentID")
					FROM 
						"Group_Many" gm
						JOIN "Group" gr ON gr."ID" = gm."GroupID"
						JOIN "Schedule" sch ON sch."ID" = gm."ScheduleID"
						JOIN "Group_Table" gt ON gt."GroupID" = gm."GroupID"
					GROUP BY 
						gr."ID",
						gr."Type",
						sch."StartTime",
						sch."EndTime",
						gr."BranchID",
						gr."TeacherID",
						gr."SupportTeacherID"
					`
	row, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var starttime, endtime sql.NullTime
		scheduleMonth := &schedule_service.SchMonth{}
		if err := row.Scan(
			&scheduleMonth.GroupID,
			&scheduleMonth.GroupType,
			&starttime,
			&endtime,
			&scheduleMonth.BranchName,
			&scheduleMonth.TeacherName,
			&scheduleMonth.SupportTeacherName,
			&scheduleMonth.StudentCount); err != nil {
			return nil, err
		}
		scheduleMonth.StartTime = helper.NullTimeStampToString(starttime)
		scheduleMonth.EndTime = helper.NullTimeStampToString(endtime)
		branch, err := c.services.BranchService().GetByID(ctx, &user_service.BranchPrimaryKey{Id: scheduleMonth.BranchName})
		if err != nil {
			return nil, errors.New("wrong Branch Id")
		}
		teacher, err := c.services.TeacherService().GetByID(ctx, &user_service.TeacherPrimaryKey{Id: scheduleMonth.TeacherName})
		if err != nil {
			return nil, errors.New("wrong teacher Id")
		}
		supportteacher, err := c.services.SupportTeacherService().GetByID(ctx, &user_service.SupportTeacherPrimaryKey{Id: scheduleMonth.SupportTeacherName})
		if err != nil {
			return nil, errors.New("wrong support teacher Id")
		}
		scheduleMonth.BranchName = branch.Address
		scheduleMonth.TeacherName = teacher.Fullname
		scheduleMonth.SupportTeacherName = supportteacher.Fullname

		resp.SMonth = append(resp.SMonth, scheduleMonth)
	}

	return resp, nil
}
