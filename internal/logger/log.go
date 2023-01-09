package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// Logger is to use logger
type Logger struct {
	TAG string
}

func (l Logger) String() string {
	return fmt.Sprintf("%s %s", time.Now().Format("20060102 15:04:05"), l.TAG)
}

// Fatal log
func (l Logger) Fatal(logMsg string) {
	logger.Printf("%s %8s %s%s", l, "FATAL:", lineNumber(), logMsg)
	os.Exit(2)
}

// Error log
func (l Logger) Error(logMsg string) {
	logger.Printf("%s %8s %s%s", l, "ERROR:", lineNumber(), logMsg)
}

// Warn log
func (l Logger) Warn(logMsg string) {
	logger.Printf("%s %8s %s", l, "WARN:", logMsg)
}

// Info log
func (l Logger) Info(logMsg string) {
	logger.Printf("%s %8s %s", l, "INFO:", logMsg)
}

// Verbose log with levels
func (l Logger) Verbose(logMsg string, level int) {
	if verbose >= level {
		logger.Printf("%s %8s %s", l, "VERBOSE:", logMsg)
	}
}

func lineNumber() string {
	if !doDebug {
		return " "
	}
	_, fn, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" [%s:%d] ", fn, line)
}
