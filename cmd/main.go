package main

import (
	rest "go-todo-list/internal/delivery/http"
	"go-todo-list/internal/infrastructure/container"
)

func main() {
	container.Initcont(".env")
	contConf := container.InitContainer()

	defer contConf.Mysqldb.Close()

	rest.HTTPRouteInit(contConf)
}
