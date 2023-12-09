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

	sse, err := NewSse(w, r)
	if err != nil {
		response.HandleErrorResponseErr(w, err)
		return
	}

	AddSSEheaders(w)
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
			sse.Publish("status", msg)
			keepAlive.Reset(tickerTime)
		case err := <-sub.errors:
			sse.Publish("close", err.Error())
			return
		case <-keepAlive.C:
			sse.PublishComment(":keep-alive")
		case <-r.Context().Done():
			return
		}
	}
}

func AddSSEheaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

type Sse struct {
	w       http.ResponseWriter
	r       *http.Request
	flusher http.Flusher
}

func NewSse(w http.ResponseWriter, r *http.Request) (*Sse, error) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("flusher is not supported. Unable to generate SSE")
	}

	return &Sse{
		w:       w,
		r:       r,
		flusher: flusher,
	}, nil
}

func (sse *Sse) Publish(event string, msg string) {
	fmt.Fprintf(sse.w, "event: %v\ndata: %v\n\n", event, msg)
	sse.flusher.Flush()
}

func (sse *Sse) PublishComment(msg string) {
	fmt.Fprintf(sse.w, "%v\n\n", msg)
	sse.flusher.Flush()
}
