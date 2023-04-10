package repository_mysql

import (
	"database/sql"
)

type (
	FilterActivities struct {
		Limit, Page int
		Short       string
	}
	Activities struct {
		ActivityID int64        `json:"activity_id"`
		Title      string       `json:"title"`
		Email      string       `json:"email"`
		CreatedAt  sql.NullTime `json:"created_at"`
		UpdatedAt  sql.NullTime `json:"updated_at"`
	}
)

type (
	FilterTodos struct {
		Limit, Page int
		Short       string
	}
	Todos struct {
		TodoID          int64        `json:"todo_id"`
		ActivityGroupID int64        `json:"activity_group_id"`
		Title           string       `json:"title"`
		Priority        int64        `json:"priority"`
		CreatedAt       sql.NullTime `json:"created_at"`
		UpdatedAt       sql.NullTime `json:"updated_at"`
	}
)
