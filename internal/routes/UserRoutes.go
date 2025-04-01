package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/bebi-delivery-server/internal/handlers"
	"github.com/jpeccia/bebi-delivery-server/internal/middlewares"
)

func UserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.Login)
		userGroup.PUT("/:id/upgrade", middlewares.AuthMiddleware(), userHandler.UpgradeToStoreOwner)
	}
}
