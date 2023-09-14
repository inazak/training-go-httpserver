package api

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	"github.com/inazak/training-go-httpserver/common/mock"
	_ "github.com/inazak/training-go-httpserver/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeLogin(t *testing.T) {
	t.Parallel()

	ps := map[string]struct {
		prepareMock func(m *mock.MockService)
		request     string
		response    string
		status      int
	}{
		"ok": {
			prepareMock: func(m *mock.MockService) {
				m.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return(string("1"), nil)
			},
			request:  `{ "name": "abc", "password": "xyz" }`,
			response: `{ "token": "1" }`,
			status:   http.StatusOK,
		},
		"badRequest": {
			prepareMock: func(m *mock.MockService) {
				// AddTaskが呼ばれないテストなので、EXPECT()しない
			},
			request: `{ "xxx": "ng" }`,
			response: `{ "message": "Key: 'Name' Error:Field validation for 'Name' failed on the 'required' tag` + "\\n" +
				`Key: 'Password' Error:Field validation for 'Password' failed on the 'required' tag" }`,
			status: http.StatusBadRequest,
		},
	}

	for n, p := range ps {
		p := p
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/login",
				bytes.NewReader([]byte(p.request)),
			)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockService(ctrl)
			p.prepareMock(m)

			h := NewHandler(m)
			h.ServeLogin(w, r)

			resp := w.Result()
			jsonhelper.AssertResponse(t, resp, p.status, []byte(p.response))
		})
	}
}
