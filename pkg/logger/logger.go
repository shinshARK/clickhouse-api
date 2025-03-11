// pkg/logger/logger.go
package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

func New(level string) *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	l.log(l.infoLogger, msg, keyvals...)
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	l.log(l.errorLogger, msg, keyvals...)
}

func (l *Logger) Fatal(msg string, keyvals ...interface{}) {
	l.log(l.fatalLogger, msg, keyvals...)
	os.Exit(1)
}

func (l *Logger) log(logger *log.Logger, msg string, keyvals ...interface{}) {
	args := []interface{}{msg}
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			args = append(args, keyvals[i], keyvals[i+1])
		}
	}
	logger.Println(args...)
}
