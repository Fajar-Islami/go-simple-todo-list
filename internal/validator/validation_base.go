package validator

import (
	"bytes"
	"strings"

	v10 "github.com/go-playground/validator/v10"
)

type ValidationHandler struct {
	Tag          string
	Func         v10.Func
	ErrorMessage string
}

type ValidationError struct {
	Message []string
}

func (v ValidationError) Error() string {
	buff := bytes.NewBufferString("")

	for i := 0; i < len(v.Message); i++ {
		errMessage := v.Message[i]
		buff.WriteString(errMessage)
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}
