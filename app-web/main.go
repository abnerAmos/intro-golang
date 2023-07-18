package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// aponta aonde estão todos os HTML's para integração
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	// atende uma requisição após o barra "/"
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil) // cria uma conexão web
}

// interage com o HTML
func index(w http.ResponseWriter, r *http.Request) {

	db := conectionDataBase() // Abre conexão

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

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close() // Fecha conexão após execuçã total dos métodos. PS. defer = AfterAll
}

// Cria conexão com o Banco de Dados
func conectionDataBase() *sql.DB {

	// String de conexão
	// "root:root@/golang"
	// conection := "user=root dbname=golang password=root host=localhost sslmode=disable"
	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}
