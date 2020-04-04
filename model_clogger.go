package badgerwrap

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Custom logger with info and debug. Although not necessary besides l, keeping info for being available. To add log level.
type CustomLog struct {
	l         *log.Logger
	logLevel  string
	spoolJSON bool // if we want JSON format
}

var thelogger CustomLog
var logLevels = make(map[string]string)
var delim = ": "

func NewLogger(level string) (Customlogger, error) {
	if thelogger.l != nil {
		return thelogger, nil
	}
	logLevels["debug"] = "DEBUG"
	logLevels["info"] = "INFO"

	_, exists := logLevels[level]
	if !exists {
		return nil, errors.New("log level passed is not supported")
	}

	thelogger = CustomLog{
		l: log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
	}
	return thelogger, nil
}

// not using pointer receiver as type contains pointer already.
func (l CustomLog) Infof(format string, args ...interface{}) {
	l.l.Output(2, logLevels["info"]+delim+fmt.Sprintf(format, args...))
}

// not using pointer receiver as type contains pointer already.
func (l CustomLog) Info(args ...interface{}) {
	l.l.Output(2, logLevels["info"]+delim+fmt.Sprint(args...))
}

// not using pointer receiver as type contains pointer already.
func (l CustomLog) Debugf(format string, args ...interface{}) {
	if l.logLevel == logLevels["debug"] {
		l.l.Output(2, logLevels["debug"]+delim+fmt.Sprintf(format, args...))
	}
}

// not using pointer receiver as type contains pointer already.
func (l CustomLog) Debug(args ...interface{}) {
	if l.logLevel == logLevels["debug"] {
		l.l.Output(2, logLevels["debug"]+delim+fmt.Sprint(args...))
	}
}
