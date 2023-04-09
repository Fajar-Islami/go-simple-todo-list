package utils

import (
	"os"

	"github.com/spf13/cast"
)

func EnvString(key string) string {
	return os.Getenv(key)
}

func EnvInt(key string) int {
	return cast.ToInt(os.Getenv(key))
}

func EnvBool(key string) bool {
	return cast.ToBool(os.Getenv(key))
}
