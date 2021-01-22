package config

import (
	"fmt"
	"os"
	"Atlantis-Backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitConnection connection databases
func InitConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env files")
	}

	dbUSER := os.Getenv("DB_USER")
	dbPASS := os.Getenv("DB_PASS")
	dbHOST := os.Getenv("DB_HOST")
	dbNAME := os.Getenv("DB_NAME")
	dbPORT := os.Getenv("DB_PORT")

	connectionName := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHOST, dbPORT, dbUSER, dbNAME, dbPASS)
	fmt.Println(connectionName)
	db, err := gorm.Open(postgres.Open(connectionName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect into database")
	}

	db.AutoMigrate(&models.User{})
	return db
}

// CloseConnection databases
func CloseConnection(db *gorm.DB) {
	dbPOSTGRE, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}

	dbPOSTGRE.Close()
}
