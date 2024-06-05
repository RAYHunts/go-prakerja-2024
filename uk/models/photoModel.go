package models

import (
	"log"
	"uk/database"
)

type Photo struct {
	Id uint `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"column:title"`
	Caption string `json:"caption" gorm:"column:caption"`
	PhotoUrl string `json:"photo_url" gorm:"column:photo_url"`
	UserID uint `json:"user_id" gorm:"column:user_id"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

func GetPhotos() []Photo{
	var photos []Photo
	result := database.GetDB().Find(&photos)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return photos
}

func GetPhoto(id int) Photo {
	var photo Photo
	result := database.GetDB().First(&photo, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return photo
}

func CreatePhoto(photo Photo) Photo {
	result := database.GetDB().Create(&photo)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return photo
}

func UpdatePhoto(id int, photo Photo) Photo {
	result := database.GetDB().Model(&Photo{}).Where("id = ?", id).Updates(photo)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	updatedPhoto := GetPhoto(id)

	return updatedPhoto
}

func DeletePhoto(id int) {
	result := database.GetDB().Delete(&Photo{}, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}