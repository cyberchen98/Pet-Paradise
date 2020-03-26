package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

var logger Logger

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

func ConfigureLogger(logLevel, logDir, fileName string, maxSize int64) error {
	lvl, err := parseLogLevel(logLevel)
	if err != nil {
		return err
	}
	if lvl == DEBUG {
		debugLogger.Level.SetLogLevel("debug")
		logger = debugLogger
	}
	fileLogger = &FileLogger{
		Level:       lvl,
		fileDir:     logDir,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
	logger = fileLogger
	return nil
}

func Log() Logger {
	return logger
}

type LogLevel int64

func (l *LogLevel) GetLogLevel() string {
	return l.String()
}

func (l *LogLevel) SetLogLevel(s string) error {
	lvl, err := parseLogLevel(s)
	if err != nil {
		return err
	}
	*l = lvl
	return nil
}

func (l *LogLevel) String() string {
	switch *l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "invalid"
}

func parseLogLevel(levelstr string) (LogLevel, error) {
	switch strings.ToLower(levelstr) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	}
	return 0, fmt.Errorf("invalid log level '%s' (debug, info, warn, error, fatal)", levelstr)
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getRuntimeInfo(skip int) (string, string, int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return "", "", -1
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}
