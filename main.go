package main

import (
	"github.com/jpeccia/bebi-delivery-server/config"
	"github.com/jpeccia/bebi-delivery-server/internal/database"
)

func main() {
	config.LoadEnvVariables()
	database.Connect()
	database.MigrateDB()
}
