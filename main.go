package main

import (
	"fmt"
	"net/http"
	"os"

	v1 "github.com/Benyam-S/dostuff/api/v1"
	"github.com/gorilla/mux"
)

// main is a function that runs the server and related services
func main() {

	// Creating new http request router
	router := mux.NewRouter()

	v1.Start(router)

	fmt.Println("Web server has been successfully started!")

	http.ListenAndServeTLS(":"+os.Getenv("http_client_server_port"),
		os.Getenv("server.pem"), os.Getenv("server.key"), router)
}
