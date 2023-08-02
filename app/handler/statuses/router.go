package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	sr repository.Status
}

func NewRouter(ar repository.Account, sr repository.Status) http.Handler {
	r := chi.NewRouter()
	r.Use(auth.Middleware(ar)) //アカウント認証を行うミドルウェアを設定

	h := &handler{sr}
	r.Post("/", h.Create)
	return r
}
