package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	Success    = "Success"
	NotFound   = "Not Found"
	BadRequest = "Bad Request"
)

var EmptyMap = fiber.Map{}

func BuildResponse(ctx *fiber.Ctx, status bool, message string, data interface{}, code int) error {
	var statusStr string
	if status {
		statusStr = Success
		message = Success
	} else {
		switch code {
		case 400:
			statusStr = BadRequest
		case 404:
			statusStr = NotFound
		}
	}

	return ctx.Status(code).JSON(&Response{
		Status:  statusStr,
		Message: message,
		Data:    data,
	})
}
