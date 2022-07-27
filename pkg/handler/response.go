package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Error ..
type Error struct {
	Message string
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger, _ := zap.NewProduction()
	logger.Error(message)

	c.AbortWithStatusJSON(statusCode, Error{message})
}
