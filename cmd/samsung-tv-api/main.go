package main

import (
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/internal/app/samsung-tv-api/helpers"
	samsung_tv_api "github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api"
	"log"
)

func main() {
	config := helpers.LoadConfiguration()

	c := samsung_tv_api.NewSamsungTvWebSocket("192.168.1.188", config.Token, 8002, 0, 1, "", true)
	device, err := c.Rest.GetDeviceInfo()

	if err != nil {
		log.Fatalln(err)
	}

	updatedToken := c.GetToken()

	if updatedToken != "" && updatedToken != config.Token {
		config.Token = c.GetToken()
	}

	config.Mac = device.Device.WifiMac
	_ = helpers.SaveConfiguration(&config)

	muted, err := c.Upnp.GetCurrentMuteStatus()

	if err != nil {
		log.Fatalln(muted)
	}

	fmt.Printf("device mute status is currently: %v\n", muted)

}
