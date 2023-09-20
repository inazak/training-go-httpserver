package api

import (
	"encoding/json"
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	"github.com/inazak/training-go-httpserver/model"
	"net/http"
)

// FIXME これいまいち
// レスポンスで使う用のJSON構造
type task struct {
	ID     model.TaskID     `json:"id"`
	Title  string           `json:"title"`
	Status model.TaskStatus `json:"status"`
}

func (h *Handler) ServeAddTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Title string `json:"title" validate:"required"`
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

	id, ok := ctx.Value(contextKeyUserID).(model.UserID)
	if !ok {
		_ = jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: "context has no userid"},
			http.StatusInternalServerError)
		return
	}

	// Serviceから新規タスクの登録処理
	t, err := h.backend.AddTask(ctx, id, body.Title)
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
		ID model.TaskID `json:"id"`
	}{
		ID: t.ID,
	}

	_ = jsonhelper.WriteJSONResponse(ctx, w, rsp, http.StatusOK)
}

func (h *Handler) ServeGetTaskList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ここは単純なGETと呼ばれることを想定しており
	// bodyのチェックが何もなくて、いきなり応答を作成する

	tasklist, err := h.backend.GetTaskList(ctx)
	if err != nil {
		_ = jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: err.Error()},
			http.StatusInternalServerError)
		return
	}

	rsp := []task{}
	for _, t := range tasklist {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	_ = jsonhelper.WriteJSONResponse(ctx, w, rsp, http.StatusOK)
}
