package user

import (
	"github.com/labstack/echo"
	"go-notes/logic/user"
)

type UserHandlerInterface interface {
	GetUser(ctx echo.Context) error
}

type User struct {
	userLogic user.UserLogicInterface
}

func NewUserHandler(userLogic user.UserLogicInterface) UserHandlerInterface {
	return &User{
		userLogic: userLogic,
	}
}
