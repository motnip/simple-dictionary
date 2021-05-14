package system

import (
	"log"
)

type SermoLog struct {
	logErr *log.Logger
}

func NewLog() *SermoLog {
	initLogger := log.Default()
	initLogger.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	initLogger.SetPrefix("ERROR - ")
	return &SermoLog{
		logErr: initLogger,
	}
}

func (l *SermoLog) LogErr(message string) {
	l.logErr.Print(message)
}
