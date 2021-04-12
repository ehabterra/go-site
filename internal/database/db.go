package database

import (
	"fmt"
	"log"

	"github.com/ehabterra/go-site/internal/models"

	"gorm.io/driver/postgres"

	"github.com/ehabterra/go-site/internal/environment"

	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	username := environment.GetEnv("DB_USER", "")
	password := environment.GetEnv("DB_PASS", "")
	dbName := environment.GetEnv("DB_NAME", "")
	dbHost := environment.GetEnv("DB_HOST", "")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	log.Println(dsn)

	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	db := conn
	db.Debug().AutoMigrate(&models.GoSite{}, &models.GoSiteAttribute{}) //Database migration

	return db
}
