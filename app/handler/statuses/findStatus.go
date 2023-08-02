package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FindStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//URLからstatusIDを取得
	statusIDstr := chi.URLParam(r, "id")

	statusID, err := strconv.ParseInt(statusIDstr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//statusIDからステータスを取得
	status, err := h.sr.FindStatusByID(ctx, statusID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if status == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
