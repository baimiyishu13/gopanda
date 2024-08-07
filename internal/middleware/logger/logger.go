package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", // 设置时间戳格式
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// GetLogger 返回初始化的 logger 实例
func GetLogger() *logrus.Logger {
	return log
}
