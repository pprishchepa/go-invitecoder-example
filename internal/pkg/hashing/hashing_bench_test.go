package hashing_test

import (
	"sync"
	"testing"

	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/hashing"
)

func BenchmarkHashStringKey(b *testing.B) {
	const key = "foo@example.com"
	const buckets = 3

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for g := 0; g < 1000; g++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				hashing.HashStringKey(key, buckets)
			}()
		}
		wg.Wait()
	}
}
