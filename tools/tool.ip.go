package tools

import (
	"net/http"
	"strings"
)

// GetClientIP the IP address of the request sender
func GetClientIP(r *http.Request) string {
	var ip string
	if len(r.Header.Get("CF-Connecting-IP")) > 1 {
		ip = r.Header.Get("CF-Connecting-IP")
	} else if len(r.Header.Get("X-Forwarded-For")) > 1 {
		ip = r.Header.Get("X-Forwarded-For")
	} else if len(r.Header.Get("X-Real-IP")) > 1 {
		ip = r.Header.Get("X-Real-IP")
	} else {
		ip = r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip = strings.Split(ip, ":")[0]
		}
	}

	return ip
}
