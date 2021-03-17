package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"go-notes/config"
	myLog "go-notes/log"
)

/**
    config
	log
	gateway
	http server
	rpc server
	mysql
	reload
*/

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

func main() {
	log.Info(Conf.DB)
	log.Errorf("test error,%+v", Conf.DB)
}
