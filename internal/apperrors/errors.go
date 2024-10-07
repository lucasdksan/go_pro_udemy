package apperrors

type StatusError struct {
	error
	status int
}

func NewWithStatus(err error, status int) error {
	return StatusError{
		error:  err,
		status: status,
	}
}

func (se StatusError) HTTPStatus() int {
	return se.status
}
