package main

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jmoiron/sqlx"

	"github.com/flag3/SNS/database"
	"github.com/flag3/SNS/router"
)

var (
	db *sqlx.DB
)

func main() {
	fmt.Println("string")
	database.Connect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(database.Store))

	router.NewServer(e)
}
