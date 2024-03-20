package di

import (
	server "github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/api"
	"github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/config"
	"github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/db"
	"github.com/rahulchacko7/GO-GRPC-ORDER-SVC/pkg/repository"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(&cfg)
	productClient := client.NewProductClient(&cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)

	orderServiceServer := service.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
