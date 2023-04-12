package usecase

import (
	"fmt"
	"go-todo-list/internal/helper"
	"go-todo-list/internal/validator"
	"net/http"
	"reflect"

	v10 "github.com/go-playground/validator/v10"
)

func usecaseValidation(params any) *helper.ErrorStruct {

	if err := helper.Validate.Struct(params); err != nil {
		errs := err.(v10.ValidationErrors)
		var newErr validator.ValidationError

		for _, val := range errs {
			// fmt.Printf("error :  %#v \n\n ", val)
			// fmt.Printf("error :  %#v \n\n ", val.StructNamespace() )

			getField, _ := reflect.TypeOf(params).FieldByName(val.Field())
			jsonTag := getField.Tag.Get("json")

			var message string
			switch val.Tag() {
			case "required":
				message = fmt.Sprintf("%s cannot be null", jsonTag)
			default:
				message = fmt.Sprintf("validation error for '%s', Tag: %s", jsonTag, val.Tag())
			}

			newErr.Message = append(newErr.Message, message)
		}

		return &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  newErr,
		}
	}

	return nil
}
