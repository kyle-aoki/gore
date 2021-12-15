package db

import (
	"database/sql"
	"fmt"
	"gore/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func OpenDb() {
	db, err := sql.Open("mysql", GetConnString())
	if err != nil {
		log.Fatalf("Error opening DB: %+v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging DB: %+v", err)
	}

	Db = db
}

func GetConnString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Gore.Username,
		config.Gore.Password,
		config.Gore.Host,
		config.Gore.Port,
		config.Gore.Subdatabase,
	)
}
