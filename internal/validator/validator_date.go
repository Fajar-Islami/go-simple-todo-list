package validator

import (
	"regexp"
	"time"

	"go-todo-list/internal/utils"

	v10 "github.com/go-playground/validator/v10"
)

func IsDdMmYyyyValidator(fl v10.FieldLevel) bool {
	ISDdMmYyyyString := "^\\d{4}\\-(0?[1-9]|1[012])\\-(0?[1-9]|[12][0-9]|3[01])$"
	regex := regexp.MustCompile(ISDdMmYyyyString)

	return regex.MatchString(fl.Field().String())
}

func AfterTodayValidator(fl v10.FieldLevel) bool {
	fieldValue, _ := fl.Field().Interface().(string)
	now := time.Now()
	expireTime, err := utils.ShortDateFromString(fieldValue)
	if err != nil {
		return false
	}
	if expireTime.Before(now) {
		return false
	}

	return true
}
