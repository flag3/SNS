package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"

	_ "github.com/go-sql-driver/mysql"
)

type Retweet struct {
	RetweetID int `json:"retweetID,omitempty"  db:"RetweetID"  form:"retweetID"`
	TweetID   int `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID    int `json:"userID,omitempty"  db:"UserID"  form:"userID"`
}

func postRetweetHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	username := c.Get("username").(string)
	userID := usernameToUserID(username)

	retweet := Retweet{}
	retweetState := "INSERT INTO retweet (TweetID, UserID) VALUES (?, ?)"

	// ふぁぼしてるかチェック
	var count int

	err := database.DB.Get(&count, "SELECT COUNT(*) FROM retweet WHERE TweetID=? AND UserID=?", tweetID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "ツイートを既にリツイートしています")
	}

	database.DB.Exec(retweetState, tweetID, userID)
	return c.JSON(http.StatusCreated, retweet)
}

func deleteRetweetHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	username := c.Get("username").(string)
	userID := usernameToUserID(username)

	retweetState := "DELETE FROM retweet WHERE TweetID = ? AND UserID = ?"
	database.DB.Exec(retweetState, tweetID, userID)
	return c.NoContent(http.StatusOK)
}

func getRetweetUsersHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")

	users := []User{}
	database.DB.Select(&users, "SELECT user.UserID, user.Username, user.DisplayName, user.Bio FROM user JOIN retweet ON user.UserID = retweet.UserID WHERE retweet.TweetID = ?", tweetID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}
