package main

import (
	"html/template"
	"intro-golang/app-web/model"
	"net/http"
)

// aponta aonde estão todos os HTML's para integração
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// atende uma requisição após o barra "/"
	http.HandleFunc("/", index)

	// cria uma conexão web
	http.ListenAndServe(":8000", nil)
}

// interage com o HTML
func index(w http.ResponseWriter, r *http.Request) {

	findAllProducts := model.FindAllProducts()

	// adiciona itens diretamente no HTML.
	temp.ExecuteTemplate(w, "Index", findAllProducts)
}
