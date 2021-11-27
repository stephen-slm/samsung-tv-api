package samsung_tv_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type samsungRestClient struct {
	baseUrl *url.URL
}

func (s *samsungRestClient) makeRestRequest(endpoint, method string, output interface{}) error {
	u := fmt.Sprintf("%s/%s", s.baseUrl.String(), endpoint)

	client := &http.Client{}

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
