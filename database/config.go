package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/helpers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	var path string
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "testing" {
		path = "../.env"
	}
	if stage != "testing" {
		path = ".env"
	}

	// comment this line for production ready app (with container)
	helpers.LoadEnv(path)

	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		helpers.GetAsString("DB_USER", "postgres"),
		helpers.GetAsString("DB_PASSWORD", "BROWN%20100"),
		helpers.GetAsString("DB_HOST", "localhost"),
		helpers.GetAsInt("DB_PORT", 5432),
		helpers.GetAsString("DB_NAME", "pbifinal"),
	)

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MigrateDB() {
	stage := helpers.GetAsString("STAGE", "development")

	if stage == "development" ||
		stage == "production" {
		DB.Debug().AutoMigrate(models.User{}, models.Photo{})
	}
}

func GetDB() *gorm.DB {
	return DB
}
