package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"product-service/internal/entity"
	"product-service/internal/infrastructure/redis"
	"product-service/internal/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
	redisRepo  redis.RedisRepository
}

func NewProductUsecase(r repository.ProductRepository, redisRepo redis.RedisRepository) *ProductUsecase {
	return &ProductUsecase{
		repository: r,
		redisRepo:  redisRepo,
	}
}

func (u *ProductUsecase) GetProducts() ([]entity.Product, error) {
	products, err := u.repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *ProductUsecase) GetByID(id uint) (entity.Product, error) {
	productKey := "product:" + fmt.Sprint(id)

	cached, err := u.redisRepo.GetByKey(productKey)
	if err != nil {
		log.Println("Error get cache product: ", err.Error())
	}

	if cached != "" {
		var product entity.Product
		err = json.Unmarshal([]byte(cached), &product)

		if err != nil {
			return entity.Product{}, err
		}

		log.Println("Get product from redis")

		return product, nil
	}

	product, err := u.repository.GetByID(id)
	if err != nil {
		return entity.Product{}, err
	}

	strProduct, err := json.Marshal(product)
	if err != nil {
		log.Println(err.Error())
	}

	u.redisRepo.SetByKey(productKey, string(strProduct))

	log.Println("Get product from db")

	return product, nil
}
