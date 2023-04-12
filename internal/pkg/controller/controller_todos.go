package controller

import (
	"go-todo-list/internal/helper"
	todosdto "go-todo-list/internal/pkg/dto"
	todosusecase "go-todo-list/internal/pkg/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodosController interface {
	GetAllTodos(ctx *fiber.Ctx) error
	GetTodosByID(ctx *fiber.Ctx) error
	CreateTodos(ctx *fiber.Ctx) error
	UpdateTodosByID(ctx *fiber.Ctx) error
	DeleteTodosByID(ctx *fiber.Ctx) error
}

type TodosControllerImpl struct {
	todosusecase todosusecase.TodosUseCase
}

func NewTodosController(todosusecase todosusecase.TodosUseCase) TodosController {
	return &TodosControllerImpl{
		todosusecase: todosusecase,
	}
}

func (tc *TodosControllerImpl) GetAllTodos(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(todosdto.TodosFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, nil, http.StatusBadRequest)
	}

	res, err := tc.todosusecase.GetAllTodos(c, todosdto.TodosFilter{
		Short:           filter.Short,
		Limit:           filter.Limit,
		Page:            filter.Page,
		ActivityGroupID: filter.ActivityGroupID,
	})

	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (tc *TodosControllerImpl) GetTodosByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	todosid := ctx.Params("id_todos")
	if todosid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(todosid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, nil, http.StatusBadRequest)
	}

	res, err := tc.todosusecase.GetTodosByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, res, http.StatusOK)
}

func (tc *TodosControllerImpl) CreateTodos(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(todosdto.TodosReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, nil, http.StatusBadRequest)
		}
	}

	res, err := tc.todosusecase.CreateTodos(c, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, res, http.StatusOK)
}

func (tc *TodosControllerImpl) UpdateTodosByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	todosid := ctx.Params("id_todos")
	if todosid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, nil, http.StatusBadRequest)
	}

	data := new(todosdto.TodosReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		if err != nil {
			return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, nil, http.StatusBadRequest)
		}
	}

	activId, errConv := strconv.Atoi(todosid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, nil, http.StatusBadRequest)
	}

	res, err := tc.todosusecase.UpdateTodosByID(c, int64(activId), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDUPDATEDATA, res, http.StatusOK)
}

func (tc *TodosControllerImpl) DeleteTodosByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	todosid := ctx.Params("id_todos")
	if todosid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, nil, http.StatusBadRequest)
	}

	activId, errConv := strconv.Atoi(todosid)
	if errConv != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, nil, http.StatusBadRequest)
	}

	res, err := tc.todosusecase.DeleteTodosByID(c, int64(activId))
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDDELETEDATA, nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDDELETEDATA, res, http.StatusOK)
}
