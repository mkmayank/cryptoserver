package hitbtc

type tick struct {
	T  int64  `json:"t"`
	A  string `json:"a"`
	A0 string `json:"A"`
	B  string `json:"b"`
	B0 string `json:"B"`
	C  string `json:"c"`
	O  string `json:"o"`
	H  string `json:"h"`
	L  string `json:"l"`
	V  string `json:"v"`
	Q  string `json:"q"`
	P  string `json:"p"`
	P0 string `json:"P"`
	L0 int    `json:"L"`
}

type socketResponse struct {
	Result struct {
		Ch            string   `json:"ch"`
		Subscriptions []string `json:"subscriptions"`
	} `json:"result"`
	ID int `json:"id"`
}

type socketNotification struct {
	Ch   string          `json:"ch"`
	Data map[string]tick `json:"data"`
}

// tickerResponse is the data received on ticker
type tickerResponse struct {
	Ch   string          `json:"ch"`
	Data map[string]tick `json:"data"`
}

type symbolInfo struct {
	Type              string `json:"type"`
	BaseCurrency      string `json:"base_currency"`
	QuoteCurrency     string `json:"quote_currency"`
	Status            string `json:"status"`
	QuantityIncrement string `json:"quantity_increment"`
	TickSize          string `json:"tick_size"`
	TakeRate          string `json:"take_rate"`
	MakeRate          string `json:"make_rate"`
	FeeCurrency       string `json:"fee_currency"`
}

type currencyInfo struct {
	FullName          string `json:"full_name"`
	Crypto            bool   `json:"crypto"`
	PayinEnabled      bool   `json:"payin_enabled"`
	PayoutEnabled     bool   `json:"payout_enabled"`
	TransferEnabled   bool   `json:"transfer_enabled"`
	PrecisionTransfer string `json:"precision_transfer"`
	Networks          []struct {
		Network            string `json:"network"`
		Protocol           string `json:"protocol"`
		Default            bool   `json:"default"`
		PayinEnabled       bool   `json:"payin_enabled"`
		PayoutEnabled      bool   `json:"payout_enabled"`
		PrecisionPayout    string `json:"precision_payout"`
		PayoutFee          string `json:"payout_fee"`
		PayoutIsPaymentID  bool   `json:"payout_is_payment_id"`
		PayinPaymentID     bool   `json:"payin_payment_id"`
		PayinConfirmations int    `json:"payin_confirmations"`
	} `json:"networks"`
}
