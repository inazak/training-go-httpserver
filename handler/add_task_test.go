package handler

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/go-playground/validator/v10"
  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/testutil"
  "github.com/inazak/training-go-httpserver/mock"
  "github.com/golang/mock/gomock"
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


// gomock を使ったテストの構築
// interface に go generate で mockパッケージの生成
// mockパッケージをインポートして
//      ctrl := gomock.NewController(t)
//      defer ctrl.Finish()
//      m := mock.NewMockXXXX(ctrl)
//      m.EXPECT().Func(1, 2, gomock.Any()).Return("ok")
// のようにするが、複数ケースをループで起動するテストの場合は
// テストケース自体にモックのセットアップ関数を組み込んで
//  tests := map[string]struct {
//    ...
//    prepareMock func(m *mock.MockXXXX)
// 関数内でEXPECTを実行するようにテストケースを定義し、
//    "case1": {
//      prepareMock: func(m *mock.MockXXXX){
//        m.EXPECT().Func(1,2,gomock.Any()).Return("ok")
//      },
// テストループでそれを呼ぶ際に、生成したモックにprepare経由でEXPECTを適用する
//      m := mock.NewXXXX(ctrl)
//      tt.prepareMock(m)
// ただ、EXPECTしたのに呼ばれないとmissing call(s)で失敗するので
// 呼ぶことが確実なモックだけセットアップする
//

func TestAddTask(t *testing.T) {
  t.Parallel()

  type want struct {
    status int
    rspFile string
  }

  tests := map[string]struct {
    prepareMock func(m *mock.MockAddTaskService)
    reqFile string
    want    want
  }{
    "ok": {
      prepareMock: func(m *mock.MockAddTaskService){
        m.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(&entity.Task{ID: 1}, nil)
      },
      reqFile: "testdata/add_task/ok_req.json.golden",
      want: want {
        status: http.StatusOK,
        rspFile: "testdata/add_task/ok_rsp.json.golden",
      },
    },
    "badRequest": {
      prepareMock: func(m *mock.MockAddTaskService){
        // AddTaskが呼ばれないテストなので、EXPECT()しない
      },
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

      ctrl := gomock.NewController(t)
      defer ctrl.Finish()

      m := mock.NewMockAddTaskService(ctrl)
      tt.prepareMock(m)

      sut := AddTask{
        Service: m,
        Validator: validator.New(),
      }
      sut.ServeHTTP(w, r)

      resp := w.Result()
      testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
    })
  }
}

