package reader

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	Error       func(err error)
	Fatal       func(err string)
	Info        func(info string)
}

func NewLogger() (l Logger) {
	mod := os.FileMode(0777)
	_ = os.MkdirAll(ConfigDir, mod)
	openLogfile, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %s, logs will not be recorded in the filesystem\n", err.Error())
		l.Error = func(lerr error) {
			log.Println(lerr.Error())
		}
		l.Fatal = func(lerr string) {
			log.Fatal(lerr)
		}
		l.Info = func(info string) {
			log.Println(info)
		}
		return l
	}
	l.errorLogger = log.New(openLogfile, "Info:\t", log.Ldate|log.Ltime|log.Lshortfile)
	l.infoLogger = log.New(openLogfile, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
	l.Error = func(lerr error) {
		l.errorLogger.Println(lerr)
	}
	l.Fatal = func(lerr string) {
		l.errorLogger.Fatal(lerr)
	}
	l.Info = func(info string) {
		l.infoLogger.Println(info)
	}
	return l
}
