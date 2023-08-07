package utils

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

func UpdateFollowCounts(ctx context.Context, rr repository.Relationship, account *object.Account) error {
	var err error
	account.FollowerCount, err = rr.GetFollowerCount(ctx, account.ID)
	if err != nil {
		return err
	}
	account.FollowingCount, err = rr.GetFollowingCount(ctx, account.ID)
	if err != nil {
		return err
	}
	return nil
}
