//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hymkor/go-iterator"
)

func NewSequence(start, end int) *iterator.Iterator[int] {
	return iterator.New(func() (int, error) {
		if start > end {
			return -1, io.EOF
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
