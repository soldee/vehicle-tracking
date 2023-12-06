package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"vehicle-maps/models"
)

type RandomConsumer struct{}

func NewRandomConsumer() *RandomConsumer {
	return &RandomConsumer{}
}

func (consumer *RandomConsumer) Read(ctx context.Context, broker *Broker, errorCh chan error) {

	for {
		routeId, vehicleId, userId := generateRouteUserVehicle()

		msg := models.Message{
			Timestamp:   generateTimestamp(),
			Coordinates: generateCoordinates(),
			Speed:       rand.Float64() * 10,
			RouteId:     routeId,
			VehicleId:   vehicleId,
			UserId:      userId,
		}

		fmt.Printf("New message: %v\n", msg)
		b, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		broker.Publish(string(b))
		fmt.Println("published")
		time.Sleep(time.Second * 10)
	}
	fmt.Println("Random consumer done")
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
