package errorsclass

import "errors"

var ErrInvalidQueryParam = errors.New("invalid query param error")
var ErrInvalidRequestBody = errors.New("invalid request body error")

var ErrBusinessRule = errors.New("business rule violation error")

var ErrInternal = errors.New("internal server error")

func FakeInvalidQueryParamError() error {
	return NewBadRequestError(ErrInvalidQueryParam, "'start' must be a date before 'end'")
}

func FakeInvalidRequestBodyError() error {
	return NewBadRequestError(ErrInvalidRequestBody, "'email' must be a valid e-mail address")
}

func FakeInvalidRequestError() error {
	return NewBadRequestError(errors.Join(ErrInvalidQueryParam, ErrInvalidRequestBody), "multiple input errors")
}

func FakeUnprocessableEntityError() error {
	return NewUnprocessableEntityError(ErrBusinessRule, "bill amount must be paid in full")
}

func FakeInternalServerError() error {
	return NewInternalServerError(ErrInternal, "SQL Statement error: 'SELCT' is an invalid statement")
}
