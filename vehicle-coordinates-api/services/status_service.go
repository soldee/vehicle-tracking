package services

import (
	"context"
	"errors"
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

	routeID, userID, vehicleID, dateFrom, dateTo, err := GetValidatedFilters(r)
	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}

	var statusList []models.Status
	var status *models.Status

	if routeIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteId(ctx, routeID)
	} else if routeIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdBetween(ctx, routeID, dateFrom, dateTo)
	} else if userIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByUserId(ctx, userID)
	} else if userIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByUserIdBetween(ctx, userID, dateFrom, dateTo)
	} else if vehicleIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByVehicleId(ctx, vehicleID)
	} else if vehicleIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByVehicleIdBetween(ctx, vehicleID, dateFrom, dateTo)
	} else if routeIDAndUserIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndUserIdBetween(ctx, routeID, userID, time.Time{}, time.Now())
	} else if routeIDAndVehicleIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndVehicleIdBetween(ctx, routeID, vehicleID, time.Time{}, time.Now())
	} else if userIDAndVehicleIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByUserIdAndVehicleIdBetween(ctx, userID, vehicleID, time.Time{}, time.Now())
	} else if routeIDAndUserIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndUserIdBetween(ctx, routeID, userID, dateFrom, dateTo)
	} else if userIDAndVehicleIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		statusList, err = service.statusRepo.FindByUserIdAndVehicleIdBetween(ctx, userID, vehicleID, dateFrom, dateTo)
	} else if routeIDAndVehicleIDAndDatesSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndVehicleIdBetween(ctx, routeID, vehicleID, dateFrom, dateTo)
	} else if routeIDAnduserIDAndVehicleIDSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndUserIdAndVehicleIdBetween(ctx, routeID, userID, vehicleID, time.Time{}, time.Now())
	} else if allSet(routeID, userID, vehicleID, dateFrom, dateTo) {
		status, err = service.statusRepo.FindByRouteIdAndUserIdAndVehicleIdBetween(ctx, routeID, userID, vehicleID, dateFrom, dateTo)
	} else {
		response.HandleErrorResponseErr(w, errors.New("invalid filters provided"))
		return
	}

	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}
	if statusList == nil {
		statusList = []models.Status{*status}
	}

	response.HandleJsonResponse(w, http.StatusOK, statusList)
}

func GetValidatedFilters(r *http.Request) (string, string, string, time.Time, time.Time, error) {
	errors := make([]string, 0)

	var routeID string = r.URL.Query().Get("route_id")
	var userID string = r.URL.Query().Get("user_id")
	var vehicleID string = r.URL.Query().Get("vehicle_id")
	var dateFromIn string = r.URL.Query().Get("date[gt]")
	var dateToIn string = r.URL.Query().Get("date[lt]")

	if routeID == "" && userID == "" && vehicleID == "" {
		return "", "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: "At least one of the following query parameters required: route_id, vehicle_id, user_id"}
	}

	var dateFrom time.Time
	var dateTo time.Time
	var err error
	if dateFromIn != "" && dateToIn != "" {
		dateFrom, dateTo, err = isDateRangeValid(&dateFromIn, &dateToIn, time.RFC3339)
		if err != nil {
			return "", "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: err.Error()}
		}
	}
	if (dateFromIn != "" && dateToIn == "") || (dateFromIn == "" && dateToIn != "") {
		return "", "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: "Invalid date range specified, date[gt] and date[lt] are required when specifying a date range"}
	}

	if len(errors) > 0 {
		var errorMessage = fmt.Sprintf("Invalid query parameters provided: %s", strings.Join(errors, ","))
		return "", "", "", time.Time{}, time.Time{}, &response.InvalidInput{Msg: errorMessage}
	}

	return routeID, userID, vehicleID, dateFrom, dateTo, nil
}

func isDateRangeValid(dateFromStr *string, dateToStr *string, format string) (time.Time, time.Time, error) {
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
			errorStr += err.Error() + ", "
		}
		errorStr = errorStr[:len(errorStr)-2]
		return time.Time{}, time.Time{}, fmt.Errorf("invalid date range specified: %v", errorStr)
	} else if dateFrom.After(dateTo) || dateFrom.Equal(dateTo) {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid date range specified: lower date bound must be lower than the higher date bound")
	}
	return dateFrom, dateTo, nil
}

func routeIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID == "" && vehicleID == "" && dateFrom.IsZero() && dateTo.IsZero()
}

func routeIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID == "" && vehicleID == "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func userIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID != "" && vehicleID == "" && dateFrom.IsZero() && dateTo.IsZero()
}

func userIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID != "" && vehicleID == "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func vehicleIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID == "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero()
}

func vehicleIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID == "" && vehicleID != "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func routeIDAndUserIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID != "" && vehicleID == "" && dateFrom.IsZero() && dateTo.IsZero()
}

func routeIDAndVehicleIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID == "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero()
}

func userIDAndVehicleIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID != "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero()
}

func routeIDAndUserIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID != "" && vehicleID == "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func routeIDAndVehicleIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID == "" && vehicleID != "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func userIDAndVehicleIDAndDatesSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID == "" && userID != "" && vehicleID != "" && !dateFrom.IsZero() && !dateTo.IsZero()
}

func routeIDAnduserIDAndVehicleIDSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID != "" && vehicleID != "" && dateFrom.IsZero() && dateTo.IsZero()
}

func allSet(routeID string, userID string, vehicleID string, dateFrom time.Time, dateTo time.Time) bool {
	return routeID != "" && userID != "" && vehicleID != "" && !dateFrom.IsZero() && !dateTo.IsZero()
}
