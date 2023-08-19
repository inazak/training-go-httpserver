package handler

import (
  "encoding/json"
  "net/http"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/entity"
)

// Service を使うように変更
type AddTask struct {
  Service   AddTaskService
  Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()

  // ここでバリデーション内容を決めているが、このなぜメソッドの中で
  // 定義しているのかは不明、またvalidatorはAddTask構造に含まれているが
  // 構造毎に生成する必要もないのでは
  var b struct {
    // ここでタイプミスがあり vali'c'ate としていたが、コンパイルは通ってしまう
    // そのため、ミスに気付くのが遅くなった、このあたりをなんとかできないのか
    Title string `json:"title" validate:"required"`
  }

  // POSTされたJSONのパース
  if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
    RespondJSON(ctx, w, &ErrResponse{
      Message: err.Error(),
    }, http.StatusInternalServerError)
    return
  }

  // JSONが要件を満たすかチェック
  if err := at.Validator.Struct(b) ; err != nil {
    RespondJSON(ctx, w, &ErrResponse{
      Message: err.Error(),
    }, http.StatusBadRequest)
    return
  }

  // Serviceから新規タスクの登録処理
  t, err := at.Service.AddTask(ctx, b.Title)
  if err != nil {
    RespondJSON(ctx, w, &ErrResponse{
      Message: err.Error(),
    }, http.StatusInternalServerError)
    return
  }

  // 正常な場合はIDを返す
  rsp := struct{
    ID entity.TaskID `json:"id"`
  }{ ID: t.ID }
  RespondJSON(ctx, w ,rsp, http.StatusOK)
}

