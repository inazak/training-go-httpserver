package handler

import (
  "net/http"
  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/store"
)

// タスク操作をする store.TaskStore をラップしてハンドラを定義
type ListTask struct {
  Store     *store.TaskStore
}

// レスポンスで使う用のJSON構造
type task struct {
  ID     entity.TaskID     `json:"id"`
  Title  string            `json:"title"`
  Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()

  // ここは単純なGETと呼ばれることを想定しており
  // bodyのチェックが何もなくて、いきなり応答を作成する

  tasks := lt.Store.All()
  rsp := []task{}

  for _, t := range tasks {
    rsp = append(rsp, task{
      ID:     t.ID,
      Title:  t.Title,
      Status: t.Status,
    })
  }

  RespondJSON(ctx, w, rsp, http.StatusOK)
}



