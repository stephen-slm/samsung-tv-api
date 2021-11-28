package samsung_tv_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSslTestClient() *SamsungTvClient {
	return NewSamsungTvWebSocket(
		"2.2.2.2",
		"",
		8002,
		0,
		1,
		"ssl.client",
		false)
}

func getTestClient() *SamsungTvClient {
	return NewSamsungTvWebSocket(
		"1.1.1.1",
		"",
		8001,
		0,
		1,
		"standard.client",
		false)
}

func TestFormatWebSocketUrl(t *testing.T) {
	client := getTestClient()

	url := client.formatWebSocketUrl("standard.client").String()
	expected := "ws://1.1.1.1:8001/api/v2/channels/standard.client?name=standard.client"

	assert.Equal(t, expected, url)
}

func TestSslFormatWebSocketUrl(t *testing.T) {
	client := getSslTestClient()

	url := client.formatWebSocketUrl("ssl.client").String()
	expected := "wss://2.2.2.2:8002/api/v2/channels/ssl.client?name=ssl.client&token="

	assert.Equal(t, expected, url)
}

func TestFormatRestUrl(t *testing.T) {
	client := getTestClient()

	url := client.formatRestUrl("standard.endpoint").String()
	expected := "http://1.1.1.1:8001/api/v2/standard.endpoint"

	assert.Equal(t, expected, url)
}

func TestSslFormatRestUrl(t *testing.T) {
	client := getSslTestClient()

	url := client.formatRestUrl("ssl.endpoint").String()
	expected := "https://2.2.2.2:8002/api/v2/ssl.endpoint"

	assert.Equal(t, expected, url)
}
