package main

import (
	"aulas/basics/tests"
	"errors"
	"os"
	"time"
)

type FakeTimer struct {
}

func (f FakeTimer) Now() time.Time {
	return time.Now()
}

func main() {
	cErr := *tests.NewCustomError("Error while estabilishing connection", tests.InternalErrorKind, errors.New("ENOTFOUND"))
	JSONLogger := tests.NewLogger(FakeTimer{}, os.Stdout, tests.JSONFormat)
	rawLogger := tests.NewLogger(FakeTimer{}, os.Stdout, tests.RawFormat)

	JSONLogger.Error("An error occurred while making HTTP request", cErr)
	rawLogger.Error("An error occurred while making HTTP request", cErr)

	JSONLogger.Info("A new log")
	rawLogger.Info("A new log")
}
