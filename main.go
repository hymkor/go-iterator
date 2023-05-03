package iterator

import (
	"errors"
	"io"
)

var End = errors.New("Iterator End")

type Iterator[T any] struct {
	next  func() (T, error)
	err   error
	Value T
}

// New returns a new Iterator.
func New[T any](f func() (T, error)) *Iterator[T] {
	return &Iterator[T]{next: f}
}

// Next advances the Iterator[T] to the next element.
func (I *Iterator[T]) Next() bool {
	I.Value, I.err = I.next()
	return I.err == nil
}

// Err returns the first non-EOF error that was encountered by the Iterator[T]
func (I *Iterator[T]) Err() error {
	if I.err == io.EOF || I.err == End {
		return nil
	}
	return I.err
}
