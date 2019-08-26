package main

import "testing"

func Benchmark_login(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		login(nil)
	}
}

func Benchmark_division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		division(4,5)
	}
}