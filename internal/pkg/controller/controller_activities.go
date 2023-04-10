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
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.activitiesusecase.GetAllActivities(c, activitiesdto.ActivitiesFilter{
		Short: filter.Short,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *ActivitiesControllerImpl) GetActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "id_activities is required", nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, errConv.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.GetActivitiesByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *ActivitiesControllerImpl) CreateActivities(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(activitiesdto.ActivitiesReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, http.StatusBadRequest)
		}
	}

	res, err := uc.activitiesusecase.CreateActivities(c, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}

func (uc *ActivitiesControllerImpl) UpdateActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "id_activities is required", nil, http.StatusBadRequest)
	}

	data := new(activitiesdto.ActivitiesReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, http.StatusBadRequest)
		}
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, errConv.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.UpdateActivitiesByID(c, int64(activId), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDUPDATEDATA, "", res, fiber.StatusOK)
}

func (uc *ActivitiesControllerImpl) DeleteActivitiesByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	activitiesid := ctx.Params("id_activities")
	if activitiesid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, "id_activities is required", nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(activitiesid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, errConv.Error(), nil, http.StatusBadRequest)
	}

	res, err := uc.activitiesusecase.DeleteActivitiesByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDDELETEDATA, "", res, fiber.StatusOK)
}
