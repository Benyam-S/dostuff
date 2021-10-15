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
)

// GeoIPService is a type that defines Geo IP service
type GeoIPService struct {
}

// NewURLService is a function that returns a new Geo IP service
func NewGeoIPService() geoip.IService {
	return &GeoIPService{}
}

// GetGeoIPLocation is a function that return the geo ip location of a give ip address
func (service *GeoIPService) GetGeoIPLocation(ip string) (*entity.GeoIPLocation, error) {

	geoIPLocation := new(entity.GeoIPLocation)
	data, _ := json.Marshal(map[string]string{"ip": ip})
	client := &http.Client{}

	req, err := http.NewRequest("POST", os.Getenv("geo_ip_service_base_url"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("geo_ip_service_api_key")))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, geoIPLocation)
	if err != nil {
		return nil, err
	}

	if geoIPLocation.IP != ip || geoIPLocation.Country == "" {
		return nil, errors.New("unable to find location of the ip address")
	}

	return geoIPLocation, nil
}
