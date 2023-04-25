package dbsetup

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"scanner-web-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db_hostname := "192.168.1.23" //os.Getenv("POSTGRES_HOST")
	db_name := "inventory"        //os.Getenv("POSTGRES_DB")
	db_user := "postgres"         //os.Getenv("POSTGRES_USER")
	db_pass := "pepito"           //os.Getenv("POSTGRES_PASSWORD")
	db_port := "5432"             //os.Getenv("POSTGRES_PORT")

	dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)
	database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Product{})

	DB = database
}
