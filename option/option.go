// package option provides option types to handle case where a type can either have a value or no value at all.
package option

type (
	// Option is a type that either hold a value ([Some]) or hold no value ([None]).
	// It enforces user to check if there is a value before using it, and therefore is made to avoid `nil pointer dereference` errors.
	Option[T any] interface {
		UnwrapOrDefault(def T) T
		UnwrapOrElse(def func() T) T
	}
	// optionBase implement the [Option] type with default implementation for [Some] and [None].
	optionBase[T any] struct {
		value  T
		isSome bool
	}
	// Some is a generic type, parts of the [Option] type used when a value is defined.
	Some[T any] struct {
		optionBase[T]
	}
	// None is a generic type, parts of the [Option] type used when no value is defined.
	None[T any] struct {
		optionBase[T]
	}
)

// UnwrapOrDefault will unwrap the option type and provide the held value if called on a [Some] or the provided value if called on [None].
func (o optionBase[T]) UnwrapOrDefault(def T) T {
	if o.isSome {
		return o.value
	}
	return def
}

// UnwrapOrDefault will unwrap the option type and provide the held value if called on a [Some] or call the provided closure if called on [None].
func (o optionBase[T]) UnwrapOrElse(def func() T) T {
	if o.isSome {
		return o.value
	}
	return def()
}

// NewSome create a [Some] with the provided value.
func NewSome[T any](value T) Some[T] {
	return Some[T]{optionBase: optionBase[T]{value: value, isSome: true}}
}

// NewNone create a [None].
func NewNone[T any]() None[T] {
	return None[T]{optionBase: optionBase[T]{isSome: false}}
}
