package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	tr repository.Timeline
}

func NewRouter(tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{tr}
	r.Get("/public", h.public)
	return r
}
