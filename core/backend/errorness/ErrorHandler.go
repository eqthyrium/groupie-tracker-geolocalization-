package errorness

import (
	"html/template"
	"net/http"
)

type Errorness struct {
	Status_meesage string
	Error_message  string
}

func ErrorHandler(writer http.ResponseWriter, status int) { // Internal Server Error, Bad request, Not Found, Method Not Allowed,
	var object Errorness = Errorness{}

	if status == 400 {
		object.Status_meesage = "400"
		object.Error_message = "Bad Request"
		writer.WriteHeader(http.StatusBadRequest)
	} else if status == 404 {
		object.Status_meesage = "404"
		object.Error_message = "Not Found"
		writer.WriteHeader(http.StatusNotFound)

	} else if status == 405 {
		object.Status_meesage = "405"
		object.Error_message = "Method Not Allowed"
		writer.WriteHeader(http.StatusMethodNotAllowed)
	} else if status == 500 {
		object.Status_meesage = "500"
		object.Error_message = "Internal Server Error"
		writer.WriteHeader(http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("../core/frontend/errorness/goback.html")
	if err != nil {
		http.Error(writer, "(500) Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(writer, object)
	if err != nil {
		http.Error(writer, "(500) Internal Server Error", http.StatusInternalServerError)
		return
	}
}
