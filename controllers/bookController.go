package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jamwitk/crud-books/initializers"
	"github.com/jamwitk/crud-books/models"
	"gorm.io/gorm/utils"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	initializers.DB.Find(&books)
	c.JSON(200, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initializers.DB.First(&book, id)
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	// Parse the request body
	var body struct {
		Title       string `json:"Title" validate:"required"`
		Author      string `json:"Author" validate:"required"`
		Description string `json:"Description" validate:"required"`
		ReleaseDate string `json:"ReleaseDate" validate:"required" validate:"datetime=2006-01-02"`
	}
	err := c.Bind(&body)
	if err != nil {
		return
	}

	validate := validator.New()

	err = validate.Struct(body)
	if err != nil {
		LogEvent(c, "Create Book", "Book", "Validation", err.Error())
		LogRequest(c, "400")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx := initializers.DB.Begin()

	book := models.Book{Title: &body.Title, Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate}
	result := tx.Create(&book)

	if result.Error != nil {
		LogEvent(c, "Create Book", "Book", "Error", result.Error.Error())
		LogRequest(c, "400")
		tx.Rollback()
		c.JSON(400, result.Error)
		return
	}
	tx.Commit()
	LogDatabase(c, "INSERT", "Books", "INSERT INTO books (title, author, description, release_date) VALUES ('"+body.Title+"', '"+body.Author+"', '"+body.Description+"', '"+body.ReleaseDate+"')", utils.ToString(result.RowsAffected))
	LogEvent(c, "Create Book", "Book", "Success", "Book created successfully")
	LogRequest(c, "200")

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title       string
		Author      string
		Description string
		ReleaseDate string
	}
	err := c.Bind(&body)
	if err != nil {
		return
	}
	var book models.Book
	initializers.DB.First(&book, id)
	result := initializers.DB.Model(&book).Updates(models.Book{Title: &body.Title, Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate})
	if result.Error != nil {
		LogEvent(c, "Update Book", "Book", "Error", result.Error.Error())
		LogRequest(c, "400")
		c.JSON(400, result.Error)
		return
	}
	LogDatabase(c, "UPDATE", "Books", "UPDATE books SET title = '"+body.Title+"', author = '"+body.Author+"', description = '"+body.Description+"', release_date = '"+body.ReleaseDate+"' WHERE id = "+id, "1")
	LogEvent(c, "Update Book", "Book", "Success", "Book updated successfully")
	LogRequest(c, "200")
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	result := initializers.DB.Delete(&models.Book{}, id)
	if result != nil {
		LogEvent(c, "Delete Book", "Book", "Error", result.Error.Error())
		LogRequest(c, "400")
		c.JSON(400, result.Error)
		return
	}
	LogDatabase(c, "DELETE", "Books", "DELETE FROM books WHERE id = "+id, utils.ToString(result.RowsAffected))
	LogEvent(c, "Delete Book", "Book", "Success", "Book deleted successfully")
	c.JSON(200, id)
}
