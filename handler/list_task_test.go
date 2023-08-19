package handler

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/inazak/training-go-httpserver/entity"
  "github.com/inazak/training-go-httpserver/testutil"
  "github.com/inazak/training-go-httpserver/mock"
  "github.com/golang/mock/gomock"
)

func TestListTask(t *testing.T) {

  type want struct {
    status int
    rspFile string
  }

  // これ名前をtestsにしない方がよくないか
  tests := map[string]struct {
    prepareMock func(m *mock.MockListTasksService)
    tasks map[entity.TaskID]*entity.Task
    want  want
  }{
    "ok": {
      prepareMock: func(m *mock.MockListTasksService){
        // 結果がこの順序通りかどうかは、保証はできないはず
        m.EXPECT().ListTasks(gomock.Any()).Return(entity.Tasks{
          { ID: 1, Title: "test1", Status: entity.TaskStatusTodo, },
          { ID: 2, Title: "test2", Status: entity.TaskStatusDone, },
        }, nil)
      },
      tasks: map[entity.TaskID]*entity.Task{
        1: {
          ID:     1,
          Title:  "test1",
          Status: entity.TaskStatusTodo,
        },
        2: {
          ID:     2,
          Title:  "test2",
          Status: entity.TaskStatusDone,
        },
      },
      want: want {
        status: http.StatusOK,
        rspFile: "testdata/list_task/ok_rsp.json.golden",
      },
    },
    "empty": {
      prepareMock: func(m *mock.MockListTasksService){
        m.EXPECT().ListTasks(gomock.Any()).Return(entity.Tasks{}, nil)
      },
      tasks: map[entity.TaskID]*entity.Task{},
      want: want {
        status: http.StatusOK,
        rspFile: "testdata/list_task/empty_rsp.json.golden",
      },
    },
  }

  for n, tt := range tests {
    tt := tt
    t.Run(n, func(t *testing.T) {
      t.Parallel()

      w := httptest.NewRecorder()
      r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

      ctrl := gomock.NewController(t)
      defer ctrl.Finish()

      m := mock.NewMockListTasksService(ctrl)
      tt.prepareMock(m)

      sut := ListTask{ Service: m }
      sut.ServeHTTP(w, r)

      resp := w.Result()
      testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
    })
  }
}

