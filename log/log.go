package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func InitLog() error{
	var stdFormatter *prefixed.TextFormatter  // 命令行输出格式
	var fileFormatter *prefixed.TextFormatter // 文件输出格式

	stdFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     true,
		DisableColors:   false,
	}

	fileFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}
	log.SetFormatter(stdFormatter)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	logName := fmt.Sprintf("%s/access_log.", "log")
	writer, err := rotatelogs.New(logName + "%Y%m%d")
	if err != nil {
		return err
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel:  writer,
		log.DebugLevel: writer,
		log.ErrorLevel: writer,
	}, fileFormatter)
	log.SetOutput(os.Stdout)
	log.AddHook(lfHook)
	return nil
}
