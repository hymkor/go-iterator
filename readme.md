go-iterator
============

tiny iterator supporter

sample
------

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
