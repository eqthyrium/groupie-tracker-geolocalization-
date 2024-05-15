package service

import (
	"net/http"
	"strconv"
	"text/template"

	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
)

func ToString(transit container.LocationName, ID int) string {
	var result string
	var count int
	for _, value := range transit.LocationNameField {
		result += "[" + value[0] + " " + value[1] + "]"
		count++
	}
	// fmt.Println(result)
	// fmt.Println("Second:", count)
	return result
}

func AccountPageHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		errorness.ErrorHandler(writer, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("../core/frontend/service/Accountpage.html")
	if err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}

	ImageIndex := request.FormValue("Image_Index_Value")
	if len(ImageIndex) >= 2 && ImageIndex[0] == '0' {
		errorness.ErrorHandler(writer, http.StatusNotFound)
		return
	}
	result, err := strconv.Atoi(ImageIndex)
	if err != nil || (result >= 52 || result < 0) {
		errorness.ErrorHandler(writer, http.StatusBadRequest)
		return
	}

	artist := container.Mainobject.Artist.Index[result]
	// location := container.Mainobject.Location.Index[result]
	// dates := container.Mainobject.Dates.Index[result]
	relation := container.Mainobject.Relation.Index[result]

	// fmt.Println("First:", container.MapAddressLocation.LocationMainField[result])

	data := struct {
		Image         string
		Name          string
		Members       []string
		CreationDate  int
		FirstAlbum    string
		DatesLocation map[string][]string
		Coordinate    string
	}{
		Image:         artist.Image,
		Name:          artist.Name,
		Members:       artist.Members,
		CreationDate:  artist.CreationDate,
		FirstAlbum:    artist.FirstAlbum,
		DatesLocation: relation.DatesLocations,
		Coordinate:    ToString(container.MapAddressLocation.LocationMainField[result], result),
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}
}
