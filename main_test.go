package main

import (
	"os"
	"testing"
)

func BenchmarkJson(b *testing.B) {
	data, err := os.ReadFile("ticker.txt")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Run function to benchmark here
		JsonUnmarshal(&data)
	}
}

func BenchmarkByteSlice(b *testing.B) {
	data, err := os.ReadFile("ticker.txt")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Run function to benchmark here
		SliceByteMap(&data)
	}
}
