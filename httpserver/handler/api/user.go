package api

import (
	"encoding/json"
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	"github.com/inazak/training-go-httpserver/model"
	"net/http"
)

func (h *Handler) ServeAddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role"     validate:"required"`
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
	id, err := h.backend.AddUser(ctx, body.Username, body.Password, body.Role)
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
		ID model.UserID `json:"id"`
	}{
		ID: id,
	}

	_ = jsonhelper.WriteJSONResponse(ctx, w, rsp, http.StatusOK)
}
