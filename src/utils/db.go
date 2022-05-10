package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/eas-kaine/discord-bot/models"
)

var DB *gorm.DB

func SetupDB() {
  
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")

	// Capture connection properties
	// cfg := mysql.Config{
	// 	User:   user,
	// 	Passwd: pass,
	// 	Net:    "tcp",
	// 	Addr:   "mysql:3306",
	// 	DBName: "discord-bot",
	// 	AllowNativePasswords: true,
	// 	ParseTime: true,
	// }
	
	
	// sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	panic(err)
	// }
	// import (
		// "gorm.io/driver/mysql"
		// "gorm.io/gorm"
	//   )
	  
	//   func main() {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		dsn := fmt.Sprintf(`%s:%s@tcp(mysql:3306)/discord-bot?charset=utf8mb4&parseTime=True&loc=Local`, user, pass)
		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	//   }

	// sqlDB.SetConnMaxLifetime(time.Minute * 3)
	// sqlDB.SetMaxOpenConns(10)
	// sqlDB.SetMaxIdleConns(10)

	// DB, err = gorm.Open(gormsql.New(gormsql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	
	// pingErr := sqlDB.Ping()
	// if pingErr != nil {
	// 	log.Println(pingErr)
	// }
	// fmt.Println("Connected!")

	// Migrate the schemas
	DB.AutoMigrate(&models.User{}, &models.Action{})
}