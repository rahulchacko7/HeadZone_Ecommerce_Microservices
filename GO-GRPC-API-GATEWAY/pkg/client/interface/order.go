package interfaces

import "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

type OrderClient interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (models.OrderSuccessResponse, error)
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
