package config

import (
	"fmt"
	"os"
)

// HitBtc config
type HitBtc struct {
	RestAPI   string `json:"rest_api"`
	WsAPIHost string `json:"ws_api_host"`
	WsAPIPath string `json:"ws_api_path"`
}

func (h HitBtc) checkSanity() {

	if h.RestAPI == "" {
		fmt.Println("hitbtc.rest_api is not defined in given config file")
		os.Exit(2)
	}

	if h.WsAPIHost == "" {
		fmt.Println("hitbtc.ws_api_host is not defined in given config file")
		os.Exit(2)
	}

	if h.WsAPIPath == "" {
		fmt.Println("hitbtc.ws_api_path is not defined in given config file")
		os.Exit(2)
	}
}
