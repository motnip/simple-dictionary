package system

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type SermoLog struct {
	logErr  *log.Logger
	logInfo *log.Logger
}

func NewLog() *SermoLog {

	return &SermoLog{
		logInfo: initInfo(),
		logErr:  initError(),
	}
}

func initDefaultLog() *log.Logger {
	initLogger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)
	return initLogger
}

func initInfo() *log.Logger {
	newLogger := initDefaultLog()
	newLogger.SetPrefix("INFO - ")
	return newLogger
}

func initError() *log.Logger {
	newLogger := initDefaultLog()
	newLogger.SetPrefix("ERROR - ")
	return newLogger
}

func (l *SermoLog) LogErr(message string) {
	fileName, line := getCaller()
	l.logErr.Print(fmt.Sprintf("[%s %d] - %s", fileName, line, message))
}

func (l *SermoLog) LogInfo(message string) {
	fileName, line := getCaller()
	l.logInfo.Print(fmt.Sprintf("[%s %d] - %s", fileName, line, message))
}

func getCaller() (string, int) {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).FileLine(pc)
}
