package router

import (
  "fmt"
  "net/http"

  "github.com/labstack/echo-contrib/session"
  "github.com/labstack/echo/v4"
  "golang.org/x/crypto/bcrypt"

  //_ "github.com/go-sql-driver/mysql

  "github.com/flag3/SNS/database"

)

type Me struct {
	Username string `json:"username,omitempty"  db:"username"`
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

func getWhoAmIHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Me{
		Username: c.Get("userID").(string),
	})
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

  err = database.Db.Get(&count, "SELECT COUNT(*) FROM user WHERE UserID=?", req.UserID)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }

  if count > 0 {
    return c.String(http.StatusConflict, "ユーザーが既に存在しています")
  }

  _, err = database.Db.Exec("INSERT INTO user (UserID, Username, HashedPass) VALUES (?, ?, ?)", req.UserID, req.UserID, hashedPass)
  if err != nil {
    return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
  }
  return c.NoContent(http.StatusCreated)
}

func postLoginHandler(c echo.Context) error {
  req := LoginRequestBody{}
  c.Bind(&req)

  user := User{}
  err := database.Db.Get(&user, "SELECT * FROM user WHERE UserID=?", req.UserID)
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

func getLogoutHandler(c echo.Context) error {
  sess, err := session.Get("sessions", c)
  if err != nil {
    fmt.Println(err)
    return c.String(http.StatusInternalServerError, "something wrong in getting session")
  }
  sess.Values["userID"] = nil
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
