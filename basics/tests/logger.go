package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const logLevelEnvVarName = "LOG_LEVEL"

type Logger struct {
	timer  ITimer
	target io.Writer
	format LogFormat
}

func NewLogger(timer ITimer, target io.Writer, format LogFormat) *Logger {
	return &Logger{
		timer:  timer,
		target: target,
		format: format,
	}
}

func (l *Logger) Debug(message string) {
	if !l.canLog(DebugLogLevel) {
		return
	}
	fmt.Fprintln(l.target, l.formatLog(DebugLogLevel, message))
}

func (l *Logger) Info(message string) {
	if !l.canLog(InfoLogLevel) {
		return
	}
	fmt.Fprintln(l.target, l.formatLog(InfoLogLevel, message))
}

func (l *Logger) Warn(message string) {
	if !l.canLog(WarnLogLevel) {
		return
	}
	fmt.Fprintln(l.target, l.formatLog(WarnLogLevel, message))
}

func (l *Logger) Error(message string, err error) {
	if !l.canLog(ErrorLogLevel) {
		return
	}
	fmt.Fprintln(l.target, l.formatErrorLog(message, err))
}

func (l *Logger) canLog(currentLevel LogLevel) bool {
	logLevel := parseLogLevel(os.Getenv(logLevelEnvVarName))
	return logLevel <= currentLevel
}

func (l *Logger) formatLog(level LogLevel, message string) string {
	log := NewLog(l.timer.Now().UTC(), level, message)
	switch l.format {
	case JSONFormat:
		return l.formatToJSON(*log)
	default:
		return log.String()
	}
}

func (l *Logger) formatErrorLog(message string, err error) string {
	var cErr CustomError
	var ok bool
	if cErr, ok = errors.AsType[CustomError](err); !ok {
		cErr = *NewCustomError(err.Error(), UnknownKind, err)
	}

	log := NewErrorLog(l.timer.Now().UTC(), message, cErr)
	switch l.format {
	case JSONFormat:
		return l.formatToJSON(*log)
	default:
		return log.String()
	}
}

func (l *Logger) formatToJSON[L Log | ErrorLog](log L) string {
	result, _ := json.Marshal(log)
	return string(result)
}
