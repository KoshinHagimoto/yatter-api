package utils

import (
	"context"
	"yatter-backend-go/app/domain/repository"
)

func FetchFollowRelationship(ctx context.Context, rr repository.Relationship, sourceID, targetID int64) (bool, bool, error) {
	isFollowing, err := rr.IsFollowing(ctx, sourceID, targetID)
	if err != nil {
		return false, false, err
	}
	isFollowedBy, err := rr.IsFollowing(ctx, targetID, sourceID)
	if err != nil {
		return false, false, err
	}
	return isFollowing, isFollowedBy, nil
}
