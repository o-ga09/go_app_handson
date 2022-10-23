package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/taiti09/go_app_handson/entity"
)

func User(u *entity.User) *entity.User {
	result := &entity.User{
		ID: entity.UserID(rand.Int()),
		Name: "test" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role: "admin",
		Created_at: time.Now(),
		Modified_at: time.Now(),
	}
	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Role != "" {
		result.Role = u.Role
	}
	if !u.Created_at.IsZero() {
		result.Created_at = u.Created_at
	}
	if !u.Modified_at.IsZero() {
		result.Modified_at = u.Modified_at
	}
	return result
}