package store

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/taiti09/go_app_handson/clock"
	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/testutil"
	"github.com/taiti09/go_app_handson/testutil/fixture"
)

func TestRepository_ListTasks(t *testing.T) {
	t.Parallel()
	
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx,nil)
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}

	wantUserID, wants := prepareTasks(ctx,t,tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx,tx,wantUserID)
	if err != nil {
		t.Fatalf("unexected error: %v", err)
	}
	if d := cmp.Diff(gots,wants); len(d) != 0 {
		t.Errorf("differs: (-gots +wants)\n%s",d)
	}
}

func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	
	ctx := context.Background()
	c := clock.FixedClocker{}
	var wantID int64 = 20
	okTask := &entity.Task{
		UserID: 33,
		Title: "ok task",
		Status: "todo",
		Created_at: c.Now(),
		Modified_at: c.Now(),
	}
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close()})
	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO task (user_id, title, status, created_at, modified_at) VALUES (?, ?, ?, ?, ?);`)).
		WithArgs(okTask.UserID,okTask.Title,okTask.Status,c.Now(),c.Now()).
		WillReturnResult(sqlmock.NewResult(wantID,1))
	
	xdb := sqlx.NewDb(db,"mysql")
	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx,xdb,okTask); err != nil {
		t.Errorf("want no error, but got %v",err)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) (entity.UserID,entity.Tasks) {
	t.Helper()

	userID := prepareUser(ctx,t,con)
	otherUserID := prepareUser(ctx,t,con)
	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			UserID: userID,
			Title: "want task 1", Status: "todo",
			Created_at: c.Now(), Modified_at: c.Now(),
		},
		{
			UserID: userID,
			Title: "want task 2", Status: "todo",
			Created_at: c.Now(), Modified_at: c.Now(),
		},
	}
	tasks := entity.Tasks{
		wants[0],
		{
			UserID: otherUserID,
			Title:  "not want task", Status: "todo",
			Created_at: c.Now(), Modified_at: c.Now(),
		},
		wants[1],
	}
	result, err := con.ExecContext(ctx,`INSERT INTO task (user_id,title,status,created_at,modified_at)
										VALUES (?,?,?,?,?),
											   (?,?,?,?,?),
											   (?,?,?,?,?);`,
											tasks[0].UserID,tasks[0].Title,tasks[0].Status,tasks[0].Created_at,tasks[0].Modified_at,
											tasks[1].UserID,tasks[1].Title,tasks[1].Status,tasks[1].Created_at,tasks[1].Modified_at,
											tasks[2].UserID,tasks[2].Title,tasks[2].Status,tasks[2].Created_at,tasks[2].Modified_at,
										)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	tasks[0].ID = entity.TaskID(id)
	tasks[1].ID = entity.TaskID(id + 1)
	tasks[2].ID = entity.TaskID(id + 2)
	return userID,wants
}

func prepareUser(ctx context.Context, t *testing.T, db Execer) entity.UserID {
	t.Helper()

	u := fixture.User(nil)
	result, err := db.ExecContext(ctx,`INSERT INTO user (name,password,role,created_at,modified_at) VALUE (?,?,?,?,?);`,
					u.Name,u.Password,u.Role,u.Created_at,u.Modified_at)
	if err != nil {
		t.Fatalf("insert user: %v",err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("got user_id: %v",err)
	}
	return entity.UserID(id)
}