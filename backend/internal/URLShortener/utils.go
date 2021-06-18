package URLShortener

import (
	"net"
	"net/url"
	"strings"
)

// isURL checks the correctness of the URL
func isURL(str string) bool {
	URL, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	address := net.ParseIP(URL.Host)
	if address == nil {
		return strings.Contains(URL.Host, ".")
	}
	return true
}
