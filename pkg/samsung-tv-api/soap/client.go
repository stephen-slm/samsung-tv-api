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

func (s *SamsungSoapClient) MakeSoapRequest(action, arguments, protocol string, output interface{}) error {
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

// GetCurrentVolume returns the value of the current volume level
func (s *SamsungSoapClient) GetCurrentVolume() (string, error) {
	log.Println("Get device volume via saop api")

	output := GetDeviceVolumeResponse{}
	err := s.MakeSoapRequest("GetCurrentVolume", "<Channel>Master</Channel>", "RenderingControl", &output)

	return output.Envelope.Body.GetVolumeResponse.CurrentVolume, err
}

// SetVolume will update the current volume of the display to the provided value.
func (s *SamsungSoapClient) SetVolume(volume int) error {
	log.Printf("set the volume of the tv to %d via soap api\n", volume)

	var output interface{}

	args := fmt.Sprintf("<Channel>Master</Channel><DesiredVolume>%d</DesiredVolume>", volume)
	return s.MakeSoapRequest("SetVolume", args, "RenderingControl", &output)
}

// GetCurrentMuteStatus returns true if and only if the TV is currently muted
func (s *SamsungSoapClient) GetCurrentMuteStatus() (bool, error) {
	log.Println("Get device mute status via saop api")

	output := GetCurrentMuteStatusResponse{}
	err := s.MakeSoapRequest("GetMute", "<Channel>Master</Channel>", "RenderingControl", &output)

	if err != nil {
		return false, err
	}

	return output.Envelope.Body.GetMuteResponse.CurrentMute == "1", err
}

// SetCurrentMedia will tell the display to play the current media via the URL.
func (s *SamsungSoapClient) SetCurrentMedia(url string) error {
	args := fmt.Sprintf("<CurrentURI>%s</CurrentURI><CurrentURIMetaData></CurrentURIMetaData>", url)

	var output interface{}
	var err error

	err = s.MakeSoapRequest("SetAVTransportURI", args, "AVTransport", &output)

	if err != nil {
		return err
	}

	return s.PlayCurrentMedia()
}

// PlayCurrentMedia will attempt to play the current media already set on the display.
func (s *SamsungSoapClient) PlayCurrentMedia() error {
	var output interface{}
	return s.MakeSoapRequest("Play", "<Speed>1</Speed>", "AVTransport", &output)
}
