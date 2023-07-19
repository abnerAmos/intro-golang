package routes

import (
	controller "intro-golang/app-web/controller"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new-product", controller.NewProduct)
	http.HandleFunc("/insert", controller.Insert)
}
