//go:build ignore
// +build ignore

package main

import (
	"sort"
)

type HasOrder interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type KeyValuePair[K HasOrder, V any] struct {
	Key   K
	Value V
}

func SortedRange[K HasOrder, V any](m map[K]V) []KeyValuePair[K, V] {
	pairs := make([]KeyValuePair[K, V], 0, len(m))
	for key, val := range m {
		pairs = append(pairs, KeyValuePair[K, V]{Key: key, Value: val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})

	return pairs
}

func main() {
	sample := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	kv := SortedRange(sample)
	for _, p := range kv {
		println(p.Key, p.Value)
	}
}
