package resulto

// Result is a container type that represents either success (with a value) or failure (with an error).
type Result[T any] struct {
	Value T
	Err   error
	Ok    bool
}

// Success creates a new successful Result with the given value.
func Success[T any](value T) Result[T] {
	return Result[T]{Value: value, Ok: true}
}

// Failure creates a new failed Result with the given error.
func Failure[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

// FailureOf creates a new failed Result with the given error and value.
func FailureOf[T any](err error, _ T) Result[T] {
	return Result[T]{Err: err}
}

// IsOk returns true if the Result is successful.
func (r Result[T]) IsOk() bool {
	return r.Ok
}

// IsErr returns true if the Result is a failure.
func (r Result[T]) IsErr() bool {
	return !r.Ok
}

// Unwrap returns the value contained in a successful Result or panics if the Result is a failure.
func (r Result[T]) Unwrap() T {
	if !r.Ok {
		panic(r.Err)
	}
	return r.Value
}

// UnwrapOr returns the value contained in a successful Result or the provided default value if the Result is a failure.
func (r Result[T]) UnwrapOr(def T) T {
	if !r.Ok {
		return def
	}
	return r.Value
}

// UnwrapErr returns the error contained in a failed Result or panics if the Result is successful.
func (r Result[T]) UnwrapErr() error {
	if r.Ok {
		panic("called UnwrapErr on a successful Result")
	}
	return r.Err
}
