package usecase

import (
	"context"
	"database/sql"
	"errors"
	"go-todo-list/internal/helper"
	activitiesdto "go-todo-list/internal/pkg/dto"
	repository_mysql "go-todo-list/internal/pkg/repository/mysql"
	"go-todo-list/internal/utils"
	"log"
	"net/http"
)

type ActivitiesUseCase interface {
	GetAllActivities(ctx context.Context, params activitiesdto.ActivitiesFilter) (res []activitiesdto.ActivitiesResp, err *helper.ErrorStruct)
	GetActivitiesByID(ctx context.Context, activitiesid int64) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct)
	CreateActivities(ctx context.Context, data activitiesdto.ActivitiesReqCreate) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct)
	UpdateActivitiesByID(ctx context.Context, activitiesid int64, data activitiesdto.ActivitiesReqUpdate) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct)
	DeleteActivitiesByID(ctx context.Context, activitiesid int64) (err *helper.ErrorStruct)
}

type ActivitiesUseCaseImpl struct {
	activitiesrepository repository_mysql.ActivitiesRepository
	currentfilepath      string
}

func NewActivitiesUseCase(activitiesrepository repository_mysql.ActivitiesRepository) ActivitiesUseCase {
	return &ActivitiesUseCaseImpl{
		activitiesrepository: activitiesrepository,
		currentfilepath:      "internal/pkg/usecase/usecase_activities.go",
	}

}

func (acu *ActivitiesUseCaseImpl) GetAllActivities(ctx context.Context, params activitiesdto.ActivitiesFilter) (res []activitiesdto.ActivitiesResp, err *helper.ErrorStruct) {
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
	default:
		params.Short = "date desc"
	}

	resRepo, errRepo := acu.activitiesrepository.GetAllActivities(ctx, repository_mysql.FilterActivities{
		Limit: params.Limit,
		Page:  params.Page,
		Short: params.Short,
	})
	if errors.Is(errRepo, sql.ErrNoRows) {
		return res, &helper.ErrorStruct{
			Code: http.StatusNotFound,
			Err:  errors.New("No Data activities"),
		}
	}

	if errRepo != nil {
		helper.Logger(acu.currentfilepath, helper.LoggerLevelError, "Error at GetAllActivities", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, activitiesdto.ActivitiesResp{
			ActivityID: v.ActivityID,
			Title:      v.Title,
			Email:      v.Email,
			CreatedAt:  utils.DateFormatter(v.CreatedAt.Time),
			UpdatedAt:  utils.DateFormatter(v.UpdatedAt.Time),
		})
	}

	return res, nil
}
func (acu *ActivitiesUseCaseImpl) GetActivitiesByID(ctx context.Context, activitiesid int64) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct) {
	resRepo, errRepo := acu.activitiesrepository.GetActivitiesByID(ctx, activitiesid)

	if errRepo != nil {
		helper.Logger(acu.currentfilepath, helper.LoggerLevelError, "Error at GetActivitiesByID", errRepo)
		err = helper.HelperErrorResponse(errRepo)
		return res, err
	}

	return activitiesdto.ActivitiesResp{
		ActivityID: resRepo.ActivityID,
		Title:      resRepo.Title,
		Email:      resRepo.Email,
		CreatedAt:  resRepo.CreatedAt.Time,
		UpdatedAt:  resRepo.UpdatedAt.Time,
	}, nil
}

func (acu *ActivitiesUseCaseImpl) CreateActivities(ctx context.Context, data activitiesdto.ActivitiesReqCreate) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct) {
	if errValidate := usecaseValidation(data); errValidate != nil {
		log.Println(errValidate)
		return res, errValidate
	}

	resRepo, errRepo := acu.activitiesrepository.CreateActivities(ctx, repository_mysql.Activities{
		Title: data.Title,
		Email: data.Email,
	})

	if errRepo != nil {
		helper.Logger(acu.currentfilepath, helper.LoggerLevelError, "Error at CreateActivities", errRepo)
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return activitiesdto.ActivitiesResp{
		ActivityID: resRepo.ActivityID,
		Title:      resRepo.Title,
		Email:      resRepo.Email,
		CreatedAt:  utils.DateFormatter(resRepo.CreatedAt.Time),
		UpdatedAt:  utils.DateFormatter(resRepo.UpdatedAt.Time),
	}, nil
}

func (acu *ActivitiesUseCaseImpl) UpdateActivitiesByID(ctx context.Context, activitiesid int64, data activitiesdto.ActivitiesReqUpdate) (res activitiesdto.ActivitiesResp, err *helper.ErrorStruct) {
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

	resRepo, errRepo := acu.activitiesrepository.UpdateActivitiesByID(ctx, activitiesid, repository_mysql.Activities{
		Title: data.Title,
		Email: data.Email,
	})

	if errRepo != nil {
		helper.Logger(acu.currentfilepath, helper.LoggerLevelError, "Error at UpdateActivitiesByID", errRepo)
		err = helper.HelperErrorResponse(errRepo)
		return res, err
	}

	return activitiesdto.ActivitiesResp{
		ActivityID: resRepo.ActivityID,
		Title:      resRepo.Title,
		Email:      resRepo.Email,
		CreatedAt:  utils.DateFormatter(resRepo.CreatedAt.Time),
		UpdatedAt:  utils.DateFormatter(resRepo.UpdatedAt.Time),
	}, nil
}
func (acu *ActivitiesUseCaseImpl) DeleteActivitiesByID(ctx context.Context, activitiesid int64) (err *helper.ErrorStruct) {
	_, errRepo := acu.activitiesrepository.DeleteActivitiesByID(ctx, activitiesid)
	if errRepo != nil {
		return helper.HelperErrorResponse(errRepo)
	}

	return nil
}
