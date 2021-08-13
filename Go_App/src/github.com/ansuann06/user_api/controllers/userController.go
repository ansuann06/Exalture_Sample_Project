package controllers

import (
	"fmt"
	"net/http"

	"github.com/ansuann06/user_api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBooks(c *gin.Context) {
	books, err := GetBooks(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (a *APIEnv) GetBook(c *gin.Context) {
	id := c.Params.ByName("id")
	book, exists, err := GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no book in db")
		return
	}

	c.JSON(http.StatusOK, book)
}

type CreateBookInput struct {
	Name      string `json:"name" binding:"required"`
	Author    string `json:"author" binding:"required"`
	PageCount int    `json:"pagecount" binding:"required" `
}

func (a *APIEnv) CreateBook(c *gin.Context) {
	// book := models.Book{}
	var input CreateBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	book := models.Book{Name: input.Name, Author: input.Author, PageCount: input.PageCount}
	if err := a.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	err = DeleteBook(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

func (a *APIEnv) UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := GetBookByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	updatedBook := models.Book{}
	err = c.BindJSON(&updatedBook)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := UpdateBook(a.DB, &updatedBook); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetBook(c)
}
