package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type JournalService struct {
	schedule_service.UnimplementedJournalServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewJournalService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *JournalService {
	return &JournalService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *JournalService) Create(ctx context.Context, req *schedule_service.CreateJournal) (*schedule_service.JournalPrimaryKey, error) {

	c.log.Info("---CreateJournal--->>>", logger.Any("req", req))

	resp, err := c.strg.Journal().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateJournal--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *JournalService) GetByID(ctx context.Context, req *schedule_service.JournalPrimaryKey) (*schedule_service.Journal, error) {
	c.log.Info("---GetByIdJournal--->>>", logger.Any("req", req))

	resp, err := c.strg.Journal().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdJournal--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *JournalService) GetList(ctx context.Context, req *schedule_service.GetListJournalRequest) (*schedule_service.GetListJournalResponse, error) {
	c.log.Info("---GetAllJournal--->>>", logger.Any("req", req))

	resp, err := c.strg.Journal().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllJournal--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *JournalService) Update(ctx context.Context, req *schedule_service.UpdateJournalRequest) (*schedule_service.JRMessage, error) {
	c.log.Info("---UpdateJournal--->>>", logger.Any("req", req))

	resp, err := c.strg.Journal().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateJournal--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *JournalService) Delete(ctx context.Context, req *schedule_service.JournalPrimaryKey) (*schedule_service.JRMessage, error) {
	c.log.Info("---DeleteJournal--->>>", logger.Any("req", req))

	resp, err := c.strg.Journal().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteJournal--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}