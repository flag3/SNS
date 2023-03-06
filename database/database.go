package database

import (
  "fmt"
  "log"
  //"net/http"
  "os"

  //"github.com/labstack/echo-contrib/session"
  //"github.com/labstack/echo/v4"
  //"github.com/labstack/echo/v4/middleware"
  "github.com/srinathgs/mysqlstore"
  //"golang.org/x/crypto/bcrypt"

  "github.com/jmoiron/sqlx"
)

var (
  Db *sqlx.DB
  Store *mysqlstore.MySQLStore
)

func Connect() {
    _db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
  if err != nil {
    log.Fatalf("Cannot Connect to Database: %s", err)
  }
  Db = _db

  Store, err = mysqlstore.NewMySQLStoreFromConnection(Db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
  if err != nil {
    panic(err)
  }
}
