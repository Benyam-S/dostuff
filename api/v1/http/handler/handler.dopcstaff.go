package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Benyam-S/dostaff/tools"
)

// HandleDoPrettyCoolStaff is a handler function that handles request for doing pretty cool staff
func (handler *APIHandler) HandleDoPrettyCoolStaff(w http.ResponseWriter, r *http.Request) {

	var apiRespone *APIRespone

	clientIP := tools.GetClientIP(r)
	// clientIP := "197.156.95.233"

	geoIPLocation, err := handler.GeoIPService.GetGeoIPLocation(clientIP)
	if err != nil {
		apiRespone = &APIRespone{
			Status:   http.StatusBadRequest,
			Result:   "Sorry we let you down",
			Location: "Oops, we couldn't find your location!",
		}
	} else {
		apiRespone = &APIRespone{
			Status:   http.StatusOK,
			Result:   "Thank You for visiting this api. If you are intersted in doing pretty cool staff check out the repository https://github.com/Benyam-S/dostaff.",
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
