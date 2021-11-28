package main

import (
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/internal/app/samsung-tv-api/helpers"
	samsung_tv_api "github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api"
)

func main() {
	config := helpers.LoadConfiguration()

	c := samsung_tv_api.NewSamsungTvWebSocket("192.168.1.188", config.Token, 8002, 0, 1, "", true)
	device, _ := c.Rest.GetDeviceInfo()
	updatedToken := c.GetToken()

	if updatedToken != "" && updatedToken != config.Token {
		config.Token = c.GetToken()
	}

	config.Mac = device.Device.WifiMac
	_ = helpers.SaveConfiguration(&config)

	result, _ := c.Upnp.GetVolume()
	fmt.Printf("volumn %s", result)
	_ = c.Upnp.SetVolume(16)

	result, _ = c.Upnp.GetVolume()
	fmt.Printf("volumn %s", result)

}
