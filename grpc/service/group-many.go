package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
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

	resp, err := c.strg.GroupMany().ScheduleM(ctx, req)
	if err != nil {
		c.log.Error("---ScheduleGroupMany--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}