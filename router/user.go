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
	IsFollowed  bool           `json:"isFollowed"  db:"IsFollowed"`
	IsFollowing bool           `json:"isFollowing"  db:"IsFollowing"`
}

type UserDetail struct {
	UserID         int            `json:"id,omitempty"  db:"UserID"`
	Username       string         `json:"username,omitempty"  db:"Username"`
	DisplayName    string         `json:"displayName,omitempty"  db:"DisplayName"`
	Bio            sql.NullString `json:"bio,omitempty"  db:"Bio"`
	Location       sql.NullString `json:"location,omitempty"  db:"Location"`
	Website        sql.NullString `json:"website,omitempty"  db:"Website"`
	FollowingCount int            `json:"followingCount"  db:"FollowingCount"`
	FollowerCount  int            `json:"followerCount"  db:"FollowerCount"`
	IsFollowed     bool           `json:"isFollowed"  db:"IsFollowed"`
	IsFollowing    bool           `json:"isFollowing"  db:"IsFollowing"`
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
