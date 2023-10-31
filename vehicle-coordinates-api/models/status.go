package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Metadata struct {
	VehicleId primitive.ObjectID `bson:"vehicle_id"`
	UserId    primitive.ObjectID `bson:"user_id"`
	RouteId   primitive.ObjectID `bson:"route_id"`
}

type Location struct {
	Type        string     `bson:"type"`
	Coordinates [2]float64 `bson:"coordinates"`
}

type Status struct {
	ID        primitive.ObjectID `bson:"_id"`
	Timestamp primitive.DateTime `bson:"ts"`
	Metadata  Metadata           `bson:"meta"`
	Location  Location           `bson:"location"`
	Speed     float32            `bson:"speed"`
}
