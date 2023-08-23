package httpserver

import (
  "context"
  "net/http"
  "github.com/go-chi/chi/v5"
  //"github.com/go-playground/validator/v10"
  //"github.com/inazak/training-go-httpserver/service"
  "github.com/inazak/training-go-httpserver/common/config"
)

// go-chi/chi を利用する理由は、http.ServeMuxの表現力の乏しさ
// 例えば /user/10 のようなパスパラメータの解釈
// GET /users と POST /users といったメソッドの違いのハンドリング
// が難しい

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, error) {
  mux := chi.NewRouter()

  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    // 静的解析のエラー回避のため、明示的に戻り値を捨てる
    _, _ = w.Write([]byte(`{status: "ok"}`))
  })

  // ここでvalidatorインスタンスを生成するメリットはあるのか
  //v := validator.New()

  /*
  at := &handler.AddTask{
    Service: &service.AddTask{ DB: db.SqlxDB, Repo: &r },
    Validator: v,
  }
  mux.Post("/tasks", at.ServeHTTP)
  */

  return mux, nil
}

