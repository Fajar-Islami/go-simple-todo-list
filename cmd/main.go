package main

import (
	"database/sql"
	"fmt"
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

	// fmt.Println("contConf.Mysqldb", fmt.Sprintf("%#v \n\n", contConf.Mysqldb.Stats()))
	// mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName
	migration(contConf.Mysqldb)

	rest.HTTPRouteInit(contConf)
}

func migration(db *sql.DB) {
	fmt.Println("db", fmt.Sprintf("%#v \n\n", db))
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println("err", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("Running migration")
	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatal("err migrate up ", err)
		}
	}
	log.Println("Migration Done!!!")
}
