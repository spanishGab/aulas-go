package tests

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

type TimerMock struct {
}

func (t *TimerMock) Now() time.Time {
	return time.Date(2026, 10, 30, 10, 42, 0, 1, time.UTC)
}

func TestDebug(t *testing.T) {
	t.Run("Debug", func(t *testing.T) {
		t.Run(fmt.Sprintf("with %s set", logLevelEnvVarName), func(t *testing.T) {

			tests := map[string]struct {
				logger         *Logger
				input          string
				setEnvironment func()
				expected       string
			}{
				"should log when level is allowed": {
					input:          "hello, world!",
					setEnvironment: func() { t.Setenv(logLevelEnvVarName, "DEBUG") },
					expected:       "timestamp=2026-10-30T10:42:00.000000001Z level=DEBUG message=\"hello, world!\"\n",
				},
				"should not log when level is not allowed": {
					input:          "hello, world!",
					setEnvironment: func() { t.Setenv(logLevelEnvVarName, "INFO") },
					expected:       "",
				},
			}
			for name, test := range tests {
				t.Run(name, func(t *testing.T) {
					logBuffer := &bytes.Buffer{}
					t.Cleanup(func() { logBuffer.Reset() })

					test.setEnvironment()
					logger := NewLogger(&TimerMock{}, logBuffer, RawFormat)
					logger.Debug(test.input)

					result := logBuffer.String()
					if result != test.expected {
						t.Errorf("expected=%q, got=%q", test.expected, result)
					}
				})
			}

		})
	})
}
