package apperrors

type RepositoryError struct {
	error
}

func NewRepositoryError(err error) error {
	return RepositoryError{error: err}
}
