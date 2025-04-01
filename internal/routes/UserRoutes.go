package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/bebi-delivery-server/internal/handlers"
	"github.com/jpeccia/bebi-delivery-server/internal/middlewares"
)

func UserRoutes(r *gin.RouterGroup, userHandler *handlers.UserHandler) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.Login)
		userGroup.PUT("/:id/upgrade", middlewares.AuthMiddleware(), userHandler.UpgradeToStoreOwner)
	}
}
