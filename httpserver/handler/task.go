package handler

import (
  "encoding/json"
  "net/http"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/model"
  "github.com/inazak/training-go-httpserver/service"
  "github.com/inazak/training-go-httpserver/common/jsonhelper"
)

type TaskHandler struct {
  serv service.Service
  vali *validator.Validate
}

func (th *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()

  var body struct {
    Title string `json:"title" validate:"required"`
  }

  // POSTされたJSONのパース
  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    jsonhelper.WriteJSONResponse(ctx, w, &jsonhelper.ErrorResponse{
      Message: err.Error(),
    }, http.StatusInternalServerError)
    return
  }

  // JSONが要件を満たすかチェック
  if err := th.vali.Struct(body) ; err != nil {
    _ = jsonhelper.WriteJSONResponse(ctx, w, &jsonhelper.ErrorResponse{
      Message: err.Error(),
    }, http.StatusBadRequest)
    return
  }

  // Serviceから新規タスクの登録処理
  t, err := th.serv.AddTask(ctx, body.Title)
  if err != nil {
    _ = jsonhelper.WriteJSONResponse(ctx, w, &jsonhelper.ErrorResponse{
      Message: err.Error(),
    }, http.StatusInternalServerError)
    return
  }

  // 正常な場合はIDを返す
  rsp := struct{
    ID model.TaskID `json:"id"`
  }{ ID: t.ID }
  _ = jsonhelper.WriteJSONResponse(ctx, w ,rsp, http.StatusOK)
}

