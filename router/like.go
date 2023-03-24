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
	loginUsername := c.Get("username").(string)
	loginUserID := usernameToUserID(loginUsername)
	paramUsername := c.Param("username")
	paramUserID := usernameToUserID(paramUsername)

	tweets := []TweetDetail{}
	database.DB.Select(&tweets,
		`SELECT 
      t.TweetID, 
      t.UserID, 
      u.Username, 
      u.DisplayName, 
      t.Content, 
      t.Reply, 
      t.Quote, 
      (SELECT COUNT(*) FROM tweet WHERE Reply = t.TweetID) AS ReplyCount,
      COUNT(DISTINCT r.UserID) AS RetweetCount, 
      (SELECT COUNT(*) FROM tweet WHERE Quote = t.TweetID) AS QuoteCount,
      COUNT(DISTINCT l.UserID) AS LikeCount, 
      COUNT(DISTINCT CASE WHEN r.UserID = ? THEN r.UserID END) AS IsRetweeted,
      COUNT(DISTINCT CASE WHEN l.UserID = ? THEN l.UserID END) AS IsLiked
		FROM 
      tweet t
		  JOIN user u ON t.UserID = u.UserID
		  LEFT JOIN retweet r ON t.TweetID = r.tweetID
		  LEFT JOIN fav l ON t.TweetID = l.TweetID
    WHERE 
      EXISTS(SELECT 1 FROM fav WHERE TweetID = t.TweetID AND UserID = ?)
		GROUP BY 
      t.TweetID`,
		loginUserID, loginUserID, paramUserID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func getLikesHandler(c echo.Context) error {
	username := c.Get("username").(string)
	userID := usernameToUserID(username)
	tweetID := c.Param("tweetID")

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
      JOIN fav l ON u.UserID = l.UserID
      LEFT JOIN follow f1 ON u.UserID = f1.FollowerUserID
      LEFT JOIN follow f2 ON u.UserID = f2.FolloweeUserID
    WHERE 
      l.TweetID = ?
    GROUP BY 
      u.UserID`,
		userID, userID, tweetID)
	if users == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}
