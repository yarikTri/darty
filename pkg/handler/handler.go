package handler

import "github.com/gin-gonic/gin"

// Handler ..
type Handler struct {
}

// InitRoutes ..
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/")

	events := router.Group("/events")
	{
		events.GET("/:id")
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	user := router.Group("/user")
	{
		user.GET("/:id")
	}

	return router
}
