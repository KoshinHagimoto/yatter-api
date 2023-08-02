package handler

import (
	"net/http"
	"time"

	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/accounts"
	"yatter-backend-go/app/handler/health"
	"yatter-backend-go/app/handler/statuses"
	"yatter-backend-go/app/handler/timelines"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// db.NewAccountで返された、arが引数に入っている
func NewRouter(ar repository.Account, sr repository.Status, tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) //各リクエストにユニークなIDを付与する
	r.Use(middleware.RealIP)    //リクエスト元のIPアドレスを取得する
	r.Use(middleware.Logger)    //リクエストの開始と終了をログに出力する
	r.Use(middleware.Recoverer) //パニックをキャッチして500エラーを返す
	r.Use(newCORS().Handler)    //CORSを許可する

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second)) //リクエストのタイムアウトを設定する

	r.Mount("/v1/accounts", accounts.NewRouter(ar))
	r.Mount("/v1/statuses", statuses.NewRouter(ar, sr))
	r.Mount("/v1/timelines", timelines.NewRouter(tr))
	r.Mount("/v1/health", health.NewRouter())

	return r
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
