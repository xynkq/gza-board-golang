package models

import (
	"time"
)

type Board struct {
	Id      int       `json:"id" binding:"required" gorm:"primary_key" `
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	UserId  string    `json:"userId" binding:"required"`
	Created time.Time `json:"created" gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP"`
	Updated time.Time `json:"updated" gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
