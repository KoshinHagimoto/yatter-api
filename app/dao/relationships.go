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

func (r *relationship) DeleteRelationship(ctx context.Context, followerID, followingID int64) error {
	_, err := r.db.ExecContext(ctx, "delete from relationship where follower_id = ? and following_id = ?", followerID, followingID)
	if err != nil {
		return fmt.Errorf("failed to delete relationship from db: %w", err)
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

func (r *relationship) GetFollowing(ctx context.Context, followerID, limit int64) ([]*object.Account, error) {
	var accounts []*object.Account
	err := r.db.SelectContext(ctx, &accounts, "select * from account where id in (select following_id from relationship where follower_id = ?) LIMIT ?", followerID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query relationship in db: %w", err)
	}

	return accounts, nil
}

func (r *relationship) GetFollowers(ctx context.Context, followingID int64, timeline *object.Timeline) ([]*object.Account, error) {
	var accounts []*object.Account

	query := "select * from account where id in (select follower_id from relationship where following_id = ?)"
	args := []interface{}{followingID}

	if timeline.MaxID != nil {
		query += " and id <= ?"
		args = append(args, *timeline.MaxID)
	}

	if timeline.SinceID != nil {
		query += " and id >= ?"
		args = append(args, *timeline.SinceID)
	}

	query += " order by id desc"

	if timeline.Limit != nil {
		query += " limit ?"
		args = append(args, *timeline.Limit)
	}

	err := r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query relationship in db: %w", err)
	}
	return accounts, nil
}

func (r *relationship) GetFollowerCount(ctx context.Context, accountID int64) (int64, error) {
	var count int64
	err := r.db.QueryRowContext(ctx, "select count(*) from relationship where following_id = ?", accountID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to query relationship in db: %w", err)
	}

	return count, nil
}

func (r *relationship) GetFollowingCount(ctx context.Context, accountID int64) (int64, error) {
	var count int64
	err := r.db.QueryRowContext(ctx, "select count(*) from relationship where follower_id = ?", accountID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to query relationship in db: %w", err)
	}

	return count, nil
}
