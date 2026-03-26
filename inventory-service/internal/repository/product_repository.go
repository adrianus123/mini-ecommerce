package repository

import "inventory-service/internal/entity"

type ProductRepository interface {
	GetProductByIdIn(ids []uint) ([]entity.Product, error)
	SaveProducts(products []entity.Product) error
}
