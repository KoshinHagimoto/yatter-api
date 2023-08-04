package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/utils"
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// クエリパラメータを取得
	timeline, err := utils.GetTimelineParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statuses, err := h.tr.GetHomeTimeline(ctx, account.ID, timeline)
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
