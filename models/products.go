package models

import (
	"github.com/nicodemos234/cursoweb/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Total       int
}

func SearchAllProducts() []Product {
	database := db.ConnectDatabase()
	selectAllProducts, err := database.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id int
		var name, description string
		var price float64
		var total int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &total)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Total = total
		products = append(products, p)
	}

	defer database.Close()
	return products
}

func CreateNewProduct(name string, description string, price float64, total int) {
	database := db.ConnectDatabase()
	insertData, err := database.Prepare("INSERT INTO products(name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, total)
	defer database.Close()
}

func DeleteProduct(productId string) {
	database := db.ConnectDatabase()
	deleteProduct, err := database.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(productId)
	defer database.Close()
}

func GetProduct(productId string) Product {
	database := db.ConnectDatabase()
	productDb, err := database.Query("SELECT * FROM products WHERE id=$1", productId)
	if err != nil {
		panic(err.Error())
	}
	productToUpdate := Product{}
	for productDb.Next() {
		var id, total int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &total)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Total = total
	}
	defer database.Close()
	return productToUpdate
}

func UpdateProduct(id int, name string, description string, price float64, total int) {
	database := db.ConnectDatabase()
	updateProduct, err := database.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, total, id)
	defer database.Close()
}
