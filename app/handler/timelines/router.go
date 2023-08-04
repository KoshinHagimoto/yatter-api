package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	tr repository.Timeline
	rr repository.Relationship
}

func NewRouter(ar repository.Account, tr repository.Timeline, rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		h := &handler{tr, rr}
		r.Get("/home", h.home)
	})

	r.Group(func(r chi.Router) {
		h := &handler{tr, rr}
		r.Get("/public", h.public)
	})
	return r
}
