package utils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

// WriteLogToFolder 将日志写入指定目录下
// 默认写入logs目录下
func WriteLogToFolder(logger *logrus.Logger, args ...string) {
	if args != nil {
		for i := 0; i < len(args); i++ {
			WriteLogToPath(logger, args[i])
		}
	} else {
		WriteLogToPath(logger, "logs")
	}
}

func WriteLogToPath(logger *logrus.Logger, p string) {
	writer, err := rotatelogs.New(
		path.Join(p, "%Y-%m-%d.log"),
		// 日志保存周期
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 切割日志周期
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.WithError(err).Error("error when writeLogToPath")
		return
	}
	logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.DebugLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}
