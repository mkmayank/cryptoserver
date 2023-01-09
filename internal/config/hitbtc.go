package config

// HitBtc config
type HitBtc struct {
	RestAPI       string `json:"rest_api"`
	TickerChannel string `json:"ticker_channel"`
}

func (h HitBtc) checkSanity() {

	// todo
	// this can be used to check if all the config values are fine
}
