package service

import (
	"context"
	"schedule_service/config"
	"schedule_service/genproto/genproto/schedule_service"
	"schedule_service/grpc/client"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type EventService struct {
	schedule_service.UnimplementedEventServiceServer
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewEventService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventService {
	return &EventService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *EventService) Create(ctx context.Context, req *schedule_service.CreateEvent) (*schedule_service.EventPrimaryKey, error) {

	c.log.Info("---CreateEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *EventService) GetByID(ctx context.Context, req *schedule_service.EventPrimaryKey) (*schedule_service.Event, error) {
	c.log.Info("---GetByIdEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *EventService) GetList(ctx context.Context, req *schedule_service.GetListEventRequest) (*schedule_service.GetListEventResponse, error) {
	c.log.Info("---GetAllEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *EventService) Update(ctx context.Context, req *schedule_service.UpdateEventRequest) (*schedule_service.EVMessage, error) {
	c.log.Info("---UpdateEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *EventService) Delete(ctx context.Context, req *schedule_service.EventPrimaryKey) (*schedule_service.EVMessage, error) {
	c.log.Info("---DeleteEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *EventService) RegisterEvent(ctx context.Context, req *schedule_service.RegisterEv) (*schedule_service.EVMessage, error) {
	c.log.Info("---RegisterEvent--->>>", logger.Any("req", req))

	resp, err := c.strg.Event().RegisterEvent(ctx, req)
	if err != nil {
		c.log.Error("---RegisterEvent--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}