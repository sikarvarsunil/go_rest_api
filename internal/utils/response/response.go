package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "ok"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		var msg string
		switch err.ActualTag() {
		case "required":
			msg = fmt.Sprintf("field %s is a required field", err.Field())
		case "email":
			msg = fmt.Sprintf("field %s must be a valid email address", err.Field())
		case "min":
			msg = fmt.Sprintf("field %s must be at least %s characters", err.Field(), err.Param())
		case "max":
			msg = fmt.Sprintf("field %s must be at most %s characters", err.Field(), err.Param())
		default:
			msg = fmt.Sprintf("field %s is invalid", err.Field())
		}
		errMsgs = append(errMsgs, msg)
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
