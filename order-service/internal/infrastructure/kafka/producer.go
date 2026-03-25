package kafka

import (
	"context"
	"encoding/json"
	"order-service/internal/entity"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &KafkaProducer{
		writer: writer,
	}
}

func (k *KafkaProducer) Publish(event entity.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.Owner),
		Value: value,
		Time:  time.Now(),
	}

	return k.writer.WriteMessages(ctx, msg)
}
