package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"go-notes/config"
	"go-notes/handler/order"
	"go-notes/handler/user"
	myLog "go-notes/log"
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
	userHandler := user.NewUserHandler()
	orderHandler := order.NewOrderHandler()

	engine := router.Router(userHandler, orderHandler)
	err := engine.Start(":8080")
	if err != nil {
		log.Error("HTTP SERVER start failed:", err.Error())
	}
}
