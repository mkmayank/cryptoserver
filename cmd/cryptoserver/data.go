package main

import "cc/internal/data"

type currencyResponse struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

func (s *state) tickToCurrencyResponse(symbol string, tick data.Tick) currencyResponse {
	return currencyResponse{
		ID:          s.symbolsInfo[symbol].BaseCurrency,
		FullName:    s.symbolsInfo[symbol].FullName,
		Ask:         tick.Ask,
		Bid:         tick.Bid,
		Last:        tick.Last,
		Open:        tick.Open,
		Low:         tick.Low,
		High:        tick.High,
		FeeCurrency: s.symbolsInfo[symbol].FeeCurrency,
	}
}
