package periodlimiter

import (
	"testing"
	"time"
)

func TestPeriodlimiter_Limit(t *testing.T) {
	pl := New()

	key := "test1"
	duration := 200 * time.Millisecond
	burst := 2

	result := pl.Limit(key, duration, burst)
	if !result {
		t.Fatal("expected first Limit to return true")
	}

	result = pl.Limit(key, duration, burst)
	if !result {
		t.Fatal("expected second Limit to return true")
	}

	result = pl.Limit(key, duration, burst)
	if result {
		t.Fatal("expected third Limit to return false")
	}

	time.Sleep(222 * time.Millisecond)

	result = pl.Limit(key, duration, burst)
	if !result {
		t.Fatal("expected Limit after sleep to return true")
	}
}
