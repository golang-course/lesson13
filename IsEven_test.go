package main

import "testing"

func BenchmarkInserXInMap100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100000, b)
	}
}
func BenchmarkInserXInMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(1000, b)
	}
}
func BenchmarkInserXInMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMap(100, b)
	}
}

func BenchmarkInserXInSlice100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100000, b)
	}
}
func BenchmarkInserXInSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(1000, b)
	}
}
func BenchmarkInserXInSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInSlice(100, b)
	}
}
