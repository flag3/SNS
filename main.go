package main

import (
  "fmt"
  "log"
  //"net/http"
  "os"

  "github.com/labstack/echo-contrib/session"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/srinathgs/mysqlstore"
  //"golang.org/x/crypto/bcrypt"

  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
)

var (
  db *sqlx.DB
)

func main() {
  _db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
  if err != nil {
    log.Fatalf("Cannot Connect to Database: %s", err)
  }
  db = _db

  store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
  if err != nil {
    panic(err)
  }
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(session.Middleware(store))

  e.POST("/login", postLoginHandler)
  e.POST("/signup", postSignUpHandler)

  withLogin := e.Group("")
  withLogin.Use(checkLogin)

  withLogin.GET("/:userID", getTweetHandler)
  withLogin.GET("/:userID/following", getFollowingHandler)
  withLogin.GET("/:userID/followers", getFollowersHandler)
  withLogin.GET("/:userID/likes", getFavoriteHandler)
  withLogin.POST("/tweet", postTweetHandler)
  withLogin.POST("/likes", postFavoriteHandler)
  withLogin.POST("/follow", postFollowHandler)
  withLogin.GET("/whoami", getWhoAmIHandler)
  withLogin.GET("/logout", getLogoutHandler)

  e.Start(":4000")
}
