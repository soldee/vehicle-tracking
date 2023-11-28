package services

import (
	"context"
	"vehicle-maps/models"
)

type QueueConsumer interface {
	Read(ctx context.Context, msgCh chan models.Message, errorCh chan error)
}
