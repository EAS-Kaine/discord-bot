package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	gormsql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/eas-kaine/discord-bot/models"
)

var DB *gorm.DB

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")

	// Capture connection properties
	cfg := mysql.Config{
		User:   user,
		Passwd: pass,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mysql",
		AllowNativePasswords: true,
		ParseTime: true,
	}
	
	
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	DB, err = gorm.Open(gormsql.New(gormsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	
	pingErr := sqlDB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Migrate the schemas
	DB.AutoMigrate(&models.User{}, &models.Action{})
	// DB.AutoMigrate(&models.Quiz{})
}