package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/bebi-delivery-server/internal/handlers"
)

func UserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.PUT("/:id/upgrade", userHandler.UpgradeToStoreOwner)
	}
}