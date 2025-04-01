package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/bebi-delivery-server/config"
	"github.com/jpeccia/bebi-delivery-server/internal/database"
	"github.com/jpeccia/bebi-delivery-server/internal/handlers"
	"github.com/jpeccia/bebi-delivery-server/internal/repositories"
	"github.com/jpeccia/bebi-delivery-server/internal/routes"
	"github.com/jpeccia/bebi-delivery-server/internal/services"
)

func main() {
	config.LoadEnvVariables()

	database.Connect()
	database.MigrateDB()

	userRepo := repositories.NewUserRepository(database.DB)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	routes.SetupRoutes(router, userHandler)

	port := ":8080"
	log.Printf("Servidor rodando na porta %s", port)
	router.Run(port)
}
