package kafka

import (
	"context"
	"inventory-service/internal/usecase"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	usecase usecase.EventUsecase
}

func NewConsumer(broker []string, topic string, groupID string, uc usecase.EventUsecase) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: broker,
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		reader:  r,
		usecase: uc,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Println("Message received:", string(msg.Value))

		err = c.usecase.ProcessEvent(ctx, msg.Value)
		if err != nil {
			log.Println("Error processing event:", err)
		}
	}
}
