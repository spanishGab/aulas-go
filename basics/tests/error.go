package tests

import (
	"encoding/json"
	"fmt"
)

type ErrorKind int

const (
	InternalErrorKind ErrorKind = iota
	BusinessErrorKind
	UnknownKind
)

func (e ErrorKind) String() string {
	return [...]string{"InternalError", "BusinessError", "UnknownKind"}[e]
}

func (e ErrorKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

type CustomError struct {
	Message string    `json:"message"`
	Kind    ErrorKind `json:"kind"`
	Cause   error     `json:"cause"`
}

func NewCustomError(message string, kind ErrorKind, cause error) *CustomError {
	return &CustomError{
		Message: message,
		Kind:    kind,
		Cause:   cause,
	}
}

func (c CustomError) Error() string {
	return c.Message
}

func (c CustomError) Unwrap() error {
	return c.Cause
}

func (c CustomError) String() string {
	return fmt.Sprintf("%s occurred - %q: %s", c.Kind, c.Message, c.Cause)
}

func (c CustomError) MarshalJSON() ([]byte, error) {
	marshal := []byte{}
	marshal = fmt.Appendf(marshal, "{\"message\": %q, \"kind\": %q, \"cause\": %q}", c.Message, c.Kind, c.Cause.Error())
	return marshal, nil
}
