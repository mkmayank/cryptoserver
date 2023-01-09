package exchange

import "cc/internal/data"

// Exchange defines all the required methods a exchange required to have
type Exchange interface {

	// GetSymbols will return all available symbol->info map
	GetSymbols() (map[string]data.SymbolInfo, error)

	// StartTicker will start a ticker
	// may be websocket or some state if required so
	StartTicker(symbols []string, subscriberChan chan<- map[string]data.Tick) (err error)

	// if we want to allow resubscription
	// todo SubscribeTicker
	// todo UnsubscribeTicker
	// todo UnsubscribeSymbol
}
