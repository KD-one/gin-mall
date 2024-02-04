package util

import (
	"gin-mall/conf"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	logFilePath := SetLogPath()
	if LogrusObj != nil {
		LogrusObj.Out = logFilePath
		return
	}
	logger := logrus.New()
	logger.Out = logFilePath
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
}

func SetLogPath() *os.File {
	var logPath string
	fileName := time.Now().Format("2006-01-02") + ".log"
	// TODO: conf文件中的Mode字段，只有debug模式下才是开发环境可以在根目录下创建logs文件夹，除了debug模式其余都是在工作目录下创建logs文件夹
	if conf.Mode != "debug" {
		root, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		logPath = root + "/logs/"
	} else {
		logPath = "../logs/"
	}

	err := os.MkdirAll(logPath, 0777)
	if err != nil {
		panic(err)
	}

	file, _ := os.OpenFile(logPath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	return file
}
