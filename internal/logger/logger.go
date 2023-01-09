package logger

import (
	"io"
	"log"
	"os"
)

// loggers
var (
	logger  *log.Logger
	doDebug bool
	verbose int
)

// Init logs on file and if verbose then on console too
var Init = func() func(logFile string, consoleFlag bool, verboseLevel int, debugFlag bool) {

	f := false

	return func(logFile string, consoleFlag bool, verboseLevel int, debugFlag bool) {

		if f {
			log.Fatalln("logger is initialized more than once")
		}

		f = true

		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open log file", err.Error())
		}

		var out io.Writer

		if consoleFlag {
			out = io.MultiWriter(file, os.Stdout)
		} else {
			out = file
		}

		doDebug = debugFlag
		verbose = verboseLevel

		logger = log.New(out, "", 0)
		logger.SetFlags(0)

		startLogger := Logger{TAG: "START"}
		startLogger.Warn("logs started")
	}
}()
