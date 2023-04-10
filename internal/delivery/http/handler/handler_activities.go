package handler

import (
	"go-todo-list/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	activitiescontroller "go-todo-list/internal/pkg/controller"

	activitiesrepository "go-todo-list/internal/pkg/repository/mysql"

	activitiesusecase "go-todo-list/internal/pkg/usecase"
)

func ActivitiesRoute(r fiber.Router, containerConf *container.Container) {
	repo := activitiesrepository.NewActivitiesRepository(containerConf.Mysqldb)
	usecase := activitiesusecase.NewActivitiesUseCase(repo)
	controller := activitiescontroller.NewActivitiesController(usecase)

	r.Get("", controller.GetAllActivities)
	r.Get("/:id_activities", controller.GetActivitiesByID)
	r.Post("", controller.CreateActivities)
	r.Patch("/:id_activities", controller.UpdateActivitiesByID)
	r.Delete("/:id_activities", controller.DeleteActivitiesByID)
}
