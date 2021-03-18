package user

import (
	"context"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (user *User) GetUser(c echo.Context) error {
	log.Info("getUser handler")
	userInfo, err := user.userLogic.GetUserInfo(context.TODO(),1)
	if err != nil {
		log.Error("GetUser-GetUserInfo-err :",err.Error())
		return err
	}
	return c.JSON(http.StatusOK, userInfo)
}
