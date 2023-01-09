package hitbtc

import (
	"cc/internal/config"
	"cc/internal/exchange"
	"cc/internal/logger"
)

var log = logger.Logger{TAG: "hitbtc"}

type HitBtc struct {
	cfg config.HitBtc

	ticker *ticker
}

func Exchange() exchange.Exchange {
	return &HitBtc{
		cfg: config.HitBtcConfig(),
	}
}
