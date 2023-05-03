package iterator_test

import (
	"testing"

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

func TestIterator(t *testing.T) {
	seq := NewSequence(1, 10)
	expect := 1
	for seq.Next() {
		if seq.Value != expect {
			t.Fatalf("expect %v but %v", expect, seq.Value)
		}
		expect++
	}
	if err := seq.Err(); err != nil {
		t.Fatal(err.Error())
	}
}
