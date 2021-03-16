package main

import (
	"log"
	"net/http"
	controller "server/controller/api"

	"github.com/julienschmidt/httprouter"

	"server/configs"
)

func main() {
	println("Starting server...")

	r := httprouter.New()
	routes(r)

	var configs = configs.InitServerParams()

	err := http.ListenAndServe(configs.Adress, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/static/*filepath", http.Dir("static"))

	r.GET("/", controller.StartPage)
}
