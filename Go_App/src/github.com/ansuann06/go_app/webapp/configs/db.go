package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	controllers "github.com/ansuann06/go_app/webapp/controllers"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres123"
	dbname   = "first_db"
)

var DB *gorm.DB

func Open() error {
	var err error
	psqlconn := fmt.Sprintf("host = %s port = %d user= %s password= %s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = gorm.Open("postgres", psqlconn)
	controllers.CreateUserTable()
	if err != nil {
		return err
	}
	return err
}

func Close() error {
	return DB.Close()
}
