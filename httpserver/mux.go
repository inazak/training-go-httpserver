package httpserver

import (
  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/inazak/training-go-httpserver/service"
  "github.com/inazak/training-go-httpserver/httpserver/handler/api"
)

// go-chi/chi を利用する理由は、http.ServeMuxの表現力の乏しさ
// 例えば /user/10 のようなパスパラメータの解釈
// GET /users と POST /users といったメソッドの違いのハンドリング
// が難しい

func NewMux(svc service.Service) http.Handler {
  mux := chi.NewRouter()

  apiHandler := api.NewHandler(svc)
  mux.HandleFunc("/health", apiHandler.ServeHealthCheck)
  mux.Get("/task", apiHandler.ServeGetTaskList)
  mux.Post("/task", apiHandler.ServeAddTask)

  return mux
}

