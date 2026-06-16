package exc

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}

func (e *InternalError) Unwrap() error {
	return e
}

func NewInternalError(msg string) error {
	return &InternalError{Message: msg}
}
