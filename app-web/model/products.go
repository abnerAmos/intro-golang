package model

import (
	"intro-golang/app-web/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConectionDataBase() // Abre conexão

	selectAll, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAll.Next() {
		var id, quantity int
		var name, description string
		var price float64

		// Verifica os dados no Banco de dados. PS. Precisa estar na mesma ordem das colunas no DB
		err = selectAll.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}

	defer db.Close() // Fecha conexão após execuçã total dos métodos. PS. defer = AfterAll
	return products
}

func NewProduct(name, description string, price float64, quantity int) {
	db := db.ConectionDataBase()

	insert, err := db.Prepare("insert into products(name, description, price, quantity) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectionDataBase()

	deleteProduct, err := db.Prepare("delete from products where id=?")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConectionDataBase()

	selectProduct, err := db.Query("select * from products where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for selectProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}
	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConectionDataBase()

	updateProduct, err := db.Prepare("update products set name=?, description=?, price=?, quantity=? where id=?")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
