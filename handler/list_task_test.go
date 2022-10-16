package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/testutil"
)

func TestListTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspfile string
	}
	tests := map[string]struct {
		tasks []*entity.Task
		want want
	}{
		"ok": {
			tasks: []*entity.Task{
				{
					ID: 1,
					Title: "test1",
					Status: entity.TaskStatusTodo,
				},
				{
					ID: 2,
					Title: "test2",
					Status: entity.TaskStatusDone,
				},
			},
			want: want{
				status: http.StatusOK,
				rspfile: "testdata/listtask/ok_rsp.json.golden",
			},
		},
		"empty": {
			tasks: []*entity.Task{},
			want: want{
				status: http.StatusOK,
				rspfile: "testdata/listtask/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n,func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodGet,
				"/tasks",
				nil,
			)
			moq := &ListTasksServiceMock{}
			moq.ListTasksFunc = func(ctx context.Context) (entity.Tasks, error) {
				if tt.tasks != nil {
					return tt.tasks, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := ListTask{
				Service: moq,
			}
			sut.ServeHTTP(w,r)

			resp := w.Result()
			testutil.AssertResponse(t,resp,tt.want.status,testutil.LoadFile(t,tt.want.rspfile))
		})
	}
}