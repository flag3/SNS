package main

import (
  //"database/sql"
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/labstack/echo-contrib/session"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/srinathgs/mysqlstore"
  "golang.org/x/crypto/bcrypt"

  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"

)

var (
  db *sqlx.DB
)

type Tweet struct {
  TweetID int     `json:"id,omitempty"  db:"TweetID"  form:"tweetid"`
  UserID  string  `json:"userid,omitempty"  db:"UserID"  form:"userid"`
  Body    string  `json:"body,omitempty"  db:"Body"  form:"body"`
}

func main() {
  _db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
  if err != nil {
    log.Fatalf("Cannot Connect to Database: %s", err)
  }
  db = _db

  store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
  if err != nil {
    panic(err)
  }
  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(session.Middleware(store))

  e.POST("/login", postLoginHandler)
  e.POST("/signup", postSignUpHandler)

  withLogin := e.Group("")
  withLogin.Use(checkLogin)

  withLogin.GET("/:userID", getTweetHandler)
  withLogin.POST("/:userID/post", postTweetHandler)

  e.Start(":4000")
}

type LoginRequestBody struct {
  UserID    string `json:"userid,omitempty" form:"userid"`
  Password  string `json:"password,omitempty" form:"password"`
}

type User struct {
  UserID     string `json:"userid,omitempty"  db:"UserID"`
  Username   string `json:"username,omitempty"  db:"Username"`
  HashedPass string `json:"-"  db:"HashedPass"`
}

func postSignUpHandler(c echo.Context) error {
  req := LoginRequestBody{}
  c.Bind(&req)

  // もう少し真面目にバリデーションするべき
  if req.Password == "" || req.UserID == "" {
    // エラーは真面目に返すべき
    return c.String(http.StatusBadRequest, "項目が空です")
  }

  hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("bcrypt generate error: %v", err))
  }

  // ユーザーの存在チェック
  var count int

  err = db.Get(&count, "SELECT COUNT(*) FROM user WHERE UserID=?", req.UserID)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }

  if count > 0 {
    return c.String(http.StatusConflict, "ユーザーが既に存在しています")
  }

  _, err = db.Exec("INSERT INTO user (UserID, Username, HashedPass) VALUES (?, ?, ?)", req.UserID, req.UserID, hashedPass)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }
  return c.NoContent(http.StatusCreated)
}

func postLoginHandler(c echo.Context) error {
  req := LoginRequestBody{}
  c.Bind(&req)

  user := User{}
  err := db.Get(&user, "SELECT * FROM user WHERE UserID=?", req.UserID)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(req.Password))
  if err != nil {
    if err == bcrypt.ErrMismatchedHashAndPassword {
      return c.NoContent(http.StatusForbidden)
    } else {
      return c.NoContent(http.StatusInternalServerError)
    }
  }

  sess, err := session.Get("sessions", c)
  if err != nil {
    fmt.Println(err)
    return c.String(http.StatusInternalServerError, "something wrong in getting session")
  }
  sess.Values["userID"] = req.UserID
  sess.Save(c.Request(), c.Response())

  return c.NoContent(http.StatusOK)
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    sess, err := session.Get("sessions", c)
    if err != nil {
      fmt.Println(err)
      return c.String(http.StatusInternalServerError, "something wrong in getting session")
    }

    if sess.Values["userID"] == nil {
      return c.String(http.StatusForbidden, "please login")
    }
    c.Set("userID", sess.Values["userID"].(string))

    return next(c)
  }
}

func getTweetHandler(c echo.Context) error {
  userID := c.Param("userID")

  tweets := []Tweet{}
  db.Select(&tweets, "SELECT * FROM tweet WHERE UserID=?", userID)
  if tweets == nil {
    return c.NoContent(http.StatusNotFound)
  }

  return c.JSON(http.StatusOK, tweets)
}


func postTweetHandler(c echo.Context) error {
  userID := c.Param("userID")

  tweet := Tweet{}
  tweetState := "INSERT INTO tweet (UserID, Body) VALUES (?, ?)"

  if err := c.Bind(&tweet); err != nil {
    return c.JSON(http.StatusBadRequest, tweet)
  }

  if tweet.Body == "" {
    return c.String(http.StatusBadRequest, "empty string")
  }

  db.Exec(tweetState, userID, tweet.Body)
  return c.JSON(http.StatusOK, tweet)
}
