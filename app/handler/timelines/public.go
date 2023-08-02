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
	maxID, err := strconv.ParseInt(maxIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid max_id parameter", http.StatusBadRequest)
		return
	}
	sinceID, err := strconv.ParseInt(sinceIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid since_id parameter", http.StatusBadRequest)
		return
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

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
