package support

import (
	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

const apikey = "831cb471-5969-49ce-bf40-76aefe8f3cad"

func CleanString1() {
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

func GetCoordinate(coordinate string) (string, string) {
	var latitude, longitude string
	for i := 0; i < len(coordinate); i++ {
		if coordinate[i] == ' ' {
			latitude = coordinate[:i]
			longitude = coordinate[i+1:]
			break
		}
	}
	return latitude, longitude
}

func SetLocationCoordinate(address string, location string, ID int) error {
	response, err := http.Get(address)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// fmt.Println(string(result))
	var object *container.Geolocalization = &container.Geolocalization{}
	err = json.Unmarshal(result, object)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	latitude, longitude := GetCoordinate(object.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos)
	container.MapAddressLocation.LocationMainField[ID].LocationNameField[location][0] = latitude
	container.MapAddressLocation.LocationMainField[ID].LocationNameField[location][1] = longitude

	return nil
}

func GeolocalizationGetCoordinate() {
	CleanString1()
	var address string
	response := httptest.NewRecorder()
	var mapping_existence map[string]bool = make(map[string]bool)
	for i := 0; i < len(container.Mainobject.Location.Index); i++ {
		container.MapAddressLocation.LocationMainField = append(container.MapAddressLocation.LocationMainField, container.LocationName{})
		container.MapAddressLocation.LocationMainField[i].LocationNameField = make(map[string][]string)
		for j := 0; j < len(container.Mainobject.Location.Index[i].Locations); j++ {
			_, ok := mapping_existence[container.Mainobject.Location.Index[i].Locations[j]]

			if !ok {
				mapping_existence[container.Mainobject.Location.Index[i].Locations[j]] = true
				fmt.Println(container.Mainobject.Location.Index[i].Locations[j])

				container.MapAddressLocation.LocationMainField[i].LocationNameField[container.Mainobject.Location.Index[i].Locations[j]] = make([]string, 2)

				address = "https://geocode-maps.yandex.ru/1.x/?apikey=" + apikey + "&geocode=" + container.Mainobject.Location.Index[i].Locations[j] + "&lang=en_US&format=json"
				err := SetLocationCoordinate(address, container.Mainobject.Location.Index[i].Locations[j], i)
				if err != nil {
					log.Println("There is an error:", err)
					errorness.ErrorHandler(response, http.StatusInternalServerError)

				}
			}
		}
	}

	fmt.Println(container.MapAddressLocation)
	fileName := "/home/student/Golang/Project/Geolocalization/core/backend/container/LocationCoordinate.json"

	jsonData, err := json.Marshal(container.MapAddressLocation)
	if err != nil {
		// handle error
		fmt.Println("Error marshalling data:", err)
		return
	}
	fmt.Println(string(jsonData))

	err = ioutil.WriteFile(fileName, jsonData, 0o644) // Adjust file permissions as needed
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file successfully!")

	result, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(result))
	var fromjson *container.LocationMain = &container.MapAddressLocation

	err = json.Unmarshal(result, fromjson)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println(fromjson.LocationMainField[0].LocationNameField["georgia_usa"])
}

//  func Geolocalization() {

// 	// var object *container.LocationData = &container.LocationData{}
// 	// err = json.Unmarshal(result, object)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// container.MapAddressCoordinate = *object

// 	// // fmt.Println("Zero:", *object)
// 	// container.MapLocationAddress = make([][]map[string]string, len(container.Mainobject.Location.Index))

// 	// for i := 0; i < len(container.Mainobject.Location.Index); i++ {
// 	// 	container.MapAddressCoordinate[("Usa")]
// 	// 	container.MapLocationAddress[i] = make([]map[string]string, 0)
// 	// 	for j := 0; j < len(container.Mainobject.Location.Index[i].Locations); j++ {
// 	// 		container.MapLocationAddress[i] = append(container.MapLocationAddress[i])
// 	// 	}
// 	// }
// }
