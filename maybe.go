package maybe

import "fmt"

type Maybe[T any] struct {
	value    T
	hasValue bool
	checked  bool // HasValue() vas called (experimental)
}

func (m Maybe[T]) String() string {
	if m.HasValue() {
		v, _ := m.FromJust()
		return fmt.Sprintf("Just %v", v)
	} else {
		return "Nothing"
	}
}

func (m *Maybe[T]) HasValue() bool {
	m.checked = true
	return m.hasValue
}

func (m *Maybe[T]) FromJust() (T, error) {
	if m.hasValue {
		return m.value, nil
	} else {
		m.checked = false
		return *new(T), fmt.Errorf("Nothing")
	}
}

func Just[T any](val T) Maybe[T] { return Maybe[T]{value: val, hasValue: true} }

func Nothing[T any]() Maybe[T] { return Maybe[T]{hasValue: false} }
