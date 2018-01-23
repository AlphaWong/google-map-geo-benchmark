# Objective
Store the source of the issue of the google map library
https://github.com/googlemaps/google-maps-services-go/issues/125

```
Dropped 1 node (cum <= 0.02s)
      flat  flat%   sum%        cum   cum%
     1.84s 46.00% 46.00%      1.84s 46.00%  github.com/google-map-geo-benchmark.NewGoogleMapInstance
     1.18s 29.50% 75.50%      1.18s 29.50%  github.com/google-map-geo-benchmark.BenchmarkMartinlindheGoogleMapInstance
     0.96s 24.00% 99.50%      2.80s 70.00%  github.com/google-map-geo-benchmark.BenchmarkGoogleMapInstance
```
My benchmark code
```go
package main

import (
	"testing"

	geo "github.com/martinlindhe/google-geolocate"
	"googlemaps.github.io/maps"
)

var (
	googleMapClient *maps.Client
	g               *geo.GoogleGeo
)

func NewGoogleMapInstance() {
	if nil == googleMapClient {
		c, _ := maps.NewClient(maps.WithAPIKey("Insert-API-Key-Here"))
		googleMapClient = c
	}
}

func NewMartinlindheGoogleMapInstance() {
	if nil == g {
		client := geo.NewGoogleGeo("api-key")
		g = client
	}
}

func BenchmarkGoogleMapInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGoogleMapInstance()
	}
}

func BenchmarkMartinlindheGoogleMapInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewMartinlindheGoogleMapInstance()
	}
}
```

Run
```
go test -benchmem -bench=. -cpuprofile=cpu.out
```

Result
```sh
goos: linux
goarch: amd64
pkg: github.com/google-map-geo-benchmark
BenchmarkGoogleMapInstance-4               	1000000000	         2.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkMartinlindheGoogleMapInstance-4   	2000000000	         0.53 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/google-map-geo-benchmark	3.840s
```
