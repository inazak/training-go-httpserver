package api

import (
	"encoding/json"
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	_ "github.com/inazak/training-go-httpserver/model"
	"net/http"
)

func (h *Handler) ServeLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Name     string `json:"name"     validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// POSTされたJSONのパース
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: err.Error()},
			http.StatusInternalServerError)
		return
	}

	// JSONが要件を満たすかチェック
	if err := validation.Struct(body); err != nil {
		_ = jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: err.Error()},
			http.StatusBadRequest)
		return
	}

	// Serviceから新規ユーザーの登録処理
	token, err := h.backend.Login(ctx, body.Name, body.Password)
	if err != nil {
		_ = jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: err.Error()},
			http.StatusInternalServerError)
		return
	}

	// 正常な場合はIDを返す
	rsp := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	_ = jsonhelper.WriteJSONResponse(ctx, w, rsp, http.StatusOK)
}
