package tests

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ITimer interface {
	Now() time.Time
}

type LogFormat int

const (
	JSONFormat LogFormat = iota
	RawFormat
)

type LogLevel int

const (
	DebugLogLevel LogLevel = iota
	InfoLogLevel
	WarnLogLevel
	ErrorLogLevel
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

func (l LogLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func parseLogLevel(logLevel string) LogLevel {
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		return DebugLogLevel
	case "INFO":
		return InfoLogLevel
	case "WARN":
		return WarnLogLevel
	case "ERROR":
		return ErrorLogLevel
	default:
		return DebugLogLevel
	}
}

type Log struct {
	Timestamp_ time.Time `json:"timestamp"`
	Level      LogLevel  `json:"level"`
	Message    string    `json:"message"`
}

func NewLog(timestamp_ time.Time, level LogLevel, message string) *Log {
	return &Log{
		Timestamp_: timestamp_,
		Level:      level,
		Message:    message,
	}
}

func (l Log) String() string {
	return fmt.Sprintf("timestamp=%s level=%s message=%q", l.Timestamp_.Format(time.RFC3339Nano), l.Level, l.Message)
}

type ErrorLog struct {
	Log
	Err CustomError `json:"error"`
}

func NewErrorLog(timestamp_ time.Time, message string, err CustomError) *ErrorLog {
	return &ErrorLog{
		Timestamp_: timestamp_,
		Level:      ErrorLogLevel,
		Message:    message,
		Err:        err,
	}
}

func (l ErrorLog) String() string {
	return fmt.Sprintf("%s error=%s", l.Log.String(), l.Err.String())
}
