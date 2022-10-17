package service

import (
	"context"

	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskListner
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskListner interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}