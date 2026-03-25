package repository

import (
	"errors"
	"product-service/internal/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) GetProducts() ([]entity.Product, error) {
	var products []entity.Product
	result := r.db.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepo) GetByID(id uint) (entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Product{}, errors.New("Product not found")
		}

		return entity.Product{}, result.Error
	}

	return product, nil
}
