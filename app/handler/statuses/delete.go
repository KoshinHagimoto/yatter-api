package statuses

import (
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//URLからstatusIDを取得
	statusIDstr := chi.URLParam(r, "id")

	statusID, err := strconv.ParseInt(statusIDstr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.sr.DeleteStatus(ctx, statusID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
