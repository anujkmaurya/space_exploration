package models

//AppError : defines application errors
type AppError struct {
	ErrorMessage    string
	HTTPResposeCode int
}

func (m *AppError) Error() string {
	return m.ErrorMessage
}

//CreateAppError : create App error
func CreateAppError(msg string, code int) *AppError {
	return &AppError{
		ErrorMessage:    msg,
		HTTPResposeCode: code,
	}
}
