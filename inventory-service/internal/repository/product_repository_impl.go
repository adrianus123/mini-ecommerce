package repository

import (
	"inventory-service/internal/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) GetProductByIdIn(ids []uint) ([]entity.Product, error) {
	var products []entity.Product

	result := r.db.Where("id IN ?", ids).Find(&products)

	return products, result.Error
}

func (r *ProductRepo) SaveProducts(products []entity.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			if err := tx.Save(&product).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
