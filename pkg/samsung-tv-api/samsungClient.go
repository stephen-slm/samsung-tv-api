package samsung_tv_api

import (
	"log"
)

type SamsungTvClient struct {
	Rest      samsungRestClient
	Websocket samsungWebsocket

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

	client.Rest = samsungRestClient{
		baseUrl: client.formatRestUrl(""),
	}

	client.Websocket = samsungWebsocket{
		baseUrl:       client.formatWebSocketUrl("samsung.remote.control"),
		keyPressDelay: keyPressDelay,
	}

	wsResp, err := client.Websocket.openConnection()

	if err != nil {
		log.Fatalln(err)
	}

	if len(wsResp.Data.Clients) > 0 && wsResp.Data.Clients[0].Attributes.Token != "" {
		client.token = wsResp.Data.Clients[0].Attributes.Token
	}

	return client
}

// isSslConnection returns true if and only if the port is the SSL port for the
// connection otherwise it is not configured for SSL.
func (s *SamsungTvClient) isSslConnection() bool {
	return s.port == 8002
}
