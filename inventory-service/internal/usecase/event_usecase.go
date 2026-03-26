package usecase

import (
	"context"
	"encoding/json"
	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
)

type EventUsecase interface {
	ProcessEvent(ctx context.Context, data []byte) error
}

type eventUsecase struct {
	repo              repository.EventRepository
	productRepository repository.ProductRepository
}

func NewEventUseCase(repo repository.EventRepository, productRepository repository.ProductRepository) *eventUsecase {
	return &eventUsecase{
		repo:              repo,
		productRepository: productRepository,
	}
}

func (u *eventUsecase) ProcessEvent(ctx context.Context, data []byte) error {
	var event entity.Event

	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	if event.Action != "UPDATE" {
		return nil
	}

	var productIds []uint
	substractItem := map[uint]uint{}

	for _, e := range event.Value {
		productIds = append(productIds, e.ProductID)
		substractItem[e.ProductID] = e.Qty
	}

	products, err := u.productRepository.GetProductByIdIn(productIds)
	if err != nil {
		return err
	}

	var updatedProducts []entity.Product
	for _, product := range products {
		product.Qty -= substractItem[product.ID]
		updatedProducts = append(updatedProducts, product)
	}

	err = u.productRepository.SaveProducts(updatedProducts)
	if err != nil {
		return err
	}

	return nil
}
