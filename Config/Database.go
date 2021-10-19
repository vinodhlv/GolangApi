package Config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var err error

func Connect() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "admin",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "go_api",
		ParseTime: true, /// Parse time values to time.Time
	}

	// Get a database handle.
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic("The given DB format is wrong")

	}

	pingErr := DB.Ping()
	if pingErr != nil {
		panic("The dbconnection is unsucessfull")

	}
	fmt.Println("Connected!")
}

func GetDB() *sql.DB {
	return DB
}
