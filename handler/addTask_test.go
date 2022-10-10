package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/store"
	"github.com/taiti09/go_app_handson/testutil"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status int
		rspfile string
	}
	tests := map[string]struct {
		reqFile string
		want want
	}{
		"ok": {
			reqFile: "testdata/addtask/ok_req.json.golden",
			want: want{
				status: http.StatusOK,
				rspfile: "testdata/addtask/ok_rsp.json.golden",
			},
		},
		"badrequest": {
			reqFile: "testdata/addtask/bad_req.json.golden",
			want: want{
				status: http.StatusOK,
				rspfile: "testdata/addtask/bad_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n,func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t,tt.reqFile)),
			)

			sut := AddTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{},
				},
				Validator: validator.New(),
			}
			sut.ServerHTTP(w,r)

			resp := w.Result()
			testutil.AssertResponse(t,resp,tt.want.status,testutil.LoadFile(t,tt.want.rspfile))
		})
	}
}