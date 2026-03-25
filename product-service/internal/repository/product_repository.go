package repository

import "product-service/internal/entity"

type ProductRepository interface {
	GetProducts() ([]entity.Product, error)
	GetByID(id uint) (entity.Product, error)
}
