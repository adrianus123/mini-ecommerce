package usecase

import (
	"order-service/internal/application/port"
	"order-service/internal/entity"
)

type PublishEventUsecase struct {
	publisher port.EventPublisher
}

func NewPublishEventUsecase(p port.EventPublisher) *PublishEventUsecase {
	return &PublishEventUsecase{
		publisher: p,
	}
}

func (u *PublishEventUsecase) Execute(event entity.Event) error {
	return u.publisher.Publish(event)
}
