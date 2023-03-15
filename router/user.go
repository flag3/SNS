package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"
	_ "github.com/go-sql-driver/mysql"
)

func getUsersHandler(c echo.Context) error {
	accounts := []Account{}

	database.DB.Select(&accounts, "SELECT UserID, Username FROM user")
	if accounts == nil {
		return c.NoContent(http.StatusNotFound)
	}
	fmt.Println(accounts)

	return c.JSON(http.StatusOK, accounts)
}
