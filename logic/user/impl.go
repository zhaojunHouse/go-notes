package user

import (
	"context"
	"go-notes/model"
)

func (u *userLogic) GetUserInfo(ctx context.Context, userID int64) (*model.User, error) {
	return &model.User{
		ID:      1,
		Name:    "john",
		Age:     20,
		Address: "上海市徐汇区",
	} ,nil
}

