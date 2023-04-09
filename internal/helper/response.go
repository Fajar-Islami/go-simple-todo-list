package helper

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func BuildResponse(ctx echo.Context, status bool, message string, err string, data interface{}, code int) error {
	var splittedError []string

	if len(err) > 0 {
		splittedError = strings.Split(err, "\n")
	}

	return ctx.JSON(code, Response{
		Status:  status,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	})
}
