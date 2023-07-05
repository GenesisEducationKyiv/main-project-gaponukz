package logger

import (
	"fmt"
	"time"
)

type consoleLogger struct{}

func NewConsoleLogger() consoleLogger {
	return consoleLogger{}
}

func (l consoleLogger) log(lvl, message string) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%s | %s | %s\n", formattedTime, lvl, message)
}

func (l consoleLogger) Info(message string) {
	l.log("INFO", message)
}

func (l consoleLogger) Warn(message string) {
	l.log("WARNING", message)
}

func (l consoleLogger) Error(message string) {
	l.log("ERROR", message)
}

func (l consoleLogger) Debug(message string) {
	l.log("DEBUG", message)
}
