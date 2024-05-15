package support

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Adlet/core/backend/container"
)

func Getting(link string, object interface{}) error {
	response, err := http.Get(link)
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
	err = json.Unmarshal(result, object)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func Jsonhandler() {
	object := &struct {
		Artists   string
		Locations string
		Dates     string
		Relation  string
	}{}

	if err := Getting("https://groupietrackers.herokuapp.com/api", object); err != nil {
		return
	}

	var objectdate *container.Datestruct = &container.Datestruct{}
	if err := Getting(object.Dates, objectdate); err != nil {
		return
	}

	var objectlocation *container.Locationstruct = &container.Locationstruct{}
	if err := Getting(object.Locations, objectlocation); err != nil {
		return
	}

	var objectrelation *container.Relationstruct = &container.Relationstruct{}
	if err := Getting(object.Relation, objectrelation); err != nil {
		return
	}

	var objectarirst *container.Artiststruct = &container.Artiststruct{}
	if err := Getting(object.Artists, &((*objectarirst).Index)); err != nil {
		return
	}

	container.Mainobject.Dates = objectdate
	container.Mainobject.Location = objectlocation
	container.Mainobject.Relation = objectrelation
	container.Mainobject.Artist = objectarirst
}
