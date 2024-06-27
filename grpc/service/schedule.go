package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ScheduleService struct {
	schedule_service.UnimplementedScheduleServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewScheduleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScheduleService {
	return &ScheduleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *ScheduleService) Create(ctx context.Context, req *schedule_service.CreateSchedule) (*schedule_service.SchedulePrimaryKey, error) {

	c.log.Info("---CreateSchedule--->>>", logger.Any("req", req))

	resp, err := c.strg.Schedule().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateSchedule--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleService) GetByID(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (*schedule_service.Schedule, error) {
	c.log.Info("---GetByIdSchedule--->>>", logger.Any("req", req))

	resp, err := c.strg.Schedule().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdSchedule--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleService) GetList(ctx context.Context, req *schedule_service.GetListScheduleRequest) (*schedule_service.GetListScheduleResponse, error) {
	c.log.Info("---GetAllSchedule--->>>", logger.Any("req", req))

	resp, err := c.strg.Schedule().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllSchedule--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleService) Update(ctx context.Context, req *schedule_service.UpdateScheduleRequest) (*schedule_service.SHMessage, error) {
	c.log.Info("---UpdateSchedule--->>>", logger.Any("req", req))

	resp, err := c.strg.Schedule().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateSchedule--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *ScheduleService) Delete(ctx context.Context, req *schedule_service.SchedulePrimaryKey) (*schedule_service.SHMessage, error) {
	c.log.Info("---DeleteSchedule--->>>", logger.Any("req", req))

	resp, err := c.strg.Schedule().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteSchedule--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}