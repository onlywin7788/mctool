package log

import (
	"fmt"
	"time"
)

const (
	LOGTYPE_INFO  string = "[INFO ]"
	LOGTYPE_ERROR string = "[ERROR]"
	LOGTYPE_DEBUG string = "[DEBUG]"
	LOGTYPE_TRACE string = "[TRACE]"
)

type CommonLogger struct {
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

func (logger CommonLogger) Debug(contents string) bool {

	msg := getFormatString(LOGTYPE_DEBUG, contents)
	writeLog(msg)
	return true
}

func (logger CommonLogger) Error(contents string) bool {

	msg := getFormatString(LOGTYPE_ERROR, contents)
	writeLog(msg)
	return true
}

func (logger CommonLogger) Trace(contents string) bool {

	msg := getFormatString(LOGTYPE_TRACE, contents)
	writeLog(msg)
	return true
}

func (logger CommonLogger) BasicPrint(contents string) bool {

	writeLog(contents)
	return true
}


func (logger CommonLogger) Dummy() bool {

	return true
}
