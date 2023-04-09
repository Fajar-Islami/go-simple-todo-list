package utils

import "github.com/labstack/echo/v4"

type MiddlewareType func(next echo.HandlerFunc) echo.HandlerFunc
