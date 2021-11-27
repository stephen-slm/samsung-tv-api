package main

import (
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api"
)

func main() {
	c := samsung_tv_api.NewSamsungTvWebSocket("192.168.1.188", "13992102", 8002, 0, 1, "")
	resp, _ := c.Websocket.GetApplicationsList()

	fmt.Printf("getting application details for %s\n", resp.Data.Applications[0].Name)
	appDetails, err := c.Rest.GetApplicationStatus(resp.Data.Applications[0].AppID)

	fmt.Println(appDetails, err)

	fmt.Printf("runing application %s\n", resp.Data.Applications[0].Name)
	runErr := c.Websocket.RunApplication(resp.Data.Applications[0].AppID, "", "")

	fmt.Println(runErr)

}
