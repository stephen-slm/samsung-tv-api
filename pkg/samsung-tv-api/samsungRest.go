package samsung_tv_api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type samsungRestClient struct {
	baseUrl *url.URL
}

func (s *samsungRestClient) makeRestRequest(endpoint, method string, output interface{}) error {
	u := fmt.Sprintf("%s/%s", s.baseUrl.String(), endpoint)

	fmt.Println(u)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(strings.ToUpper(method), u, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, clientErr := client.Do(req)

	if clientErr != nil {
		return clientErr
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(&output)
}

// GetDeviceInfo returns the related Tv information via the rest api.
func (s *samsungRestClient) GetDeviceInfo() (RestDeviceResponse, error) {
	log.Println("Get device info via rest api")

	output := RestDeviceResponse{}
	err := s.makeRestRequest("", "get", &output)

	return output, err
}

func (s *samsungRestClient) GetApplicationStatus(appId string) (RestApplicationResponse, error) {
	log.Println("Get application info via rest api")

	var output RestApplicationResponse
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "get", &output)

	return output, err
}

func (s *samsungRestClient) RunApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "post", &output)

	return output, err
}
func (s *samsungRestClient) CloseApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "delete", &output)

	return output, err
}

func (s *samsungRestClient) InstallApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "PUT", &output)

	return output, err
}
