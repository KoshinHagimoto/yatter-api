package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/utils"

	"github.com/go-chi/chi/v5"
)

type FollowResponse struct {
	ID         int64 `json:"id"`
	Following  bool  `json:"following"`
	FollowedBy bool  `json:"followed_by"`
}

func (h *handler) FollowAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	//認証されたユーザーアカウントを取得
	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//フォローするアカウントを取得
	targetAccount, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//自分自身をフォローしようとしていないことを確認
	if account.ID == targetAccount.ID {
		http.Error(w, "You can't follow yourself", http.StatusForbidden)
		return
	}

	relationship := &object.Relationship{
		FollowerID:  account.ID,
		FollowingID: targetAccount.ID,
	}

	err = h.rr.SaveRelationship(ctx, relationship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isFollowing, isFollowedBy, err := utils.FetchFollowRelationship(ctx, h.rr, account.ID, targetAccount.ID)
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
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
