package hitbtc

import (
	"cc/internal/data"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

// todo
// need auto-reconnect
// and checking for ping
// handling of errors

type tickerInput struct {
	Method string `json:"method"`
	Ch     string `json:"ch"`
	Params struct {
		Symbols []string `json:"symbols"`
	} `json:"params"`
	ID int `json:"id"`
}

type ticker struct {
	Conn *websocket.Conn
	id   int

	url url.URL

	subscribedSymbols []string
	subscriberChan    chan<- map[string]data.Tick
}

func (t *ticker) connect() (err error) {

	log.Info("trying to connect")

	conn, _, err := websocket.DefaultDialer.Dial(t.url.String(), nil)
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("started a new connection %p", conn))

	// Assign the current connection to the instance.
	t.Conn = conn

	err = t.subscribe()
	if err != nil {
		return
	}
	go t.readMessage()
	return
}

func (t *ticker) subscribe() (err error) {

	out, err := json.Marshal(tickerInput{
		Method: "subscribe",
		Ch:     "ticker/1s",
		Params: struct {
			// todo, not a good practice to have anonymous struct here
			Symbols []string `json:"symbols"`
		}{
			Symbols: t.subscribedSymbols,
		},
		ID: t.id,
	})

	if err != nil {
		log.Error(err.Error())
		return
	}

	return t.Conn.WriteMessage(websocket.TextMessage, out)
}

// readMessage reads the data in a loop.
func (t *ticker) readMessage() {
	for {
		mType, msg, err := t.Conn.ReadMessage()
		if err != nil {
			log.Error(err.Error())
			return
		}

		if mType == websocket.TextMessage {
			t.processTextMessage(msg)
		} else {
			log.Error("got binary message")
		}
	}
}

func (t *ticker) processTextMessage(inp []byte) {

	// todo sanity check
	msg := socketNotification{}

	if err := json.Unmarshal(inp, &msg); err != nil {
		log.Error(fmt.Sprintf("processing text message %s, error: %s", inp, err.Error()))
		return
	}

	dataToSend := map[string]data.Tick{}
	for symbol, tick := range msg.Data {
		dataToSend[symbol] = data.Tick{
			Ask:  tick.A,
			Bid:  tick.B,
			Last: tick.C,
			Open: tick.O,
			High: tick.H,
			Low:  tick.L,
		}
	}

	t.subscriberChan <- dataToSend
}

func (h *HitBtc) StartTicker(symbols []string, subscriberChan chan<- map[string]data.Tick) (err error) {

	if h.ticker != nil {
		err = fmt.Errorf("asked to start ticker again")
		return
	}

	if len(symbols) == 0 {
		err = fmt.Errorf("symbol list is empty")
		return
	}

	h.ticker = &ticker{
		url: url.URL{Scheme: "wss", Host: h.cfg.WsAPIHost, Path: h.cfg.WsAPIPath},
		// can be made configurable
		id:                123,
		subscribedSymbols: symbols,
		subscriberChan:    subscriberChan,
	}

	err = h.ticker.connect()
	return
}
