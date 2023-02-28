package main

import (
  "net/http"

  "github.com/labstack/echo/v4"

  _ "github.com/go-sql-driver/mysql"
)

type Favorite struct {
  FavoriteID int     `json:"favoriteid,omitempty"  db:"FavoriteID"  form:"favoriteid"`
  TweetID    int     `json:"tweetid,omitempty"  db:"TweetID"  form:"tweetid"`
  UserID     string  `json:"userid,omitempty"  db:"UserID"  form:"userid"`
}

func getFavoriteHandler(c echo.Context) error {
  userID := c.Param("userID")

  tweets := []Tweet{}

  db.Select(&tweets, "SELECT tweet.TweetID, tweet.UserID, tweet.Body FROM tweet JOIN favorite ON tweet.TweetID = favorite.TweetID WHERE favorite.UserID = ?", userID)
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

  db.Exec(favoriteState, favorite.TweetID, userID)
  return c.JSON(http.StatusOK, favorite)
}