package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/taiti09/go_app_handson/clock"
	"github.com/taiti09/go_app_handson/entity"
)

func TestRepository_ListTasks_withMock(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	c := clock.FixedClocker{}
	var wantID int64 = 20

	okTask := entity.Task{
		Title: "ok Task",
		Status: "todo",
		Created_at: c.Now(),
		Modified_at: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(
		`INSERT INTO task \(title,status,created_at,modified_at\) VALUES \(\?,\?,\?,\?\)`,
	).WithArgs(okTask.Title,okTask.Status,okTask,c.Now(),c.Now()).WillReturnResult(sqlmock.NewResult(wantID,1))
	if err != nil {
		t.Fatal(err)
	}

	xdb := sqlx.NewDb(db,"mysql")
	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx,xdb,&okTask); err != nil {
		t.Errorf("want no error, but got %v",err)
	}
}