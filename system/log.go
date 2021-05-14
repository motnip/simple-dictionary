package system

import (
	"log"
)

type SermoLog struct {
	logErr  *log.Logger
	logInfo *log.Logger
}

func NewLog() *SermoLog {

	return &SermoLog{
		logErr:  initError(),
		logInfo: initInfo(),
	}
}

func initDefaultLog() *log.Logger {
	initLogger := log.Default() //available from 1.16 version
	initLogger.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	return initLogger
}

func initInfo() *log.Logger {
	initLogger := initDefaultLog()
	initLogger.SetPrefix("INFO - ")
	return initLogger
}

func initError() *log.Logger {
	initLogger := initDefaultLog()
	initLogger.SetPrefix("ERROR - ")
	return initLogger
}

func (l *SermoLog) LogErr(message string) {
	l.logErr.Print(message)
}

func (l *SermoLog) LogInfo(message string) {
	l.logErr.Print(message)
}
