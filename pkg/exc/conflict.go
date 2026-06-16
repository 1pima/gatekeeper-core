package exc

type AppConflict struct {
	Message string
}

func (e *AppConflict) Error() string {
	return e.Message
}

func (e *AppConflict) Unwrap() error {
	return e
}

func NewConflict(msg string) error {
	return &AppConflict{Message: msg}
}
