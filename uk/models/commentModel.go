package models

import (
	"log"
	"uk/database"
)

type Comment struct {
	Id uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId uint `json:"user_id" gorm:"column:user_id"`
	PhotoId uint `json:"photo_id" gorm:"column:photo_id"`
	Message string `json:"message" gorm:"column:message"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func CreateComment(comment Comment) Comment {
	result := database.GetDB().Create(&comment)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return comment
}