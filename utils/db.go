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
)

type User struct {
	gorm.Model
	Name string
	Permission int
}

func Db() {
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
		Addr:   ":3306",
		DBName: "bot",
	}
	
	
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	db, err := gorm.Open(gormsql.New(gormsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	
	pingErr := sqlDB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Migrate the schema
	db.AutoMigrate(&User{})
}