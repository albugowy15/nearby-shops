package usecase

type UseCaseError struct {
	Message string
	Code    int
}

func NewUseCaseError(msg string, code int) *UseCaseError {
	return &UseCaseError{
		Message: msg,
		Code:    code,
	}
}

func (ucr UseCaseError) Error() string {
	return ucr.Message
}
