package main

import (
	"crud/models"
	"crud/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupDB()
	db.AutoMigrate(&models.Book{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	
	v1 := r.Group("/api/v1/books")
	{
		v1.GET("/", controllers.FindBooks)
		v1.POST("/create", controllers.CreateBook)
		v1.GET("/:id", controllers.FindBook)
		v1.PUT("/:id", controllers.UpdateBook)
		v1.DELETE("/:id", controllers.DeleteBook)
	}
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
	r.Run()
}
