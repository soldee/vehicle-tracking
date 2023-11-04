package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"vehicle-maps/configs"
	"vehicle-maps/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var coordinatesCollection *mongo.Collection = configs.GetCollection("vehicle-status")

func GetStatusByRouteId(w http.ResponseWriter, r *http.Request) {
	var routeID string = r.URL.Query().Get("route_id")
	if routeID == "" {
		response.HandleErrorResponse(w, http.StatusUnprocessableEntity, errors.New("query parameter route_id required"))
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.route_id", Value: routeID},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "null"},
			{Key: "ts", Value: bson.D{{Key: "$push", Value: "$ts"}}},
			{Key: "coordinates", Value: bson.D{{Key: "$push", Value: "$location.coordinates"}}},
			{Key: "speed", Value: bson.D{{Key: "$push", Value: bson.D{{Key: "$trunc", Value: bson.A{"$speed", 2}}}}}},
			{Key: "route_id", Value: bson.D{{Key: "$first", Value: "$meta.route_id"}}},
			{Key: "vehicle_id", Value: bson.D{{Key: "$first", Value: "$meta.vehicle_id"}}},
			{Key: "user_id", Value: bson.D{{Key: "$first", Value: "$meta.user_id"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
		}}},
	})
	if err != nil {
		response.HandleErrorResponse(w, 503, err)
		return
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		response.HandleErrorResponse(w, 503, err)
		return
	}

	if len(results) == 0 {
		response.HandleErrorResponse(w, http.StatusNotFound, fmt.Errorf("route_id not found: %v", routeID))
		return
	}
	response.HandleJsonResponse(w, http.StatusOK, results[0])
}
