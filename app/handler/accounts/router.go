package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
}

// Create Handler for `/v1/accounts/`
// repository.Accountインターフェースを満たすオブジェクトを受け取り、ハンドラーはアカウント関連の処理を行うメソッドを使用可能
func NewRouter(ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar}
	r.Post("/", h.Create)
	r.Get("/{username}", h.FindAccount)
	return r
}
