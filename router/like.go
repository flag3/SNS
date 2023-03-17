package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"

	_ "github.com/go-sql-driver/mysql"
)

type Like struct {
	LikeID  int `json:"likeID,omitempty"  db:"LikeID"  form:"likeID"`
	TweetID int `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID  int `json:"userID,omitempty"  db:"UserID"  form:"userID"`
}

func postLikeHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	username := c.Get("username").(string)
	userID := usernameToUserID(username)

	like := Like{}
	likeState := "INSERT INTO fav (TweetID, UserID) VALUES (?, ?)"

	// ふぁぼしてるかチェック
	var count int

	err := database.DB.Get(&count, "SELECT COUNT(*) FROM fav WHERE TweetID=? AND UserID=?", tweetID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "ツイートを既にふぁぼしています")
	}

	database.DB.Exec(likeState, tweetID, userID)
	return c.JSON(http.StatusCreated, like)
}

func deleteLikeHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	username := c.Get("username").(string)
	userID := usernameToUserID(username)

	likeState := "DELETE FROM fav WHERE TweetID = ? AND UserID = ?"
	database.DB.Exec(likeState, tweetID, userID)
	return c.NoContent(http.StatusOK)
}

func getUserLikesHandler(c echo.Context) error {
	username := c.Param("username")
	userID := usernameToUserID(username)

	tweets := []Tweet{}
	database.DB.Select(&tweets, "SELECT tweet.* FROM tweet JOIN fav ON tweet.TweetID = fav.TweetID WHERE fav.UserID = ?", userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func getLikeUsersHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")

	users := []User{}
	database.DB.Select(&users, "SELECT user.UserID, user.Username, user.DisplayName, user.Bio FROM user JOIN fav ON user.UserID = fav.UserID WHERE fav.TweetID = ?", tweetID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}
