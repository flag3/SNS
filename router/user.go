package router

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserID      int            `json:"id,omitempty"  db:"UserID"`
	Username    string         `json:"username,omitempty"  db:"Username"`
	DisplayName string         `json:"displayName,omitempty"  db:"DisplayName"`
	Bio         sql.NullString `json:"bio,omitempty"  db:"Bio"`
}

type UserDetail struct {
	UserID      int            `json:"id,omitempty"  db:"UserID"`
	Username    string         `json:"username,omitempty"  db:"Username"`
	DisplayName string         `json:"displayName,omitempty"  db:"DisplayName"`
	Bio         sql.NullString `json:"bio,omitempty"  db:"Bio"`
	Location    sql.NullString `json:"location,omitempty"  db:"Location"`
	Website     sql.NullString `json:"website,omitempty"  db:"Website"`
}

func getUsersHandler(c echo.Context) error {
	users := []User{}

	database.DB.Select(&users, "SELECT UserID, Username, DisplayName, Bio FROM user")
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func getUserHandler(c echo.Context) error {
	username := c.Param("username")

	users := []UserDetail{}

	database.DB.Select(&users, "SELECT UserID, Username, DisplayName, Bio, Location, Website FROM user WHERE Username = ?", username)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}
