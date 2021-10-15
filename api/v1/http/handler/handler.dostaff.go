package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Benyam-S/dostuff/tools"
)

// HandleDoStuff is a handler function that handles request for doing stuff
func (handler *APIHandler) HandleDoStuff(w http.ResponseWriter, r *http.Request) {

	/* ---------------------------- Logging ---------------------------- */
	handler.logger.Log("New request, doing stuff...", handler.logger.Logs.DebugLogFile)

	var status int
	var location string
	var apiRespone *APIRespone
	var translatedStuff string

	// Getting client ip address
	clientIP := tools.GetClientIP(r)

	// Getting the location of the client using its IP address
	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		status = http.StatusBadRequest
		translatedStuff = "Did nothing"
		location = "Oops, we couldn't find your location!"

	} else {

		// Translating stuff to client's location language
		translatedStuff, err = handler.TranslationService.TranslateStuff(geoIPLocation.Country)
		if err != nil {
			status = http.StatusBadRequest
			translatedStuff = "Oops, we couldn't find your location language!"
			location = geoIPLocation.Country

		} else {
			status = http.StatusOK
			location = geoIPLocation.Country

		}

	}

	apiRespone = &APIRespone{
		Status:   status,
		Result:   translatedStuff,
		Location: location,
	}

	result, err := json.MarshalIndent(apiRespone, "", "")
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.logger.Log(fmt.Sprintf("Error: unable to marshal response to client "+
			"{ Client IP : %s, Geo IP Location : %s, Translated Stuff : %s }, %s",
			clientIP, geoIPLocation.ToString(), translatedStuff, err.Error()), handler.logger.Logs.ErrorLogFile)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(apiRespone.Status)
	w.Write([]byte(result))
}
