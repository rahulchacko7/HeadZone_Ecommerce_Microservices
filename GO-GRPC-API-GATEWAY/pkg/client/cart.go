package client

import (
	"context"
	"errors"
	"fmt"
	"rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

	"google.golang.org/grpc"
	"rahulchacko7/GO-GRPC-API-GATEWAY/pkg/config"

	pb "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/pb/cart"

	interfaces "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/client/interface"
)

type cartClient struct {
	Client pb.CartClient
}

func NewCartClient(cfg config.Config) interfaces.CartClient {

	grpcConnection, err := grpc.Dial(cfg.CartSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewCartClient(grpcConnection)

	return &cartClient{
		Client: grpcClient,
	}

}
func (c *cartClient) AddToCart(product_id int, user_id int, quantity int) (models.CartResponse, error) {
	res, err := c.Client.AddToCart(context.Background(), &pb.AddToCartRequest{
		ProductID: int64(product_id),
		UserID:    int64(user_id),
		Quantity:  int64(quantity),
	})
	if err != nil {
		return models.CartResponse{}, err
	}
	if res.Error != "" {
		return models.CartResponse{}, errors.New(res.Error)
	}
	var carts []models.Cart
	for _, cartDetails := range res.Cart {
		cart := models.Cart{
			ProductID:  uint(cartDetails.ProductID),
			Quantity:   float64(cartDetails.Quantity),
			TotalPrice: float64(cartDetails.TotalPrice),
		}
		carts = append(carts, cart)
	}

	return models.CartResponse{
		TotalPrice: float64(res.Price),
		Cart:       carts,
	}, nil
}
func (c *cartClient) GetCart(user_id int) (models.CartResponse, error) {
	res, err := c.Client.GetCart(context.Background(), &pb.GetCartRequest{
		UserID: int64(user_id),
	})
	if err != nil {
		return models.CartResponse{}, err
	}
	if res.Error != "" {
		return models.CartResponse{}, errors.New(res.Error)
	}
	var carts []models.Cart
	for _, cartDetails := range res.Cart {
		cart := models.Cart{
			ProductID:  uint(cartDetails.ProductID),
			Quantity:   float64(cartDetails.Quantity),
			TotalPrice: float64(cartDetails.TotalPrice),
		}
		carts = append(carts, cart)
	}

	return models.CartResponse{
		TotalPrice: float64(res.Price),
		Cart:       carts,
	}, nil
}
