### Limit number of events over a set time period

This Go package provides a generic way to limit the number of occurences of an event over time.

It allows a burst of hits within a specific time period and then blocks until the next time period.


#### Install
```shell
$ go get github.com/andys/periodlimiter
```


#### Usage

```go
import (
  "time"
  "github.com/andys/periodlimiter"
)

// Create a long-lived handle
periodLimiter := periodlimiter.New()

// Limits bob to a burst of 5 requests (at any speed) over 1 hour of elapsed time
allowed := periodLimiter.Limit("bob", time.Hour, 5)

// returns true to allow the event, or false to block it
```

