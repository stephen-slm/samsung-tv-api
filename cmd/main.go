package main

import (
	"encoding/json"
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/samsung-tv-api"
	"os"
)

func main() {
	client := samsung_tv_api.NewSamsungTvWebSocket("192.168.1.188", "", 8001, 0, 1, "")
	resp, err := client.GetDeviceInfo()

	b, _ := json.MarshalIndent(resp, "", "\t")
	os.Stdout.Write(b)

	fmt.Println(resp.Device, err)
}
