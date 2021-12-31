package soap

type GetDeviceVolumeResponse struct {
	Envelope struct {
		Body struct {
			GetVolumeResponse struct {
				CurrentVolume string `json:"CurrentVolume"`
			} `json:"GetVolumeResponse"`
		} `json:"Body"`
	} `json:"Envelope"`
}

type GetCurrentMuteStatusResponse struct {
	Envelope struct {
		Body struct {
			GetMuteResponse struct {
				CurrentMute string `json:"CurrentMute"`
			} `json:"GetMuteResponse"`
		} `json:"Body"`
	} `json:"Envelope"`
}
