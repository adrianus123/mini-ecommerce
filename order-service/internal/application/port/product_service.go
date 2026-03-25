package port

import (
	"order-service/internal/delivery/http/dto"
)

type ProductService interface {
	GetProduct(productID int) (*dto.Product, error)
}
