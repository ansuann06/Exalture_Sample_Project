package controllers

import (
	"log"

	"gorm.io/gorm"
)

func CreateTodoTable(db) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Todo{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *gorm.DB

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

// opts := &orm.CreateTableOptions{
// IfNotExists: true,
// }
// createError := db.CreateTable(&Todo{}, opts)
// if createError != nil {
// log.Printf("Error while creating todo table, Reason: %v\n", createError)
// return createError
// }
// log.Printf("Todo table created")
// 	return nil
// }
