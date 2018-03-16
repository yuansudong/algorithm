package main

import (
	"SDAlgorithm/golang/brtree"
	"strconv"
	"testing"
)

/*
goos: windows
goarch: amd64
pkg: SDAlgorithm/golang/test
Benchmark_Insert-4   	 1000000	      1583 ns/op	      95 B/op	       3 allocs/op
PASS
ok  	SDAlgorithm/golang/test	2.199s
*/
func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	BRT := brtree.Create()
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		brtree.Insert(BRT, strconv.Itoa(i), i)
	}
}

func Benchmark_Get(B *testing.B) {
	B.StopTimer()
	BRT := brtree.Create()
	for i := 0; i < B.N; i++ { //use b.N for looping
		brtree.Insert(BRT, strconv.Itoa(i), i)
	}
	B.StartTimer()
	brtree.Get(BRT, "51452")
}
