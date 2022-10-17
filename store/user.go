package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/taiti09/go_app_handson/entity"
)

func (r *Repository) RegisterUser(ctx context.Context, db Execer, u *entity.User) error {
	u.Created_at = r.Clocker.Now()
	u.Modified_at = r.Clocker.Now()
	sql := `INSERT INTO user (name,password,role,created_at,modified_at) VALUES (?,?,?,?,?)`
	result, err := db.ExecContext(ctx,sql,u.Name,u.Password,u.Role,u.Created_at,u.Modified_at)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err,&mysqlErr) && mysqlErr.Number == ErrCodeMYSQLDuplicateEntry {
			return fmt.Errorf("cannnot create sama name user: %w", ErrAlreadyEntry)
		}
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = entity.UserID(id)
	return nil
}