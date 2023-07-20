package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Cria conexão com o Banco de Dados
func ConectionDataBase() *sql.DB {

	// String de conexão
	// conection := "user=root dbname=golang password=root host=localhost sslmode=disable"
	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err.Error())
	}
	return db
}
