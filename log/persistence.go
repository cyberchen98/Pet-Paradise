package log

import (
	"fmt"
	"os"
	"path"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

var fileLogger *FileLogger

type FileLogger struct {
	Level       LogLevel
	fileDir     string
	fileName    string
	file        *os.File
	errFile     *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

func configureFileLogger(levelStr, fileDir, fileName string, maxSizeMB int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		fileDir:     fileDir,
		fileName:    fileName,
		maxFileSize: maxSizeMB,
		logChan:     make(chan *logMsg, maxSizeMB),
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.fileDir, f.fileName)
	file, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	errFile, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.file = file
	f.errFile = errFile
	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format(TIME_FORMAT)
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	logName := path.Join(f.fileDir, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	file.Close()
	os.Rename(logName, newLogName)
	file, err = os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f *FileLogger) writeLogBackground() {

	for {
		if f.checkSize(f.file) {
			newFile, err := f.splitFile(f.file)
			if err != nil {
				return
			}
			f.file = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.file, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFile) {
					newFile, err := f.splitFile(f.errFile)
					if err != nil {
						return
					}
					f.errFile = newFile
				}
				fmt.Fprintf(f.errFile, logInfo)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getRuntimeInfo(3)
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format(TIME_FORMAT),
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)

}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)

}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)

}

func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
