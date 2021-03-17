package user

import "context"

type UserHandlerInterface interface {
	GetUser(ctx context.Context, userID int64) interface{}
	UpdateUser(ctx context.Context, userID int64) interface{}
}

type User struct {
}

func NewUserHandler() UserHandlerInterface {
	return &User{}
}
