package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/utils"

	"github.com/go-chi/chi/v5"
)

// GETでユーザーネームからアカウントを取得する
func (h *handler) FindAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//URLからusernameを取得
	username := chi.URLParam(r, "username")

	//usernameからアカウントを取得
	account, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//フォロー数とフォロワー数を更新
	err = utils.UpdateFollowCounts(ctx, h.rr, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
