// Package log provides an interface and a few helper functions/constants.
package log

// Fields is a map of key value pairs of extra information worthwhile logging.
type Fields map[string]interface{}

// Logger is the logger interface that should be used by libraries to log.
type Logger interface {
	Trace(msg string, fields Fields)
	Debug(msg string, fields Fields)
	Info(msg string, fields Fields)
	Warn(msg string, fields Fields)
	Error(msg string, fields Fields)
}

// LogLevel defines the log level constants.
type LogLevel uint

// LogLevels constants
const (
	TRACE LogLevel = 10
	DEBUG LogLevel = 20
	INFO  LogLevel = 30
	WARN  LogLevel = 40
	ERROR LogLevel = 50
)

// Log is an internal helper function to be used by this library.
func Log(log Logger, level LogLevel, msg string, fields Fields) {
	if log != nil {
		switch level {
		case TRACE:
			log.Trace(msg, fields)
		case DEBUG:
			log.Debug(msg, fields)
		case INFO:
			log.Info(msg, fields)
		case WARN:
			log.Warn(msg, fields)
		case ERROR:
			log.Error(msg, fields)
		}
	}
}
