package errorsclass

import (
	"errors"
	"testing"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func handleError(err error) Response {
	return Response{}
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		err error
		res Response
	}{
		{
			err: FakeInvalidQueryParamError(),
			res: Response{StatusCode: 400, Message: "'start' must be a date before 'end'"},
		},
		{
			err: FakeUnprocessableEntityError(),
			res: Response{StatusCode: 422, Message: "bill amount must be paid in full"},
		},
		{
			err: FakeTooManyRequestsError(),
			res: Response{StatusCode: 409, Message: "rate limit exceeded"},
		},
		{
			err: FakeInternalServerError(),
			res: Response{StatusCode: 500, Message: "internal error occurred"},
		},
		{
			err: errors.New("any error"),
			res: Response{StatusCode: 500, Message: "unexpected error occurred"},
		},
	}
}
