package routes

import (
	"vehicle-maps/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var coordinatesCollection *mongo.Collection = configs.GetCollection("vehicle-status")

func GetCoordinates() {
}