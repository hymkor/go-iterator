//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/hymkor/go-iterator"
)

type HasOrder interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type KeyValuePair[K HasOrder, V any] struct {
	Key   K
	Value V
}

func SortedRange[K HasOrder, V any](m map[K]V) *iterator.Iterator[KeyValuePair[K, V]] {
	keys := make([]K, 0, len(m))
	for key1 := range m {
		keys = append(keys, key1)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return iterator.New(func() (KeyValuePair[K, V], error) {
		if len(keys) <= 0 {
			return KeyValuePair[K, V]{}, iterator.End
		}
		value := KeyValuePair[K, V]{Key: keys[0], Value: m[keys[0]]}
		keys = keys[1:]
		return value, nil
	})
}

func main() {
	sample := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}

	fmt.Println("--- for-range ---")
	for key, val := range sample {
		fmt.Println(key, val)
	}

	fmt.Println("--- SortedRange ---")
	p := SortedRange(sample)
	for p.Next() {
		fmt.Println(p.Value.Key, p.Value.Value)
	}
	if err := p.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
