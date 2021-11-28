package main

import (
	"encoding/json"
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api"
)

func main() {
	c := samsung_tv_api.NewSamsungTvWebSocket("192.168.1.188", "13992102", 8002, 0, 1, "")
	resp, _ := c.Websocket.GetApplicationsList()

	device, _ := c.Rest.GetDeviceInfo()
	deviceJson, _ := json.MarshalIndent(device, "", "\t")

	fmt.Printf("getting application details for %s\n", resp.Data.Applications[0].Name)
	appDetails, _ := c.Rest.GetApplicationStatus(resp.Data.Applications[0].AppID)

	appJson, _ := json.MarshalIndent(appDetails, "", "\t")
	fmt.Println(string(deviceJson))
	fmt.Println(string(appJson))

	c.Websocket.PowerOff()

	//fmt.Printf("runing application %s\n", resp.Data.Applications[0].Name)
	//runErr := c.Websocket.RunApplication(resp.Data.Applications[0].AppID, "", "")
}
