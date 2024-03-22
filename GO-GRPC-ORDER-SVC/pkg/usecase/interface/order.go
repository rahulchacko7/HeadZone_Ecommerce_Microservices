package interfaces

import (
	"github.com/rahulchacko7/pkg/domain"
	"github.com/rahulchacko7/pkg/util/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
