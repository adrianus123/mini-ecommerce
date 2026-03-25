package repository

import (
	"context"
	"inventory-service/internal/entity"
	"log"

	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) Save(ctx context.Context, event entity.Event) error {
	log.Println("Saving event: ", event)
	return nil
}
