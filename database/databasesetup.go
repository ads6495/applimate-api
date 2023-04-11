package database

import (
	"log"
	"os"

	"github.com/ads6495/applimate-api/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	db, err := gorm.Open(mysql.Open("secret string"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("connected to the Database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Add migrations
	db.Migrator().DropTable(&models.Job{}, &models.User{})
	db.AutoMigrate(&models.Job{}, &models.User{})

	Database = DbInstance{Db: db}
}
