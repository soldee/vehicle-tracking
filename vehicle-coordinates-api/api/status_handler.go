package api

import (
	"context"
	"net/http"
	"time"
	"vehicle-maps/services"
)

type StatusHandler struct {
	StatusService    *services.StatusService
	SubscribeService *services.SubscribeService
}

func (handler StatusHandler) HandleGetStatus(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	handler.StatusService.GetRouteById(w, r, ctx)
}

func (handler StatusHandler) HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	var ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	handler.SubscribeService.Subscribe(w, r, ctx)
}
