package samsung_tv_api

import (
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
	timeout       int
	keyPressDelay int
	name          string
}

func NewSamsungTvWebSocket(host, token string, port, timeout, keyPressDelay int, name string, autoConnect bool) *SamsungTvClient {
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

	client.Rest = http.SamsungRestClient{
		BaseUrl: client.formatRestUrl(""),
	}

	client.Websocket = websocket.SamsungWebsocket{
		BaseUrl:       client.formatWebSocketUrl("samsung.remote.control"),
		KeyPressDelay: keyPressDelay,
	}

	client.Upnp = soap.SamsungSoapClient{
		BaseUrl: client.formatUpnpUrl(""),
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

	if len(wsResp.Data.Clients) > 0 && wsResp.Data.Clients[0].Attributes.Token != "" {
		s.token = wsResp.Data.Clients[0].Attributes.Token
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

// formatUpnpUrl returns the formatted api url for connecting to
// the tv soap service
func (s *SamsungTvClient) formatUpnpUrl(endpoint string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.host, 9197),
		Path:   fmt.Sprintf("upnp/control/%s", endpoint),
	}
}

// GetToken returns the current Auth token used by the client.
func (s *SamsungTvClient) GetToken() string {
	return s.token
}
