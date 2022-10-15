package store

import (
	"context"

	"github.com/taiti09/go_app_handson/entity"
)

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT id, title, status, created_at, modified_at FROM task;`
	if err := db.SelectContext(ctx,&tasks, sql); err != nil{
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, db Execer, t *entity.Task) error {
	t.Created_at = r.Clocker.Now()
	t.Modified_at = r.Clocker.Now()
	sql := `INSERT INTO task (title,status,created_at,modified_at) VALUES(?,?,?,?)`
	result, err := db.ExecContext(ctx,sql,t.Title,t.Status,t.Created_at,t.Modified_at)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	t.ID = entity.TaskID(id)
	return nil
}