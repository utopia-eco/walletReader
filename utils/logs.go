package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	Logger *logger
)

type logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

func InitLogger() {
	Logger = &logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stdout, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *logger) Info(format string, a ...interface{}) {
	println(fmt.Sprintf(format, a))
	Logger.infoLogger.Println(fmt.Sprintf(format, a))
}

func (l *logger) Warn(format string, a ...interface{}) {
	Logger.warnLogger.Println(fmt.Sprintf(format, a))
}

func (l *logger) Error(format string, a ...interface{}) {
	Logger.errorLogger.Println(fmt.Sprintf(format, a))
}

func (l *logger) Fatal(format string, a ...interface{}) {
	Logger.fatalLogger.Println(fmt.Sprintf(format, a))
}
