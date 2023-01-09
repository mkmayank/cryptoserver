package data

// Tick defines the tick for a given symbol
// todo string fields can be done to float64
// if needed
// but will need implementation on ticker side for conversion
type Tick struct {
	Ask  string
	Bid  string
	Last string
	Open string
	High string
	Low  string
}

type SymbolInfo struct {
	FullName     string
	BaseCurrency string `json:"base_currency"`
	FeeCurrency  string `json:"fee_currency"`
}
