package main

import (
	"cc/internal/data"
	"cc/internal/exchange"
	"cc/internal/exchange/hitbtc"
	"fmt"
	"strings"
)

type stateReq struct {
	symbol  string
	resChan chan map[string]data.Tick
}

type stateRes struct {
	ticks map[string]data.Tick
	err   error
}

// it will hold all the tickers data
type state struct {
	tickChan    chan map[string]data.Tick
	symbolsInfo map[string]data.SymbolInfo

	reqChan chan stateReq
}

func (s *state) init(symbols []string) {

	var exchange exchange.Exchange = hitbtc.Exchange()

	var err error
	s.symbolsInfo, err = exchange.GetSymbols()
	if err != nil {
		log.Fatal(err.Error())
	}

	// checking if asked symbols are valid or not
	for _, symbol := range symbols {
		if _, ok := s.symbolsInfo[symbol]; !ok {

			validSymbols := []string{}
			for symbolName := range s.symbolsInfo {
				validSymbols = append(validSymbols, symbolName)
			}
			log.Fatal(fmt.Sprintf("unknown symbol: %s, valid symbols: %s", symbol, strings.Join(validSymbols, " ")))
		}
	}

	s.tickChan = make(chan map[string]data.Tick)

	go s.run()

	exchange.StartTicker()
	exchange.SubscribeTicker(symbols, s.tickChan)

}

func (s *state) run() {
	for {
		select {
		case tick := <-s.tickChan:
			fmt.Println(tick)
		}
	}
}
