package services

import (
	"context"
	"net/http"
	"vehicle-maps/db"
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

	var routeID string = r.URL.Query().Get("route_id")
	if routeID == "" {
		response.HandleErrorResponseErr(w, &response.InvalidInput{Msg: "query parameter route_id required"})
		return
	}

	status, err := service.statusRepo.GetByRouteId(ctx, routeID)
	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}

	response.HandleJsonResponse(w, http.StatusOK, status)
}
