package db

import (
	"context"
	"fmt"
	"time"
	"vehicle-maps/models"
	"vehicle-maps/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatusRepo interface {
	FindByRouteId(ctx context.Context, RouteId string) (*models.Status, error)
	FindByRouteIdBetween(ctx context.Context, RouteId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error)
	FindByUserId(ctx context.Context, UserId string) ([]models.Status, error)
	FindByUserIdBetween(ctx context.Context, UserId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error)
	FindByVehicleId(ctx context.Context, VehicleId string) ([]models.Status, error)
	FindByVehicleIdBetween(ctx context.Context, VehicleId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error)
	FindByRouteIdAndUserIdBetween(ctx context.Context, RouteId string, UserId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error)
	FindByRouteIdAndVehicleIdBetween(ctx context.Context, RouteId string, VehicleId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error)
	FindByUserIdAndVehicleIdBetween(ctx context.Context, UserId string, VehicleId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error)
	FindByRouteIdAndUserIdAndVehicleIdBetween(ctx context.Context, RouteId string, UserId string, VehicleId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error)
}

type MongoStatusRepo struct {
	coordinatesCollection *mongo.Collection
}

func NewMongoStatusRepo(client *mongo.Client) *MongoStatusRepo {
	return &MongoStatusRepo{
		coordinatesCollection: client.Database("VEHICLE-TRACKING").Collection("vehicle-status"),
	}
}

func (repo *MongoStatusRepo) FindByRouteId(ctx context.Context, RouteId string) (*models.Status, error) {

	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.route_id", Value: RouteId},
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
		return nil, &response.InternalError{Msg: "GetByRouteId aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "GetByRouteId error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found", RouteId)}
	}
	return &results[0], nil
}

func (repo *MongoStatusRepo) FindByRouteIdBetween(ctx context.Context, RouteId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error) {
	status, err := repo.FindByRouteId(ctx, RouteId)
	if err != nil {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found", RouteId)}
	}

	routeStartTime := status.Timestamp[0].Time()

	if !(DateBetween(routeStartTime, DateFrom, DateTo)) {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v'", RouteId, DateFrom, DateTo)}
	}
	return status, nil
}

func (repo *MongoStatusRepo) FindByUserId(ctx context.Context, UserId string) ([]models.Status, error) {

	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.user_id", Value: UserId},
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
		return nil, &response.InternalError{Msg: "FindByUserId aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "FindByUserId error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("user_id %v not found", UserId)}
	}

	return results, nil
}

func (repo *MongoStatusRepo) FindByUserIdBetween(ctx context.Context, UserId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error) {
	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.user_id", Value: UserId},
			{Key: "ts", Value: bson.D{{Key: "$gt", Value: DateFrom}, {Key: "$lt", Value: DateTo}}},
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
		return nil, &response.InternalError{Msg: "FindByUserIdBetween aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "FindByUserIdBetween error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("No routes found for user_id %v between the specified dates %v-%v", UserId, DateFrom, DateTo)}
	}
	return results, nil
}

func (repo *MongoStatusRepo) FindByVehicleId(ctx context.Context, VehicleId string) ([]models.Status, error) {
	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.vehicle_id", Value: VehicleId},
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
		return nil, &response.InternalError{Msg: "FindByVehicleId aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "FindByVehicleId error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("vehicle_id %v not found", VehicleId)}
	}

	return results, nil
}

func (repo *MongoStatusRepo) FindByVehicleIdBetween(ctx context.Context, VehicleId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error) {
	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.vehicle_id", Value: VehicleId},
			{Key: "ts", Value: bson.D{{Key: "$gt", Value: DateFrom}, {Key: "$lt", Value: DateTo}}},
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
		return nil, &response.InternalError{Msg: "FindByVehicleIdBetween aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "FindByVehicleIdBetween error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("No routes found for vehicle_id %v between the specified dates %v-%v", VehicleId, DateFrom, DateTo)}
	}
	return results, nil
}

func (repo *MongoStatusRepo) FindByRouteIdAndUserIdBetween(ctx context.Context, RouteId string, UserId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error) {
	status, err := repo.FindByRouteId(ctx, RouteId)
	if err != nil {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found", RouteId)}
	}

	routeStartTime := status.Timestamp[0].Time()

	if !(DateBetween(routeStartTime, DateFrom, DateTo)) {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v'", RouteId, DateFrom, DateTo)}
	}
	if status.UserId != UserId {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v' for the user '%v'", RouteId, DateFrom, DateTo, UserId)}
	}
	return status, nil
}

func (repo *MongoStatusRepo) FindByRouteIdAndVehicleIdBetween(ctx context.Context, RouteId string, VehicleId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error) {
	status, err := repo.FindByRouteId(ctx, RouteId)
	if err != nil {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found", RouteId)}
	}

	routeStartTime := status.Timestamp[0].Time()

	if !(DateBetween(routeStartTime, DateFrom, DateTo)) {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v'", RouteId, DateFrom, DateTo)}
	}
	if status.VehicleId != VehicleId {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v' for the vehicle '%v'", RouteId, DateFrom, DateTo, VehicleId)}
	}
	return status, nil
}

func (repo *MongoStatusRepo) FindByUserIdAndVehicleIdBetween(ctx context.Context, UserId string, VehicleId string, DateFrom time.Time, DateTo time.Time) ([]models.Status, error) {
	cursor, err := repo.coordinatesCollection.Aggregate(ctx, mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "meta.user_id", Value: UserId},
			{Key: "ts", Value: bson.D{{Key: "$gt", Value: DateFrom}, {Key: "$lt", Value: DateTo}}},
			{Key: "meta.vehicle_id", Value: VehicleId},
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
		return nil, &response.InternalError{Msg: "FindByUserIdAndVehicleIdBetween aggregation generated error:" + err.Error()}
	}

	defer cursor.Close(ctx)

	var results []models.Status
	if err = cursor.All(ctx, &results); err != nil {
		return nil, &response.InternalError{Msg: "FindByUserIdAndVehicleIdBetween error unmarshalling bson response into Status model: " + err.Error()}
	}

	if len(results) == 0 {
		return nil, &response.NotFound{Msg: fmt.Sprintf("No routes found for vehicle_id %v and user_id %v between the specified dates %v-%v", VehicleId, UserId, DateFrom, DateTo)}
	}
	return results, nil
}

func (repo *MongoStatusRepo) FindByRouteIdAndUserIdAndVehicleIdBetween(ctx context.Context, RouteId string, UserId string, VehicleId string, DateFrom time.Time, DateTo time.Time) (*models.Status, error) {
	status, err := repo.FindByRouteId(ctx, RouteId)
	if err != nil {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found", RouteId)}
	}

	routeStartTime := status.Timestamp[0].Time()

	if !(DateBetween(routeStartTime, DateFrom, DateTo)) {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v'", RouteId, DateFrom, DateTo)}
	}
	if status.VehicleId != VehicleId {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v' for the vehicle '%v'", RouteId, DateFrom, DateTo, VehicleId)}
	}
	if status.UserId != UserId {
		return nil, &response.NotFound{Msg: fmt.Sprintf("route_id %v not found in the specified date range '%v' '%v' for the vehicle '%v' and user %v", RouteId, DateFrom, DateTo, VehicleId, UserId)}
	}
	return status, nil
}

func DateBetween(t, dateFrom, dateTo time.Time) bool {
	return (t.Equal(dateFrom) || t.After(dateFrom)) && (t.Equal(dateTo) || t.Before(dateTo))
}
