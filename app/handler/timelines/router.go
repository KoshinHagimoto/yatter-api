package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	tr repository.Timeline
	rr repository.Relationship
}

func NewRouter(tr repository.Timeline, rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	h := &handler{tr, rr}
	r.Get("/public", h.public)
	return r
}
