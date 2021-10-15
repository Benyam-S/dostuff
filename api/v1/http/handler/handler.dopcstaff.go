package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Benyam-S/dostuff/tools"
)

// HandleDoPrettyCoolStuff is a handler function that handles request for doing pretty cool stuff
func (handler *APIHandler) HandleDoPrettyCoolStuff(w http.ResponseWriter, r *http.Request) {

	var apiRespone *APIRespone

	// Getting client ip address
	clientIP := tools.GetClientIP(r)

	// Getting the location of the client using its IP address
	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		apiRespone = &APIRespone{
			Status:   http.StatusBadRequest,
			Result:   "Sorry! we let you down",
			Location: "We couldn't find your location",
		}
	} else {
		apiRespone = &APIRespone{
			Status: http.StatusOK,
			Result: "Thank You for using this api. \n" +
				"If you are intersted in doing pretty cool stuff check out \n" +
				"the repository https://github.com/Benyam-S/dostuff. \n" +
				"Let as collaberate in doing pretty cool stuff!",
			Location: geoIPLocation.Country,
		}
	}

	result, err := json.MarshalIndent(apiRespone, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(apiRespone.Status)
	w.Write([]byte(result))
}
