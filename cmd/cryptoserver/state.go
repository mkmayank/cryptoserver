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
	resChan chan stateRes
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

	closeChan chan struct{}
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
	s.reqChan = make(chan stateReq)
	s.closeChan = make(chan struct{})

	go s.run()

	err = exchange.StartTicker(symbols, s.tickChan)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to start ticker, error: %s", err.Error()))
	}
}

func (s *state) run() {

	ticks := map[string]data.Tick{}

	for {
		select {
		case tick, ok := <-s.tickChan:
			if !ok {
				log.Error("ticker is stopped")
				s.closeChan <- struct{}{}
				return
			}

			for symbol, tick := range tick {
				ticks[symbol] = tick
			}

		case req := <-s.reqChan:

			res := stateRes{
				ticks: map[string]data.Tick{},
			}

			if req.symbol == "" {
				for symbol, tick := range ticks {
					res.ticks[symbol] = tick
				}
			} else {
				if _, ok := ticks[req.symbol]; !ok {
					res.err = fmt.Errorf("unknown symbol: %s", req.symbol)
				} else {
					res.ticks[req.symbol] = ticks[req.symbol]
				}
			}

			req.resChan <- res
		}
	}
}
