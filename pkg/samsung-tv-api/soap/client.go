package soap

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	xj "github.com/basgys/goxml2json"
)

// This package covers the support for the Universal Plug & Play (UPNP)

type SamsungSoapClient struct {
	BaseUrl *url.URL
}

func (s *SamsungSoapClient) MakeRestRequest(action, arguments, protocol string, output interface{}) error {
	u := fmt.Sprintf("%s%s1", s.BaseUrl.String(), protocol)

	body := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"+
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n"+
		"<s:Body>\n"+
		"<u:%s xmlns:u=\"urn:schemas-upnp-org:service:%s:1\">\n"+
		"<InstanceID>0</InstanceID>\n"+
		"%s\n"+
		"</u:%s>\n"+
		"</s:Body>\n"+
		"</s:Envelope>", action, protocol, arguments, action)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Duration(200) * time.Millisecond,
	}

	req, err := http.NewRequest("POST", u, strings.NewReader(body))
	req.Header.Set("SOAPAction", fmt.Sprintf("\"urn:schemas-upnp-org:service:%s:1#%s\"", protocol, action))

	if err != nil {
		return err
	}

	resp, clientErr := client.Do(req)

	if clientErr != nil {
		return clientErr
	}

	defer resp.Body.Close()

	content, convertErr := xj.Convert(resp.Body)

	fmt.Printf(content.String())

	if convertErr != nil {
		return convertErr
	}

	return json.Unmarshal(content.Bytes(), &output)
}

func (s *SamsungSoapClient) GetVolume() (string, error) {
	log.Println("Get device info via saop api")

	output := GetVolumeResponse{}
	err := s.MakeRestRequest("GetVolume", "<Channel>Master</Channel>", "RenderingControl", &output)

	return output.Envelope.Body.GetVolumeResponse.CurrentVolume, err
}

func (s *SamsungSoapClient) SetVolume(volume int) error {
	log.Println("Get device info via soap api")

	var output interface{}

	args := fmt.Sprintf("<Channel>Master</Channel><DesiredVolume>%d</DesiredVolume>", volume)
	return s.MakeRestRequest("SetVolume", args, "RenderingControl", &output)
}
