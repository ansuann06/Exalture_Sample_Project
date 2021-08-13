package database

import (
	"github.com/ansuann06/crud_api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbname := "first_db"
	user := "postgres"
	password := "postgres123"

	db, err := gorm.Open("postgres",
			"host="+host+
				" port="+port+
				" user="+user+
				" dbname="+dbname+
				" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate([]models.Book{})
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}