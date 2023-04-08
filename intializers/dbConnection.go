package intializers

import (
	"go-crud-fiber/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes a new database connection and auto-migrates the PostModel table.
func ConnectDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	log.Println("Connected to database")

	log.Println("Running database migrations")
	if err := DB.AutoMigrate(&models.PostModel{}); err != nil {
		log.Fatalf("Error running database migrations: %s", err)
	}
}
