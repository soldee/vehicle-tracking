package services

import (
	"context"
	"net/http"
	"vehicle-maps/db"
	"vehicle-maps/response"
)

type SubscribeService struct {
	statusRepo db.StatusRepo
}

func NewSubscribeService(statusRepo db.StatusRepo) *SubscribeService {
	return &SubscribeService{
		statusRepo: statusRepo,
	}
}

func (s *SubscribeService) Subscribe(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	response.HandleJsonResponse(w, 501, "Not implemented")
}
