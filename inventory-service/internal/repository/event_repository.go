package repository

import (
	"context"
	"inventory-service/internal/entity"
)

type EventRepository interface {
	Save(ctx context.Context, event entity.Event) error
}
