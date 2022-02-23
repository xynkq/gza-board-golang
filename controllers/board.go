package controllers

import (
	"net/http"
	"time"

	"gza/board/models"

	"github.com/gin-gonic/gin"
)

func GetBoards(c *gin.Context) {
	var boards []models.Board
	db := models.GetDatabase()
	r := db.Model(boards).Find(&boards)
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	// c.JSON(http.StatusOK, &boards)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   boards,
	})
}

func GetBoard(c *gin.Context) {
	var board models.Board
	db := models.GetDatabase()
	r := db.Model(board).First(&board, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   board,
	})
}

func CreateBoard(c *gin.Context) {
	var board models.Board
	c.BindJSON(&board)
	board.Created = time.Now()
	board.Updated = time.Now()
	db := models.GetDatabase()
	db.Create(&board)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   board,
	})
}

func UpdateBoard(c *gin.Context) {
	var data models.Board
	c.BindJSON(&data)
	db := models.GetDatabase()
	var board models.Board
	r := db.Find(&board, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	db.Model(&board).Updates(&data)
	//models.DB.Model(&info).Update("Name", "Lee")
	//models.DB.Model(&info).Update(models.Info{Name: "Lee", Email: "rhdtha01@gmail.com"})
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func DeleteBoard(c *gin.Context) {
	var board models.Board
	db := models.GetDatabase()
	r := db.Find(&board, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	db.Delete(&board)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   board,
	})
}
