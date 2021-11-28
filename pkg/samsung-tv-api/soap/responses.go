package soap

type GetVolumeResponse struct {
	Envelope struct {
		Body struct {
			GetVolumeResponse struct {
				CurrentVolume string `json:"CurrentVolume"`
			} `json:"GetVolumeResponse"`
		} `json:"Body"`
	} `json:"Envelope"`
}
