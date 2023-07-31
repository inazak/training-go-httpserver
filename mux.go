package main

import (
  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/handler"
  "github.com/inazak/training-go-httpserver/store"
)

func NewMux() http.Handler {
  mux := chi.NewRouter()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    // 静的解析のエラー回避のため、明示的に戻り値を捨てる
    _, _ = w.Write([]byte(`{status: "ok"}`))
  })

  v := validator.New()
  at := &handler.AddTask{ Store: store.Tasks, Validator: v }
  mux.Post("/tasks", at.ServeHTTP)

  lt := &handler.ListTask{ Store: store.Tasks }
  mux.Get("/tasks", lt.ServeHTTP)

  return mux
}

// go-chi/chi を利用する理由は、http.ServeMuxの表現力の乏しさ
// 例えば /user/10 のようなパスパラメータの解釈
// GET /users と POST /users といったメソッドの違いのハンドリング
// が難しい

