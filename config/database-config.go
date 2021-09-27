package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lexuanquynh/go_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}
	// errDB := db.AutoMigrate(&entity.Book{}, &entity.User{}).Error
	// if errDB != nil {
	// 	fmt.Println("Error HandleMigrate:" + errDB())
	// }
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	// if errDb != nil {
	// 	panic(err)
	// }
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Faild to close database connection")
	}
	dbSQL.Close()
}
