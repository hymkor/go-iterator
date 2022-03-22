package iterator

type Iterator[T any] struct {
	next  func() (T, error)
	Err   error
	Value T
}

func New[T any](f func() (T, error)) *Iterator[T] {
	return &Iterator[T]{next: f}
}

func (I *Iterator[T]) Next() bool {
	I.Value, I.Err = I.next()
	return I.Err == nil
}
