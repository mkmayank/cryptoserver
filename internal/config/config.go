package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type cfg struct {
	HitBtc HitBtc `json:"hitbtc"`
	Paths  paths  `json:"paths"`
}

var rawConfig cfg

// Init initialized config
var Init = func() func(configFile string) {

	f := false

	return func(configFile string) {

		if f {
			fmt.Println("config is initialized more than once")
			os.Exit(2)
		}

		file, err := os.Open(configFile)
		if err != nil {
			fmt.Printf("Please check config file : %s\n", configFile)
			os.Exit(2)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&rawConfig)
		if err != nil {
			fmt.Printf("Unable to parse config file : %s\n", configFile)
			os.Exit(2)
		}

		checkSanity()
	}
}()

func checkSanity() {
	rawConfig.HitBtc.checkSanity()
	rawConfig.Paths.checkSanity()
}
