package config

import (
  "fmt"
  "testing"
)

func TestNew(t *testing.T) {
  wantPort := 3333
  t.Setenv("PORT", fmt.Sprintf("%d", wantPort))

  got, err := New()

  if err != nil {
    t.Fatalf("cannot create config: %v", err)
  }

  if got.Port != wantPort {
    t.Errorf("want %d, but got %d", wantPort, got.Port)
  }

  // デフォルトで設定される値のチェック
  wantEnv := "dev"
  if got.Env != wantEnv {
    t.Errorf("want %s, but got %s", wantEnv, got.Env)
  }
}

