package hashing

import (
	"hash"
	"hash/fnv"
	"io"
	"sync"

	"github.com/dgryski/go-jump"
)

var hash64Pool = sync.Pool{New: func() any {
	return fnv.New64a()
}}

func HashStringKey(key string, buckets int) int32 {
	h := hash64Pool.Get().(hash.Hash64)
	h.Reset()
	defer hash64Pool.Put(h)

	_, _ = io.WriteString(h, key)

	return jump.Hash(h.Sum64(), buckets)
}
