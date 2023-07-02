package logger

import "fmt"

type consoleLogger struct{}

func NewConsoleLogger() consoleLogger {
	return consoleLogger{}
}

func (l consoleLogger) log(lvl, message string) {
	fmt.Printf("%s: %s\n", lvl, message)
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
