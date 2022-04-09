go-iterator
============

tiny iterator supporter

Example-1
---------

```go
package main

import (
	"io"

	"github.com/hymkor/go-iterator"
)

func Seq(start, end int) *iterator.Iterator[int] {
	return iterator.New(func() (int, error) {
		if start > end {
			return -1, io.EOF
		}
		start++
		return start - 1, nil
	})
}

func main() {
	seq := Seq(1, 10)
	for seq.Next() {
		println(seq.Value)
	}
}
```

```
$ go run example.go
1
2
3
4
5
6
7
8
9
10
```

Example2: sorted keys
---------------------

```go
package main

import (
	"io"
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
			return KeyValuePair[K, V]{}, io.EOF
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

	println("--- for-range ---")
	for key, val := range sample {
		println(key, val)
	}

	println("--- SortedRange ---")
	p := SortedRange(sample)
	for p.Next() {
		println(p.Value.Key, p.Value.Value)
	}
}
```

```
$ go run example2.go
--- for-range ---
C 3
D 4
A 1
B 2
--- SortedRange ---
A 1
B 2
C 3
D 4
```
