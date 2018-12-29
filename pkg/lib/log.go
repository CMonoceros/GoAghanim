package lib

import (
	"os"
	"path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// LogFile 日志文件名称
type LogFile = string

// LogFileConfig 日志文件配置
type LogFileConfig struct {
	BaseLevel  LogFile
	DebugLevel LogFile
	InfoLevel  LogFile
	WarnLevel  LogFile
	ErrorLevel LogFile
	FatalLevel LogFile
	PanicLevel LogFile
}

// Logger 默认日志记录器
var Logger = logrus.New()

// InitLogger 初始化日志记录器
func InitLogger(config LogFileConfig) {
	configLocalFileSystemLogger(config)
}

func configLocalFileSystemLogger(config LogFileConfig) {
	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: getLogWriter(config.DebugLevel, config.BaseLevel),
			logrus.InfoLevel:  getLogWriter(config.InfoLevel, config.BaseLevel),
			logrus.WarnLevel:  getLogWriter(config.WarnLevel, config.BaseLevel),
			logrus.ErrorLevel: getLogWriter(config.ErrorLevel, config.BaseLevel),
			logrus.FatalLevel: getLogWriter(config.FatalLevel, config.BaseLevel),
			logrus.PanicLevel: getLogWriter(config.PanicLevel, config.BaseLevel),
		},
		&logrus.TextFormatter{DisableColors: true},
	)
	Logger.AddHook(lfHook)
	Logger.SetReportCaller(true)
}

func getLogWriter(logFileName LogFile, defaultFileName LogFile) *rotatelogs.RotateLogs {
	fileName := logFileName
	if fileName == "" {
		fileName = defaultFileName
	}

	logPath, err := os.Getwd()
	CheckAndThrowError(err, LibraryCode)
	baseLogPath := path.Join(logPath, "logs")
	_, err = os.Stat(baseLogPath)
	if err != nil && os.IsNotExist(err) {
		if os.Mkdir(baseLogPath, os.ModePerm) != nil {
			CheckAndThrowError(err, LibraryCode)
		}
	} else if err != nil {
		CheckAndThrowError(err, LibraryCode)
	}

	baseLogPath = path.Join(baseLogPath, fileName)
	maxAge := time.Duration(time.Hour * 48)
	rotationTime := time.Duration(time.Hour * 24)

	writer, err := rotatelogs.New(
		baseLogPath+"%Y-%m-%d.log",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	CheckAndThrowError(err, LibraryCode)
	return writer
}
