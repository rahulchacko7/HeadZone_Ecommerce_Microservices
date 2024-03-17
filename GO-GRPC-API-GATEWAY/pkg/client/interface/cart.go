package interfaces

import "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

type CartClient interface {
	AddToCart(product_id, user_id, quantity int) (models.CartResponse, error)
	GetCart(user_id int) (models.CartResponse, error)
}
