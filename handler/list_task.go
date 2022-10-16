package handler

import (
	"net/http"

	"github.com/taiti09/go_app_handson/entity"
)

type ListTask struct {
	Service ListTasksService
}

type task struct {
	ID entity.TaskID
	Title string
	Status entity.TaskStatus
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Service.ListTasks(ctx)
	if err != nil {
		RespondJSON(ctx,w,&ErrResponse{
			Message: err.Error(),
		},http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp,task {
			ID: t.ID,
			Title: t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx,w,rsp,http.StatusOK)
}