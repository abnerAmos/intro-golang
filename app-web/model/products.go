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

		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}

	defer db.Close() // Fecha conexão após execuçã total dos métodos. PS. defer = AfterAll
	return products
}
