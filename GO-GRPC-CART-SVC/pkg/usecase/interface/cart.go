package interfaces

import "github.com/rahulchacko7/GO-GRPC-CART-SVC/pkg/utils/models"

type CartUseCase interface {
	AddToCart(product_id, user_id, quantity int) (models.CartResponse, error)
	DisplayCart(user_id int) (models.CartResponse, error)
	GetAllItemsFromCart(userID int) ([]models.Cart, error)
	DoesCartExist(userID int) (bool, error)
	TotalAmountInCart(userID int) (float64, error)
	UpdateCartAfterOrder(userID, productID int, quantity float64) error
}
