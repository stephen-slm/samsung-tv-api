package http

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SamsungRestClient struct {
	BaseUrl func(string) *url.URL
}

// makeRestRequest will send a API http call to the given endpoint (base url + endpoint)
// based on the given method. output will be the binding JSON output of the request.
//
// TODO
// 	* support binding to a non 200 response or determine the error message returned and use it in the error response
func (s *SamsungRestClient) makeRestRequest(endpoint, method string, output interface{}) error {
	u := s.BaseUrl(endpoint).String()

	log.Printf("rest url %s\n", u)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(200) * time.Millisecond,
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

// GetDeviceInfo returns the related Tv information via the rest api
//
// TODO
// 	* This has to been tested with any bad input, should be regarded as not stable.
func (s *SamsungRestClient) GetDeviceInfo() (DeviceResponse, error) {
	log.Println("Get device info via rest api")

	output := DeviceResponse{}
	err := s.makeRestRequest("", "get", &output)

	return output, err
}

// GetApplicationStatus returns basic information regrading a given installed application.
// It will use the provided application id to send via the rest api.
//
// TODO
// 	* This has to been tested with any bad input, should be regarded as not stable.
func (s *SamsungRestClient) GetApplicationStatus(appId string) (ApplicationResponse, error) {
	log.Println("Get application info via rest api")

	var output ApplicationResponse
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "get", &output)

	return output, err
}

// RunApplication will tell the TV via the rest api to run a given application
// by using the provided application id.
//
// TODO
// 	* This has to been tested with any bad input, should be regarded as not stable.
func (s *SamsungRestClient) RunApplication(appId string) (interface{}, error) {
	log.Println("Run application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "post", &output)

	return output, err
}

// CloseApplication will tell the TV via the rest api to close a given application
// by using the provided application id.
//
// TODO
// 	* This requires to be tested, it has not been ran to close any applications yet.
// 	* This has to been tested with any bad input, should be regarded as not stable.
func (s *SamsungRestClient) CloseApplication(appId string) (interface{}, error) {
	log.Println("Close application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "delete", &output)

	return output, err
}

// InstallApplication will tell the TV via the rest api to install a given application
// by using the provided application id.
//
// TODO
// 	* This requires to be tested, it has not been ran to install any applications yet.
// 	* This has to been tested with any bad input, should be regarded as not stable.
func (s *SamsungRestClient) InstallApplication(appId string) (interface{}, error) {
	log.Println("Install application via rest api")

	var output interface{}
	err := s.makeRestRequest(fmt.Sprintf("applications/%s", appId), "PUT", &output)

	return output, err
}
