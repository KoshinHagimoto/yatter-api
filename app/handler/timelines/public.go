package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
)

func (h *handler) public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Query parameters from the request
	maxIDParam := r.URL.Query().Get("max_id")
	sinceIDParam := r.URL.Query().Get("since_id")
	limitParam := r.URL.Query().Get("limit")

	// Convert query parameters to int64
	var err error
	var maxID int64 = object.DefaultMaxID
	if maxIDParam != "" {
		maxID, err = strconv.ParseInt(maxIDParam, 10, 64)
		if err != nil {
			http.Error(w, "Invalid max_id parameter", http.StatusBadRequest)
			return
		}
	}

	var sinceID int64 = object.DefaultSinceID
	if sinceIDParam != "" {
		sinceID, err = strconv.ParseInt(sinceIDParam, 10, 64)
		if err != nil {
			http.Error(w, "Invalid since_id parameter", http.StatusBadRequest)
			return
		}
	}

	var limit int64 = object.DefaultLimit
	if limitParam != "" {
		limit, err = strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
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

	statuses, err := h.tr.GetPublicTimeline(ctx, timeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, status := range statuses {
		status.Account.FollowerCount, err = h.rr.GetFollowerCount(ctx, status.Account.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		status.Account.FollowingCount, err = h.rr.GetFollowingCount(ctx, status.Account.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
