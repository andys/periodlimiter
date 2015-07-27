package periodlimiter

import (
	"sync"
	"time"
)

type client struct {
	burst   int
	touched time.Time
}

type Periodlimiter struct {
	clients map[string]*client
	mutex   sync.Mutex
}

func New() *Periodlimiter {
	return &Periodlimiter{clients: make(map[string]*client), mutex: sync.Mutex{}}
}

func (pl *Periodlimiter) Limit(key string, period time.Duration, burst int) bool {
	pl.mutex.Lock()
	c, exists := pl.clients[key]
	if !exists {
		c = &client{burst: burst, touched: time.Now()}
		pl.clients[key] = c
	}

	retval := false
	if time.Now().Sub(c.touched) >= period {
		c.touched = time.Now()
		c.burst = burst
	}

	if c.burst > 0 {
		c.burst -= 1
		retval = true
	}

	pl.mutex.Unlock()
	return retval
}
