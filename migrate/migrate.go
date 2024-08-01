package main

import (
	"github.com/jamwitk/crud-books/initializers"
	"github.com/jamwitk/crud-books/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.Connect()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Book{})
	if err != nil {
		return
	}
}
