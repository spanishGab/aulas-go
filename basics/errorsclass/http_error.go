package errorsclass

type HTTPErrorKind int

const (
	InternalError HTTPErrorKind = iota
	BusinessError
	ExternalDependencyError
)

func (h HTTPErrorKind) String() string {
	return [...]string{"InternalError", "BusinessError", "ExternalDependencyError"}[h]
}

type HTTPError struct {
	message string
	code    int
	kind    HTTPErrorKind
	err     error
}

func (h *HTTPError) Code() int {
	return h.code
}

func (h *HTTPError) Kind() string {
	return h.kind.String()
}

func (h *HTTPError) Error() string {
	return h.message
}

func (h *HTTPError) Unwrap() error {
	return h.err
}

func NewBadRequestError(err error, message string) *HTTPError {
	return &HTTPError{
		message: message,
		code:    400,
		kind:    BusinessError,
		err:     err,
	}
}

func NewUnprocessableEntityError(err error, message string) *HTTPError {
	return &HTTPError{
		message: message,
		code:    422,
		kind:    BusinessError,
		err:     err,
	}
}

func NewInternalServerError(err error, message string) *HTTPError {
	return &HTTPError{
		message: message,
		code:    500,
		kind:    InternalError,
		err:     err,
	}
}
