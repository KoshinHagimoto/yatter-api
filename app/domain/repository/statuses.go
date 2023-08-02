package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	SaveStatus(ctx context.Context, status *object.Status) error
	FindStatusByID(ctx context.Context, id int64) (*object.Status, error)
}