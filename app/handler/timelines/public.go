package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/utils"
)

func (h *handler) public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Query parameters from the request
	timeline, err := utils.GetTimelineParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statuses, err := h.tr.GetPublicTimeline(ctx, timeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, status := range statuses {
		//フォロー数とフォロワー数を更新
		err = utils.UpdateFollowCounts(ctx, h.rr, status.Account)
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
