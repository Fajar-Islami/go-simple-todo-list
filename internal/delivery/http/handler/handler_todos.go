package handler

import (
	"go-todo-list/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	activitiescontroller "go-todo-list/internal/pkg/controller"

	activitiesrepository "go-todo-list/internal/pkg/repository/mysql"

	activitiesusecase "go-todo-list/internal/pkg/usecase"
)

func TodosRoute(r fiber.Router, containerConf *container.Container) {
	repo := activitiesrepository.NewTodosRepository(containerConf.Mysqldb)
	usecase := activitiesusecase.NewTodosUseCase(repo)
	controller := activitiescontroller.NewTodosController(usecase)

	r.Get("", controller.GetAllTodos)
	r.Get("/:id_todos", controller.GetTodosByID)
	r.Post("", controller.CreateTodos)
	r.Patch("/:id_todos", controller.UpdateTodosByID)
	r.Delete("/:id_todos", controller.DeleteTodosByID)
}
