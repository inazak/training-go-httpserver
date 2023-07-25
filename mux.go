package main

import (
  "net/http"
)

func NewMux() http.Handler {
  mux := http.NewServeMux()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    // 静的解析のエラー回避のため、明示的に戻り値を捨てる
    _, _ = w.Write([]byte(`{status: "ok"}`))
  })
  return mux
}

// そもそも http.ListenAndServe(":80", nil) の場合は
// ルーティングハンドラとして net/http のパッケージ変数 DefaultServeMux が使われる
// DefaultServeMux は ServeMux のポインタ型
// ServeMux は [URLパターン]と[ハンドラ] の map を持っている
// 上記は自前で ServeMux 構造体を作って [URLパターン]と[ハンドラ] を登録しているだけ
// ServeMux.ServeHTTP(w, r) が可能だがその場合のポートはどうなるのか

