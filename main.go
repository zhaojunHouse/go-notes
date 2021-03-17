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
	// init log module
	err := myLog.InitLog()
	if err != nil {
		log.Errorln(err.Error())
		return
	}

	// init config module
	err = configor.Load(&Conf, "config/conf.json")
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

func main() {
	log.Info(Conf.DB)
}
