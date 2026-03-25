package port

import (
	"order-service/internal/entity"
)

type EventPublisher interface {
	Publish(event entity.Event) error
}
