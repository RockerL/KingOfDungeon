package shared

import "time"

type UserTicker struct {
	ticker   *time.Ticker
	closeSig chan bool
}

type OnTickerFunc func()

func NewUserTicker(d time.Duration, f OnTickerFunc) *UserTicker {
	userTicker := &UserTicker{
		ticker:   time.NewTicker(d),
		closeSig: make(chan bool, 0),
	}

	go func() {
		for {
			select {
			case stop := <-userTicker.closeSig:
				if stop {
					return
				}
			case <-userTicker.ticker.C:
				f()
			}
		}
	}()

	return userTicker
}

func (t *UserTicker) Stop() {
	t.ticker.Stop()
	t.closeSig <- true
}
