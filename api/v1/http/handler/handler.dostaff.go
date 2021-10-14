package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Benyam-S/dostaff/tools"
)

// HandleCoolStaff is a handler function that handles request for doing staff
func (handler *APIHandler) HandleDoStaff(w http.ResponseWriter, r *http.Request) {

	var apiRespone *APIRespone

	clientIP := tools.GetClientIP(r)
	// clientIP := "197.156.95.233"

	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		apiRespone = &APIRespone{
			Status:   http.StatusBadRequest,
			Result:   "Didn't Do Staff",
			Location: "Oops, we couldn't find your location!",
		}
	} else {
		apiRespone = &APIRespone{
			Status:   http.StatusOK,
			Result:   "Did Staff",
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
