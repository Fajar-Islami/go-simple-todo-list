package controller

import (
	"go-todo-list/internal/helper"
	activitiesdto "go-todo-list/internal/pkg/dto"
	activitiesusecase "go-todo-list/internal/pkg/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActivitiesController interface {
	GetAllActivities(ctx *fiber.Ctx) error
	GetActivitiesByID(ctx *fiber.Ctx) error
	CreateActivities(ctx *fiber.Ctx) error
	UpdateActivitiesByID(ctx *fiber.Ctx) error
	DeleteActivitiesByID(ctx *fiber.Ctx) error
}

type ActivitiesControllerImpl struct {
	activitiesusecase activitiesusecase.ActivitiesUseCase
}

func NewActivitiesController(activitiesusecase activitiesusecase.ActivitiesUseCase) ActivitiesController {
	return &ActivitiesControllerImpl{
		activitiesusecase: activitiesusecase,
	}
}

func (uc *ActivitiesControllerImpl) GetAllActivities(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(activitiesdto.ActivitiesFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.GetAllActivities(c, activitiesdto.ActivitiesFilter{
		Short: filter.Short,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (uc *ActivitiesControllerImpl) GetActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, "id_activities cant be empty", nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, errConv.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.GetActivitiesByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (uc *ActivitiesControllerImpl) CreateActivities(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(activitiesdto.ActivitiesReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
		}
	}

	res, err := uc.activitiesusecase.CreateActivities(c, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, res, http.StatusCreated)
}

func (uc *ActivitiesControllerImpl) UpdateActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, "id_activities cant be empty", nil, http.StatusBadRequest)
	}

	data := new(activitiesdto.ActivitiesReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
		}
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, errConv.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.UpdateActivitiesByID(c, int64(activId), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDUPDATEDATA, res, http.StatusOK)
}

func (uc *ActivitiesControllerImpl) DeleteActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, "id_activities cant be empty", nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, errConv.Error(), nil, http.StatusBadRequest)
	}

	err := uc.activitiesusecase.DeleteActivitiesByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDDELETEDATA, helper.EmptyMap, http.StatusOK)
}
