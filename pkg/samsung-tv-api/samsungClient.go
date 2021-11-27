package samsung_tv_api

import (
	"fmt"
	"log"
)

type SamsungTvClient struct {
	rest samsungRestClient

	host          string
	token         string
	port          int
	timeout       int
	keyPressDelay int
	name          string
}

func NewSamsungTvWebSocket(host, token string, port, timeout, keyPressDelay int, name string) *SamsungTvClient {
	if keyPressDelay == 0 {
		keyPressDelay = 1
	}

	if name == "" {
		name = "SamsungTvRemote"

	}

	client := &SamsungTvClient{
		host:          host,
		token:         token,
		port:          port,
		timeout:       timeout,
		keyPressDelay: keyPressDelay,
		name:          name,
	}

	client.rest = samsungRestClient{
		baseUrl: client.formatRestUrl(""),
	}

	return client
}

// isSslConnection returns true if and only if the port is the SSL port for the
// connection otherwise it is not configured for SSL.
func (s *SamsungTvClient) isSslConnection() bool {
	return s.port == 8002
}

// GetDeviceInfo returns the related Tv information via the rest api.
func (s *SamsungTvClient) GetDeviceInfo() (DeviceResponse, error) {
	log.Println("Get device info via rest api")

	output := DeviceResponse{}
	err := s.rest.makeRestRequest("", "get", &output)

	return output, err
}

func (s *SamsungTvClient) getApplicationStatus(appId string) (interface{}, error) {
	log.Println("Get application info via rest api")

	var output interface{}
	err := s.rest.makeRestRequest(fmt.Sprintf("applications/%s", appId), "get", &output)

	return output, err
}

func (s *SamsungTvClient) runApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.rest.makeRestRequest(fmt.Sprintf("applications/%s", appId), "post", &output)

	return output, err
}
func (s *SamsungTvClient) closeApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.rest.makeRestRequest(fmt.Sprintf("applications/%s", appId), "delete", &output)

	return output, err
}

func (s *SamsungTvClient) installApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.rest.makeRestRequest(fmt.Sprintf("applications/%s", appId), "PUT", &output)

	return output, err
}
