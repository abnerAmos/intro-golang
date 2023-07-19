package main

import (
	"intro-golang/app-web/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	// cria uma conexão web
	http.ListenAndServe(":8000", nil)
}
