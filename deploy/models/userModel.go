package models

import (
	"fmt"
	"sesi7/database"
	"sesi7/helper"
)

type User struct {
	Id       uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func Register(user User) User {
	hashedPassword, _ := helper.HashPassword(user.Password)
	user.Password = hashedPassword
	result := database.GetDB().Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return user
}

func Login(user User) (User, error) {
	var userDB User
	result := database.GetDB().Where("email = ?", user.Email).First(&userDB)

	// Check if user not found
	if result.RowsAffected == 0 {
		return userDB, fmt.Errorf("User not found")
	}

	// Check if password not match
	if !helper.CheckPasswordHash(user.Password, userDB.Password) {
		return userDB, fmt.Errorf("password not match")
	}

	return userDB, nil
}

func GetUsers() []User {
	var users []User
	result := database.GetDB().Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return users
}
