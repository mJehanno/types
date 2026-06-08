// package result provides result types to handle case where a return type can either have a value or an error.
package result

type (
	// Result is a type that either hold a value ([Ok]) or an error ([Err]).
	// It enforces user to check if a funcction returns a value or an error before using it.
	Result[T any] interface {
		UnwrapOrDefault(def T) T
		UnwrapOrElse(def func() T) T
	}

	// resultBase implement the [Result] type with default implementation for [Ok] and [Err].
	resultBase[T any] struct {
		value T
		err   error
	}

	// Ok is a generic type, parts of the [Result] type used when a function returns a value.
	Ok[T any] struct {
		resultBase[T]
	}
	// Err is a generic type, parts of the [Result] type used when a function returns an error.
	Err[T any] struct {
		resultBase[T]
	}
)

// UnwrapOrDefault will unwrap the result type and provide the held value if called on a [Ok] or the provided value if called on [Err].
func (r resultBase[T]) UnwrapOrDefault(def T) T {
	if r.err != nil {
		return def
	}
	return r.value
}

// UnwrapOrDefault will unwrap the option type and provide the held value if called on a [Ok] or call the provided closure if called on [Err].
func (r resultBase[T]) UnwrapOrElse(def func() T) T {
	if r.err != nil {
		return def()
	}
	return r.value
}

// GetError returns the underlying error of the [Err] type.
func (e Err[T]) GetError() error {
	return e.err
}

// NewOk create a [Ok] with the provided value.
func NewOk[T any](value T) Ok[T] {
	return Ok[T]{resultBase: resultBase[T]{value: value}}
}

// NewErr create a [Err].
func NewErr[T any](err error) Err[T] {
	return Err[T]{resultBase: resultBase[T]{err: err}}
}
