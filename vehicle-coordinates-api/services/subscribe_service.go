package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"vehicle-maps/db"
	"vehicle-maps/response"
)

type SubscribeService struct {
	statusRepo db.StatusRepo
	broker     *Broker
}

func NewSubscribeService(statusRepo db.StatusRepo, broker *Broker) *SubscribeService {
	return &SubscribeService{
		statusRepo: statusRepo,
		broker:     broker,
	}
}

func (s *SubscribeService) Subscribe(w http.ResponseWriter, r *http.Request, ctx context.Context) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		response.HandleErrorResponseErr(w, fmt.Errorf("flusher is not supported. Unable to generate SSE"))
	}
	tickerTime := 3 * time.Second
	keepAlive := time.NewTicker(tickerTime)
	sub := s.broker.Subscribe()

	defer func() {
		keepAlive.Stop()
		err := s.broker.Unsubscribe(sub.id)
		if err != nil {
			fmt.Printf("Error unsubscribing: %v", err)
		}
	}()

	for {
		select {
		case msg := <-sub.msgs:
			fmt.Fprintf(w, "event: status\ndata: %v\n\n", msg)
			keepAlive.Reset(tickerTime)
			flusher.Flush()
		case <-keepAlive.C:
			fmt.Fprint(w, ":keep-alive\n\n")
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
