package utils

import (
	"encoding/json"
	"net/http"

	"github.com/akash-scalent/gotodo/apimodel"
)

const (
	STATUS_SUCCESS               = "success"
	STATUS_FAILED                = "failed"
	STATUS_INTERNAL_SERVER_ERROR = "internal server error"
)

type name struct {
	
}

func SendResponseWithData(w http.ResponseWriter, statusCode int, msg string, payload interface{}) {
	res := apimodel.Response{
		StatusCode: statusCode,
		Status:     STATUS_SUCCESS,
		Message:    msg,
		Data:       payload,
	}
	response, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(response)
}

func SendResponseWithError(w http.ResponseWriter,err apimodel.ErrorResponse, payload interface{}) {
	res := apimodel.Response{
		StatusCode: err.StatusCode(),
		Status:     STATUS_FAILED,
		Message:    err.Error(),
		Data:       payload,
	}
	json.NewEncoder(w).Encode(res)
	w.WriteHeader(err.StatusCode())
}
