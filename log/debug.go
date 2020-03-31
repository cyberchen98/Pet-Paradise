package log

import (
	"fmt"
	"time"
)

var debugLogger = &DebugLogger{}

type DebugLogger struct {
	Level LogLevel
}

func (c *DebugLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c *DebugLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getRuntimeInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format(TIME_FORMAT), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c *DebugLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c *DebugLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c *DebugLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c *DebugLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c *DebugLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
