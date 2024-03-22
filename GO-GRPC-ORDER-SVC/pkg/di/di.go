package di

import (
	server "github.com/rahulchacko7/pkg/api"
	"github.com/rahulchacko7/pkg/api/service"
	"github.com/rahulchacko7/pkg/client"
	"github.com/rahulchacko7/pkg/config"
	"github.com/rahulchacko7/pkg/db"
	"github.com/rahulchacko7/pkg/repository"
	"github.com/rahulchacko7/pkg/usecase"
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
