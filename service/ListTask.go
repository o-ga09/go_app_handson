package service

import (
	"context"
	"fmt"

	"github.com/taiti09/go_app_handson/auth"
	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/store"
)

type ListTask struct {
	DB store.Queryer
	Repo TaskListner
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil,fmt.Errorf("user_id not found")
	}
	ts, err := l.Repo.ListTasks(ctx,l.DB,id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
} 