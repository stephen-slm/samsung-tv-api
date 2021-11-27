package samsung_tv_api

import (
	"fmt"
	"net/url"
)

// formatWebSocketUrl returns the formatted web socket url for connecting
func (s *SamsungTvClient) formatWebSocketUrl(endpoint string) *url.URL {
	u := &url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprintf("%s:%d", s.host, s.port),
		Path:     fmt.Sprintf("api/v2/channels/%s", endpoint),
		RawQuery: fmt.Sprintf("name=%s", s.name),
	}

	if s.isSslConnection() {
		u.Scheme += "s"
		u.RawQuery += fmt.Sprintf("&token=%s", s.token)
		return u
	}

	return u
}

// formatRestUrl returns the formatted rest api url for connecting to
// the tv rest service
func (s *SamsungTvClient) formatRestUrl(endpoint string) *url.URL {
	u := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.host, s.port),
		Path:   fmt.Sprintf("api/v2/%s", endpoint),
	}

	if s.isSslConnection() {
		u.Scheme += "s"
	}

	return u
}
