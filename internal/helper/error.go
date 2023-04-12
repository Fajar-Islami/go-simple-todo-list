package helper

import (
	"net/http"
	"strings"
)

func HelperErrorResponse(err error) *ErrorStruct {
	if strings.Contains(err.Error(), "Not Found") {
		return &ErrorStruct{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}
	return &ErrorStruct{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}
