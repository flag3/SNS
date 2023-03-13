package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/flag3/SNS/database"

	_ "github.com/go-sql-driver/mysql"
)

type Favorite struct {
	FavoriteID int    `json:"favoriteID,omitempty"  db:"FavoriteID"  form:"favoriteID"`
	TweetID    int    `json:"tweetID,omitempty"  db:"TweetID"  form:"tweetID"`
	UserID     string `json:"userID,omitempty"  db:"UserID"  form:"userID"`
}

func getFavoriteHandler(c echo.Context) error {
	userID := c.Param("userID")

	tweets := []Tweet{}

	database.DB.Select(&tweets, "SELECT tweet.* FROM tweet JOIN favorite ON tweet.TweetID = favorite.TweetID WHERE favorite.UserID = ?", userID)
	if tweets == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, tweets)
}

func postFavoriteHandler(c echo.Context) error {
	userID := c.Get("userID").(string)

	favorite := Favorite{}
	favoriteState := "INSERT INTO favorite (TweetID, UserID) VALUES (?, ?)"

	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadRequest, favorite)
	}

	// ふぁぼしてるかチェック
	var count int

	err := database.DB.Get(&count, "SELECT COUNT(*) FROM favorite WHERE TweetID=? AND UserID=?", favorite.TweetID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "ツイートを既にふぁぼしています")
	}

	database.DB.Exec(favoriteState, favorite.TweetID, userID)
	return c.JSON(http.StatusOK, favorite)
}

func deleteFavoriteHandler(c echo.Context) error {
	tweetID := c.Param("tweetID")
	userID := c.Get("userID").(string)

	favoriteState := "DELETE FROM favorite WHERE TweetID = ? AND UserID = ?"

	database.DB.Exec(favoriteState, tweetID, userID)
	return c.NoContent(http.StatusOK)
}
