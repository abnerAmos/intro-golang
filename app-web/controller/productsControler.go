package controler

import (
	"html/template"
	"intro-golang/app-web/model"
	"log"
	"net/http"
	"strconv"
)

// aponta aonde estão todos os HTML's para integração
var temp = template.Must(template.ParseGlob("templates/*.html"))

// interage com o HTML
func Index(w http.ResponseWriter, r *http.Request) {

	findAllProducts := model.FindAllProducts()

	// adiciona itens diretamente no HTML.
	temp.ExecuteTemplate(w, "Index", findAllProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		model.NewProduct(name, description, priceConvert, quantityConvert)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	model.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := model.EditProduct(idProduct)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na Conversão do Id para int:", err)
		}

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na Conversão do Preço para float64:", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na Conversão da Quantidade para int:", err)
		}

		model.UpdateProduct(idConvert, name, description, priceConvert, quantityConvert)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
