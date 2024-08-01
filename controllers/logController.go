package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jamwitk/crud-books/initializers"
	"github.com/jamwitk/crud-books/models"
	"time"
)

func LogRequest(c *gin.Context, responseStatus string) {
	logBody := models.RequestLog{
		UserAgent: c.Request.UserAgent(),
		SessionID: c.Request.Header.Get("SessionID"),
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Request:   c.Request.Method + " " + c.Request.URL.Path,
		Response:  responseStatus,
	}

	tx := initializers.DB.Begin()

	result := tx.Create(&logBody)

	if result.Error != nil {
		tx.Rollback()
		c.JSON(400, result.Error)
		return
	}

	tx.Commit()
}
func LogEvent(c *gin.Context, eventName string, source string, tags string, description string) {
	logBody := models.EventLog{
		EventName:   eventName,
		Source:      source,
		Tags:        tags,
		Description: description,
		UserAgent:   c.Request.UserAgent(),
		SessionID:   c.Request.Header.Get("SessionID"),
		DateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}

	tx := initializers.DB.Begin()

	result := tx.Create(&logBody)

	if result.Error != nil {
		tx.Rollback()
		c.JSON(400, result.Error)
		return
	}

	tx.Commit()
}

func LogDatabase(c *gin.Context, action string, table string, query string, response string) {
	logBody := models.DataBaseLog{
		DateTime:    time.Now().Format("2006-01-02 15:04:05"),
		TableAction: action,
		Table:       table,
		Query:       query,
		Response:    response,
		UserAgent:   c.Request.UserAgent(),
	}

	tx := initializers.DB.Begin()

	result := tx.Create(&logBody)

	if result.Error != nil {
		tx.Rollback()
		c.JSON(400, result.Error)
		return
	}
	tx.Commit()
}
