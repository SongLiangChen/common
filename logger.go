package common

import (
	"path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var DefaultLoggerFormat = &log.JSONFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
	FieldMap: log.FieldMap{
		log.FieldKeyMsg: "message",
	},
}

var (
	_appName = ""
)

//配置业务日志系统
func ConfigLogger(curMode string, appName, dir string, logName string, keepDays int, rotateHours int) {
	if curMode != "online" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	_appName = appName
	log.SetFormatter(DefaultLoggerFormat)
	configLocalFs(dir, logName, time.Hour*time.Duration(keepDays*24), time.Hour*time.Duration(rotateHours))
}

//配置本地文件系统并按周期分割
func configLocalFs(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, _ := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, DefaultLoggerFormat)
	log.AddHook(lfHook)
}

// ---------------------------------------------------------------------------------------------------------------------
func ErrorLog(action string, mates interface{}, msg string) {
	log.WithFields(log.Fields{
		"app":    _appName,
		"action": action,
		"mates":  mates,
	}).Error(msg)
}

func InfoLog(action string, mates interface{}, msg string) {
	log.WithFields(log.Fields{
		"app":    _appName,
		"action": action,
		"mates":  mates,
	}).Info(msg)
}

func DebugLog(action string, mates interface{}, msg string) {
	log.WithFields(log.Fields{
		"app":    _appName,
		"action": action,
		"mates":  mates,
	}).Debug(msg)
}

func WarningLog(action string, mates interface{}, msg string) {
	log.WithFields(log.Fields{
		"app":    _appName,
		"action": action,
		"mates":  mates,
	}).Warn(msg)
}
