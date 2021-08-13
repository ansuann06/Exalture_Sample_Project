package controllers

import (
	"github.com/ansuann06/user_api/models"
	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	query := db.Select("books.*").
		Group("books.id")
	if err := query.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func GetBookByID(id string, db *gorm.DB) (models.Book, bool, error) {
	b := models.Book{}

	query := db.Select("books.*")
	query = query.Group("books.id")
	err := query.Where("books.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var b models.Book
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *models.Book) error {
	if err := db.Model(&b).Updates(&b).Error; err != nil {
		return err
	}
	return nil
}
