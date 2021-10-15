package entity

import (
	"encoding/json"
	"fmt"
)

// ToString is a method that converts a GeoIPLocation Entry struct to readable JSON string format
func (geoIPLocation *GeoIPLocation) ToString() string {
	output, err := json.Marshal(geoIPLocation)
	if err != nil {
		return fmt.Sprint(geoIPLocation)
	}

	return string(output)
}
