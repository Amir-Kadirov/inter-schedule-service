package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type LessonService struct {
	schedule_service.UnimplementedLessonServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewLessonService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *LessonService {
	return &LessonService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *LessonService) Create(ctx context.Context, req *schedule_service.CreateLesson) (*schedule_service.LessonPrimaryKey, error) {

	c.log.Info("---CreateLesson--->>>", logger.Any("req", req))

	resp, err := c.strg.Lesson().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateLesson--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *LessonService) GetByID(ctx context.Context, req *schedule_service.LessonPrimaryKey) (*schedule_service.Lesson, error) {
	c.log.Info("---GetByIdLesson--->>>", logger.Any("req", req))

	resp, err := c.strg.Lesson().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdLesson--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *LessonService) GetList(ctx context.Context, req *schedule_service.GetListLessonRequest) (*schedule_service.GetListLessonResponse, error) {
	c.log.Info("---GetAllLesson--->>>", logger.Any("req", req))

	resp, err := c.strg.Lesson().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllLesson--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *LessonService) Update(ctx context.Context, req *schedule_service.UpdateLessonRequest) (*schedule_service.LSMessage, error) {
	c.log.Info("---UpdateLesson--->>>", logger.Any("req", req))

	resp, err := c.strg.Lesson().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateLesson--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *LessonService) Delete(ctx context.Context, req *schedule_service.LessonPrimaryKey) (*schedule_service.LSMessage, error) {
	c.log.Info("---DeleteLesson--->>>", logger.Any("req", req))

	resp, err := c.strg.Lesson().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteLesson--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}