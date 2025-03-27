package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=" + os.Getenv("DB_HOST") + "user=" + os.Getenv("DB_USER") + "password=" + os.Getenv("DB_PASSWORD") + "dbname=" + os.Getenv("DB_NAME") + "port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	DB = db
}