package helper

import "strings"

// Response body
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// EmptyObj response
type EmptyObj struct{}

// ResponseSucces template if success
func ResponseSucces(status bool, msg string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: msg,
		Errors:  nil,
		Data:    data,
	}

	return res
}

// ResponseFailed template if error
func ResponseFailed(msg string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: msg,
		Errors:  splitError,
		Data:    data,
	}

	return res
}
