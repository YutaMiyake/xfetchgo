package xfetch

import (
	"math"
	"math/rand"
	"time"
)

type options struct {
	// delta is the time for recomputation, the higher the delta, the earlier the recomputation
	delta time.Duration

	// beta is a scaling (1.0 by default), the higher the beta, the earlier the recomputation
	beta float64

	// ttl is a time-to-live of the cache, 0 means no expiration, negative value means expired
	ttl time.Duration
}

type CacheEntryOption func(*options)

func WithDelta(delta time.Duration) CacheEntryOption {
	return func(o *options) {
		o.delta = delta
	}
}

func WithBeta(beta float64) CacheEntryOption {
	return func(o *options) {
		o.beta = beta
	}
}

func WithTTL(ttl time.Duration) CacheEntryOption {
	return func(o *options) {
		o.ttl = ttl
	}
}

const defaultBeta = 1.0

// NewCacheEntry creates a new CacheEntry instance.
func NewCacheEntry[T any](valueFunc func() T, opts ...CacheEntryOption) *CacheEntry[T] {
	// Perform expensive computation e.g., query database, network request, etc.
	start := time.Now()
	value := valueFunc()
	recomputeTime := time.Since(start)

	o := options{
		delta: recomputeTime,
		beta:  defaultBeta,
		ttl:   0,
	}
	for _, opt := range opts {
		opt(&o)
	}

	c := &CacheEntry[T]{
		value: value,
		delta: o.delta,
		beta:  o.beta,
	}

	if o.ttl != 0 {
		c.expiry = time.Now().Add(o.ttl)
	}

	return c
}

// CacheEntry represents a cache entry with probabilistic early expiration mechanism.
type CacheEntry[T any] struct {
	value  T
	delta  time.Duration
	beta   float64
	expiry time.Time
}

// IsExpired checks whether the cache entry has expired.
func (e CacheEntry[T]) IsExpired() bool {
	return e.isExpiredWithRNG()
}

func (e CacheEntry[T]) isExpiredWithRNG() bool {
	if !e.expiry.IsZero() {
		now := time.Now()
		delta := float64(e.delta.Nanoseconds())
		xfetchTime := time.Duration(delta * e.beta * -math.Log(rand.Float64()))
		return now.Add(xfetchTime).After(e.expiry)
	}
	return false
}

// Get returns the value stored in the cache entry.
func (e CacheEntry[T]) Get() T {
	return e.value
}
