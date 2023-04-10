package main

import (
	"database/sql"
	rest "go-todo-list/internal/delivery/http"
	"go-todo-list/internal/infrastructure/container"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	container.Initcont(".env")
	contConf := container.InitContainer()

	defer contConf.Mysqldb.Close()

	migration(contConf.Mysqldb)

	rest.HTTPRouteInit(contConf)
}

func migration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("err", err)
	}

	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("Running migration")
	if err := m.Up(); err != nil {
		log.Fatal("err migrate up ", err)
	}
	log.Println("Migration Done!!!")
}
