package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	SaveRelationship(ctx context.Context, relationship *object.Relationship) error
	DeleteRelationship(ctx context.Context, followerID, followingID int64) error
	IsFollowing(ctx context.Context, followerID, followingID int64) (bool, error)
	GetFollowing(ctx context.Context, followerID, limit int64) ([]*object.Account, error) //followerIDを渡して、フォローしている人を返す
	GetFollowers(ctx context.Context, followingID int64) ([]*object.Account, error)       //followingIDを渡して、フォローされている人を返す
}
