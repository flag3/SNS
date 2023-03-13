package database

import (
	"log"
	"time"

	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/srinathgs/mysqlstore"

	"github.com/jmoiron/sqlx"
)

var (
	DB    *sqlx.DB
	Store *mysqlstore.MySQLStore
)

func Connect() {
	jst, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	c := mysql.Config{
		DBName:    os.Getenv("DB_DATABASE"),
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_HOSTNAME") + ":" + os.Getenv("DB_PORT"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	_DB, err := sqlx.Connect("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	DB = _DB

	Store, err = mysqlstore.NewMySQLStoreFromConnection(DB.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
}
