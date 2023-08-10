package log

import (
	"fmt"
	"time"
	"sync"
)

const (
	LOGTYPE_INFO  string = "[INFO ]"
	LOGTYPE_ERROR string = "[ERROR]"
	LOGTYPE_DEBUG string = "[DEBUG]"
	LOGTYPE_TRACE string = "[TRACE]"
)

const (
	LOGLEVEL_INFO  int = 1
	LOGLEVEL_DEBUG int = 2
	LOGLEVEL_TRACE int = 3
)

type CommonLogger struct {
	loglevel int
}

var loggerInstance *CommonLogger
var once sync.Once

func GetLogger() *CommonLogger {
    once.Do(func() {
        loggerInstance = &CommonLogger{
        }
    })
    return loggerInstance
}

func (logger CommonLogger) SetLogLevel(level string) {

	logger.loglevel = LOGLEVEL_INFO

	if level == "debug"{
		logger.loglevel = LOGLEVEL_DEBUG
	}
	if level == "trace"{
		logger.loglevel = LOGLEVEL_TRACE
	}
}

func getTimeStamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05.000")
}

func getFormatString(logType string, contents string) string {
	var formatString = fmt.Sprintf("[%s]%s %s", getTimeStamp(), logType, contents)
	return formatString
}


func writeLog(msg string) bool {
	
	fmt.Println(msg)
	return true
}

func (logger CommonLogger) Info(contents string) bool {

	msg := getFormatString(LOGTYPE_INFO, contents)
	writeLog(msg)
	return true
}

func (logger CommonLogger) Error(contents string) bool {

	msg := getFormatString(LOGTYPE_ERROR, contents)
	writeLog(msg)
	return true
}

func (logger CommonLogger) Debug(contents string) bool {

	if logger.loglevel == LOGLEVEL_DEBUG || logger.loglevel == LOGLEVEL_TRACE{
		msg := getFormatString(LOGTYPE_DEBUG, contents)
		writeLog(msg)
	}

	return true
}

func (logger CommonLogger) Trace(contents string) bool {

	if logger.loglevel == LOGLEVEL_TRACE{
		msg := getFormatString(LOGTYPE_TRACE, contents)
		writeLog(msg)
	}

	return true
}

func (logger CommonLogger) BasicPrint(contents string) bool {

	writeLog(contents)
	return true
}


func (logger CommonLogger) Dummy() bool {

	return true
}
