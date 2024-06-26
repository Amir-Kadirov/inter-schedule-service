package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type GroupService struct {
	schedule_service.UnimplementedGroupServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewGroupService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupService {
	return &GroupService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *GroupService) Create(ctx context.Context, req *schedule_service.CreateGroup) (*schedule_service.GroupPrimaryKey, error) {

	c.log.Info("---CreateGroup--->>>", logger.Any("req", req))

	resp, err := c.strg.Group().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateGroup--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *GroupService) GetByID(ctx context.Context, req *schedule_service.GroupPrimaryKey) (*schedule_service.Group, error) {
	c.log.Info("---GetByIdGroup--->>>", logger.Any("req", req))

	resp, err := c.strg.Group().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdGroup--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *GroupService) GetList(ctx context.Context, req *schedule_service.GetListGroupRequest) (*schedule_service.GetListGroupResponse, error) {
	c.log.Info("---GetAllGroup--->>>", logger.Any("req", req))

	resp, err := c.strg.Group().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllGroup--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *GroupService) Update(ctx context.Context, req *schedule_service.UpdateGroupRequest) (*schedule_service.GRMessage, error) {
	c.log.Info("---UpdateGroup--->>>", logger.Any("req", req))

	resp, err := c.strg.Group().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateGroup--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *GroupService) Delete(ctx context.Context, req *schedule_service.GroupPrimaryKey) (*schedule_service.GRMessage, error) {
	c.log.Info("---DeleteGroup--->>>", logger.Any("req", req))

	resp, err := c.strg.Group().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteGroup--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}