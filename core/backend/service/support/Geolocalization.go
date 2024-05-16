package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
)

// const apikey = "831cb471-5969-49ce-bf40-76aefe8f3cad"

func CleanString() {
	for i := 0; i < len(container.Mainobject.Location.Index); i++ {
		for j := 0; j < len(container.Mainobject.Location.Index[i].Locations); j++ {
			var subsitute string = ""
			var flag bool = false
			var turtle int = 0
			for z := 0; z < len(container.Mainobject.Location.Index[i].Locations[j]); z++ {
				if container.Mainobject.Location.Index[i].Locations[j][z] == '-' {
					flag = true
					subsitute += container.Mainobject.Location.Index[i].Locations[j][turtle:z] + "_"
					turtle = z + 1
				}

				if flag && z+1 == len(container.Mainobject.Location.Index[i].Locations[j]) {
					subsitute += container.Mainobject.Location.Index[i].Locations[j][turtle : z+1]
				}
			}

			if flag {
				container.Mainobject.Location.Index[i].Locations[j] = subsitute
			}

		}
	}
}

func Geolocalization() error {
	CleanString()
	response := httptest.NewRecorder()

	fileName := "/home/student/Golang/Project/Geolocalization/core/backend/container/LocationCoordinate.json"

	result, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file:", err)
		errorness.ErrorHandler(response, http.StatusInternalServerError)
		return err
	}

	// fmt.Println(string(result))
	var fromjson *container.LocationMain = &container.MapAddressLocation

	err = json.Unmarshal(result, fromjson)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		errorness.ErrorHandler(response, http.StatusInternalServerError)
		return err
	}

	return nil
}
