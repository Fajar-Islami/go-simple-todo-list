package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"go-todo-list/internal/helper"
	"go-todo-list/internal/utils"
)

type MysqlConf struct {
	Username           string `env:"MYSQL_USER"`
	Password           string `env:"MYSQL_PASSWORD"`
	DbName             string `env:"MYSQL_DBNAME"`
	Host               string `env:"MYSQL_HOST"`
	Port               int    `env:"MYSQL_PORT"`
	Schema             string `env:"MYSQL_schema"`
	LogMode            bool   `env:"MYSQL_logMode"`
	MaxLifetime        int    `env:"MYSQL_maxLifetime"`
	MinIdleConnections int    `env:"MYSQL_minIdleConnections"`
	MaxOpenConnections int    `env:"MYSQL_maxOpenConnections"`
}

const currentfilepath = "internal/infrastructure/mysql/mysql.go"

func DatabaseInit() *sql.DB {
	var mysqlConfig = MysqlConf{
		Username:           utils.EnvString("MYSQL_USER"),
		Password:           utils.EnvString("MYSQL_PASSWORD"),
		DbName:             utils.EnvString("MYSQL_DBNAME"),
		Host:               utils.EnvString("MYSQL_HOST"),
		Port:               utils.EnvInt("MYSQL_PORT"),
		Schema:             utils.EnvString("MYSQL_schema"),
		LogMode:            utils.EnvBool("MYSQL_logMode"),
		MaxLifetime:        utils.EnvInt("MYSQL_maxLifetime"),
		MinIdleConnections: utils.EnvInt("MYSQL_minIdleConnections"),
		MaxOpenConnections: utils.EnvInt("MYSQL_maxOpenConnections"),
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
