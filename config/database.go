package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUSER, dbPASS, dbHOST, dbNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect into database")
	}

	// db.AutoMigrate(&model.Book{}, &model.User{})
	return db
}

// CloseConnection databases
func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()
}
