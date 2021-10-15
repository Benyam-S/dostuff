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

// SystemConfig is a type that defines a server system configuration file
type SystemConfig struct {
	HTTPDomainAddress    string            `json:"http_domain_address"`
	HTTPClientServerPort string            `json:"http_client_server_port"`
	LogsPath             string            `json:"logs_path"`
	Logs                 map[string]string `json:"logs"`
}
