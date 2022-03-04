package database

import (
	"main/models"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=104.198.158.95 user=postgres dbname=gorm password=1q2w3e4r5t sslmode=disable",
	}))

	if err != nil {
		panic("cloud not")
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.UserContact{})
}
