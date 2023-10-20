package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola desde gorilla mux")
}

func About(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Page about")
}

func Parametro(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID: "+vars["id"]+" | Slug: "+vars["slug"])
}

func ParametroQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Println(request.URL.RawQuery)
	fmt.Println(request.URL.Query().Get("id"))
}
