package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

func (h *handler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := chi.URLParam(r, "username")

	targetAccount, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Query parameters from the request
	maxIDParam := r.URL.Query().Get("max_id")
	sinceIDParam := r.URL.Query().Get("since_id")
	limitParam := r.URL.Query().Get("limit")

	// Convert query parameters to int64
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

	accounts, err := h.rr.GetFollowers(ctx, targetAccount.ID, timeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, account := range accounts {
		account.FollowerCount, err = h.rr.GetFollowerCount(ctx, account.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		account.FollowingCount, err = h.rr.GetFollowingCount(ctx, account.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
