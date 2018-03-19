package main

import (
	"SDAlgorithm/golang/slice"
	"testing"
)

func Benchmark_Append(b *testing.B) {
	b.StopTimer()
	S := slice.Create()
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		S.Append(i)
	}
}

func Benchmark_Delete(b *testing.B) {
	b.StopTimer()
	S := slice.Create()
	b.Log(b.N)
	for i := 0; i < b.N; i++ { //use b.N for looping
		S.Append(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		S.Delete(i)
	}
}
