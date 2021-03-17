package user

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (user *User) GetUser(ctx context.Context, userID int64) interface{} {
	log.Info("getUser method")
	panic("implement me")
}

func (user *User) UpdateUser(ctx context.Context, userID int64) interface{} {
	log.Info("update user Method")
	panic("implement me")
}
