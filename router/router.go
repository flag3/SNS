package router

import (
  //"encoding/gob"
	//"os"

	//"github.com/google/uuid"
	//"github.com/gorilla/sessions"
	//"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	//"go.uber.org/zap"
)

func NewServer(e *echo.Echo) {
  e.POST("/login", postLoginHandler)
  e.POST("/signup", postSignUpHandler)

  withLogin := e.Group("")
  withLogin.Use(checkLogin)

  withLogin.GET("/:userID", getTweetHandler)
  withLogin.GET("/:userID/following", getFollowingHandler)
  withLogin.GET("/:userID/followers", getFollowersHandler)
  withLogin.GET("/:userID/likes", getFavoriteHandler)
  withLogin.GET("/home", getHomeTweetHandler)
  withLogin.POST("/tweet", postTweetHandler)
  withLogin.POST("/like", postFavoriteHandler)
  withLogin.POST("/follow", postFollowHandler)
  withLogin.GET("/whoami", getWhoAmIHandler)
  withLogin.GET("/logout", getLogoutHandler)

  withLogin.DELETE("/delete/tweet/:tweetID", deleteTweetHandler)
  withLogin.DELETE("/delete/follow/:followeeUserID", deleteFollowHandler)
  withLogin.DELETE("/delete/like/:tweetID", deleteFavoriteHandler)

  e.Start(":4000")
}
