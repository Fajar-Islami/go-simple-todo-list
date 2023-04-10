package container

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"go-todo-list/internal/helper"
	"go-todo-list/internal/utils"

	mysqlclient "go-todo-list/internal/infrastructure/mysql"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"golang.org/x/sync/errgroup"
)

const currentfilepath = "internal/infrastructure/container/container.go"

type (
	Container struct {
		Mysqldb *sql.DB
		Apps    *Apps
		Logger  *Logger
	}

	Logger struct {
		Log     zerolog.Logger
		Path    string `env:"log_path"`
		LogFile string `env:"log_file"`
	}

	Apps struct {
		Name           string `env:"apps_appName"`
		Host           string `env:"apps_host"`
		Version        string `env:"apps_version"`
		SwaggerAddress string `env:"apps_swagger_address"`
		HttpPort       int    `env:"apps_httpport"`
		SecretJwt      string `env:"apps_secretJwt"`
		CtxTimeout     int    `env:"apps_timeout"`
	}
)

func Initcont(filename string) {
	err := godotenv.Load(fmt.Sprintf("%s/%s", helper.ProjectRootPath, filename))
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, "", fmt.Errorf("error when loadenv : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read environment variable", nil)
}

func AppsInit() Apps {
	var appsConf = Apps{
		Name:           utils.EnvString("apps_appName"),
		Host:           utils.EnvString("apps_host"),
		Version:        utils.EnvString("apps_version"),
		SwaggerAddress: utils.EnvString("apps_swagger_address"),
		HttpPort:       utils.EnvInt("apps_httpport"),
		SecretJwt:      utils.EnvString("apps_secretJwt"),
		CtxTimeout:     utils.EnvInt("apps_timeout"),
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read appsConf", nil)
	return appsConf
}

func LoggerInit() Logger {
	var loggerConf = Logger{
		Path:    utils.EnvString("log_path"),
		LogFile: utils.EnvString("log_file"),
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when read loggerConf", nil)

	var stdout io.Writer = os.Stdout
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if loggerConf.LogFile == "ON" {
		path := fmt.Sprintf("%s%s", helper.ProjectRootPath, loggerConf.Path)
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664)
		if err != nil {
			helper.Logger(currentfilepath, helper.LoggerLevelError, "", fmt.Errorf("error when setting loggerConf : %s", err.Error()))
		}
		// Create a multi writer with both the console and file writers
		stdout = zerolog.MultiLevelWriter(os.Stdout, file)

	}

	loggerConf.Log = zerolog.New(stdout).With().Caller().Timestamp().Logger()
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read loggerConf", nil)
	return loggerConf
}

// containters = apps,mysql,logger,redis
func InitContainer(containters ...string) *Container {
	newStrContainer := strings.Join(containters, ",")
	var cont Container
	errGroup, _ := errgroup.WithContext(context.Background())

	errGroup.Go(func() (err error) {
		if strings.Contains(newStrContainer, "apps") || len(containters) == 0 {
			apps := AppsInit()
			cont.Apps = &apps
			return
		}
		return nil
	})

	errGroup.Go(func() (err error) {
		if strings.Contains(newStrContainer, "mysql") || len(containters) == 0 {
			mysqldb := mysqlclient.DatabaseInit()
			cont.Mysqldb = mysqldb
			return
		}
		return nil
	})

	errGroup.Go(func() (err error) {
		if strings.Contains(newStrContainer, "log") || len(containters) == 0 {
			logger := LoggerInit()
			cont.Logger = &logger
			return
		}
		return nil
	})

	_ = errGroup.Wait()

	return &cont
}
