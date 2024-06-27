package client

import (
	"schedule_service/config"
	"schedule_service/genproto/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	TeacherService() user_service.TeacherServiceClient
	BranchService() user_service.BranchServiceClient
	SupportTeacherService() user_service.SupportTeacherServiceClient
}

type grpcClients struct {
	teacherService user_service.TeacherServiceClient
	branchService user_service.BranchServiceClient
	supportteacherService user_service.SupportTeacherServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// Правильное формирование адреса сервиса
	connTeacherService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort), // Исправлено
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		teacherService: user_service.NewTeacherServiceClient(connTeacherService),
		branchService: user_service.NewBranchServiceClient(connTeacherService),
		supportteacherService: user_service.NewSupportTeacherServiceClient(connTeacherService),
	}, nil
}

func (g *grpcClients) TeacherService() user_service.TeacherServiceClient {
	return g.teacherService
}

func (g *grpcClients) BranchService() user_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) SupportTeacherService() user_service.SupportTeacherServiceClient {
	return g.supportteacherService
}