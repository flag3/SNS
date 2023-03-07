package router

import (
	"github.com/labstack/echo/v4"
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

	withLogin.DELETE("/tweet/:tweetID", deleteTweetHandler)
	withLogin.DELETE("/follow/:followeeUserID", deleteFollowHandler)
	withLogin.DELETE("/like/:tweetID", deleteFavoriteHandler)

	e.Start(":4000")
}
