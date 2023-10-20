package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/notcuddle/go-http/routes"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/about", routes.About)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", routes.Parametro)
	mux.HandleFunc("/parametro-querystring", routes.ParametroQueryString)

	server := &http.Server{
		Addr:         "localhost:8081",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Corriendo servidor desde localhost:8081")
	log.Fatal(server.ListenAndServe())

}

/*
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:8081", nil))

}*/
