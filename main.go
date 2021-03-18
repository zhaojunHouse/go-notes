package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"go-notes/config"
	"go-notes/handler/order"
	"go-notes/handler/user"
	myLog "go-notes/log"
	logicUser "go-notes/logic/user"
	"go-notes/router"
)

var Conf config.Config

func init() {
	err := myLog.InitLog()
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	err = configor.Load(&Conf, "config/conf.json")
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

/**
http server
rpc server
mysql
reload
gateway
tracing
test
*/

func main() {
	// logic 层
	userLogic := logicUser.NewUserLogicInterface()

	// handler 层
	userHandler := user.NewUserHandler(userLogic)
	orderHandler := order.NewOrderHandler()

	// router 层
	engine := router.Router(userHandler, orderHandler)
	err := engine.Start(":8080")
	if err != nil {
		log.Error("HTTP SERVER start failed:", err.Error())
	}
}
