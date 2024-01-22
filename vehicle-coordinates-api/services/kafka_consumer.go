package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"vehicle-maps/models"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

func NewKafkaConsumer(reader *kafka.Reader) *KafkaConsumer {
	return &KafkaConsumer{
		Reader: reader,
	}
}

type KafkaMessage struct {
	Timestamp string `json:"ts"`
	Meta      struct {
		RouteId   string `json:"route_id"`
		UserId    string `json:"user_id"`
		VehicleId string `json:"vehicle_id"`
	} `json:"meta"`
	Speed    float64    `json:"speed"`
	Location [2]float64 `json:"location"`
}

func (consumer *KafkaConsumer) Run(ctx context.Context, broker *Broker) {

	defer func() {
		if err := consumer.Reader.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()

	for {
		kafkaMsg, err := ReadMessage(consumer)
		if err != nil {
			log.Printf("skipping message; error reading message from kafka: %v", err)
			break
		}

		err = PublishMessage(broker, kafkaMsg)
		if err != nil {
			log.Printf("skipping message; error publishing message to broker: %v", err)
		}
	}
}

func ReadMessage(consumer *KafkaConsumer) (*KafkaMessage, error) {
	msg, err := consumer.Reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}
	log.Printf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))

	kafkaMsg := KafkaMessage{}
	err = json.Unmarshal(msg.Value, &kafkaMsg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling kafka message to struct %v; error is: %v", msg.Value, err)
	}
	return &kafkaMsg, nil
}

func PublishMessage(broker *Broker, msg *KafkaMessage) error {
	dateFormat := time.RFC3339
	ts, err := time.Parse(dateFormat, msg.Timestamp)
	if err != nil {
		return fmt.Errorf("error parsing date '%v' using format '%v': error is: %v", msg.Timestamp, dateFormat, err)
	}

	brokerMsg := models.Message{
		Timestamp:   ts,
		Coordinates: msg.Location,
		Speed:       msg.Speed,
		RouteId:     msg.Meta.RouteId,
		VehicleId:   msg.Meta.VehicleId,
		UserId:      msg.Meta.UserId,
	}
	msgJson, err := json.Marshal(brokerMsg)
	if err != nil {
		return fmt.Errorf("error marshaling struct %v; error is: %v", msg, err)
	}
	broker.Publish(string(msgJson))
	return nil
}
