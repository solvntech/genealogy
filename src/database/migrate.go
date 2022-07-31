package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

func Migration(DB *sql.DB) error {
	dbName := os.Getenv("MYSQL_DATABASE")
	driver, driverError := mysql.WithInstance(DB, &mysql.Config{})

	if driverError != nil {
		fmt.Println(driverError)
		return driverError
	}

	m, mError := migrate.NewWithDatabaseInstance(
		"file://./src/database/migrations",
		dbName,
		driver,
	)

	if mError != nil {
		fmt.Println(mError)
		return mError
	}

	if downError := m.Down(); downError != nil {
		fmt.Println(downError)
	}

	if upError := m.Up(); upError != nil {
		fmt.Println(upError)
		return upError
	}

	return nil
}
