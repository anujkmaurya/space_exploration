package models

//Response http response
type Response struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}

type Header struct {
	ProcessTime  string `json:"process_time"`
	ErrorMessage string `json:"err_msg"`
}
