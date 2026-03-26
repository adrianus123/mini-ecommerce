package repository

import (
	"context"
	"inventory-service/internal/entity"
	"log"
)

type eventRepository struct{}

func NewEventRepository() EventRepository {
	return &eventRepository{}
}

func (r *eventRepository) Save(ctx context.Context, event entity.Event) error {
	log.Println("Saving event: ", event)
	return nil
}
