package main

import (
	"encoding/json"
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

	resp, _ := c.Websocket.GetApplicationsList()

	deviceJson, _ := json.MarshalIndent(device, "", "\t")

	fmt.Printf("getting application details for %s\n", resp.Data.Applications[0].Name)
	appDetails, _ := c.Rest.GetApplicationStatus(resp.Data.Applications[0].AppID)

	appJson, _ := json.MarshalIndent(appDetails, "", "\t")
	fmt.Println(string(deviceJson))
	fmt.Println(string(appJson))
}
