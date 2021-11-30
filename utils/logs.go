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

func init() {
	Logger.infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.warnLogger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.errorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.fatalLogger = log.New(os.Stdout, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *logger) Info(format string, a ...interface{}) {
	l.infoLogger.Println(fmt.Sprint(format, a))
}

func (l *logger) Warn(format string, a ...interface{}) {
	l.warnLogger.Println(fmt.Sprint(format, a))
}

func (l *logger) Error(format string, a ...interface{}) {
	l.errorLogger.Println(fmt.Sprint(format, a))
}

func (l *logger) Fatal(format string, a ...interface{}) {
	l.fatalLogger.Println(fmt.Sprint(format, a))
}
