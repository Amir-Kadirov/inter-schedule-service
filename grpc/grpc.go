package grpc

import (
	"schedule_service/config"
	"schedule_service/genproto/genproto/user_service"
	"schedule_service/grpc/client"
	"schedule_service/grpc/service"
	"schedule_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	user_service.RegisterTeacherServiceServer(grpcServer, service.NewTeacherService(cfg, log, strg, srvc))

	user_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg, srvc))

	user_service.RegisterSupportTeacherServiceServer(grpcServer, service.NewSupportTeacherService(cfg, log, strg, srvc))

	user_service.RegisterAdminServiceServer(grpcServer, service.NewAdminService(cfg, log, strg, srvc))

	user_service.RegisterManagerServiceServer(grpcServer, service.NewManagerService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
