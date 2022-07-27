package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yarikTri/darty/pkg/service"
)

// Handler is the wrap for service.Service struct
type Handler struct {
	services *service.Service
}

// NewHandler inits new Handler
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes ..
// In the development process, the amount of routes will increase
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// No handler func yet for
	// router.GET("/")

	events := router.Group("/events")
	{
		events.GET("/", h.getAllEvents)
		events.POST("/", h.createEvent)

		events.GET("/:id", h.getEventByID)
		events.PUT("/:id", h.updateEvent)
		events.DELETE("/:id", h.deleteEvent)
		// ...
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		// ...
	}

	// No handler func yet for
	// user := router.Group("/user")
	// {
	// 	user.GET("/:id")
	// 	// ...
	// }

	// ...

	return router
}
