package apimodel

type Response struct {
	StatusCode int         `json:"code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	statusCode int
	msg string
}

func (err ErrorResponse) StatusCode() int {
	return err.statusCode
}

func (err ErrorResponse) Error() string {
	return err.msg
}

// generator function for the general http error
func NewResponseError(statusCode int, msg string) ErrorResponse {
	return ErrorResponse{
		statusCode: statusCode,
		msg:        msg,
	}
}