package handler

import "github.com/Benyam-S/dostaff/geoip"

// APIHandler is a type that defines a handler for api client
type APIHandler struct {
	GeoIPService geoip.IService
}

// APIRespone is a type that defines what the API end points return
type APIRespone struct {
	Status   int // The response status code
	Result   string
	Location string
}

// NewAPIHandler is a function that returns a new API handler
func NewAPIHandler(locationService geoip.IService) *APIHandler {
	return &APIHandler{GeoIPService: locationService}
}
