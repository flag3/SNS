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
	TweetID int           `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID  int           `json:"userID,omitempty"  db:"UserID"  form:"userID"`
	Content string        `json:"content,omitempty"  db:"Content"  form:"content"`
	Reply   sql.NullInt64 `json:"reply,omitempty"  db:"Reply"  form:"reply"`
	Quote   sql.NullInt64 `json:"quote,omitempty"  db:"Quote"  form:"quote"`
}

type TweetDetail struct {
	TweetID      int           `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID       int           `json:"userID,omitempty"  db:"UserID"  form:"userID"`
	Username     string        `json:"username,omitempty"  db:"Username"  form:"username"`
	DisplayName  string        `json:"displayName,omitempty"  db:"DisplayName"  form:"displayName"`
	Content      string        `json:"content,omitempty"  db:"Content"  form:"content"`
	Reply        sql.NullInt64 `json:"reply,omitempty"  db:"Reply"  form:"reply"`
	Quote        sql.NullInt64 `json:"quote,omitempty"  db:"Quote"  form:"quote"`
	ReplyCount   int           `json:"replyCount"  db:"ReplyCount"  form:"replyCount"`
	RetweetCount int           `json:"retweetCount"  db:"RetweetCount"  form:"retweetCount"`
	QuoteCount   int           `json:"quoteCount"  db:"QuoteCount"  form:"quoteCount"`
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
	tweets := []TweetDetail{}

	database.DB.Select(&tweets,
		`SELECT t.TweetID, t.UserID, u.Username, u.DisplayName, t.Content, t.Reply, t.Quote, COUNT(DISTINCT t.Reply) as ReplyCount, COUNT(DISTINCT r.UserID) as RetweetCount, COUNT(DISTINCT t.Quote) as QuoteCount, COUNT(DISTINCT fa.UserID) as LikeCount
		FROM tweet t
		JOIN user u ON t.UserID = u.UserID
		LEFT JOIN retweet r ON t.TweetID = r.tweetID
		LEFT JOIN fav fa ON t.TweetID = fa.TweetID
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

	tweets := []TweetDetail{}

	database.DB.Select(&tweets,
		`SELECT t.TweetID, t.UserID, u.Username, u.DisplayName, t.Content, t.Reply, t.Quote, COUNT(DISTINCT t.Reply) as ReplyCount, COUNT(DISTINCT r.UserID) as RetweetCount, COUNT(DISTINCT t.Quote) as QuoteCount, COUNT(DISTINCT fa.UserID) as LikeCount
		FROM tweet t
		JOIN user u ON t.UserID = u.UserID
		LEFT JOIN fav fa ON t.TweetID = fa.TweetID
		LEFT JOIN retweet r ON t.TweetID = r.tweetID 
    WHERE t.TweetID = ? 
		GROUP BY t.TweetID`,
		tweetID)
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

	tweets := []TweetDetail{}

	userID := usernameToUserID(username)
	database.DB.Select(&tweets,
		`SELECT t.TweetID, t.UserID, u.Username, u.DisplayName, t.Content, t.Reply, t.Quote, COUNT(DISTINCT t.Reply) as ReplyCount, COUNT(DISTINCT r.UserID) as RetweetCount, COUNT(DISTINCT t.Quote) as QuoteCount, COUNT(DISTINCT fa.UserID) as LikeCount
		FROM tweet t
		JOIN user u ON t.UserID = u.UserID
		LEFT JOIN fav fa ON t.TweetID = fa.TweetID
		LEFT JOIN retweet r ON t.TweetID = r.tweetID 
    WHERE t.userID = ? 
		GROUP BY t.TweetID`,
		userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func getHomeTweetHandler(c echo.Context) error {
	username := c.Get("username").(string)

	tweets := []TweetDetail{}

	userID := usernameToUserID(username)
	database.DB.Select(&tweets,
		`SELECT t.TweetID, t.UserID, u.Username, u.DisplayName, t.Content, t.Reply, t.Quote, COUNT(DISTINCT t.Reply) as ReplyCount, COUNT(DISTINCT r.UserID) as RetweetCount, COUNT(DISTINCT t.Quote) as QuoteCount, COUNT(DISTINCT fa.UserID) as LikeCount
    FROM tweet t
    JOIN user u ON t.UserID = u.UserID
		LEFT JOIN fav fa ON t.TweetID = fa.TweetID
		LEFT JOIN retweet r ON t.TweetID = r.tweetID 
    LEFT JOIN follow fo ON t.UserID = fo.FolloweeUserID AND fo.FollowerUserID = ? 
    WHERE t.UserID = ? OR fo.FolloweeUserID IS NOT NULL
    GROUP BY t.TweetID`,
		userID, userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}
