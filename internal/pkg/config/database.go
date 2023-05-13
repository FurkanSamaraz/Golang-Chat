package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Bağlantı, veritabanına bağlanan bir işlevdir.
func Connection() *gorm.DB {
	godotenv.Load()
	dsn := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
