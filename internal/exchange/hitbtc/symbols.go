package hitbtc

import (
	"bytes"
	"cc/internal/config"
	"cc/internal/data"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetSymbols make HTTP GET request to /public/symbol and return symbol info
func (h *HitBtc) GetSymbols() (symbols map[string]data.SymbolInfo, err error) {

	currencyInfo, err := h.getCurrencyInfo()
	if err != nil {
		log.Error(fmt.Sprintf("get symbols, currency fetch error: %s", err.Error()))
		err = nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/public/symbol", config.HitBtcConfig().RestAPI), nil)
	if err != nil {
		err = fmt.Errorf("get symbols new request error: %s", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("get symbols making request error: %s", err.Error())
		return
	}

	defer res.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(res.Body, &buf)

	symbolInfo := map[string]symbolInfo{}

	err = json.NewDecoder(tee).Decode(&symbolInfo)
	if err != nil {
		var bodyBytes []byte
		bodyBytes, err = ioutil.ReadAll(&buf)
		if err != nil {
			err = fmt.Errorf("get symbols dumping errored getSymbols response, error: %s", err.Error())
			return
		}

		err = fmt.Errorf("get symbols decode error: %s, response: %s", err.Error(), string(bodyBytes))
		return
	}

	symbols = map[string]data.SymbolInfo{}
	for symbol, info := range symbolInfo {

		fullName := "unknown"
		if _, ok := currencyInfo[info.BaseCurrency]; ok {
			fullName = currencyInfo[info.BaseCurrency].FullName

		}
		symbols[symbol] = data.SymbolInfo{
			FullName:     fullName,
			BaseCurrency: info.BaseCurrency,
			FeeCurrency:  info.FeeCurrency,
		}
	}

	return
}

// getCurrencyInfo returns currencyInfo map
func (h *HitBtc) getCurrencyInfo() (currencyInfo map[string]currencyInfo, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/public/currency", config.HitBtcConfig().RestAPI), nil)
	if err != nil {
		err = fmt.Errorf("get currency new request error: %s", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("get currency making request error: %s", err.Error())
		return
	}

	defer res.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(res.Body, &buf)

	err = json.NewDecoder(tee).Decode(&currencyInfo)
	if err != nil {
		var bodyBytes []byte
		bodyBytes, err = ioutil.ReadAll(&buf)
		if err != nil {
			err = fmt.Errorf("get currency dumping errored currency response, error: %s", err.Error())
			return
		}

		err = fmt.Errorf("get currency decode error: %s, response: %s", err.Error(), string(bodyBytes))
		return
	}

	return
}
