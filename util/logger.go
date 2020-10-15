package util

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var logger *logrus.Logger

// BuildLogger 构建logger
func BuildLogger(level string) {
	l := Log()
	c, _ := logrus.ParseLevel(level)
	l.SetLevel(c)
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetReportCaller(true)

	path := "storage/go.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	logger.SetOutput(writer)
}

// Log 返回日志对象
func Log() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}
