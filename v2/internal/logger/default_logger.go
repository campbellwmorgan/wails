package logger

import (
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// LogLevel is an alias for the public LogLevel
type LogLevel = logger.LogLevel

// Logger is a utlility to log messages to a number of destinations
type Logger struct {
	output         logger.Logger
	logLevel       LogLevel
	showLevelInLog bool
}

// New creates a new Logger. You may pass in a number of `io.Writer`s that
// are the targets for the logs
func New(output logger.Logger) *Logger {
	result := &Logger{
		logLevel:       logger.INFO,
		showLevelInLog: true,
		output:         output,
	}
	return result
}

// CustomLogger creates a new custom logger that prints out a name/id
// before the messages
func (l *Logger) CustomLogger(name string) CustomLogger {
	return newcustomLogger(l, name)
}

// HideLogLevel removes the loglevel text from the start of each logged line
func (l *Logger) HideLogLevel() {
	l.showLevelInLog = true
}

// SetLogLevel sets the minimum level of logs that will be output
func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

// Writeln writes directly to the output with no log level
// Appends a carriage return to the message
func (l *Logger) Writeln(message string) error {
	return l.output.Print(message)
}

// Write writes directly to the output with no log level
func (l *Logger) Write(message string) error {
	return l.output.Print(message)
}

// Print writes directly to the output with no log level
// Appends a carriage return to the message
func (l *Logger) Print(message string) error {
	return l.Write(message)
}

// Trace level logging. Works like Sprintf.
func (l *Logger) Trace(format string, args ...interface{}) error {
	if l.logLevel <= logger.TRACE {
		return l.output.Trace(fmt.Sprintf(format, args...))
	}
	return nil
}

// Debug level logging. Works like Sprintf.
func (l *Logger) Debug(format string, args ...interface{}) error {
	if l.logLevel <= logger.DEBUG {
		return l.output.Debug(fmt.Sprintf(format, args...))
	}
	return nil
}

// Info level logging. Works like Sprintf.
func (l *Logger) Info(format string, args ...interface{}) error {
	if l.logLevel <= logger.INFO {
		return l.output.Info(fmt.Sprintf(format, args...))
	}
	return nil
}

// Warning level logging. Works like Sprintf.
func (l *Logger) Warning(format string, args ...interface{}) error {
	if l.logLevel <= logger.WARNING {
		return l.output.Warning(fmt.Sprintf(format, args...))
	}
	return nil
}

// Error level logging. Works like Sprintf.
func (l *Logger) Error(format string, args ...interface{}) error {
	if l.logLevel <= logger.ERROR {
		return l.output.Error(fmt.Sprintf(format, args...))
	}
	return nil
}

// Fatal level logging. Works like Sprintf.
func (l *Logger) Fatal(format string, args ...interface{}) {
	err := l.output.Fatal(fmt.Sprintf(format, args...))
	// Not much we can do but print it out before exiting
	if err != nil {
		println(err.Error())
	}
	os.Exit(1)
}
