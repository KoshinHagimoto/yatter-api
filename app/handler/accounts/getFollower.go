package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/utils"

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

	// クエリパラメータを取得
	timeline, err := utils.GetTimelineParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//フォロワーを取得
	accounts, err := h.rr.GetFollowers(ctx, targetAccount.ID, timeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//フォロー数とフォロワー数を更新
	for _, account := range accounts {
		err = utils.UpdateFollowCounts(ctx, h.rr, account)
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
