go-iterator
============

[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-iterator.svg)](https://pkg.go.dev/github.com/hymkor/go-iterator)

A simple iterator supporter

Example
-------

```example.go
package main

import (
    "fmt"
    "os"

    "github.com/hymkor/go-iterator"
)

func NewSequence(start, end int) *iterator.Iterator[int] {
    return iterator.New(func() (int, error) {
        if start > end {
            return -1, iterator.End
        }
        start++
        return start - 1, nil
    })
}

func main() {
    seq := NewSequence(1, 10)
    for seq.Next() {
        fmt.Println(seq.Value)
    }
    if err := seq.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
    }
}
```

```go run example.go|
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
