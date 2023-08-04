package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
)

func GetTimelineParams(r *http.Request) (*object.Timeline, error) {
	maxIDParam := r.URL.Query().Get("max_id")
	sinceIDParam := r.URL.Query().Get("since_id")
	limitParam := r.URL.Query().Get("limit")

	// Convert query parameters to int64
	var maxID int64 = object.DefaultMaxID
	if maxIDParam != "" {
		var err error
		maxID, err = strconv.ParseInt(maxIDParam, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid max_id parameter: %w", err)
		}
	}

	var sinceID int64 = object.DefaultSinceID
	if sinceIDParam != "" {
		var err error
		sinceID, err = strconv.ParseInt(sinceIDParam, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid since_id parameter: %w", err)
		}
	}

	var limit int64 = object.DefaultLimit
	if limitParam != "" {
		var err error
		limit, err = strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid limit parameter: %w", err)
		}
	}

	if limit > object.MaxLimit {
		limit = object.MaxLimit
	}

	timeline := &object.Timeline{
		MaxID:   &maxID,
		SinceID: &sinceID,
		Limit:   &limit,
	}

	return timeline, nil
}
