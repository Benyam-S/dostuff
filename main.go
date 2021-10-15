package main

import (
	"fmt"
	"net/http"

	v1 "github.com/Benyam-S/dostaff/api/v1"
	"github.com/gorilla/mux"
)

// main is a function that runs the server and related services
func main() {

	// Creating new http request router
	router := mux.NewRouter()

	v1.Start(router)

	fmt.Println("Web server has been successfully started!")

	http.ListenAndServe(":8080", router)
}
