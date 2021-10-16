package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Benyam-S/dostuff/tools"
)

// HandleDoPrettyCoolStuff is a handler function that handles request for doing pretty cool stuff
func (handler *APIHandler) HandleDoPrettyCoolStuff(w http.ResponseWriter, r *http.Request) {

	/* ---------------------------- Logging ---------------------------- */
	handler.logger.Log("New request, doing pretty cool stuff...", handler.logger.Logs.DebugLogFile)

	var status int
	var location string
	var apiRespone *APIRespone
	var result string

	// Getting client ip address
	clientIP := tools.GetClientIP(r)

	// Getting the location of the client using its IP address
	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		status = http.StatusBadRequest
		result = "Sorry! we let you down"
		location = "We couldn't find your location"

	} else {
		status = http.StatusOK
		result = "Thank You for using this api. " +
			"If you are intersted in doing pretty cool stuff check out " +
			"the repository https://github.com/Benyam-S/dostuff. " +
			"Let as collaberate in doing pretty cool stuff!"
		location = geoIPLocation.Country

	}

	apiRespone = &APIRespone{
		Status:   status,
		Result:   result,
		Location: location,
	}

	response, err := json.MarshalIndent(apiRespone, "", "")
	if err != nil {
		/* ---------------------------- Logging ---------------------------- */
		handler.logger.Log(fmt.Sprintf("Error: unable to marshal response to client "+
			"{ Client IP : %s, Geo IP Location : %s, Result : %s }, %s",
			clientIP, geoIPLocation.ToString(), result, err.Error()), handler.logger.Logs.ErrorLogFile)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(apiRespone.Status)
	w.Write([]byte(response))
}
