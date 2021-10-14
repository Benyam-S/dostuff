package entity

// GeoIPLocation is a type that holds the Geo IP location values
type GeoIPLocation struct {
	IP        string  `json:"ip"`
	ASN       int64   `json:"asn"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Continent string  `json:"continent"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
}
