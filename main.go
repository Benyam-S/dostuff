package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Benyam-S/dostaff/api/v1/http/handler"
	"github.com/Benyam-S/dostaff/geoip/service"
	"github.com/gorilla/mux"
)

// main is a function that runs the server and related services
func main() {

	pwd, _ := os.Getwd()

	// Reading data from config.api.json file
	apiConfig := make(map[string]interface{})
	apiConfigDir := filepath.Join(pwd, "config/config.api.json")
	apiConfigData, err := ioutil.ReadFile(apiConfigDir)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(apiConfigData, &apiConfig)
	if err != nil {
		panic(err)
	}

	baseURL, ok1 := apiConfig["geo_ip_service_base_url"].(string)
	apiKey, ok2 := apiConfig["geo_ip_service_api_key"].(string)

	if !ok1 || !ok2 {
		panic(errors.New("unable to parse api config data"))
	}

	os.Setenv("geo_ip_service_base_url", baseURL)
	os.Setenv("geo_ip_service_api_key", apiKey)

	// Creating Geo IP location service
	geoIPLocationService := service.NewGeoIPService()

	// Creating new API handler
	apiHandler := handler.NewAPIHandler(geoIPLocationService)

	// Creating new http request router
	router := mux.NewRouter()

	router.HandleFunc("/do_staff", apiHandler.HandleDoStaff).Methods("GET")
	router.HandleFunc("/do_pretty_cool_staff", apiHandler.HandleDoPrettyCoolStaff).Methods("GET")

	fmt.Println("Web server has been successfully started!")

	http.ListenAndServe(":8080", router)
}
