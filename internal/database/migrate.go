package database

import (
	"log"

	"github.com/jpeccia/bebi-delivery-server/internal/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Store{},
		&models.Order{},
		&models.OrderItem{},
		&models.StoreProduct{},
		&models.Rating{},
	)

	if err != nil {
		log.Printf("Erro ao aplicar as migrations: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso!")
}
