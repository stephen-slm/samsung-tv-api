package samsung_tv_api

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSslTestClient() *SamsungTvClient {
	return NewSamsungTvWebSocket(
		"2.2.2.2",
		"",
		8002,
		1,
		"ssl.client",
		false)
}

func getTestClient() *SamsungTvClient {
	return NewSamsungTvWebSocket(
		"1.1.1.1",
		"",
		8001,
		1,
		"standard.client",
		false)
}

func TestFormatWebSocketUrl(t *testing.T) {
	client := getTestClient()

	url := client.formatWebSocketUrl("standard.client").String()
	name := base64.StdEncoding.EncodeToString([]byte("standard.client"))

	expected := fmt.Sprintf("ws://1.1.1.1:8001/api/v2/channels/standard.client?name=%s", name)

	assert.Equal(t, expected, url)
}

func TestSslFormatWebSocketUrl(t *testing.T) {
	client := getSslTestClient()

	url := client.formatWebSocketUrl("ssl.client").String()
	name := base64.StdEncoding.EncodeToString([]byte("ssl.client"))

	expected := fmt.Sprintf("wss://2.2.2.2:8002/api/v2/channels/ssl.client?name=%s&token=", name)

	assert.Equal(t, expected, url)
}

func TestFormatRestUrl(t *testing.T) {
	client := getTestClient()

	url := client.formatRestUrl("standard.endpoint").String()
	expected := "http://1.1.1.1:8001/api/v2/standard.endpoint/"

	assert.Equal(t, expected, url)
}

func TestSslFormatRestUrl(t *testing.T) {
	client := getSslTestClient()

	url := client.formatRestUrl("ssl.endpoint").String()
	expected := "https://2.2.2.2:8002/api/v2/ssl.endpoint/"

	assert.Equal(t, expected, url)
}
