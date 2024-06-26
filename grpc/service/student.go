package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type StudentService struct {
	schedule_service.UnimplementedStudentServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *StudentService) Create(ctx context.Context, req *schedule_service.CreateStudent) (*schedule_service.StudentPrimaryKey, error) {

	c.log.Info("---CreateStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *StudentService) GetByID(ctx context.Context, req *schedule_service.StudentPrimaryKey) (*schedule_service.Student, error) {
	c.log.Info("---GetByIdStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *StudentService) GetList(ctx context.Context, req *schedule_service.GetListStudentRequest) (*schedule_service.GetListStudentResponse, error) {
	c.log.Info("---GetAllStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *StudentService) Update(ctx context.Context, req *schedule_service.UpdateStudentRequest) (*schedule_service.STMessage, error) {
	c.log.Info("---UpdateStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *StudentService) Delete(ctx context.Context, req *schedule_service.StudentPrimaryKey) (*schedule_service.STMessage, error) {
	c.log.Info("---DeleteStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *StudentService) GetByGmail(ctx context.Context, req *schedule_service.StudentGmail) (*schedule_service.StudentPrimaryKey, error) {
	c.log.Info("---GetByGmailStudent--->>>", logger.Any("req", req))

	resp, err := c.strg.Student().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailStudent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
