package usecase

import (
	"context"
	"database/sql"
	"errors"
	"go-todo-list/internal/helper"
	todosdto "go-todo-list/internal/pkg/dto"
	repository_mysql "go-todo-list/internal/pkg/repository/mysql"
	"go-todo-list/internal/utils"
	"log"
	"net/http"
)

type TodosUseCase interface {
	GetAllTodos(ctx context.Context, params todosdto.TodosFilter) (res []todosdto.TodosResp, err *helper.ErrorStruct)
	GetTodosByID(ctx context.Context, todosid int64) (res todosdto.TodosResp, err *helper.ErrorStruct)
	CreateTodos(ctx context.Context, data todosdto.TodosReqCreate) (res int64, err *helper.ErrorStruct)
	UpdateTodosByID(ctx context.Context, todosid int64, data todosdto.TodosReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteTodosByID(ctx context.Context, todosid int64) (res string, err *helper.ErrorStruct)
}

type TodosUseCaseImpl struct {
	todosrepository repository_mysql.TodosRepository
	currentfilepath string
}

func NewTodosUseCase(todosrepository repository_mysql.TodosRepository) TodosUseCase {
	return &TodosUseCaseImpl{
		todosrepository: todosrepository,
		currentfilepath: "internal/pkg/usecase/usecase_todos.go",
	}

}

func (tuc *TodosUseCaseImpl) GetAllTodos(ctx context.Context, params todosdto.TodosFilter) (res []todosdto.TodosResp, err *helper.ErrorStruct) {
	if errValidate := usecaseValidation(params); errValidate != nil {
		log.Println(errValidate)
		return res, errValidate
	}

	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	switch params.Short {
	case "title-asc":
		params.Short = "title asc"
	case "title-desc":
		params.Short = "title desc"
	case "date-asc":
		params.Short = "date asc"
	case "date-desc":
		params.Short = "date desc"
	case "done":
		params.Short = "is_active DESC, created_at DESC"
	default:
		params.Short = "date desc"
	}

	resRepo, errRepo := tuc.todosrepository.GetAllTodos(ctx, repository_mysql.FilterTodos{
		Limit:           params.Limit,
		Page:            params.Page,
		Short:           params.Short,
		ActivityGroupID: params.ActivityGroupID,
	})
	if errors.Is(errRepo, sql.ErrNoRows) {
		return res, &helper.ErrorStruct{
			Code: http.StatusNotFound,
			Err:  errors.New("No Data todos"),
		}
	}

	if errRepo != nil {
		helper.Logger(tuc.currentfilepath, helper.LoggerLevelError, "Error at GetAllTodos", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, todosdto.TodosResp{
			TodoID:          v.TodoID,
			ActivityGroupID: v.ActivityGroupID,
			Title:           v.Title,
			Priority:        v.Priority,
			IsActive:        *v.IsActive,
			CreatedAt:       utils.DateFormatter(v.CreatedAt.Time),
			UpdatedAt:       utils.DateFormatter(v.UpdatedAt.Time),
		})
	}

	return res, nil
}
func (tuc *TodosUseCaseImpl) GetTodosByID(ctx context.Context, todosid int64) (res todosdto.TodosResp, err *helper.ErrorStruct) {
	resRepo, errRepo := tuc.todosrepository.GetTodosByID(ctx, todosid)

	if errRepo != nil {
		helper.Logger(tuc.currentfilepath, helper.LoggerLevelError, "Error at GetTodosByID", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return todosdto.TodosResp{
		TodoID:          resRepo.TodoID,
		ActivityGroupID: resRepo.ActivityGroupID,
		Title:           resRepo.Title,
		Priority:        resRepo.Priority,
		IsActive:        *resRepo.IsActive,
		CreatedAt:       resRepo.CreatedAt.Time,
		UpdatedAt:       resRepo.UpdatedAt.Time,
	}, nil
}
func (tuc *TodosUseCaseImpl) CreateTodos(ctx context.Context, data todosdto.TodosReqCreate) (res int64, err *helper.ErrorStruct) {
	if errValidate := usecaseValidation(data); errValidate != nil {
		log.Println(errValidate)
		return res, errValidate
	}

	if data.Priority == "" {
		data.Priority = "very-high"
	}

	resRepo, errRepo := tuc.todosrepository.CreateTodos(ctx, repository_mysql.Todos{
		ActivityGroupID: int64(data.ActivityGroupID),
		Title:           data.Title,
		Priority:        data.Priority,
		IsActive:        data.IsActive,
	})

	if errRepo != nil {
		helper.Logger(tuc.currentfilepath, helper.LoggerLevelError, "Error at CreateTodos", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (tuc *TodosUseCaseImpl) UpdateTodosByID(ctx context.Context, todosid int64, data todosdto.TodosReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := usecaseValidation(data); errValidate != nil {
		log.Println(errValidate)
		return res, errValidate
	}

	if errValidate := data.Validate(); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errValidate,
		}
	}

	resRepo, errRepo := tuc.todosrepository.UpdateTodosByID(ctx, todosid, repository_mysql.Todos{
		ActivityGroupID: int64(data.ActivityGroupID),
		Title:           data.Title,
		Priority:        data.Priority,
		IsActive:        &data.IsActive,
	})

	if errRepo != nil {
		helper.Logger(tuc.currentfilepath, helper.LoggerLevelError, "Error at UpdateTodosByID", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (tuc *TodosUseCaseImpl) DeleteTodosByID(ctx context.Context, todosid int64) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := tuc.todosrepository.DeleteTodosByID(ctx, todosid)
	if errRepo != nil {
		helper.Logger(tuc.currentfilepath, helper.LoggerLevelError, "Error at DeleteTodosByID", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
