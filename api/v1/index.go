package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Benyam-S/dostuff/api/v1/http/handler"
	"github.com/Benyam-S/dostuff/entity"
	geoService "github.com/Benyam-S/dostuff/geoip/service"
	"github.com/Benyam-S/dostuff/log"
	tranService "github.com/Benyam-S/dostuff/translate/service"
	"github.com/gorilla/mux"
)

// Start is a function that starts the v1 API
func Start(router *mux.Router) {

	pwd, _ := os.Getwd()

	// Reading data from config.server.json file and creating the system config object
	sysConfig := new(entity.SystemConfig)
	sysConfigDir := filepath.Join(pwd, "config/config.server.json")
	sysConfigData, err := ioutil.ReadFile(sysConfigDir)
	if err != nil {
		fmt.Println("Error: Unable to read the server configuration file. "+
			"Configuration file: ./config/config.server.json, %s", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(sysConfigData, &sysConfig)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Unable to unmarshal server configuration data. "+
			"Server Config Data => %s, %s", string(sysConfigData), err.Error()))
		os.Exit(1)
	}

	logContainer := &log.LogContainer{
		DebugLogFile: filepath.Join(pwd, sysConfig.LogsPath, sysConfig.Logs["debug_log_file"]),
		ErrorLogFile: filepath.Join(pwd, sysConfig.LogsPath, sysConfig.Logs["error_log_file"]),
	}

	// Creating new logger
	logger := log.NewLogger(logContainer, log.Normal)

	// Reading data from config.api.json file for acquiring the api information
	apiConfig := make(map[string]interface{})
	apiConfigDir := filepath.Join(pwd, "config/config.api.json")
	apiConfigData, err := ioutil.ReadFile(apiConfigDir)
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to read api configuration file. "+
			"Configuration file: ./config/config.api.json, %s", err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	err = json.Unmarshal(apiConfigData, &apiConfig)
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to unmarshal api configuration data. "+
			"API Config Data => %s, %s", string(apiConfigData), err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	baseURL, ok1 := apiConfig["geo_ip_service_base_url"].(string)
	apiKey, ok2 := apiConfig["geo_ip_service_api_key"].(string)

	if !ok1 || !ok2 {
		logger.Log(fmt.Sprintf("Error: Unable to parse api configuration data. %s",
			err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	os.Setenv("geo_ip_service_base_url", baseURL)
	os.Setenv("geo_ip_service_api_key", apiKey)
	os.Setenv("http_client_server_port", sysConfig.HTTPClientServerPort)
	// setting ssl server public and private key path
	os.Setenv("server.pem", filepath.Join(pwd, "config/server.pem"))
	os.Setenv("server.key", filepath.Join(pwd, "config/server.key"))

	// Checking for ssl server public key or certificate file
	_, err = os.Stat(os.Getenv("server.pem"))
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to find ssl server certificate file. "+
			"Certificate file: ./config/server.pem, %s", err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	// Checking for ssl server private key file
	_, err = os.Stat(os.Getenv("server.key"))
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to find ssl server private key file. "+
			"Private key file: ./config/server.key, %s", err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	// Reading data from repository.store.json file
	store := make(map[string]string)
	storeDir := filepath.Join(pwd, "translate/repository/repository.store.json")
	storeData, err := ioutil.ReadFile(storeDir)
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to read translated 'stuff' store file. "+
			"Store file: ./translate/repository/repository.store.json, %s",
			err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	err = json.Unmarshal(storeData, &store)
	if err != nil {
		logger.Log(fmt.Sprintf("Error: Unable to unmarshal translated 'stuff' store data. "+
			"Translated Data => %s, %s", string(storeData), err.Error()), logger.Logs.ErrorLogFile)
		os.Exit(1)
	}

	// Creating translation service
	translationService := tranService.NewTranslateService(store, logger)

	// Creating Geo IP location service
	geoIPLocationService := geoService.NewGeoIPService(logger)

	// Creating new API handler
	apiHandler := handler.NewAPIHandler(geoIPLocationService, translationService, logger)

	router.HandleFunc("/do_stuff", apiHandler.HandleDoStuff).Methods("GET")
	router.HandleFunc("/do_pretty_cool_stuff", apiHandler.HandleDoPrettyCoolStuff).Methods("GET")
}
