package dbsetup

import (
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	db_hostname := os.Getenv("POSTGRES_HOST")
	db_name := os.Getenv("POSTGRES_DB")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_port := os.Getenv("POSTGRES_PORT")

	//dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)
	dbURl := db_user + ":" + db_pass + "@tcp" + "(" + db_hostname + ":" + db_port + ")/" + db_name + "?" + "parseTime=true&loc=Local"
	//database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})
	database, err := gorm.Open(mysql.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	//database.AutoMigrate(&models.Product{})
	DB = database
}
