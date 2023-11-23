package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"vehicle-maps/db"
	"vehicle-maps/models"
	"vehicle-maps/response"
)

type StatusService struct {
	statusRepo db.StatusRepo
}

func NewStatusService(statusRepo db.StatusRepo) *StatusService {
	return &StatusService{
		statusRepo: statusRepo,
	}
}

func (service *StatusService) GetRouteById(w http.ResponseWriter, r *http.Request, ctx context.Context) {

	routeID, vehicleID, dateFrom, dateTo, err := GetValidatedFilters(r)
	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}

	var statusList []*models.Status

	if routeID != "" && vehicleID == "" && dateFrom.IsZero() && dateTo.IsZero() {
		var status *models.Status
		status, err = service.statusRepo.FindByRouteId(ctx, routeID)
		statusList = []*models.Status{status}
	} else if routeID == "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero() {
		statusList, err = service.statusRepo.FindByVehicleId(ctx, vehicleID)
	} else if routeID != "" && vehicleID == "" && !dateFrom.IsZero() && !dateTo.IsZero() {
		var status *models.Status
		status, err = service.statusRepo.FindByRouteIdBetween(ctx, routeID, dateFrom, dateTo)
		statusList = []*models.Status{status}
	} else if routeID == "" && vehicleID != "" && !dateFrom.IsZero() && !dateTo.IsZero() {
		statusList, err = service.statusRepo.FindByVehicleIdBetween(ctx, vehicleID, dateFrom, dateTo)
	} else if routeID != "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero() {
		var status *models.Status
		status, err = service.statusRepo.FindByRouteIdAndVehicleIdBetween(ctx, routeID, vehicleID, time.Time{}, time.Now())
		statusList = []*models.Status{status}
	} else {
		var status *models.Status
		status, err = service.statusRepo.FindByRouteIdAndVehicleIdBetween(ctx, routeID, vehicleID, dateFrom, dateTo)
		statusList = []*models.Status{status}
	}

	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}

	response.HandleJsonResponse(w, http.StatusOK, statusList)
}

func GetValidatedFilters(r *http.Request) (string, string, time.Time, time.Time, error) {
	errors := make([]string, 0)

	var routeID string = r.URL.Query().Get("route_id")
	var vehicleID string = r.URL.Query().Get("vehicle_id")
	var dateFromIn string = r.URL.Query().Get("date[gt]")
	var dateToIn string = r.URL.Query().Get("date[lt]")

	if routeID == "" && vehicleID == "" && dateFromIn == "" && dateToIn == "" {
		return "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: "At least one of the following query parameters required: route_id, vehicle_id, date range (date[gt] and date[lt])"}
	}

	var dateFrom time.Time
	var dateTo time.Time
	var err error
	if dateFromIn != "" && dateToIn != "" {
		dateFrom, dateTo, err = isDateRangeValid(&dateFromIn, &dateToIn)
		if err != nil {
			return "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: err.Error()}
		}
	}
	if (dateFromIn != "" && dateToIn == "") || (dateFromIn == "" && dateToIn != "") {
		return "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: "Invalid date range specified, date[gt] and date[lt] are required when specifying a date range"}
	}

	if len(errors) > 0 {
		var errorMessage = fmt.Sprintf("Invalid query parameters provided: %s", strings.Join(errors, ","))
		return "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: errorMessage}
	}

	return routeID, vehicleID, dateFrom, dateTo, nil
}

func isDateRangeValid(dateFromStr *string, dateToStr *string) (time.Time, time.Time, error) {
	format := time.RFC3339

	errors := make([]error, 0)

	dateFrom, err := time.Parse(format, *dateFromStr)
	if err != nil {
		fmt.Printf("Error parsing date '%v' using format '%v'. Error is: %v", dateFrom, format, err)
		errors = append(errors, fmt.Errorf("invalid date specified '%v', expected '%v' format", dateFrom, format))
	}

	dateTo, err := time.Parse(format, *dateToStr)
	if err != nil {
		fmt.Printf("Error parsing date '%v' using format '%v'. Error is: %v", dateTo, format, err)
		errors = append(errors, fmt.Errorf("invalid date specified '%v', expected '%v' format", dateTo, format))
	}

	if len(errors) > 0 {
		var errorStr string
		for _, err := range errors {
			errorStr += err.Error()
		}
		return time.Time{}, time.Time{}, fmt.Errorf("invalid date range specified: %v", errorStr)
	}
	return dateFrom, dateTo, nil
}
