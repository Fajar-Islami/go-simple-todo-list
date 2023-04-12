package repository_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"go-todo-list/internal/helper"
	"log"
)

type ActivitiesRepository interface {
	GetAllActivities(ctx context.Context, params FilterActivities) (res []Activities, err error)
	GetActivitiesByID(ctx context.Context, activitiesid int64) (res Activities, err error)
	CreateActivities(ctx context.Context, data Activities) (res Activities, err error)
	UpdateActivitiesByID(ctx context.Context, activitiesid int64, data Activities) (res Activities, err error)
	DeleteActivitiesByID(ctx context.Context, activitiesid int64) (res string, err error)
}

type ActivitiesRepositoryImpl struct {
	db        *sql.DB
	tablename string
}

func NewActivitiesRepository(db *sql.DB) ActivitiesRepository {
	return &ActivitiesRepositoryImpl{
		db:        db,
		tablename: "activities",
	}
}
func (acr *ActivitiesRepositoryImpl) GetAllActivities(ctx context.Context, params FilterActivities) (res []Activities, err error) {
	db := acr.db
	sql := fmt.Sprintf("select activity_id,title,email,created_at,updated_at from %s order by ? limit ? offset ?", acr.tablename)

	rows, err := db.QueryContext(ctx, sql, params.Short, params.Limit, params.Page)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var activity Activities
		err := rows.Scan(&activity.ActivityID, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt)
		if err != nil {
			log.Println("Error scanning get activity : ", err)
		}
		res = append(res, activity)
	}

	return res, nil
}

func (acr *ActivitiesRepositoryImpl) GetActivitiesByID(ctx context.Context, activitiesid int64) (res Activities, err error) {
	db := acr.db
	sqlQuery := fmt.Sprintf("select activity_id,title,email,created_at,updated_at from %s where activity_id=? limit 1", acr.tablename)

	rows, err := db.QueryContext(ctx, sqlQuery, activitiesid)
	if err != nil {
		return res, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&res.ActivityID, &res.Title, &res.Email, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			log.Println("Error scanning get one activity : ", err)
			return res, err
		}
		return res, nil
	} else {
		return res, fmt.Errorf("Activity with ID %d Not Found", activitiesid)
	}

}

func (acr *ActivitiesRepositoryImpl) CreateActivities(ctx context.Context, data Activities) (res Activities, err error) {
	db := acr.db
	sql := fmt.Sprintf("insert into %s(title,email) values (?,?)", acr.tablename)

	result, err := db.ExecContext(ctx, sql, data.Title, data.Email)
	if err != nil {
		return res, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res, err = acr.GetActivitiesByID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (acr *ActivitiesRepositoryImpl) UpdateActivitiesByID(ctx context.Context, activitiesid int64, data Activities) (res Activities, err error) {
	db := acr.db

	_, err = acr.GetActivitiesByID(ctx, activitiesid)
	if err != nil {
		return res, err
	}

	sql := fmt.Sprintf("update %s set ", acr.tablename)
	args := []any{}

	if data.Email != "" {
		sql = sql + "email = ?, "
		args = append(args, data.Email)
	}
	if data.Title != "" {
		sql = sql + "title = ?, "
		args = append(args, data.Title)
	}

	sql = sql + "updated_at = now() where activity_id = ?"
	args = append(args, activitiesid)

	_, err = db.ExecContext(ctx, sql, args...)
	if err != nil {
		return res, err
	}

	res, err = acr.GetActivitiesByID(ctx, activitiesid)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (acr *ActivitiesRepositoryImpl) DeleteActivitiesByID(ctx context.Context, activitiesid int64) (res string, err error) {
	db := acr.db

	_, err = acr.GetActivitiesByID(ctx, activitiesid)
	if err != nil {
		return res, err
	}

	sql := fmt.Sprintf("delete from %s where activity_id=?", acr.tablename)

	_, err = db.ExecContext(ctx, sql, activitiesid)
	if err != nil {
		return helper.DeleteDataFailed, err
	}

	return helper.DeleteDataSucceed, nil
}
