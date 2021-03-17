package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"go-notes/config"
	log "github.com/sirupsen/logrus"
)

/**
    config
	log
	http server
	rpc server
	mysql
	reload
 */

var Conf config.Config

// init config and log module
func init(){
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		TimestampFormat:"2006-01-02 15:04:05",
	})
	log.SetReportCaller(true)

	err := configor.Load(&Conf, "config/conf.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	log.Infof("%+v",Conf.Mongo)
}