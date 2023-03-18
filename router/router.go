package router

import (
	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) {
	e.POST("/login", postLoginHandler)
	e.POST("/signup", postSignUpHandler)

	withLogin := e.Group("")
	withLogin.Use(checkLogin)

	withLogin.GET("/users", getUsersHandler)
	withLogin.GET("/users/:username", getUserHandler)

	withLogin.GET("/tweets", getTweetsHandler)
	withLogin.POST("/tweets", postTweetsHandler)
	withLogin.GET("/tweets/:tweetID", getTweetHandler)
	//withLogin.PUT("/tweets/:tweetID", postTweetHandler)
	withLogin.DELETE("/tweets/:tweetID", deleteTweetHandler)
	withLogin.GET("/users/:username/tweets", getUserTweetsHandler)
	withLogin.GET("/home", getHomeTweetHandler)

	withLogin.POST("/tweets/:tweetID/likes", postLikeHandler)
	withLogin.DELETE("/tweets/:tweetID/likes", deleteLikeHandler)
	withLogin.GET("/users/:username/likes", getUserLikesHandler)
	withLogin.GET("/tweets/:tweetID/likes/users", getLikeUsersHandler)

	withLogin.POST("/users/:username/follows", postFollowHandler)
	withLogin.DELETE("/users/:username/follows", deleteFollowHandler)
	withLogin.GET("/users/:username/following", getFollowingHandler)
	withLogin.GET("/users/:username/followers", getFollowersHandler)

	withLogin.POST("/tweets/:tweetID/retweets", postRetweetHandler)
	withLogin.DELETE("/tweets/:tweetID/retweets", deleteRetweetHandler)
	withLogin.GET("/tweets/:tweetID/retweets/users", getRetweetUsersHandler)

	withLogin.GET("/whoami", getWhoAmIHandler)
	withLogin.GET("/logout", getLogoutHandler)

	e.Start(":4000")
}
