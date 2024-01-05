package db

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func KafkaInstance() *kafka.Reader {
	godotenv.Load(".env")

	brokersStr := os.Getenv("KAFKA_BROKERS")
	if brokersStr == "" {
		log.Fatal("KAFKA_BROKERS env variable not found")
	}
	brokers := strings.Split(brokersStr, ",")

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		log.Fatal("KAFKA_TOPIC env variable not found")
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		Partition:   0,
		MinBytes:    1e3,  // 1KB
		MaxBytes:    10e6, // 10MB
		StartOffset: 0,
	})
}
