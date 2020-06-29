package models

//AppError : defines application errors
type AppError struct {
	ErrorMessage    string
	HTTPResposeCode int
}

func (m *AppError) Error() string {
	return m.ErrorMessage
}
