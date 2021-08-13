package models

// import (
//   "github.com/jinzhu/gorm"
// )

type Book struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Author    string `json:"author"`
	Name      string `json:"name"`
	PageCount int    `json:"pagecount"`
}
