package repository_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type TodosRepository interface {
	GetAllTodos(ctx context.Context, params FilterTodos) (res []Todos, err error)
	GetTodosByID(ctx context.Context, todosid int64) (res Todos, err error)
	CreateTodos(ctx context.Context, data Todos) (res Todos, err error)
	UpdateTodosByID(ctx context.Context, todosid int64, data Todos) (res Todos, err error)
	DeleteTodosByID(ctx context.Context, todosid int64) (err error)
}

type TodosRepositoryImpl struct {
	db              *sql.DB
	tablename       string
	activitiiesRepo ActivitiesRepository
}

func NewTodosRepository(db *sql.DB) TodosRepository {
	return &TodosRepositoryImpl{
		db:              db,
		tablename:       "todos",
		activitiiesRepo: NewActivitiesRepository(db),
	}
}
func (acr *TodosRepositoryImpl) GetAllTodos(ctx context.Context, params FilterTodos) (res []Todos, err error) {
	db := acr.db
	sql := fmt.Sprintf("select todo_id, activity_group_id, title, priority, is_active, created_at, updated_at from %s where 1=1 ", acr.tablename)
	args := []any{}

	// %d order by ? limit ? offset ?"

	if params.ActivityGroupID != 0 {
		sql = sql + "and activity_group_id = ? "
		args = append(args, params.ActivityGroupID)
	}

	sql = sql + "order by ? limit ? offset ?"
	args = append(args, params.Short, params.Limit, params.Page)

	rows, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var todo Todos
		err := rows.Scan(&todo.TodoID, &todo.ActivityGroupID, &todo.Title, &todo.Priority, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Println("Error scanning get GetAllTodos : ", err)
		}
		res = append(res, todo)
	}

	return res, nil
}

func (acr *TodosRepositoryImpl) GetTodosByID(ctx context.Context, todosid int64) (res Todos, err error) {
	db := acr.db
	sqlQuery := fmt.Sprintf("select todo_id, activity_group_id, title, priority, is_active, created_at, updated_at from %s where todo_id=? limit 1", acr.tablename)

	rows, err := db.QueryContext(ctx, sqlQuery, todosid)
	if err != nil {
		return res, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&res.TodoID, &res.ActivityGroupID, &res.Title, &res.Priority, &res.IsActive, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			log.Println("Error scanning GetTodosByID : ", err)
			return res, err
		}
		return res, nil
	} else {
		return res, fmt.Errorf("Todo with ID %d Not Found", todosid)
	}

}

func (acr *TodosRepositoryImpl) CreateTodos(ctx context.Context, data Todos) (res Todos, err error) {
	db := acr.db

	// Check if activity is exist
	_, err = acr.activitiiesRepo.GetActivitiesByID(ctx, data.ActivityGroupID)
	if err != nil {
		return res, err
	}

	sql := fmt.Sprintf("insert into %s(activity_group_id,title,priority,is_active) values (?,?,?,?)", acr.tablename)

	result, err := db.ExecContext(ctx, sql, data.ActivityGroupID, data.Title, data.Priority, data.IsActive)
	if err != nil {
		return res, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	return acr.GetTodosByID(ctx, id)
}

func (acr *TodosRepositoryImpl) UpdateTodosByID(ctx context.Context, todosid int64, data Todos) (res Todos, err error) {
	db := acr.db

	// Check if todo exist
	resultGet, err := acr.GetTodosByID(ctx, todosid)
	if err != nil {
		return res, err
	}

	actGroupId := data.ActivityGroupID
	if data.ActivityGroupID < 1 {
		actGroupId = resultGet.ActivityGroupID
	}

	// Check if activity is exist
	_, err = acr.activitiiesRepo.GetActivitiesByID(ctx, actGroupId)
	if err != nil {
		return res, err
	}

	sql := fmt.Sprintf("update %s set ", acr.tablename)
	args := []any{}

	if data.ActivityGroupID != 0 {
		sql = sql + "activity_group_id = ?, "
		args = append(args, data.ActivityGroupID)
	}
	if data.Title != "" {
		sql = sql + "title = ?, "
		args = append(args, data.Title)
	}
	if data.Priority != "" {
		sql = sql + "priority = ?, "
		args = append(args, data.Priority)
	}
	if data.IsActive != nil {
		sql = sql + "is_active = ?, "
		args = append(args, data.IsActive)
	}

	sql = sql + "updated_at = now() where todo_id = ?"
	args = append(args, todosid)

	_, err = db.ExecContext(ctx, sql, args...)
	if err != nil {
		return res, err
	}

	return acr.GetTodosByID(ctx, todosid)
}

func (acr *TodosRepositoryImpl) DeleteTodosByID(ctx context.Context, todosid int64) (err error) {
	db := acr.db

	_, err = acr.GetTodosByID(ctx, todosid)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("delete from %s where todo_id=?", acr.tablename)

	_, err = db.ExecContext(ctx, sql, todosid)
	if err != nil {
		return err
	}

	return nil
}
