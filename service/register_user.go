package service

import (
	"context"

	"github.com/taiti09/go_app_handson/entity"
	"github.com/taiti09/go_app_handson/store"
)

type RegisterUser struct {
	DB store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, name string, password string, role string) (*entity.User, error) {
	
}