package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	template.Execute(response, nil)
}

func About(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/about.html")
	if err != nil {
		panic(err)
	}

	template.Execute(response, nil)
}

func Parametro(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/parametros.html")
	vars := mux.Vars(request)

	data := map[string]string{
		"id":   vars["id"],
		"slug": vars["slug"],
	}

	if err != nil {
		panic(err)
	}
	template.Execute(response, data)
}

func ParametroQueryString(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/parametros2.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(request.URL)
	//fmt.Println(request.URL.RawQuery)
	//fmt.Println(request.URL.Query().Get("id"))

	data := make(map[string]string)

	data["id"] = request.URL.Query().Get("id")
	data["slug"] = request.URL.Query().Get("slug")

	template.Execute(response, data)
}

type Datos struct {
	Nombre string
	Edad   int
	Perfil int
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/estructuras.html")
	if err != nil {
		panic(err)
	}

	template.Execute(response, Datos{"Elias", 11, 1})
}

/*func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola desde gorilla mux")
}*/

// func About(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, "Page about")
// }

// func Parametro(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	fmt.Fprintln(response, "ID: "+vars["id"]+" | Slug: "+vars["slug"])
// }

// func ParametroQueryString(response http.ResponseWriter, request *http.Request) {
// 	fmt.Println(request.URL)
// 	fmt.Println(request.URL.RawQuery)
// 	fmt.Println(request.URL.Query().Get("id"))
// }
