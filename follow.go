package main

import (
  "fmt"
  "net/http"

  "github.com/labstack/echo/v4"

  _ "github.com/go-sql-driver/mysql"
)

type Follow struct {
  FollowID        int     `json:"followid,omitempty"  db:"FollowID"  form:"followid"`
  FollowerUserID  string  `json:"followeruserid,omitempty"  db:"FollowerUserID"  form:"followeruserid"`
  FolloweeUserID  string  `json:"followeeuserid,omitempty"  db:"FolloweeUserID"  form:"followeeuserid"`
}

type Account struct{
  UserID     string `json:"userid,omitempty"  db:"UserID"`
  Username   string `json:"username,omitempty"  db:"Username"`
}

func getFollowingHandler(c echo.Context) error {
  userID := c.Param("userID")

  accounts := []Account{}

  db.Select(&accounts, "SELECT UserID, Username FROM user JOIN follow ON UserID = FolloweeUserID WHERE FollowerUserID=?", userID)
  if accounts == nil{
    return c.NoContent(http.StatusNotFound)
  }
  fmt.Println(accounts)

  return c.JSON(http.StatusOK, accounts)
}

func getFollowersHandler(c echo.Context) error {
  userID := c.Param("userID")

  accounts := []Account{}

  db.Select(&accounts, "SELECT UserID, Username FROM user JOIN follow ON UserID = FollowerUserID WHERE FolloweeUserID=?", userID)
  if accounts == nil{
    return c.NoContent(http.StatusNotFound)
  }
  fmt.Println(accounts)

  return c.JSON(http.StatusOK, accounts)
}

func postFollowHandler(c echo.Context) error {
  userID := c.Get("userID").(string)

  follow := Follow{}
  followState := "INSERT INTO follow (FollowerUserID, FolloweeUserID) VALUES (?, ?)"

  if err := c.Bind(&follow); err != nil {
    return c.JSON(http.StatusBadRequest, follow)
  }

  // ユーザーの存在チェック
  var count int

  err := db.Get(&count, "SELECT COUNT(*) FROM follow WHERE FollowerUserID=? and FolloweeUserID=?", userID, follow.FolloweeUserID)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }

  if count > 0 {
    return c.String(http.StatusConflict, "ユーザーを既にフォローしています")
  }

  db.Exec(followState, userID, follow.FolloweeUserID)
  return c.JSON(http.StatusOK, follow)
}
