package utils

import (
	"bytes"
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

const logPath = "logs/"

var Log *logrus.Logger

type mineFormatter struct{}

func init() {
	Log = create()
}
func create() *logrus.Logger {
	if Log != nil {
		return Log
	}
	Log = logrus.New()
	Log.Formatter = new(logrus.JSONFormatter)
	Log.Formatter = new(logrus.TextFormatter)                   //default
	Log.Formatter.(*logrus.TextFormatter).DisableColors = false // remove colors
	Log.Level = logrus.TraceLevel
	Log.Out = os.Stdout
	os.Mkdir(logPath, os.ModePerm) //在运行程序的目录下创建logs目录
	writer, _ := rotatelogs.New(
		logPath+todayFileName(),
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(604800)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}
	Log.SetReportCaller(true)
	hook := lfshook.NewHook(writerMap, &mineFormatter{})

	Log.AddHook(hook)
	return Log
}

func todayFileName() string {
	return time.Now().Format(TimeFormatDay) + ".log"
}

func (s *mineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timeStamp := entry.Time.Format(TimeFormat)
	var newLog string
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s][%s][%s:%d %s] %s\n",
			timeStamp,
			entry.Level,
			fName,
			entry.Caller.Line,
			entry.Caller.Function,
			entry.Message)
		//newLog = fmt.Sprintf("%s-%s:%s\n", timeStamp, fName, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s][%s] %s\n",
			timeStamp,
			entry.Level,
			entry.Message,
		)
	}
	b.WriteString(newLog)
	return b.Bytes(), nil
}
