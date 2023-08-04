package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetPublicTimeline(ctx context.Context, timeline *object.Timeline) ([]*object.Status, error)
	GetHomeTimeline(ctx context.Context, accountID int64, timeline *object.Timeline) ([]*object.Status, error)
}
