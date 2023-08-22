package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []bool
	size   int32
}

func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{filter: make([]bool, size),
		size: size,
	}
}

var hasher hash.Hash32

func init() {
	hasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func murmurHash(key string, size int32) int {
	hasher.Write([]byte(key))
	hashValue := hasher.Sum32()
	result := hashValue % uint32(size)
	hasher.Reset()
	return int(result)
}

func (b *BloomFilter) addItem(key string) {

	idx := murmurHash(key, b.size)

	b.filter[idx] = true
}
func (b *BloomFilter) bloomContains(key string) bool {
	idx := murmurHash(key, b.size)
	return b.filter[idx]
}
func main() {
	bloom := NewBloomFilter(16)

	keys := []string{"dante", "bruise", "bilegr"}

	for _, key := range keys {
		bloom.addItem(key)
	}

	for _, key := range keys {
		r := bloom.bloomContains(key)
		fmt.Println("Subscribed to Membership", r)
		fmt.Println(key)
	}

	fmt.Println(bloom.bloomContains("Ambur"))
	fmt.Println(bloom.bloomContains("Agraz"))
	fmt.Println(bloom.bloomContains("Shubham")) // overlapping scenario key not present still showing
}
