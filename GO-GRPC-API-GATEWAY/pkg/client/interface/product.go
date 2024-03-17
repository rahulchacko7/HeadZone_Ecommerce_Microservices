package interfaces

import "rahulchacko7/GO-GRPC-API-GATEWAY/pkg/utils/models"

type ProductClient interface {
	ShowAllProducts(page int, count int) ([]models.ProductBrief, error)
	AddProducts(product models.Product) (models.Products, error)
	DeleteProduct(id int) error
	UpdateProducts(pid int, stock int) (models.ProductUpdateReciever, error)
}
