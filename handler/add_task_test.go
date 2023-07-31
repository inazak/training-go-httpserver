package handler

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/store"
  "github.com/inazak/training-go-httpserver/testutil"
)

// t.Parallel をトップレベルで書くと
// t.Parallel をトップレベルで呼び出していない他のテストが実行され、
// その後、t.Parallel を呼び出しているテストが並列に実行される、
// t.Run のサブテスト関数で t.Parallel を呼び出している場合
// そのトップレベルのテスト関数が終了して戻ったあと、(*1)
// サブテスト関数が並列して実行される
// 並列を最大にするには、トップレベルのテスト関数と、
// サブテスト関数の両方で t.Parallel を呼び出す必要がある
// t.Runによるサブテスト関数内で、t.Parallel を呼び出している場合
// defer ではなく、t.Cleanup で後処理を書く
// defer は (*1)のタイミングで実行されてしまうため
// Cleanup はテスト、サブテストの全てが完了したときに呼び出される

func TestAddTask(t *testing.T) {
  t.Parallel()

  type want struct {
    status int
    rspFile string
  }

  // これ名前をtestsにしない方がよくないか
  tests := map[string]struct {
    reqFile string
    want    want
  }{
    "ok": {
      reqFile: "testdata/add_task/ok_req.json.golden",
      want: want {
        status: http.StatusOK,
        rspFile: "testdata/add_task/ok_rsp.json.golden",
      },
    },
    "badRequest": {
      reqFile: "testdata/add_task/bad_req.json.golden",
      want: want {
        status: http.StatusBadRequest,
        rspFile: "testdata/add_task/bad_rsp.json.golden",
      },
    },
  }

  for n, tt := range tests {
    tt := tt
    t.Run(n, func(t *testing.T) {
      t.Parallel()

      w := httptest.NewRecorder()
      r := httptest.NewRequest(
        http.MethodPost,
        "/tasks",
        bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
      )

      sut := AddTask{
        Store: &store.TaskStore{
          Tasks: map[entity.TaskID]*entity.Task{},
        },
        Validator: validator.New(),
      }
      sut.ServeHTTP(w, r)

      resp := w.Result()
      testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
    })
  }
}

