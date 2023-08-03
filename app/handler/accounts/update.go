package accounts

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"yatter-backend-go/app/handler/auth"
)

type UpdateAccountRequest struct {
	DisplayName *string `json:"display_name"`
	Note        *string `json:"note"`
	Avatar      *string `json:"avatar"`
	Header      *string `json:"header"`
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//アカウント取得
	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//リクエストボディをパース
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//display_nameとnoteを取得
	var req UpdateAccountRequest
	if displayName := r.FormValue("display_name"); displayName != "" {
		req.DisplayName = &displayName
	}
	if note := r.FormValue("note"); note != "" {
		req.Note = &note
	}

	//avatarとheaderを取得
	for _, key := range []string{"avatar", "header"} {
		if file, header, err := r.FormFile(key); err == nil {
			defer file.Close()

			// 保存先ディレクトリの作成
			savePath := "./.data/media/" + key + "/" + header.Filename

			//ファイルを作成
			out, err := os.Create(savePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer out.Close()

			//ファイルを書き込み
			_, err = io.Copy(out, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if key == "avatar" {
				req.Avatar = &savePath
			} else {
				req.Header = &savePath
			}
		}
	}

	if req.DisplayName != nil {
		account.DisplayName = req.DisplayName
	}
	if req.Note != nil {
		account.Note = req.Note
	}
	if req.Avatar != nil {
		account.Avatar = req.Avatar
	}
	if req.Header != nil {
		account.Header = req.Header
	}

	err = h.ar.UpdateAccount(ctx, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
