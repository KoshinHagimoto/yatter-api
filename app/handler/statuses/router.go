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

	// 全体に適用するミドルウェアを設定（認証は含まない）

	// 認証が必要なルートグループ
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(ar)) // 認証ミドルウェアを適用
		h := &handler{sr}
		r.Post("/", h.Create) // POST /: 認証が必要
	})

	// 認証が不要なルートグループ
	r.Group(func(r chi.Router) {
		h := &handler{sr}
		r.Get("/{id}", h.FindStatus) // GET /{id}: 認証が不要
	})

	return r
}
