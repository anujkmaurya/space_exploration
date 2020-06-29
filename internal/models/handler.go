package models

//Response http response
type Response struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}
