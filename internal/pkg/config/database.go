package config

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Bağlantı, veritabanına bağlanan bir işlevdir.
func Connection() *gorm.DB {
	godotenv.Load()
	dsn := "host=localhost user=postgres password=172754 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
