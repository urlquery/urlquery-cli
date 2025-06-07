package logger

import (
	"log"
	"os"
)

// Logger levels
const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

// Logger represents a simple logger
type Logger struct {
	level  int
	logger *log.Logger
}

// New creates a new logger with the specified level
func New(level int) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// Default logger instance
var defaultLogger = New(LevelInfo)

// SetLevel sets the logging level for the default logger
func SetLevel(level int) {
	defaultLogger.level = level
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Warn logs a warning message
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Info logs an info message
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Debug logs a debug message
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level >= LevelError {
		l.logger.Printf("[ERROR] "+format, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level >= LevelWarn {
		l.logger.Printf("[WARN] "+format, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level >= LevelInfo {
		l.logger.Printf("[INFO] "+format, args...)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level >= LevelDebug {
		l.logger.Printf("[DEBUG] "+format, args...)
	}
}

// Fatal logs an error message and exits
func Fatal(format string, args ...interface{}) {
	defaultLogger.logger.Printf("[FATAL] "+format, args...)
	os.Exit(1)
}

// ParseLevel parses a string level to int
func ParseLevel(level string) int {
	switch level {
	case "error":
		return LevelError
	case "warn", "warning":
		return LevelWarn
	case "info":
		return LevelInfo
	case "debug":
		return LevelDebug
	default:
		return LevelInfo
	}
}

// EnableDebug enables debug logging if the DEBUG environment variable is set
func EnableDebug() {
	if os.Getenv("DEBUG") != "" {
		SetLevel(LevelDebug)
		Debug("Debug logging enabled")
	}
}

// LogAPIRequest logs an API request for debugging
func LogAPIRequest(method, url string) {
	Debug("API Request: %s %s", method, url)
}

// LogAPIResponse logs an API response for debugging
func LogAPIResponse(statusCode int, url string) {
	if statusCode >= 400 {
		Warn("API Response: %d %s", statusCode, url)
	} else {
		Debug("API Response: %d %s", statusCode, url)
	}
}

// LogError logs an error with context
func LogError(operation string, err error) {
	Error("Operation '%s' failed: %v", operation, err)
}

// LogSuccess logs a successful operation
func LogSuccess(operation string) {
	Info("Operation '%s' completed successfully", operation)
}
