package handler

import (
	"github.com/Benyam-S/dostuff/geoip"
	"github.com/Benyam-S/dostuff/translate"
)

// APIHandler is a type that defines a handler for api client
type APIHandler struct {
	GeoIPService       geoip.IService
	TranslationService translate.IService
}

// APIRespone is a type that defines what the API end points return
type APIRespone struct {
	Status   int // The response status code
	Result   string
	Location string
}

// NewAPIHandler is a function that returns a new API handler
func NewAPIHandler(locationService geoip.IService, translationService translate.IService) *APIHandler {
	return &APIHandler{GeoIPService: locationService, TranslationService: translationService}
}
