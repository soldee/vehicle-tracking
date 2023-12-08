package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"vehicle-maps/models"
)

type RandomConsumer struct{}

func NewRandomConsumer() *RandomConsumer {
	return &RandomConsumer{}
}

func (consumer *RandomConsumer) Read(ctx context.Context, broker *Broker) {

	for {
		if rand.Float32() < 0.1 {
			err := generateError()
			sendError(broker, err)
			sendError(broker, err)
			sendError(broker, err)
		}
		routeId, vehicleId, userId := generateRouteUserVehicle()

		msg := models.Message{
			Timestamp:   generateTimestamp(),
			Coordinates: generateCoordinates(),
			Speed:       rand.Float64() * 10,
			RouteId:     routeId,
			VehicleId:   vehicleId,
			UserId:      userId,
		}

		fmt.Printf("Generated new message: %v\n", msg)
		b, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		broker.Publish(string(b))
		time.Sleep(time.Second * 10)
	}
}

func generateError() error {
	return errors.New("unavailable sink, try reconnecting")
}

func sendError(broker *Broker, err error) {
	broker.PublishError(err)
	fmt.Printf("Generated error: %v", err.Error())
	time.Sleep(time.Second * 5)
}

func generateTimestamp() time.Time {
	min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2023, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(rand.Int63n(max-min)+min, 0)
}

func generateCoordinates() [2]float64 {
	return [2]float64{rand.Float64() * 10, rand.Float64() * 10}
}

func generateRouteUserVehicle() (string, string, string) {
	type RandomRouteUserVehicle struct {
		RouteId   string
		VehicleId string
		UserId    string
	}

	l := []RandomRouteUserVehicle{
		{"1", "1", "1"},
		{"2", "2", "2"},
		{"3", "3", "3"},
		{"4", "4", "4"},
	}

	choice := l[rand.Intn(len(l))]
	return choice.RouteId, choice.VehicleId, choice.UserId
}
