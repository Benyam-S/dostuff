package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Benyam-S/dostuff/api/v1/http/handler"
	geoService "github.com/Benyam-S/dostuff/geoip/service"
	tranService "github.com/Benyam-S/dostuff/translate/service"
	"github.com/gorilla/mux"
)

// Start is a function that starts the v1 API
func Start(router *mux.Router) {

	pwd, _ := os.Getwd()

	// Reading data from repository.store.json file
	store := make(map[string]string)
	storeDir := filepath.Join(pwd, "translate/repository/repository.store.json")
	storeData, err := ioutil.ReadFile(storeDir)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(storeData, &store)
	if err != nil {
		panic(err)
	}

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

	// Creating translation service
	translationService := tranService.NewTranslateService(store)

	// Creating Geo IP location service
	geoIPLocationService := geoService.NewGeoIPService()

	// Creating new API handler
	apiHandler := handler.NewAPIHandler(geoIPLocationService, translationService)

	router.HandleFunc("/do_stuff", apiHandler.HandleDoStuff).Methods("GET")
	router.HandleFunc("/do_pretty_cool_stuff", apiHandler.HandleDoPrettyCoolStuff).Methods("GET")
}
