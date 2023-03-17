package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"

	_ "github.com/go-sql-driver/mysql"
)

type Tweet struct {
	TweetID      int           `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID       int           `json:"userID,omitempty"  db:"UserID"  form:"userID"`
	Username     string        `json:"username,omitempty"  db:"Username"  form:"username"`
	DisplayName  string        `json:"displayName,omitempty"  db:"DisplayName"  form:"displayName"`
	Content      string        `json:"content,omitempty"  db:"Content"  form:"content"`
	Reply        sql.NullInt64 `json:"reply,omitempty"  db:"Reply"  form:"reply"`
	Quote        sql.NullInt64 `json:"quote,omitempty"  db:"Quote"  form:"quote"`
	ReplyCount   int           `json:"replyCount"  db:"ReplyCount"  form:"replyCount"`
	RetweetCount int           `json:"retweetCount"  db:"RetweetCount"  form:"retweetCount"`
	LikeCount    int           `json:"likeCount"  db:"LikeCount"  form:"likeCount"`
}

func usernameToUserID(username string) int {
	var userID int

	if err := database.DB.QueryRow("SELECT UserID FROM user WHERE Username = ?", username).Scan(&userID); err != nil {
		log.Fatal(err)
	}
	return userID
}

func getTweetsHandler(c echo.Context) error {
	tweets := []Tweet{}

	database.DB.Select(&tweets,
		`SELECT t.TweetID, t.UserID, u.Username, u.DisplayName, t.Content, t.Reply, t.Quote, COUNT(t.Reply) as ReplyCount, COUNT(r.TweetID) as RetweetCount, COUNT(f.TweetID) as LikeCount
    FROM tweet t
    JOIN user u ON t.UserID = u.UserID
    LEFT JOIN fav f ON t.TweetID = f.TweetID
    LEFT JOIN retweet r ON t.TweetID = r.tweetID
    GROUP BY t.TweetID`,
	)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func postTweetsHandler(c echo.Context) error {
	username := c.Get("username").(string)

	tweet := Tweet{}
	tweetState := "INSERT INTO tweet (UserID, Content) VALUES (?, ?)"

	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, tweet)
	}

	if tweet.Content == "" {
		return c.String(http.StatusBadRequest, "empty string")
	}

	userID := usernameToUserID(username)
	database.DB.Exec(tweetState, userID, tweet.Content)
	return c.JSON(http.StatusCreated, tweet)
}

func getTweetHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")

	tweets := []Tweet{}

	database.DB.Select(&tweets, "SELECT * FROM tweet WHERE TweetID = ?", tweetID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func deleteTweetHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	username := c.Get("username").(string)

	tweetState := "DELETE FROM tweet WHERE TweetID = ? AND userID = ?"

	userID := usernameToUserID(username)
	database.DB.Exec(tweetState, tweetID, userID)
	return c.NoContent(http.StatusOK)
}

func getUserTweetsHandler(c echo.Context) error {
	username := c.Param("username")

	tweets := []Tweet{}

	userID := usernameToUserID(username)
	database.DB.Select(&tweets, "SELECT * FROM tweet WHERE UserID=?", userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func getHomeTweetHandler(c echo.Context) error {
	username := c.Get("username").(string)

	tweets := []Tweet{}

	userID := usernameToUserID(username)
	database.DB.Select(&tweets,
		"SELECT tweet.TweetID, tweet.UserID, tweet.Content FROM tweet LEFT JOIN follow ON tweet.UserID = follow.FolloweeUserID AND follow.FollowerUserID = ? WHERE tweet.UserID = ? OR follow.FolloweeUserID IS NOT NULL",
		userID, userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}
