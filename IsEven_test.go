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

func BenchmarkInserXInMapInterface100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100000, b)
	}
}
func BenchmarkInserXInMapInterface1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(1000, b)
	}
}
func BenchmarkInserXInMapInterface100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXInMapInterface(100, b)
	}
}
