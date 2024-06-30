package client

import (
	"fmt"
	"schedule_service/config"
	"schedule_service/genproto/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	TeacherService() user_service.TeacherServiceClient
	BranchService() user_service.BranchServiceClient
	SupportTeacherService() user_service.SupportTeacherServiceClient
	Close() error
}

type grpcClients struct {
	teacherService        user_service.TeacherServiceClient
	branchService         user_service.BranchServiceClient
	supportTeacherService user_service.SupportTeacherServiceClient
	conn                  *grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	// Correct service address formation
	conn, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort), // Fixed
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		teacherService:        user_service.NewTeacherServiceClient(conn),
		branchService:         user_service.NewBranchServiceClient(conn),
		supportTeacherService: user_service.NewSupportTeacherServiceClient(conn),
		conn:                  conn,
	}, nil
}

func (g *grpcClients) TeacherService() user_service.TeacherServiceClient {
	return g.teacherService
}

func (g *grpcClients) BranchService() user_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) SupportTeacherService() user_service.SupportTeacherServiceClient {
	return g.supportTeacherService
}

func (g *grpcClients) Close() error {
	return g.conn.Close()
}
