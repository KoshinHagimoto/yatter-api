package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	relationship struct {
		db *sqlx.DB
	}
)

func NewRelationship(db *sqlx.DB) repository.Relationship {
	return &relationship{db: db}
}

func (r *relationship) SaveRelationship(ctx context.Context, relationship *object.Relationship) error {
	_, err := r.db.ExecContext(ctx, "insert into relationship (follower_id, following_id) values (?, ?)",
		relationship.FollowerID, relationship.FollowingID)
	if err != nil {
		return fmt.Errorf("failed to insert relationship into db: %w", err)
	}

	return nil
}

func (r *relationship) IsFollowing(ctx context.Context, followerID, followingID int64) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "select count(*) from relationship where follower_id = ? and following_id = ?", followerID, followingID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to query relationship in db: %w", err)
	}

	return count > 0, nil
}
