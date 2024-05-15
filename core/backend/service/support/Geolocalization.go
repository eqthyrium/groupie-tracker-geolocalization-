package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"Adlet/core/backend/container"
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

func Geolocalization() {
	CleanString()
	// response := httptest.NewRecorder()

	fileName := "/home/student/Golang/Project/Geolocalization/core/backend/container/LocationCoordinate.json"

	result, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// fmt.Println(string(result))
	var fromjson *container.LocationMain = &container.MapAddressLocation

	err = json.Unmarshal(result, fromjson)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
