package main

import (
	"log"
	"net/http"

	backend "Adlet/core/backend/service"
	"Adlet/core/backend/service/support"
)

func main() {
	support.Jsonhandler()
	support.Geolocalization()

	http.HandleFunc("/", backend.MainPageHandler)
	http.HandleFunc("/account", backend.AccountPageHandler)

	log.Println("Server started on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
