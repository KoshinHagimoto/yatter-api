package relationships

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	ar repository.Account
	rr repository.Relationship
}

func NewRouter(ar repository.Account, rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		h := &handler{ar, rr}
		r.Get("/relationships", h.GetRelationships)
		r.Post("/{username}/follow", h.FollowAccount)
		r.Post("/{username}/unfollow", h.UnfollowAccount)
	})

	r.Group(func(r chi.Router) {
		h := &handler{ar, rr}
		r.Get("/{username}/following", h.GetFollowing)
		r.Get("/{username}/followers", h.GetFollowers)
	})

	return r
}
