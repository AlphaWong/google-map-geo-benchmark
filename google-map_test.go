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
