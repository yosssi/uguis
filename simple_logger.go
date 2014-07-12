package uguis

import (
	"log"
	"os"
)

// simpleLogger represents a simple logger.
type simpleLogger struct {
	*log.Logger
	level   uint8
	logC    chan Log
	closedC chan struct{}
}

// Print sends a print signal to the print log goroutine.
func (lgr *simpleLogger) Print(lg Log) {
	// Check the log level.
	if lg.Level < lgr.level {
		return
	}

	// Send a print signal to the print log goroutine.
	lgr.logC <- lg
}

// Close closes the log channel.
func (lgr *simpleLogger) Close() error {
	// Close the log channle.
	close(lgr.logC)

	// Wait until the print log goroutine is closed.
	<-lgr.closedC

	return nil
}

// print prints a log.
func (lgr *simpleLogger) printLog() {
	for lg := range lgr.logC {
		lgr.Logger.Println(lg)
	}
	lgr.closedC <- struct{}{}
}

// NewSimpleLogger creates and returns a simple logger.
func NewSimpleLogger(opts *SimpleLoggerOptions) Logger {
	// Initialize options.
	if opts == nil {
		opts = &SimpleLoggerOptions{}
	}
	opts.setDefaults()

	// Create a logger.
	lgr := &simpleLogger{
		Logger:  log.New(os.Stdout, "", 0),
		level:   opts.level(),
		logC:    make(chan Log, opts.LogCBfSize),
		closedC: make(chan struct{}),
	}

	// Launch a goroutine for printing a log.
	go lgr.printLog()

	return lgr
}
