package main

import (
  "context"
  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/clock"
  "github.com/inazak/training-go-httpserver/config"
  "github.com/inazak/training-go-httpserver/handler"
  "github.com/inazak/training-go-httpserver/store"
  "github.com/inazak/training-go-httpserver/service"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, error) {
  mux := chi.NewRouter()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    // 静的解析のエラー回避のため、明示的に戻り値を捨てる
    _, _ = w.Write([]byte(`{status: "ok"}`))
  })

  // ここでvalidatorインスタンスを生成するメリットはあるのか
  v := validator.New()

  // MuxでDBをオープンする必要があるだろうか
  // これだとテストを作れなくないか
  db, err := store.New(ctx, cfg)
  defer db.Close() //dbはnilにならないようにしている
  if err != nil {
    return nil, err
  }

  // ここも store.Repository をここで生成するメリットが分からない
  r := store.Repository{ Clocker: clock.RealClocker{} }

  // serviceの使い方が美しくない
  at := &handler.AddTask{
    Service: &service.AddTask{ DB: db.SqlxDB, Repo: &r },
    Validator: v,
  }
  mux.Post("/tasks", at.ServeHTTP)

  lt := &handler.ListTask{
    Service: &service.ListTask{ DB: db.SqlxDB, Repo: &r },
  }
  mux.Get("/tasks", lt.ServeHTTP)

  return mux, nil
}

// go-chi/chi を利用する理由は、http.ServeMuxの表現力の乏しさ
// 例えば /user/10 のようなパスパラメータの解釈
// GET /users と POST /users といったメソッドの違いのハンドリング
// が難しい

