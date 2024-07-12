package xfetch

import (
	"testing"
	"time"
)

func TestCacheEntry_NoExpire(t *testing.T) {
	t.Parallel()

	entry := NewCacheEntry(func() interface{} {
		return "test value"
	})
	if entry.IsExpired() {
		t.Errorf("Expected entry not to be expired, but it is expired")
	}
	if entry.expiry != (time.Time{}) {
		t.Errorf("Expected expiry time to be zero, but it is not zero")
	}
}

func TestCacheEntry_IsExpired(t *testing.T) {
	t.Parallel()

	entry := NewCacheEntry(func() interface{} {
		return "test value"
	}, WithTTL(500*time.Millisecond))

	if entry.IsExpired() {
		t.Errorf("Expected entry not to be expired, but it is expired")
	}

	// Wait for the TTL to expire
	time.Sleep(500 * time.Millisecond)

	if !entry.IsExpired() {
		t.Errorf("Expected entry to be expired, but it is not expired")
	}
}

func TestCacheEntry_WithoutDelta(t *testing.T) {
	t.Parallel()

	entry := NewCacheEntry(func() interface{} {
		time.Sleep(100 * time.Millisecond)
		return "test value"
	}, WithDelta(500*time.Millisecond))

	if entry.delta < 100*time.Millisecond {
		t.Errorf("Expected delta to be greater than 100ms, but got %s", entry.delta)
	}
}

func TestCacheEntry_WithDelta(t *testing.T) {
	t.Parallel()

	entry := NewCacheEntry(func() interface{} {
		return "test value"
	}, WithDelta(500*time.Millisecond))

	if entry.delta != 500*time.Millisecond {
		t.Errorf("Expected delta to be 500ms, but got %s", entry.delta)
	}
}

func TestCacheEntry_WithBeta(t *testing.T) {
	t.Parallel()

	entry := NewCacheEntry(func() interface{} {
		return "test value"
	}, WithBeta(0.5))

	if entry.beta != 0.5 {
		t.Errorf("Expected beta to be 0.5, but got %f", entry.beta)
	}
}

func TestCacheEntry_Get(t *testing.T) {
	t.Parallel()

	value := "test value"
	entry := NewCacheEntry(func() interface{} {
		return value
	})
	result := entry.Get().(string)
	if result != value {
		t.Errorf("Retrieved value (%s) does not match expected value (%s)", result, value)
	}
}
