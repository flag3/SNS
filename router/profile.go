package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"
	_ "github.com/go-sql-driver/mysql"
)

func getWhoAmIHandler(c echo.Context) error {
	username := c.Get("username").(string)
	users := []UserDetail{}

	database.DB.Select(&users,
		`SELECT 
      u.UserID, 
      u.Username, 
      u.DisplayName, 
      u.Bio,
      u.Location,
      u.Website,
      COUNT(DISTINCT f1.FollowID) as FollowingCount, 
      COUNT(DISTINCT f2.FollowID) as FollowerCount
    FROM 
      user u
      LEFT JOIN follow f1 ON u.UserID = f1.FollowerUserID
      LEFT JOIN follow f2 ON u.UserID = f2.FolloweeUserID
    WHERE 
      u.Username = ?
    GROUP BY 
      u.UserID`,
		username,
	)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func putUserDisplayNameHandler(c echo.Context) error {
	username := c.Get("username").(string)

	user := UserDetail{}
	userState := "UPDATE user SET DisplayName = ? WHERE Username = ?"

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, user)
	}

	fmt.Println(user)

	if user.DisplayName == "" {
		return c.String(http.StatusBadRequest, "empty string")
	}

	database.DB.Exec(userState, user.DisplayName, username)
	return c.JSON(http.StatusOK, user)
}

func putUserBioHandler(c echo.Context) error {
	username := c.Get("username").(string)

	user := UserDetail{}
	userState := "UPDATE user SET Bio = ? WHERE Username = ?"

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, user)
	}

	fmt.Println(user)

	database.DB.Exec(userState, user.Bio.String, username)
	return c.JSON(http.StatusOK, user)
}

func putUserLocationHandler(c echo.Context) error {
	username := c.Get("username").(string)

	user := UserDetail{}
	userState := "UPDATE user SET Location = ? WHERE Username = ?"

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, user)
	}

	fmt.Println(user)

	database.DB.Exec(userState, user.Location.String, username)
	return c.JSON(http.StatusOK, user)
}

func putUserWebsiteHandler(c echo.Context) error {
	username := c.Get("username").(string)

	user := UserDetail{}
	userState := "UPDATE user SET Website = ? WHERE Username = ?"

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, user)
	}

	fmt.Println(user)

	database.DB.Exec(userState, user.Website.String, username)
	return c.JSON(http.StatusOK, user)
}
