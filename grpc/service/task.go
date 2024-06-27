package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type TaskService struct {
	schedule_service.UnimplementedTaskServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TaskService {
	return &TaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *TaskService) Create(ctx context.Context, req *schedule_service.CreateTask) (*schedule_service.TaskPrimaryKey, error) {

	c.log.Info("---CreateTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TaskService) GetByID(ctx context.Context, req *schedule_service.TaskPrimaryKey) (*schedule_service.Task, error) {
	c.log.Info("---GetByIdTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TaskService) GetList(ctx context.Context, req *schedule_service.GetListTaskRequest) (*schedule_service.GetListTaskResponse, error) {
	c.log.Info("---GetAllTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TaskService) Update(ctx context.Context, req *schedule_service.UpdateTask) (*schedule_service.TSMessage, error) {
	c.log.Info("---UpdateTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TaskService) Delete(ctx context.Context, req *schedule_service.TaskPrimaryKey) (*schedule_service.TSMessage, error) {
	c.log.Info("---DeleteTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *TaskService) DoTask(ctx context.Context, req *schedule_service.DoTaskReq) (*schedule_service.TSMessage, error) {
	c.log.Info("---DoTaskTask--->>>", logger.Any("req", req))

	resp, err := c.strg.Task().DoTask(ctx, req)
	if err != nil {
		c.log.Error("---DoTaskTask--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}