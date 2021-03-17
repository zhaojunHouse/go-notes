package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func InitLog() error {
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

	accessLogName := fmt.Sprintf("%s/access_log.", "log")
	accessWriter, err := rotatelogs.New(accessLogName + "%Y%m%d")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	errorLogName := fmt.Sprintf("%s/error_log.", "log")
	errorWriter, err := rotatelogs.New(errorLogName + "%Y%m%d")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel:  accessWriter,
		log.DebugLevel: accessWriter,
		log.ErrorLevel: errorWriter,
	}, fileFormatter)

	log.SetOutput(os.Stdout)
	log.AddHook(lfHook)
	return nil
}
