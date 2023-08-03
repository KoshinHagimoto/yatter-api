package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	SaveRelationship(ctx context.Context, relationship *object.Relationship) error
	DeleteRelationship(ctx context.Context, followerID, followingID int64) error
	IsFollowing(ctx context.Context, followerID, followingID int64) (bool, error)
}
