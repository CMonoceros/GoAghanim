package lib

import (
	"cmonoceros.com/GoAghanim/pkg"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"path"
	"time"
)

var Log = logrus.New()
var RequestLog = logrus.New()

type LogFile = string

const (
	DEFAULT LogFile = "gin-default-data-"
	API     LogFile = "gin-www-data-"
	ADMIN   LogFile = "gin-admin-data-"
	REQUEST LogFile = "gin-request-data-"
	ERROR   LogFile = "gin-error-data-"
)

func InitLogger(file LogFile) {
	configLocalFileSystemLogger(file)
}

func configLocalFileSystemLogger(logFileName LogFile) {
	writer := getLogWriter(logFileName)
	errorWriter := getLogWriter(ERROR)
	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: errorWriter,
			logrus.FatalLevel: errorWriter,
			logrus.PanicLevel: errorWriter,
		},
		&logrus.TextFormatter{DisableColors: true},
	)
	Log.AddHook(lfHook)

	writer = getLogWriter(REQUEST)
	lfHook = lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.TextFormatter{DisableColors: true},
	)
	RequestLog.AddHook(lfHook)
}

func getLogWriter(logFileName LogFile) *rotatelogs.RotateLogs {
	logPath, err := os.Getwd()
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	baseLogPath := path.Join(logPath, pkg.GetDefaultConfig().LogPath, logFileName)
	maxAge := time.Duration(math.MaxInt64)
	rotationTime := time.Duration(time.Hour * 24)

	writer, err := rotatelogs.New(
		baseLogPath+"%Y-%m-%d.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	return writer
}
