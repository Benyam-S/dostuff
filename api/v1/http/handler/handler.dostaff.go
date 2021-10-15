package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Benyam-S/dostuff/tools"
)

// HandleDoStuff is a handler function that handles request for doing stuff
func (handler *APIHandler) HandleDoStuff(w http.ResponseWriter, r *http.Request) {

	var apiRespone *APIRespone

	// Getting client ip address
	clientIP := tools.GetClientIP(r)

	// Getting the location of the client using its IP address
	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		apiRespone = &APIRespone{
			Status:   http.StatusBadRequest,
			Result:   "Did nothing",
			Location: "Oops, we couldn't find your location!",
		}
	} else {

		// Translating stuff to client's location language
		translatedStuff, err := handler.TranslationService.TranslateStuff(geoIPLocation.Country)
		if err != nil {
			apiRespone = &APIRespone{
				Status:   http.StatusOK,
				Result:   "Oops, we couldn't find your location language!",
				Location: geoIPLocation.Country,
			}
		} else {
			apiRespone = &APIRespone{
				Status:   http.StatusOK,
				Result:   translatedStuff,
				Location: geoIPLocation.Country,
			}
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
