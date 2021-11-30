package samsung_tv_api

import (
	"encoding/base64"
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api/http"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api/soap"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api/websocket"
	"log"
	"net/url"
)

type SamsungTvClient struct {
	Rest      http.SamsungRestClient
	Websocket websocket.SamsungWebsocket
	Upnp      soap.SamsungSoapClient

	host          string
	token         string
	port          int
	keyPressDelay int
	name          string
}

func NewSamsungTvWebSocket(host, token string, port, keyPressDelay int, name string, autoConnect bool) *SamsungTvClient {
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
		keyPressDelay: keyPressDelay,
		name:          name,
	}

	client.Rest = http.SamsungRestClient{
		BaseUrl: func(endpoint string) *url.URL {
			return client.formatRestUrl(endpoint)
		},
	}

	client.Websocket = websocket.SamsungWebsocket{
		BaseUrl: func(endpoint string) *url.URL {
			return client.formatWebSocketUrl(endpoint)
		},
		KeyPressDelay: keyPressDelay,
	}

	client.Upnp = soap.SamsungSoapClient{
		BaseUrl: func(endpoint string) *url.URL {
			return client.formatUpnpUrl(endpoint)
		},
	}

	if autoConnect {
		if err := client.ConnectionSetup(); err != nil {
			log.Fatalln(err)
		}
	}

	return client
}

// ConnectionSetup will attempt to open a connection to the websocket endpoint on
// the TV while after connecting, update the internal token to the newest value
// regardless if its the same.
func (s *SamsungTvClient) ConnectionSetup() error {
	wsResp, err := s.Websocket.OpenConnection()

	if err != nil {
		return err
	}

	if len(wsResp.Data.Clients) > 0 && wsResp.Data.Token != "" {
		s.token = wsResp.Data.Token
	}

	return nil
}

// isSslConnection returns true if and only if the port is the SSL port for the
// connection otherwise it is not configured for SSL.
func (s *SamsungTvClient) isSslConnection() bool {
	return s.port == 8002
}

// formatWebSocketUrl returns the formatted web socket url for connecting
func (s *SamsungTvClient) formatWebSocketUrl(endpoint string) *url.URL {
	if endpoint != "" && string(endpoint[0]) != "/" {
		endpoint = "/" + endpoint
	}

	name := base64.StdEncoding.EncodeToString([]byte(s.name))

	u := &url.URL{
		Scheme:   "ws",
		Host:     fmt.Sprintf("%s:%d", s.host, s.port),
		Path:     fmt.Sprintf("api/v2/channels%s", endpoint),
		RawQuery: fmt.Sprintf("?name=%s", name),
	}

	if s.isSslConnection() {
		u.Scheme += "s"
		u.RawQuery += fmt.Sprintf("&token=%s", s.token)
	}

	return u
}

// formatRestUrl returns the formatted rest api url for connecting to
// the tv rest service
func (s *SamsungTvClient) formatRestUrl(endpoint string) *url.URL {
	if endpoint != "" && string(endpoint[0]) != "/" {
		endpoint = "/" + endpoint
	}

	if endpoint == "" || string(endpoint[len(endpoint)-1]) != "/" {
		endpoint = "/"
	}

	u := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.host, s.port),
		Path:   fmt.Sprintf("api/v2%s", endpoint),
	}

	if s.isSslConnection() {
		u.Scheme += "s"
	}

	return u
}

// formatUpnpUrl returns the formatted api url for connecting to
// the tv soap service
func (s *SamsungTvClient) formatUpnpUrl(endpoint string) *url.URL {
	if endpoint != "" && string(endpoint[0]) != "/" {
		endpoint = "/" + endpoint
	}

	if endpoint == "" || string(endpoint[len(endpoint)-1]) != "/" {
		endpoint = "/"
	}

	u := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.host, 9197),
		Path:   fmt.Sprintf("upnp/control%s", endpoint),
	}

	if s.isSslConnection() {
		u.Scheme += "s"
		return u
	}

	return u
}

func (s *SamsungTvClient) Disconnect() error {
	return s.Websocket.Disconnect()
}

// GetToken returns the current Auth token used by the client.
func (s *SamsungTvClient) GetToken() string {
	return s.token
}
