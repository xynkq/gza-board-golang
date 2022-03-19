package main

import (
	"gza/board/controllers"
	"gza/board/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	boardAPI := r.Group("/board")
	{
		// GET: http://localhost:8080/board
		boardAPI.GET("/", controllers.GetBoards)
		// GET: http://localhost:8080/board/:id
		boardAPI.GET("/:id", controllers.GetBoard)
		// POST: http://localhost:8080/board
		boardAPI.POST("/", controllers.CreateBoard)
		// PUT: http://localhost:8080/board:id
		boardAPI.PUT("/:id", controllers.UpdateBoard)
		// DELETE: http://localhost:8080/board:id
		boardAPI.DELETE("/:id", controllers.DeleteBoard)
	}
	r.Use()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	DB = models.ConnectDatabase()
	r.Run(":8080")
}
