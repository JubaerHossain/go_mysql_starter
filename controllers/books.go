package controllers

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type createBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var request createBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	book := models.Book{Title: request.Title, Author: request.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}


func UpdateBook(c *gin.Context) {
	var request UpdateBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = request.Title
	book.Author = request.Author
	db.Save(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}


func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

