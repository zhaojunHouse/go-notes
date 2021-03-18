package user

import (
	"context"
	"go-notes/model"
)

type UserLogicInterface interface {
	GetUserInfo(ctx context.Context,userID int64) (*model.User, error)
}

type userLogic struct {

}

func NewUserLogicInterface() UserLogicInterface {
	return &userLogic{}
}