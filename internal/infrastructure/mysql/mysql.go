package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"go-todo-list/internal/helper"
	"go-todo-list/internal/utils"
)

type MysqlConf struct {
	Username           string `env:"mysql_username"`
	Password           string `env:"mysql_password"`
	DbName             string `env:"mysql_dbname"`
	Host               string `env:"mysql_host"`
	Port               int    `env:"mysql_port"`
	Schema             string `env:"mysql_schema"`
	LogMode            bool   `env:"mysql_logMode"`
	MaxLifetime        int    `env:"mysql_maxLifetime"`
	MinIdleConnections int    `env:"mysql_minIdleConnections"`
	MaxOpenConnections int    `env:"mysql_maxOpenConnections"`
}

const currentfilepath = "internal/infrastructure/mysql/mysql.go"

func DatabaseInit() *sql.DB {
	var mysqlConfig = MysqlConf{
		Username:           utils.EnvString("mysql_username"),
		Password:           utils.EnvString("mysql_password"),
		DbName:             utils.EnvString("mysql_dbname"),
		Host:               utils.EnvString("mysql_host"),
		Port:               utils.EnvInt("mysql_port"),
		Schema:             utils.EnvString("mysql_schema"),
		LogMode:            utils.EnvBool("mysql_logMode"),
		MaxLifetime:        utils.EnvInt("mysql_maxLifetime"),
		MinIdleConnections: utils.EnvInt("mysql_minIdleConnections"),
		MaxOpenConnections: utils.EnvInt("mysql_maxOpenConnections"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)

	db, err := sql.Open("mysql", dsn)
	// if there is an error opening the connection, handle it
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelFatal, "", fmt.Errorf("Cannot conenct to database : %s", err.Error()))
		panic(err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(mysqlConfig.MinIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxLifeTime := time.Duration(mysqlConfig.MaxLifetime) * time.Second
	db.SetConnMaxLifetime(maxLifeTime)

	if err := db.Ping(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, "⇨ MySQL status is disconnected", err)
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, fmt.Sprintf("⇨ MySQL status is connected to %s:%d database %s \n", mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName), nil)

	return db
}
