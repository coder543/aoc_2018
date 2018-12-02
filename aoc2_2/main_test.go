package main

import (
	"testing"

	"github.com/yourbasic/radix"
)

// run with `go test -bench=. -benchmem`
func BenchmarkMain(b *testing.B) {
	ids := readBoxIds()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ids := append([]string{}, ids...)
		radix.Sort(ids)
		findDiffing(ids)
	}
}
