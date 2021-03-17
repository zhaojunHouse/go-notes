package user

import (
	"context"
	"github.com/labstack/echo"
)

type UserHandlerInterface interface {
	GetUser(ctx echo.Context) error
	UpdateUser(ctx context.Context, userID int64) interface{}
}

type User struct {
}

func NewUserHandler() UserHandlerInterface {
	return &User{}
}
