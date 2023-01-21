package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	// LevelError only logs error messages
	LevelError = 0
	// LevelWarning logs warning and error messages
	LevelWarning = 1
	// LevelInfo logs info, warning, and error messages
	LevelInfo = 2
	// LevelDebug logs debug, info, warning, and error messages
	LevelDebug = 3
	// LevelTrace logs trace, debug, info, warning, and error messages
	LevelTrace = 4
	// LevelNone disables logging
	LevelNone = -1

	timeFormat = "2006-01-02T15:04:05.999Z07:00"
)

// Logger is the type used for writing log messages
type Logger struct {
	Output *os.File
	Level  int
}

// DefaultLogger logs errors to stderr
var DefaultLogger = Logger{
	Output: os.Stderr,
}

// NewLoggerFromEnv returns a logger whose level is configured by the `LOG`
// environment variable to stderr.
func NewLoggerFromEnv() Logger {
	var result Logger = Logger{
		Output: os.Stderr,
	}

	switch strings.ToLower(os.Getenv("LOG")) {
	case "trace":
		result.Level = LevelTrace
	case "debug":
		result.Level = LevelDebug
	case "warn", "warning":
		result.Level = LevelWarning
	case "info":
		result.Level = LevelInfo
	case "none":
		result.Level = LevelNone
	}

	return result
}

func (l Logger) log(level string, a ...any) {
	extraSpace := strings.Repeat(" ", 5-len(level))

	fmt.Fprint(l.Output, time.Now().Format(timeFormat))
	fmt.Fprint(l.Output, " [")
	fmt.Fprint(l.Output, level)
	fmt.Fprint(l.Output, "] ")
	fmt.Fprint(l.Output, extraSpace)
	fmt.Fprint(l.Output, a...)
	fmt.Fprint(l.Output, "\n")
}

// Info writes the arguments, prefixed [INFO], to the logger
func (l Logger) Info(a ...any) {
	if l.Level < LevelInfo {
		return
	}

	l.log("INFO", a...)
}

// Infof writes the message with standard `fmt` formatting options, prefixed
// [INFO], to the logger
func (l Logger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

// Warn writes the message, prefixed [WARN], to the logger
func (l Logger) Warn(message string) {
	if l.Level < LevelWarning {
		return
	}

	l.log("WARN", message)
}

// Warnf writes the message with standard `fmt` formatting options, prefixed
// [WARN], to the logger
func (l Logger) Warnf(format string, args ...any) {
	l.Warn(fmt.Sprintf(format, args...))
}

// Error writes the message, prefixed [ERROR], to the logger
func (l Logger) Error(a ...any) {
	if l.Level < LevelError {
		return
	}

	l.log("ERROR", a...)
}

// Errorf writes the message with standard `fmt` formatting options, prefixed
// [ERROR], to the logger
func (l Logger) Errorf(format string, args ...any) {
	l.Error(fmt.Sprintf(format, args...))
}

// Debug writes the message, prefixed [DEBUG], to the logger
func (l Logger) Debug(a ...any) {
	if l.Level < LevelDebug {
		return
	}

	l.log("DEBUG", a...)
}

// Debugf writes the message with standard `fmt` formatting options, prefixed
// [DEBUG], to the logger
func (l Logger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

// Trace writes the message, prefixed [TRACE], to the logger
func (l Logger) Trace(a ...any) {
	if l.Level < LevelTrace {
		return
	}

	l.log("TRACE", a...)
}

// Tracef writes the message with standard `fmt` formatting options, prefixed
// [TRACE], to the logger
func (l Logger) Tracef(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}
