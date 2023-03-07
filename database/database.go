package database

import (
	"fmt"
	"log"

	"os"

	"github.com/srinathgs/mysqlstore"

	"github.com/jmoiron/sqlx"
)

var (
	DB    *sqlx.DB
	Store *mysqlstore.MySQLStore
)

func Connect() {
	_DB, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	DB = _DB

	Store, err = mysqlstore.NewMySQLStoreFromConnection(DB.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
}
