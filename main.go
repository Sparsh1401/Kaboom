package main

type BloomFilter struct {
	filter []bool
}

func (*BloomFilter) NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{filter: make([]bool, size)}
}
func main() {
	bloom := NewBloomFilter(16)
}
