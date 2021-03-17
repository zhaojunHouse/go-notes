package user

import (
	"context"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (user *User) GetUser(c echo.Context) error {
	log.Info("getUser method")
	return c.JSON(http.StatusOK, "get User ")
}

func (user *User) UpdateUser(ctx context.Context, userID int64) interface{} {
	log.Info("update user Method")
	return nil
}
