package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/bebi-delivery-server/internal/handlers"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	api := router.Group("/api/v1")
	UserRoutes(api, userHandler)
}