package main

import (
	"github.com/stephensli/samsung-tv-api/internal/app/samsung-tv-api/helpers"
	samsung_tv_api "github.com/stephensli/samsung-tv-api/pkg/samsung-tv-api"
)

func main() {
	config := helpers.LoadConfiguration()

	c := samsung_tv_api.NewSamsungTvWebSocket(
		"192.168.1.188",
		config.Token,
		8002,
		5,
		"stephenLaptop",
		true)

	device, _ := c.Rest.GetDeviceInfo()
	updatedToken := c.GetToken()

	if updatedToken != "" && updatedToken != config.Token {
		config.Token = updatedToken
	}

	config.Mac = device.Device.WifiMac
	_ = helpers.SaveConfiguration(&config)

	_ = c.Disconnect()
}
