package service

import (
	"html/template"
	"net/http"

	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
)

func MainPageHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		errorness.ErrorHandler(writer, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		errorness.ErrorHandler(writer, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("../core/frontend/service/Mainpage.html")
	if err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}

	var name []string
	for i := 0; i < len(container.Mainobject.Artist.Index); i++ {
		name = append(name, container.Mainobject.Artist.Index[i].Image)
	}

	data := struct {
		ImageValue1 []string
	}{
		ImageValue1: name,
	}

	if err = tmpl.Execute(writer, data); err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}
}
