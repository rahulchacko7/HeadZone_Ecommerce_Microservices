package di

import (
	server "github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/api"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/api/service"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/db"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/repository"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/usecase"
	"github.com/rahulchacko7/GO-GRPC-USER-SVC/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(adminRepository)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
