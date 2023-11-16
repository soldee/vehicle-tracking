package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status struct {
	Timestamp   []primitive.DateTime `json:"ts" bson:"ts"`
	Coordinates [][2]float64         `json:"coordinates" bson:"coordinates"`
	Speed       []float64            `json:"speed" bson:"speed"`
	RouteId     string               `json:"route_id" bson:"route_id"`
	VehicleId   string               `json:"vehicle_id" bson:"vehicle_id"`
	UserId      string               `json:"user_id" bson:"user_id"`
}
