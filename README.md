# xfetchgo

 xfetchgo implements the XFetch algorithm of [Optimal Probabilistic Cache Stampede Prevention](https://www.vldb.org/pvldb/vol8/p886-vattani.pdf) (2015) by Vattani, Chierichetti, and Lowenstein. The XFetch is for efficient cache management in parallel computing environments. It helps prevent cache stampedes and optimizes cache recomputation without the need for coordination between processes.

## Installation

```
go get github.com/YutaMiyake/xfetchgo
```

## Usage

Creating a new cache entry:
```go
cacheEntry := xfetch.NewCacheEntry(
    func() string {
        // Your expensive computation here 
        return "cached value"
    }, 
    xfetch.WithTTL(5 * time.Minute), 
    xfetch.WithDelta(100 * time.Millisecond), 
    xfetch.WithBeta(1.5),
)
```

Checking if the cache is expired and recomputing if necessary:
```go
if cacheEntry.IsExpired() { 
    // Recompute the cache
} else {
    value := cacheEntry.Get() 
    // Use the cached value
}
```

## License
MIT License
