package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"go-notes/config"
	myLog "go-notes/log"
	"go-notes/router"
	"net/http"
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
	err := http.ListenAndServe(":8080", router.Router())
	if err != nil {
		log.Error("HTTP SERVER start failed:", err.Error())
	}
}
