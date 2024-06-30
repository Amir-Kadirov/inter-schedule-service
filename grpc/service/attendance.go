package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type AttendanceService struct {
	schedule_service.UnimplementedAttendanceServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewAttendanceService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AttendanceService {
	return &AttendanceService{
		cfg:      cfg,	
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *AttendanceService) Create(ctx context.Context, req *schedule_service.CreateAttendance) (*schedule_service.ATMessage, error) {

	c.log.Info("---CreateAttendance--->>>", logger.Any("req", req))

	resp, err := c.strg.Attendance().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateAttendance--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AttendanceService) GetByID(ctx context.Context, req *schedule_service.AttendancePrimaryKey) (*schedule_service.Attendance, error) {
	c.log.Info("---GetByIdAttendance--->>>", logger.Any("req", req))

	resp, err := c.strg.Attendance().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdAttendance--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *AttendanceService) GetList(ctx context.Context, req *schedule_service.GetListAttendanceRequest) (*schedule_service.GetListAttendanceResponse, error) {
	c.log.Info("---GetAllAttendance--->>>", logger.Any("req", req))

	resp, err := c.strg.Attendance().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllAttendance--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}