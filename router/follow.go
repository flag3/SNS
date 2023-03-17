package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"
	_ "github.com/go-sql-driver/mysql"
)

type Follow struct {
	FollowID       int `json:"followID,omitempty"  db:"FollowID"  form:"followID"`
	FollowerUserID int `json:"followerUserID,omitempty"  db:"FollowerUserID"  form:"followerUserID"`
	FolloweeUserID int `json:"followeeUserID,omitempty"  db:"FolloweeUserID"  form:"followeeUserID"`
}

type Account struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	DisplayName string `json:"displayName,omitempty"  db:"DisplayName"`
}

func postFollowHandler(c echo.Context) error {
	followerUsername := c.Get("username").(string)
	followeeUsername := c.Param("username")

	follow := Follow{}
	followState := "INSERT INTO follow (FollowerUserID, FolloweeUserID) VALUES (?, ?)"

	var followerUserID, followeeUserID int
	if err := database.DB.QueryRow("SELECT UserID FROM user WHERE Username = ?", followerUsername).Scan(&followerUserID); err != nil {
		return c.JSON(http.StatusBadRequest, followerUserID)
	}
	if err := database.DB.QueryRow("SELECT UserID FROM user WHERE Username = ?", followeeUsername).Scan(&followeeUserID); err != nil {
		return c.JSON(http.StatusBadRequest, followeeUserID)
	}

	// フォローしてるかチェック
	var count int
	err := database.DB.Get(&count, "SELECT COUNT(*) FROM follow WHERE FollowerUserID=? AND FolloweeUserID=?", followerUserID, followeeUserID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "ユーザーを既にフォローしています")
	}

	database.DB.Exec(followState, followerUserID, followeeUserID)
	return c.JSON(http.StatusOK, follow)
}

func deleteFollowHandler(c echo.Context) error {
	followerUsername := c.Get("username").(string)
	followeeUsername := c.Param("username")

	var followerUserID, followeeUserID int
	if err := database.DB.QueryRow("SELECT UserID FROM user WHERE Username = ?", followerUsername).Scan(&followerUserID); err != nil {
		return c.JSON(http.StatusBadRequest, followerUserID)
	}
	if err := database.DB.QueryRow("SELECT UserID FROM user WHERE Username = ?", followeeUsername).Scan(&followeeUserID); err != nil {
		return c.JSON(http.StatusBadRequest, followeeUserID)
	}

	followState := "DELETE FROM follow WHERE FollowerUserID = ? AND FolloweeUserID = ?"

	database.DB.Exec(followState, followerUserID, followeeUserID)
	return c.NoContent(http.StatusOK)
}

func getFollowingHandler(c echo.Context) error {
	username := c.Param("username")
	userID := usernameToUserID(username)

	users := []User{}

	database.DB.Select(&users, "SELECT UserID, Username, DisplayName, Bio FROM user JOIN follow ON UserID = FolloweeUserID WHERE FollowerUserID=?", userID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}
	fmt.Println(users)

	return c.JSON(http.StatusOK, users)
}

func getFollowersHandler(c echo.Context) error {
	username := c.Param("username")
	userID := usernameToUserID(username)

	users := []User{}

	database.DB.Select(&users, "SELECT UserID, Username FROM user JOIN follow ON UserID = FollowerUserID WHERE FolloweeUserID=?", userID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}
	fmt.Println(users)

	return c.JSON(http.StatusOK, users)
}
