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
	repo repository.EventRepository
}

func NewEventUseCase(repo repository.EventRepository) *eventUsecase {
	return &eventUsecase{
		repo: repo,
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

	return u.repo.Save(ctx, event)
}
