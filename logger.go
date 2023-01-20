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

func (l Logger) log(level, message string) {
	extraSpace := strings.Repeat(" ", 5-len(level))

	l.Output.WriteString(
		fmt.Sprintf("%s [%s]%s %s\n", time.Now().Format(timeFormat), level, extraSpace, message),
	)
}

// Info writes the message, prefixed [INFO], to the logger
func (l Logger) Info(message string) {
	if l.Level < LevelInfo {
		return
	}

	l.log("INFO", message)
}

// Infof writes the message with standard `fmt` formatting options, prefixed
// [INFO], to the logger
func (l Logger) Infof(message string, args ...any) {
	l.Info(fmt.Sprintf(message, args...))
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
func (l Logger) Warnf(message string, args ...any) {
	l.Warn(fmt.Sprintf(message, args...))
}

// Error writes the message, prefixed [ERROR], to the logger
func (l Logger) Error(message string) {
	if l.Level < LevelError {
		return
	}

	l.log("ERROR", message)
}

// Errorf writes the message with standard `fmt` formatting options, prefixed
// [ERROR], to the logger
func (l Logger) Errorf(message string, args ...any) {
	l.Error(fmt.Sprintf(message, args...))
}

// Debug writes the message, prefixed [DEBUG], to the logger
func (l Logger) Debug(message string) {
	if l.Level < LevelDebug {
		return
	}

	l.log("DEBUG", message)
}

// Debugf writes the message with standard `fmt` formatting options, prefixed
// [DEBUG], to the logger
func (l Logger) Debugf(message string, args ...any) {
	l.Debug(fmt.Sprintf(message, args...))
}
