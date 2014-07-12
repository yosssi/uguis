package uguis

// Defaults
const (
	defaultSimpleLoggerLevelStr   = logLevelStrINFO
	defaultSimpleLoggerLogCBfSize = 4096
)

// SimpleLoggerOptions represents options for a simple logger.
type SimpleLoggerOptions struct {
	// LevelStr represents a log level string.
	LevelStr string
	// LogCBfSize represents a buffer size of a log channel.
	LogCBfSize uint32
}

// setDefaults sets defaults to the simple logger options.
func (opts *SimpleLoggerOptions) setDefaults() {
	if opts.LevelStr == "" {
		opts.LevelStr = defaultSimpleLoggerLevelStr
	}

	if opts.LogCBfSize == 0 {
		opts.LogCBfSize = defaultSimpleLoggerLogCBfSize
	}
}

// level returns a log level of the simple logger options.
func (opts *SimpleLoggerOptions) level() uint8 {
	return logLevel(opts.LevelStr)
}
