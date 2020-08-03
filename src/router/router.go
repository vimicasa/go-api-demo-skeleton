package router

import (
	r "bb-server/src/handler"
	"bb-server/src/handler/login"
	"bb-server/src/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load() *gin.Engine {

	// Create the Gin engine.
	g := gin.New()

	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging())
	g.Use(middleware.RequestID())

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		r.NotFoundResponse(c)
	})

	// api for authentication functionalities
	g.POST("/login", login.Login)

	return g
}
