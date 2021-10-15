package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Benyam-S/dostuff/entity"
	"github.com/Benyam-S/dostuff/geoip"
	"github.com/Benyam-S/dostuff/log"
)

// GeoIPService is a type that defines Geo IP service
type GeoIPService struct {
	logger *log.Logger
}

// NewURLService is a function that returns a new Geo IP service
func NewGeoIPService(logger *log.Logger) geoip.IService {
	return &GeoIPService{logger: logger}
}

// GetGeoIPLocation is a function that return the geo ip location of a give ip address
func (service *GeoIPService) GetGeoIPLocation(ip string) (*entity.GeoIPLocation, error) {

	/* ---------------------------- Logging ---------------------------- */
	service.logger.Log(fmt.Sprintf("Started getting client location { Client IP : %s }",
		ip), service.logger.Logs.DebugLogFile)

	geoIPLocation := new(entity.GeoIPLocation)
	data, _ := json.Marshal(map[string]string{"ip": ip})
	client := &http.Client{}

	req, err := http.NewRequest("POST", os.Getenv("geo_ip_service_base_url"), bytes.NewBuffer(data))
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to construct request "+
			"{ Request URL : %s, Client IP : %s }, %s", os.Getenv("geo_ip_service_base_url"), ip, err.Error()),
			service.logger.Logs.ErrorLogFile)

		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("geo_ip_service_api_key")))

	res, err := client.Do(req)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to make request to Geo IP Location service "+
			"{ Request URL : %s, Client IP : %s }, %s", os.Getenv("geo_ip_service_base_url"), ip, err.Error()),
			service.logger.Logs.ErrorLogFile)

		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to read response from Geo IP Location service "+
			"{ Request URL : %s, Client IP : %s }, %s", os.Getenv("geo_ip_service_base_url"), ip, err.Error()),
			service.logger.Logs.ErrorLogFile)

		return nil, err
	}

	err = json.Unmarshal(body, geoIPLocation)
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to unmarshal response from Geo IP Location service "+
			"{ Response : %s }, %s", string(body), err.Error()), service.logger.Logs.ErrorLogFile)

		return nil, err
	}

	if geoIPLocation.IP != ip || geoIPLocation.Country == "" {
		/* ---------------------------- Logging ---------------------------- */
		service.logger.Log(fmt.Sprintf("Error: unable to determine the location from Geo IP Location service response, "+
			"Geo IP Location => %s", geoIPLocation.ToString()), service.logger.Logs.ErrorLogFile)

		return nil, errors.New("unable to find location of the ip address")
	}

	/* ---------------------------- Logging ---------------------------- */
	service.logger.Log(fmt.Sprintf("Finished getting client location, Geo IP Location => %s",
		geoIPLocation.ToString()), service.logger.Logs.DebugLogFile)

	return geoIPLocation, nil
}
