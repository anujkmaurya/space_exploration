package models

//Exception struct declaration
type Header struct {
	ProcessTime  string `json:"process_time"`
	ErrorMessage string `json:"err_msg"`
}

type AppError struct {
	ErrorMessage    string
	HttpResposeCode int
}

func (m *AppError) Error() string {
	return m.ErrorMessage
}
