package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"

	_ "github.com/go-sql-driver/mysql"
)

type Tweet struct {
	TweetID int    `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID  string `json:"userID,omitempty"  db:"UserID"  form:"userID"`
	Body    string `json:"body,omitempty"  db:"Body"  form:"body"`
}

func getTweetHandler(c echo.Context) error {
	userID := c.Param("userID")

	tweets := []Tweet{}
	database.DB.Select(&tweets, "SELECT * FROM tweet WHERE UserID=?", userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func postTweetHandler(c echo.Context) error {
	userID := c.Get("userID").(string)

	tweet := Tweet{}
	tweetState := "INSERT INTO tweet (UserID, Body) VALUES (?, ?)"

	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, tweet)
	}

	if tweet.Body == "" {
		return c.String(http.StatusBadRequest, "empty string")
	}

	database.DB.Exec(tweetState, userID, tweet.Body)
	return c.JSON(http.StatusOK, tweet)
}

func deleteTweetHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	userID := c.Get("userID").(string)

	tweetState := "DELETE FROM tweet WHERE TweetID = ? AND userID = ?"

	database.DB.Exec(tweetState, tweetID, userID)
	return c.NoContent(http.StatusOK)
}

func getHomeTweetHandler(c echo.Context) error {
	userID := c.Get("userID").(string)

	tweets := []Tweet{}
	database.DB.Select(&tweets, "SELECT tweet.* FROM tweet JOIN follow ON tweet.UserID = follow.FolloweeUserID where follow.FollowerUserID = ? OR UserID = ?", userID, userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)

}
