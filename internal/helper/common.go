package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

type ErrorStruct struct {
	Err  error
	Code int
}

var Validate = validator.New()
var mysqlErr *mysql.MySQLError

func MysqlCheckErrDuplicateEntry(err error) bool {
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return false
}
