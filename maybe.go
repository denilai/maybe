package maybe

import "fmt"

type Maybe[T any] struct {
	value    T
	hasValue bool
}

func (m Maybe[T]) String() string {
	if m.HasValue() {
		return fmt.Sprintf("Just %v", m.FromJust())
	} else {
		return "Nothing"
	}
}

func (m Maybe[T]) HasValue() bool {
	return m.hasValue
}

func (m Maybe[T]) FromJust() T {
	if m.hasValue {
		return m.value
	} else {
		panic("Nothing")
	}
}

func Just[T any](val T) Maybe[T] { return Maybe[T]{value: val, hasValue: true} }

func Nothing[T any]() Maybe[T] { return Maybe[T]{hasValue: false} }
