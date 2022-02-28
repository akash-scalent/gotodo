package validation

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

const (
	MSG_INVALID_REQUEST = "invalid request"
)

func DecodeAndVaildate(r io.Reader, requestInstance interface{}) error {
	// decode the request
	err := json.NewDecoder(r).Decode(requestInstance)
	if err != nil {
		return err
	}

	// validate the request
	validate := validator.New()
	err = validate.Struct(requestInstance)
	if err != nil {
		return err
	}

	return nil
}
