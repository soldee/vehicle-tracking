package models

import "time"

type Message struct {
	Timestamp   time.Time  `json:"ts"`
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
	Speed       float64    `json:"speed" bson:"speed"`
	RouteId     string     `json:"route_id" bson:"route_id"`
	VehicleId   string     `json:"vehicle_id" bson:"vehicle_id"`
	UserId      string     `json:"user_id" bson:"user_id"`
}
