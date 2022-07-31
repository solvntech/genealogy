package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	once       sync.Once
	MySqlDB    *sql.DB
	DBInstance *gorm.DB
)

func ConnectDB() (*sql.DB, *gorm.DB) {
	once.Do(func() {
		dbUser := os.Getenv("MYSQL_USER")
		dbPass := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST")
		dbPort := os.Getenv("MYSQL_PORT")
		dbName := os.Getenv("MYSQL_DATABASE")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPass, dbHost, dbPort, dbName)

		sqlDB, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("DB connect successful")
		}

		gormDB, gormDBErr := gorm.Open(
			mysql.New(mysql.Config{
				Conn: sqlDB,
			}),
			&gorm.Config{},
		)
		if gormDBErr != nil {
			log.Fatal(gormDBErr)
		} else {
			fmt.Println("gormDB connect successful")
		}
		DBInstance = gormDB
		MySqlDB = sqlDB
	})
	return MySqlDB, DBInstance
}
