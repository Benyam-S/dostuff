package geoip

import "github.com/Benyam-S/dostaff/entity"

// IService is a interface that defines all the functions/services that geo ip provides
type IService interface {
	GetGeoIPLocation(ip string) (*entity.GeoIPLocation, error)
}
