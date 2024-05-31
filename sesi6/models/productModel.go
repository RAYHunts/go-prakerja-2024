package models

import (
	"log"
	"sesi6/database"
)

type Product struct {
	ID    uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"column:name"`
	Price int    `json:"price" gorm:"column:price"`
}

func GetProducts() []Product {
	var products []Product
	result := database.GetDB().Find(&products)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return products
}

func GetProduct(id int) Product {
	var product Product
	result := database.GetDB().First(&product, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return product
}

func CreateProduct(product Product) Product {
	result := database.GetDB().Create(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return product
}

func UpdateProduct(product Product) Product {
	result := database.GetDB().Save(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return product
}

func DeleteProduct(id int) {
	result := database.GetDB().Delete(&Product{}, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}