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
    WHERE f2.FollowerUserID = ?
    GROUP BY 
      u.UserID`,
		userID, userID, userID)
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
    WHERE 
      f1.FolloweeUserID = ?
    GROUP BY 
      u.UserID`,
		userID, userID, userID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}
	fmt.Println(users)

	return c.JSON(http.StatusOK, users)
}
