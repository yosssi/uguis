package uguis

import (
	"fmt"
	"time"
)

// Log time layout
const logTimeLayout = "2006-01-02 15:04:05.000 -0700 MST"

// Log levels
const (
	LogLevelTRACE = iota
	LogLevelDEBUG
	LogLevelINFO
	LogLevelWARN
	LogLevelERROR
	LogLevelFATAL
)

// Log level strings
const (
	logLevelStrTRACE = "TRACE"
	logLevelStrDEBUG = "DEBUG"
	logLevelStrINFO  = "INFO"
	logLevelStrWARN  = "WARN"
	logLevelStrERROR = "ERROR"
	logLevelStrFATAL = "FATAL"
)

// Map between log levels and strings
var logLevelStrs = map[uint8]string{
	LogLevelTRACE: logLevelStrTRACE,
	LogLevelDEBUG: logLevelStrDEBUG,
	LogLevelINFO:  logLevelStrINFO,
	LogLevelWARN:  logLevelStrWARN,
	LogLevelERROR: logLevelStrERROR,
	LogLevelFATAL: logLevelStrFATAL,
}

// Log represents a log.
type Log struct {
	// Level represents a log level.
	Level uint8
	// AppName represents an application name which creates the log.
	AppName string
	// ServiceName represents a service name which creates the log.
	ServiceName string
	// Msg represents a log message.
	Msg string
	// Time represent a log time.
	Time time.Time
}

// String returns a log string.
func (lg Log) String() string {
	return fmt.Sprintf(
		"[%s][%s][%s][%s] %s",
		lg.AppName,
		lg.Time.Format(logTimeLayout),
		logLevelStrs[lg.Level],
		lg.ServiceName,
		lg.Msg,
	)
}

// NewLog creates and returns a log.
func NewLog(level uint8, appName string, serviceName string, msg string) Log {
	return Log{
		Level:       level,
		AppName:     appName,
		ServiceName: serviceName,
		Msg:         msg,
		Time:        time.Now(),
	}
}

// logLevel converts a log level string to a log level.
func logLevel(logLevelStr string) uint8 {
	for level, levelStr := range logLevelStrs {
		if logLevelStr == levelStr {
			return level
		}
	}
	return LogLevelINFO
}
