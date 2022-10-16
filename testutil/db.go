package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	if _, define := os.LookupEnv("CI"); define {
		port = 3306
	}
	
	db, err := sql.Open("mysql",fmt.Sprintf("todo:P@ssw0rd@tcp(127.0.0.1:%d)/todo?parseTime=true",port))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close()})
	return sqlx.NewDb(db,"mysql")
}