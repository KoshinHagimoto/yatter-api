package accounts

import (
	"encoding/json"
	"net/http"
	"strings"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/utils"
)

func (h *handler) GetRelationships(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	usernames := r.URL.Query().Get("username")
	splitUsernames := strings.Split(usernames, ",")

	var relationships []FollowResponse

	for _, username := range splitUsernames {
		targetAccount, err := h.ar.FindByUsername(ctx, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		isFollowing, isFollowedBy, err := utils.FetchFollowRelationship(ctx, h.rr, account.ID, targetAccount.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		relationships = append(relationships, FollowResponse{
			ID:         targetAccount.ID,
			Following:  isFollowing,
			FollowedBy: isFollowedBy,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relationships); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
