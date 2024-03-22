package di

import (
	server "github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/api"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/api/service"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/client"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/config"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/db"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/repository"
	"github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
