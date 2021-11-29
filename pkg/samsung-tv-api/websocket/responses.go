package websocket

type ConnectionResponse struct {
	Data struct {
		Clients []struct {
			Attributes struct {
				Name interface{} `json:"name"`
			} `json:"attributes"`
			ConnectTime int64  `json:"connectTime"`
			DeviceName  string `json:"deviceName"`
			ID          string `json:"id"`
			IsHost      bool   `json:"isHost"`
		} `json:"clients"`
		ID    string `json:"id"`
		Token string `json:"token"`
	} `json:"data"`
	Event string `json:"event"`
}

type ApplicationsResponse struct {
	Data struct {
		Applications []struct {
			AppID   string `json:"appId"`
			AppType int    `json:"app_type"`
			Icon    string `json:"icon"`
			IsLock  int    `json:"is_lock"`
			Name    string `json:"name"`
		} `json:"data"`
	} `json:"data"`
	Event string `json:"event"`
	From  string `json:"from"`
}
