package models

type Product struct {
	ID    uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"column:name"`
	Price int    `json:"price" gorm:"column:price"`
}

func GetProducts() []Product {
	return []Product{
		{
			ID:    1,
			Name:  "Product 1",
			Price: 1000,
		},
		{
			ID:    2,
			Name:  "Product 2",
			Price: 2000,
		},
	}
}

func GetProduct() Product {
	return Product{
		ID:    1,
		Name:  "Product 1",
		Price: 1000,
	}
}

func CreateProduct() Product {
	return Product{
		ID:    1,
		Name:  "Product 1",
		Price: 1000,
	}
}

func UpdateProduct() Product {
	return Product{
		ID:    1,
		Name:  "Product 1",
		Price: 1000,
	}
}

func DeleteProduct() Product {
	return Product{
		ID:    1,
		Name:  "Product 1",
		Price: 1000,
	}
}