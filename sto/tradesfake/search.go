package tradesfake

import (
	"time"

	"github.com/phoebetron/trades/typ/trades"
)

func (f *Fake) Search(time.Time) (*trades.Trades, error) {
	return nil, nil
}
