package relationships

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

func (h *handler) UnfollowAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := chi.URLParam(r, "username")

	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	targetAccount, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.rr.DeleteRelationship(ctx, account.ID, targetAccount.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isFollowing, err := h.rr.IsFollowing(ctx, account.ID, targetAccount.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	isFollowedBy, err := h.rr.IsFollowing(ctx, targetAccount.ID, account.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &FollowResponse{
		ID:         targetAccount.ID,
		Following:  isFollowing,
		FollowedBy: isFollowedBy,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
