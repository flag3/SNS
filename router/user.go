package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserID      int            `json:"id,omitempty"  db:"UserID"  form:"userID"`
	Username    string         `json:"username,omitempty"  db:"Username"  form:"username"`
	DisplayName string         `json:"displayName,omitempty"  db:"DisplayName"  form:"displayName"`
	Bio         sql.NullString `json:"bio,omitempty"  db:"Bio"  form:"bio"`
	IsFollowed  bool           `json:"isFollowed"  db:"IsFollowed"  form:"isFollowed"`
	IsFollowing bool           `json:"isFollowing"  db:"IsFollowing" form:"isFollowing"`
}

type UserDetail struct {
	UserID         int            `json:"id,omitempty"  db:"UserID"  form:"userID"`
	Username       string         `json:"username,omitempty"  db:"Username"  form:"username"`
	DisplayName    string         `json:"displayName,omitempty"  db:"DisplayName"  form:"displayName"`
	Bio            sql.NullString `json:"bio,omitempty"  db:"Bio"  form:"bio"`
	Location       sql.NullString `json:"location,omitempty"  db:"Location" form:"location"`
	Website        sql.NullString `json:"website,omitempty"  db:"Website"  form:"website"`
	FollowingCount int            `json:"followingCount"  db:"FollowingCount"  form:"followingCount"`
	FollowerCount  int            `json:"followerCount"  db:"FollowerCount"  form:"followerCount"`
	IsFollowed     bool           `json:"isFollowed"  db:"IsFollowed"  form:"isFollowed"`
	IsFollowing    bool           `json:"isFollowing"  db:"IsFollowing"  form:"isFollowing"`
}

func getUsersHandler(c echo.Context) error {
	username := c.Get("username").(string)
	userID := usernameToUserID(username)

	users := []User{}

	database.DB.Select(&users,
		`SELECT 
      u.UserID, 
      u.Username, 
      u.DisplayName, 
      u.Bio,
      COUNT(DISTINCT CASE WHEN f1.FolloweeUserID = ? THEN f1.FolloweeUserID END) AS IsFollowed,
      COUNT(DISTINCT CASE WHEN f2.FollowerUserID = ? THEN f2.FollowerUserID END) AS IsFollowing
    FROM 
      user u
      LEFT JOIN follow f1 ON u.UserID = f1.FollowerUserID
      LEFT JOIN follow f2 ON u.UserID = f2.FolloweeUserID
    GROUP BY 
      u.UserID`,
		userID, userID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func getUserHandler(c echo.Context) error {
	loginUsername := c.Get("username").(string)
	paramUsername := c.Param("username")
	loginUserID := usernameToUserID(loginUsername)

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
      COUNT(DISTINCT f2.FollowID) as FollowerCount, 
      COUNT(DISTINCT CASE WHEN f1.FolloweeUserID = ? THEN f1.FolloweeUserID END) AS IsFollowed,
      COUNT(DISTINCT CASE WHEN f2.FollowerUserID = ? THEN f2.FollowerUserID END) AS IsFollowing
    FROM 
      user u
      LEFT JOIN follow f1 ON u.UserID = f1.FollowerUserID
      LEFT JOIN follow f2 ON u.UserID = f2.FolloweeUserID
    WHERE 
      u.Username = ?
    GROUP BY 
      u.UserID`,
		loginUserID, loginUserID, paramUsername)

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
